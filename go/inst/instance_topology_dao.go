/*
   Copyright 2015 Shlomi Noach, courtesy Booking.com

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package inst

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/outbrain/golib/log"
	"github.com/outbrain/golib/sqlutils"
	"github.com/outbrain/orchestrator/go/config"
	"github.com/outbrain/orchestrator/go/db"
)

// Max concurrency for bulk topology operations
const topologyConcurrency = 128

var topologyConcurrencyChan = make(chan bool, topologyConcurrency)

type OperationGTIDHint string

const (
	GTIDHintDeny    OperationGTIDHint = "NoGTID"
	GTIDHintNeutral                   = "GTIDHintNeutral"
	GTIDHintForce                     = "GTIDHintForce"
)

const sqlThreadPollDuration = 400 * time.Millisecond

// ExecInstance executes a given query on the given MySQL topology instance
func ExecInstance(instanceKey *InstanceKey, query string, args ...interface{}) (sql.Result, error) {
	db, err := db.OpenTopology(instanceKey.Hostname, instanceKey.Port)
	if err != nil {
		return nil, err
	}
	res, err := sqlutils.Exec(db, query, args...)
	return res, err
}

// ExecInstanceNoPrepare executes a given query on the given MySQL topology instance, without using prepared statements
func ExecInstanceNoPrepare(instanceKey *InstanceKey, query string, args ...interface{}) (sql.Result, error) {
	db, err := db.OpenTopology(instanceKey.Hostname, instanceKey.Port)
	if err != nil {
		return nil, err
	}
	res, err := sqlutils.ExecNoPrepare(db, query, args...)
	return res, err
}

// ExecuteOnTopology will execute given function while maintaining concurrency limit
// on topology servers. It is safe in the sense that we will not leak tokens.
func ExecuteOnTopology(f func()) {
	topologyConcurrencyChan <- true
	defer func() { recover(); <-topologyConcurrencyChan }()
	f()
}

// ScanInstanceRow executes a read-a-single-row query on a given MySQL topology instance
func ScanInstanceRow(instanceKey *InstanceKey, query string, dest ...interface{}) error {
	db, err := db.OpenTopology(instanceKey.Hostname, instanceKey.Port)
	if err != nil {
		return err
	}
	err = db.QueryRow(query).Scan(dest...)
	return err
}

// EmptyCommitInstance issues an empty COMMIT on a given instance
func EmptyCommitInstance(instanceKey *InstanceKey) error {
	db, err := db.OpenTopology(instanceKey.Hostname, instanceKey.Port)
	if err != nil {
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return err
}

// RefreshTopologyInstance will synchronuously re-read topology instance
func RefreshTopologyInstance(instanceKey *InstanceKey) (*Instance, error) {
	_, err := ReadTopologyInstanceUnbuffered(instanceKey)
	if err != nil {
		return nil, err
	}

	inst, found, err := ReadInstance(instanceKey)
	if err != nil || !found {
		return nil, err
	}

	return inst, nil
}

// RefreshTopologyInstances will do a blocking (though concurrent) refresh of all given instances
func RefreshTopologyInstances(instances [](*Instance)) {
	// use concurrency but wait for all to complete
	barrier := make(chan InstanceKey)
	for _, instance := range instances {
		instance := instance
		go func() {
			// Signal completed slave
			defer func() { barrier <- instance.Key }()
			// Wait your turn to read a slave
			ExecuteOnTopology(func() {
				log.Debugf("... reading instance: %+v", instance.Key)
				ReadTopologyInstanceUnbuffered(&instance.Key)
			})
		}()
	}
	for range instances {
		<-barrier
	}
}

// RefreshInstanceSlaveHosts is a workaround for a bug in MySQL where
// SHOW SLAVE HOSTS continues to present old, long disconnected slaves.
// It turns out issuing a couple FLUSH commands mitigates the problem.
func RefreshInstanceSlaveHosts(instanceKey *InstanceKey) (*Instance, error) {
	_, _ = ExecInstance(instanceKey, `flush error logs`)
	_, _ = ExecInstance(instanceKey, `flush error logs`)

	instance, err := ReadTopologyInstanceUnbuffered(instanceKey)
	return instance, err
}

// GetSlaveRestartPreserveStatements returns a sequence of statements that make sure a slave is stopped
// and then returned to the same state. For example, if the slave was fully running, this will issue
// a STOP on both io_thread and sql_thread, followed by START on both. If one of them is not running
// at the time this function is called, said thread will be neither stopped nor started.
// The caller may provide an injected statememt, to be executed while the slave is stopped.
// This is useful for CHANGE MASTER TO commands, that unfortunately must take place while the slave
// is completely stopped.
func GetSlaveRestartPreserveStatements(instanceKey *InstanceKey, injectedStatement string) (statements []string, err error) {
	instance, err := ReadTopologyInstanceUnbuffered(instanceKey)
	if err != nil {
		return statements, err
	}
	if instance.Slave_IO_Running {
		statements = append(statements, SemicolonTerminated(`stop slave io_thread`))
	}
	if instance.Slave_SQL_Running {
		statements = append(statements, SemicolonTerminated(`stop slave sql_thread`))
	}
	if injectedStatement != "" {
		statements = append(statements, SemicolonTerminated(injectedStatement))
	}
	if instance.Slave_SQL_Running {
		statements = append(statements, SemicolonTerminated(`start slave sql_thread`))
	}
	if instance.Slave_IO_Running {
		statements = append(statements, SemicolonTerminated(`start slave io_thread`))
	}
	return statements, err
}

// FlushBinaryLogs attempts a 'FLUSH BINARY LOGS' statement on the given instance.
func FlushBinaryLogs(instanceKey *InstanceKey, count int) (*Instance, error) {
	if *config.RuntimeCLIFlags.Noop {
		return nil, fmt.Errorf("noop: aborting flush-binary-logs operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}

	for i := 0; i < count; i++ {
		_, err := ExecInstance(instanceKey, `flush binary logs`)
		if err != nil {
			return nil, log.Errore(err)
		}
	}

	log.Infof("flush-binary-logs count=%+v on %+v", count, *instanceKey)
	AuditOperation("flush-binary-logs", instanceKey, "success")

	return ReadTopologyInstanceUnbuffered(instanceKey)
}

// FlushBinaryLogsTo attempts to 'FLUSH BINARY LOGS' until given binary log is reached
func FlushBinaryLogsTo(instanceKey *InstanceKey, logFile string) (*Instance, error) {
	instance, err := ReadTopologyInstanceUnbuffered(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	distance := instance.SelfBinlogCoordinates.FileNumberDistance(&BinlogCoordinates{LogFile: logFile})
	if distance < 0 {
		return nil, log.Errorf("FlushBinaryLogsTo: target log file %+v is smaller than current log file %+v", logFile, instance.SelfBinlogCoordinates.LogFile)
	}
	return FlushBinaryLogs(instanceKey, distance)
}

// FlushBinaryLogsTo attempts to 'PURGE BINARY LOGS' until given binary log is reached
func PurgeBinaryLogsTo(instanceKey *InstanceKey, logFile string) (*Instance, error) {
	if *config.RuntimeCLIFlags.Noop {
		return nil, fmt.Errorf("noop: aborting purge-binary-logs operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}

	_, err := ExecInstanceNoPrepare(instanceKey, fmt.Sprintf("purge binary logs to '%s'", logFile))
	if err != nil {
		return nil, log.Errore(err)
	}

	log.Infof("purge-binary-logs to=%+v on %+v", logFile, *instanceKey)
	AuditOperation("purge-binary-logs", instanceKey, "success")

	return ReadTopologyInstanceUnbuffered(instanceKey)
}

// FlushBinaryLogsTo attempts to 'PURGE BINARY LOGS' until given binary log is reached
func PurgeBinaryLogsToCurrent(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstanceUnbuffered(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}
	return PurgeBinaryLogsTo(instanceKey, instance.SelfBinlogCoordinates.LogFile)
}

// StopSlaveNicely stops a slave such that SQL_thread and IO_thread are aligned (i.e.
// SQL_thread consumes all relay log entries)
// It will actually START the sql_thread even if the slave is completely stopped.
func StopSlaveNicely(instanceKey *InstanceKey, timeout time.Duration) (*Instance, error) {
	instance, err := ReadTopologyInstanceUnbuffered(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if !instance.IsSlave() {
		return instance, fmt.Errorf("instance is not a slave: %+v", instanceKey)
	}

	_, err = ExecInstanceNoPrepare(instanceKey, `stop slave io_thread`)
	_, err = ExecInstanceNoPrepare(instanceKey, `start slave sql_thread`)

	if instance.SQLDelay == 0 {
		// Otherwise we don't bother.
		startTime := time.Now()
		for upToDate := false; !upToDate; {
			if timeout > 0 && time.Since(startTime) >= timeout {
				// timeout
				return nil, log.Errorf("StopSlaveNicely timeout on %+v", *instanceKey)
			}
			instance, err = ReadTopologyInstanceUnbuffered(instanceKey)
			if err != nil {
				return instance, log.Errore(err)
			}

			if instance.SQLThreadUpToDate() {
				upToDate = true
			} else {
				time.Sleep(sqlThreadPollDuration)
			}
		}
	}
	_, err = ExecInstanceNoPrepare(instanceKey, `stop slave`)
	if err != nil {
		// Patch; current MaxScale behavior for STOP SLAVE is to throw an error if slave already stopped.
		if instance.isMaxScale() && err.Error() == "Error 1199: Slave connection is not running" {
			err = nil
		}
	}
	if err != nil {
		return instance, log.Errore(err)
	}

	instance, err = ReadTopologyInstanceUnbuffered(instanceKey)
	log.Infof("Stopped slave nicely on %+v, Self:%+v, Exec:%+v", *instanceKey, instance.SelfBinlogCoordinates, instance.ExecBinlogCoordinates)
	return instance, err
}

// StopSlavesNicely will attemt to stop all given slaves nicely, up to timeout
func StopSlavesNicely(slaves [](*Instance), timeout time.Duration) [](*Instance) {
	refreshedSlaves := [](*Instance){}

	log.Debugf("Stopping %d slaves nicely", len(slaves))
	// use concurrency but wait for all to complete
	barrier := make(chan *Instance)
	for _, slave := range slaves {
		slave := slave
		go func() {
			updatedSlave := &slave
			// Signal completed slave
			defer func() { barrier <- *updatedSlave }()
			// Wait your turn to read a slave
			ExecuteOnTopology(func() {
				StopSlaveNicely(&slave.Key, timeout)
				slave, _ = StopSlave(&slave.Key)
				updatedSlave = &slave
			})
		}()
	}
	for range slaves {
		refreshedSlaves = append(refreshedSlaves, <-barrier)
	}
	return refreshedSlaves
}

// StopSlave stops replication on a given instance
func StopSlave(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstanceUnbuffered(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if !instance.IsSlave() {
		return instance, fmt.Errorf("instance is not a slave: %+v", instanceKey)
	}
	_, err = ExecInstanceNoPrepare(instanceKey, `stop slave`)
	if err != nil {
		// Patch; current MaxScale behavior for STOP SLAVE is to throw an error if slave already stopped.
		if instance.isMaxScale() && err.Error() == "Error 1199: Slave connection is not running" {
			err = nil
		}
	}
	if err != nil {

		return instance, log.Errore(err)
	}
	instance, err = ReadTopologyInstanceUnbuffered(instanceKey)

	log.Infof("Stopped slave on %+v, Self:%+v, Exec:%+v", *instanceKey, instance.SelfBinlogCoordinates, instance.ExecBinlogCoordinates)
	return instance, err
}

// StartSlave starts replication on a given instance.
func StartSlave(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstanceUnbuffered(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if !instance.IsSlave() {
		return instance, fmt.Errorf("instance is not a slave: %+v", instanceKey)
	}

	// If async fallback is disallowed, we'd better make sure to enable slaves to
	// send ACKs before START SLAVE. Slave ACKing is off at mysqld startup because
	// some slaves (those that must never be promoted) should never ACK.
	// Note: We assume that slaves use 'skip-slave-start' so they won't
	//       START SLAVE on their own upon restart.
	if instance.SemiSyncEnforced {
		// Send ACK only from promotable instances.
		sendACK := instance.PromotionRule != MustNotPromoteRule
		// Always disable master setting, in case we're converting a former master.
		if err := EnableSemiSync(instanceKey, false, sendACK); err != nil {
			return instance, log.Errore(err)
		}
	}

	_, err = ExecInstanceNoPrepare(instanceKey, `start slave`)
	if err != nil {
		return instance, log.Errore(err)
	}
	log.Infof("Started slave on %+v", instanceKey)
	if config.Config.SlaveStartPostWaitMilliseconds > 0 {
		time.Sleep(time.Duration(config.Config.SlaveStartPostWaitMilliseconds) * time.Millisecond)
	}

	instance, err = ReadTopologyInstanceUnbuffered(instanceKey)
	return instance, err
}

// RestartSlave stops & starts replication on a given instance
func RestartSlave(instanceKey *InstanceKey) (instance *Instance, err error) {
	instance, err = StopSlave(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}
	instance, err = StartSlave(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}
	return instance, nil

}

// StartSlaves will do concurrent start-slave
func StartSlaves(slaves [](*Instance)) {
	// use concurrency but wait for all to complete
	log.Debugf("Starting %d slaves", len(slaves))
	barrier := make(chan InstanceKey)
	for _, instance := range slaves {
		instance := instance
		go func() {
			// Signal compelted slave
			defer func() { barrier <- instance.Key }()
			// Wait your turn to read a slave
			ExecuteOnTopology(func() { StartSlave(&instance.Key) })
		}()
	}
	for range slaves {
		<-barrier
	}
}

// StartSlaveUntilMasterCoordinates issuesa START SLAVE UNTIL... statement on given instance
func StartSlaveUntilMasterCoordinates(instanceKey *InstanceKey, masterCoordinates *BinlogCoordinates) (*Instance, error) {
	instance, err := ReadTopologyInstanceUnbuffered(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if !instance.IsSlave() {
		return instance, fmt.Errorf("instance is not a slave: %+v", instanceKey)
	}
	if instance.SlaveRunning() {
		return instance, fmt.Errorf("slave already running: %+v", instanceKey)
	}

	log.Infof("Will start slave on %+v until coordinates: %+v", instanceKey, masterCoordinates)

	if instance.SemiSyncEnforced {
		// Send ACK only from promotable instances.
		sendACK := instance.PromotionRule != MustNotPromoteRule
		// Always disable master setting, in case we're converting a former master.
		if err := EnableSemiSync(instanceKey, false, sendACK); err != nil {
			return instance, log.Errore(err)
		}
	}

	// MariaDB has a bug: a CHANGE MASTER TO statement does not work properly with prepared statement... :P
	// See https://mariadb.atlassian.net/browse/MDEV-7640
	// This is the reason for ExecInstanceNoPrepare
	_, err = ExecInstanceNoPrepare(instanceKey, fmt.Sprintf("start slave until master_log_file='%s', master_log_pos=%d",
		masterCoordinates.LogFile, masterCoordinates.LogPos))
	if err != nil {
		return instance, log.Errore(err)
	}

	for upToDate := false; !upToDate; {
		instance, err = ReadTopologyInstanceUnbuffered(instanceKey)
		if err != nil {
			return instance, log.Errore(err)
		}

		switch {
		case instance.ExecBinlogCoordinates.SmallerThan(masterCoordinates):
			time.Sleep(sqlThreadPollDuration)
		case instance.ExecBinlogCoordinates.Equals(masterCoordinates):
			upToDate = true
		case masterCoordinates.SmallerThan(&instance.ExecBinlogCoordinates):
			return instance, fmt.Errorf("Start SLAVE UNTIL is past coordinates: %+v", instanceKey)
		}
	}

	instance, err = StopSlave(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	return instance, err
}

// EnableSemiSync sets the rpl_semi_sync_(master|slave)_enabled variables
// on a given instance.
func EnableSemiSync(instanceKey *InstanceKey, master, slave bool) error {
	log.Infof("instance %+v rpl_semi_sync_master_enabled: %t, rpl_semi_sync_slave_enabled: %t", instanceKey, master, slave)
	_, err := ExecInstanceNoPrepare(instanceKey,
		`set global rpl_semi_sync_master_enabled = ?, global rpl_semi_sync_slave_enabled = ?`,
		master, slave)
	return err
}

// ChangeMasterCredentials issues a CHANGE MASTER TO... MASTER_USER=, MASTER_PASSWORD=...
func ChangeMasterCredentials(instanceKey *InstanceKey, masterUser string, masterPassword string) (*Instance, error) {
	instance, err := ReadTopologyInstanceUnbuffered(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}
	if masterUser == "" {
		return instance, log.Errorf("Empty user in ChangeMasterCredentials() for %+v", *instanceKey)
	}

	if instance.SlaveRunning() {
		return instance, fmt.Errorf("ChangeMasterTo: Cannot change master on: %+v because slave is running", *instanceKey)
	}
	log.Debugf("ChangeMasterTo: will attempt changing master credentials on %+v", *instanceKey)

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting CHANGE MASTER TO operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}
	_, err = ExecInstanceNoPrepare(instanceKey, fmt.Sprintf("change master to master_user='%s', master_password='%s'",
		masterUser, masterPassword))

	if err != nil {
		return instance, log.Errore(err)
	}

	log.Infof("ChangeMasterTo: Changed master credentials on %+v", *instanceKey)

	instance, err = ReadTopologyInstanceUnbuffered(instanceKey)
	return instance, err
}

// ChangeMasterTo changes the given instance's master according to given input.
func ChangeMasterTo(instanceKey *InstanceKey, masterKey *InstanceKey, masterBinlogCoordinates *BinlogCoordinates, skipUnresolve bool, gtidHint OperationGTIDHint) (*Instance, error) {
	instance, err := ReadTopologyInstanceUnbuffered(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if instance.SlaveRunning() {
		return instance, fmt.Errorf("ChangeMasterTo: Cannot change master on: %+v because slave is running", *instanceKey)
	}
	log.Debugf("ChangeMasterTo: will attempt changing master on %+v to %+v, %+v", *instanceKey, *masterKey, *masterBinlogCoordinates)
	changeToMasterKey := masterKey
	if !skipUnresolve {
		unresolvedMasterKey, nameUnresolved, err := UnresolveHostname(masterKey)
		if err != nil {
			log.Debugf("ChangeMasterTo: aborting operation on %+v due to resolving error on %+v: %+v", *instanceKey, *masterKey, err)
			return instance, err
		}
		if nameUnresolved {
			log.Debugf("ChangeMasterTo: Unresolved %+v into %+v", *masterKey, unresolvedMasterKey)
		}
		changeToMasterKey = &unresolvedMasterKey
	}

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting CHANGE MASTER TO operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}

	originalMasterKey := instance.MasterKey
	originalExecBinlogCoordinates := instance.ExecBinlogCoordinates

	changedViaGTID := false
	if instance.UsingMariaDBGTID && gtidHint != GTIDHintDeny {
		// MariaDB has a bug: a CHANGE MASTER TO statement does not work properly with prepared statement... :P
		// See https://mariadb.atlassian.net/browse/MDEV-7640
		// This is the reason for ExecInstanceNoPrepare
		// Keep on using GTID
		_, err = ExecInstanceNoPrepare(instanceKey, fmt.Sprintf("change master to master_host='%s', master_port=%d",
			changeToMasterKey.Hostname, changeToMasterKey.Port))
		changedViaGTID = true
	} else if instance.UsingMariaDBGTID && gtidHint == GTIDHintDeny {
		// Make sure to not use GTID
		_, err = ExecInstanceNoPrepare(instanceKey, fmt.Sprintf("change master to master_host='%s', master_port=%d, master_log_file='%s', master_log_pos=%d, master_use_gtid=no",
			changeToMasterKey.Hostname, changeToMasterKey.Port, masterBinlogCoordinates.LogFile, masterBinlogCoordinates.LogPos))
	} else if instance.IsMariaDB() && gtidHint == GTIDHintForce {
		// Is MariaDB; not using GTID, turn into GTID
		_, err = ExecInstanceNoPrepare(instanceKey, fmt.Sprintf("change master to master_host='%s', master_port=%d, master_use_gtid=slave_pos",
			changeToMasterKey.Hostname, changeToMasterKey.Port))
		changedViaGTID = true
	} else if instance.UsingOracleGTID && gtidHint != GTIDHintDeny {
		// Is Oracle; already uses GTID; keep using it.
		_, err = ExecInstanceNoPrepare(instanceKey, fmt.Sprintf("change master to master_host='%s', master_port=%d",
			changeToMasterKey.Hostname, changeToMasterKey.Port))
		changedViaGTID = true
	} else if instance.UsingOracleGTID && gtidHint == GTIDHintDeny {
		// Is Oracle; already uses GTID
		_, err = ExecInstanceNoPrepare(instanceKey, fmt.Sprintf("change master to master_host='%s', master_port=%d, master_log_file='%s', master_log_pos=%d, master_auto_position=0",
			changeToMasterKey.Hostname, changeToMasterKey.Port, masterBinlogCoordinates.LogFile, masterBinlogCoordinates.LogPos))
	} else if instance.SupportsOracleGTID && gtidHint == GTIDHintForce {
		// Is Oracle; not using GTID right now; turn into GTID
		_, err = ExecInstanceNoPrepare(instanceKey, fmt.Sprintf("change master to master_host='%s', master_port=%d, master_auto_position=1",
			changeToMasterKey.Hostname, changeToMasterKey.Port))
		changedViaGTID = true
	} else {
		// Normal binlog file:pos
		_, err = ExecInstanceNoPrepare(instanceKey, fmt.Sprintf("change master to master_host='%s', master_port=%d, master_log_file='%s', master_log_pos=%d",
			changeToMasterKey.Hostname, changeToMasterKey.Port, masterBinlogCoordinates.LogFile, masterBinlogCoordinates.LogPos))
	}
	if err != nil {
		return instance, log.Errore(err)
	}
	WriteMasterPositionEquivalence(&originalMasterKey, &originalExecBinlogCoordinates, changeToMasterKey, masterBinlogCoordinates)

	log.Infof("ChangeMasterTo: Changed master on %+v to: %+v, %+v. GTID: %+v", *instanceKey, masterKey, masterBinlogCoordinates, changedViaGTID)

	instance, err = ReadTopologyInstanceUnbuffered(instanceKey)
	return instance, err
}

// SkipToNextBinaryLog changes master position to beginning of next binlog
// USE WITH CARE!
// Use case is binlog servers where the master was gone & replaced by another.
func SkipToNextBinaryLog(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstanceUnbuffered(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	nextFileCoordinates, err := instance.ExecBinlogCoordinates.NextFileCoordinates()
	if err != nil {
		return instance, log.Errore(err)
	}
	nextFileCoordinates.LogPos = 4
	log.Debugf("Will skip replication on %+v to next binary log: %+v", instance.Key, nextFileCoordinates.LogFile)

	instance, err = ChangeMasterTo(&instance.Key, &instance.MasterKey, &nextFileCoordinates, false, GTIDHintNeutral)
	if err != nil {
		return instance, log.Errore(err)
	}
	AuditOperation("skip-binlog", instanceKey, fmt.Sprintf("Skipped replication to next binary log: %+v", nextFileCoordinates.LogFile))
	return StartSlave(instanceKey)
}

// ResetSlave resets a slave, breaking the replication
func ResetSlave(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstanceUnbuffered(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if instance.SlaveRunning() {
		return instance, fmt.Errorf("Cannot reset slave on: %+v because slave is running", instanceKey)
	}

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting reset-slave operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}

	// MySQL's RESET SLAVE is done correctly; however SHOW SLAVE STATUS still returns old hostnames etc
	// and only resets till after next restart. This leads to orchestrator still thinking the instance replicates
	// from old host. We therefore forcibly modify the hostname.
	// RESET SLAVE ALL command solves this, but only as of 5.6.3
	_, err = ExecInstanceNoPrepare(instanceKey, `change master to master_host='_'`)
	if err != nil {
		return instance, log.Errore(err)
	}
	_, err = ExecInstanceNoPrepare(instanceKey, `reset slave /*!50603 all */`)
	if err != nil {
		return instance, log.Errore(err)
	}
	log.Infof("Reset slave %+v", instanceKey)

	instance, err = ReadTopologyInstanceUnbuffered(instanceKey)
	return instance, err
}

// ResetMaster issues a RESET MASTER statement on given instance. Use with extreme care!
func ResetMaster(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstanceUnbuffered(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if instance.SlaveRunning() {
		return instance, fmt.Errorf("Cannot reset master on: %+v because slave is running", instanceKey)
	}

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting reset-master operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}

	_, err = ExecInstanceNoPrepare(instanceKey, `reset master`)
	if err != nil {
		return instance, log.Errore(err)
	}
	log.Infof("Reset master %+v", instanceKey)

	instance, err = ReadTopologyInstanceUnbuffered(instanceKey)
	return instance, err
}

// skipQueryClassic skips a query in normal binlog file:pos replication
func setGTIDPurged(instance *Instance, gtidPurged string) error {
	if *config.RuntimeCLIFlags.Noop {
		return fmt.Errorf("noop: aborting set-gtid-purged operation on %+v; signalling error but nothing went wrong.", instance.Key)
	}

	_, err := ExecInstance(&instance.Key, fmt.Sprintf(`set global gtid_purged := '%s'`, gtidPurged))
	return err
}

// skipQueryClassic skips a query in normal binlog file:pos replication
func skipQueryClassic(instance *Instance) error {
	_, err := ExecInstance(&instance.Key, `set global sql_slave_skip_counter := 1`)
	return err
}

// skipQueryOracleGtid skips a single query in an Oracle GTID replicating slave, by injecting an empty transaction
func skipQueryOracleGtid(instance *Instance) error {
	nextGtid, err := instance.NextGTID()
	if err != nil {
		return err
	}
	if nextGtid == "" {
		return fmt.Errorf("Empty NextGTID() in skipQueryGtid() for %+v", instance.Key)
	}
	if _, err := ExecInstanceNoPrepare(&instance.Key, fmt.Sprintf(`SET GTID_NEXT='%s'`, nextGtid)); err != nil {
		return err
	}
	if err := EmptyCommitInstance(&instance.Key); err != nil {
		return err
	}
	if _, err := ExecInstanceNoPrepare(&instance.Key, `SET GTID_NEXT='AUTOMATIC'`); err != nil {
		return err
	}
	return nil
}

// SkipQuery skip a single query in a failed replication instance
func SkipQuery(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstanceUnbuffered(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if !instance.IsSlave() {
		return instance, fmt.Errorf("instance is not a slave: %+v", instanceKey)
	}
	if instance.Slave_SQL_Running {
		return instance, fmt.Errorf("Slave SQL thread is running on %+v", instanceKey)
	}
	if instance.LastSQLError == "" {
		return instance, fmt.Errorf("No SQL error on %+v", instanceKey)
	}

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting skip-query operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}

	log.Debugf("Skipping one query on %+v", instanceKey)
	if instance.UsingOracleGTID {
		err = skipQueryOracleGtid(instance)
	} else if instance.UsingMariaDBGTID {
		return instance, log.Errorf("%+v is replicating with MariaDB GTID. To skip a query first disable GTID, then skip, then enable GTID again", *instanceKey)
	} else {
		err = skipQueryClassic(instance)
	}
	if err != nil {
		return instance, log.Errore(err)
	}
	AuditOperation("skip-query", instanceKey, "Skipped one query")
	return StartSlave(instanceKey)
}

// DetachSlave detaches a slave from replication; forcibly corrupting the binlog coordinates (though in such way
// that is reversible)
func DetachSlave(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstanceUnbuffered(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if instance.SlaveRunning() {
		return instance, fmt.Errorf("Cannot detach slave on: %+v because slave is running", instanceKey)
	}

	isDetached, _, _ := instance.ExecBinlogCoordinates.DetachedCoordinates()

	if isDetached {
		return instance, fmt.Errorf("Cannot (need not) detach slave on: %+v because slave is already detached", instanceKey)
	}

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting detach-slave operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}

	detachedCoordinates := BinlogCoordinates{LogFile: fmt.Sprintf("//%s:%d", instance.ExecBinlogCoordinates.LogFile, instance.ExecBinlogCoordinates.LogPos), LogPos: instance.ExecBinlogCoordinates.LogPos}
	// Encode the current coordinates within the log file name, in such way that replication is broken, but info can still be resurrected
	_, err = ExecInstanceNoPrepare(instanceKey, fmt.Sprintf(`change master to master_log_file='%s', master_log_pos=%d`, detachedCoordinates.LogFile, detachedCoordinates.LogPos))
	if err != nil {
		return instance, log.Errore(err)
	}

	log.Infof("Detach slave %+v", instanceKey)

	instance, err = ReadTopologyInstanceUnbuffered(instanceKey)
	return instance, err
}

// ReattachSlave restores a detached slave back into replication
func ReattachSlave(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstanceUnbuffered(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if instance.SlaveRunning() {
		return instance, fmt.Errorf("Cannot (need not) reattach slave on: %+v because slave is running", instanceKey)
	}

	isDetached, detachedLogFile, detachedLogPos := instance.ExecBinlogCoordinates.DetachedCoordinates()

	if !isDetached {
		return instance, fmt.Errorf("Cannot reattach slave on: %+v because slave is not detached", instanceKey)
	}

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting reattach-slave operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}

	_, err = ExecInstanceNoPrepare(instanceKey, fmt.Sprintf(`change master to master_log_file='%s', master_log_pos=%s`, detachedLogFile, detachedLogPos))
	if err != nil {
		return instance, log.Errore(err)
	}

	log.Infof("Reattach slave %+v", instanceKey)

	instance, err = ReadTopologyInstanceUnbuffered(instanceKey)
	return instance, err
}

// MasterPosWait issues a MASTER_POS_WAIT() an given instance according to given coordinates.
func MasterPosWait(instanceKey *InstanceKey, binlogCoordinates *BinlogCoordinates) (*Instance, error) {
	instance, err := ReadTopologyInstanceUnbuffered(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	_, err = ExecInstance(instanceKey, `select master_pos_wait(?, ?)`, binlogCoordinates.LogFile, binlogCoordinates.LogPos)
	if err != nil {
		return instance, log.Errore(err)
	}
	log.Infof("Instance %+v has reached coordinates: %+v", instanceKey, binlogCoordinates)

	instance, err = ReadTopologyInstanceUnbuffered(instanceKey)
	return instance, err
}

// Attempt to read and return replication credentials from the mysql.slave_master_info system table
func ReadReplicationCredentials(instanceKey *InstanceKey) (replicationUser string, replicationPassword string, err error) {
	query := `
		select
			ifnull(max(User_name), '') as user,
			ifnull(max(User_password), '') as password
		from
			mysql.slave_master_info
	`
	err = ScanInstanceRow(instanceKey, query, &replicationUser, &replicationPassword)
	if err != nil {
		return replicationUser, replicationPassword, err
	}
	if replicationUser == "" {
		err = fmt.Errorf("Cannot find credentials in mysql.slave_master_info")
	}
	return replicationUser, replicationPassword, err
}

// SetReadOnly sets or clears the instance's global read_only variable
func SetReadOnly(instanceKey *InstanceKey, readOnly bool) (*Instance, error) {
	instance, err := ReadTopologyInstanceUnbuffered(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting set-read-only operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}

	// If async fallback is disallowed, we're responsible for flipping the master
	// semi-sync switch ON before accepting writes. The setting is off by default.
	if instance.SemiSyncEnforced && !readOnly {
		// Send ACK only from promotable instances.
		sendACK := instance.PromotionRule != MustNotPromoteRule
		if err := EnableSemiSync(instanceKey, true, sendACK); err != nil {
			return instance, log.Errore(err)
		}
	}

	_, err = ExecInstance(instanceKey, fmt.Sprintf("set global read_only = %t", readOnly))
	if err != nil {
		return instance, log.Errore(err)
	}
	instance, err = ReadTopologyInstanceUnbuffered(instanceKey)

	// If we just went read-only, it's safe to flip the master semi-sync switch
	// OFF, which is the default value so that slaves can make progress.
	if instance.SemiSyncEnforced && readOnly {
		// Send ACK only from promotable instances.
		sendACK := instance.PromotionRule != MustNotPromoteRule
		if err := EnableSemiSync(instanceKey, false, sendACK); err != nil {
			return instance, log.Errore(err)
		}
	}

	log.Infof("instance %+v read_only: %t", instanceKey, readOnly)
	AuditOperation("read-only", instanceKey, fmt.Sprintf("set as %t", readOnly))

	return instance, err
}

// KillQuery stops replication on a given instance
func KillQuery(instanceKey *InstanceKey, process int64) (*Instance, error) {
	instance, err := ReadTopologyInstanceUnbuffered(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting kill-query operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}

	_, err = ExecInstance(instanceKey, fmt.Sprintf(`kill query %d`, process))
	if err != nil {
		return instance, log.Errore(err)
	}

	instance, err = ReadTopologyInstanceUnbuffered(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	log.Infof("Killed query on %+v", *instanceKey)
	AuditOperation("kill-query", instanceKey, fmt.Sprintf("Killed query %d", process))
	return instance, err
}
