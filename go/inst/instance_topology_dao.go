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
	"strings"
	"time"

	"github.com/github/orchestrator/go/config"
	"github.com/github/orchestrator/go/db"
	"github.com/github/orchestrator/go/util"
	"github.com/openark/golib/log"
	"github.com/openark/golib/sqlutils"
	"github.com/patrickmn/go-cache"
)

// Max concurrency for bulk topology operations
const topologyConcurrency = 128

var topologyConcurrencyChan = make(chan bool, topologyConcurrency)
var supportedAutoPseudoGTIDWriters *cache.Cache = cache.New(config.CheckAutoPseudoGTIDGrantsIntervalSeconds*time.Second, time.Second)

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
	return sqlutils.ExecNoPrepare(db, query, args...)
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
	_, err := ReadTopologyInstance(instanceKey)
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
			// Signal completed replica
			defer func() { barrier <- instance.Key }()
			// Wait your turn to read a replica
			ExecuteOnTopology(func() {
				log.Debugf("... reading instance: %+v", instance.Key)
				ReadTopologyInstance(&instance.Key)
			})
		}()
	}
	for range instances {
		<-barrier
	}
}

// RefreshInstanceSlaveHosts is a workaround for a bug in MySQL where
// SHOW SLAVE HOSTS continues to present old, long disconnected replicas.
// It turns out issuing a couple FLUSH commands mitigates the problem.
func RefreshInstanceSlaveHosts(instanceKey *InstanceKey) (*Instance, error) {
	_, _ = ExecInstance(instanceKey, `flush error logs`)
	_, _ = ExecInstance(instanceKey, `flush error logs`)

	instance, err := ReadTopologyInstance(instanceKey)
	return instance, err
}

// GetSlaveRestartPreserveStatements returns a sequence of statements that make sure a replica is stopped
// and then returned to the same state. For example, if the replica was fully running, this will issue
// a STOP on both io_thread and sql_thread, followed by START on both. If one of them is not running
// at the time this function is called, said thread will be neither stopped nor started.
// The caller may provide an injected statememt, to be executed while the replica is stopped.
// This is useful for CHANGE MASTER TO commands, that unfortunately must take place while the replica
// is completely stopped.
func GetSlaveRestartPreserveStatements(instanceKey *InstanceKey, injectedStatement string) (statements []string, err error) {
	instance, err := ReadTopologyInstance(instanceKey)
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

	return ReadTopologyInstance(instanceKey)
}

// FlushBinaryLogsTo attempts to 'FLUSH BINARY LOGS' until given binary log is reached
func FlushBinaryLogsTo(instanceKey *InstanceKey, logFile string) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
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

	_, err := ExecInstance(instanceKey, "purge binary logs to ?", logFile)
	if err != nil {
		return nil, log.Errore(err)
	}

	log.Infof("purge-binary-logs to=%+v on %+v", logFile, *instanceKey)
	AuditOperation("purge-binary-logs", instanceKey, "success")

	return ReadTopologyInstance(instanceKey)
}

// FlushBinaryLogsTo attempts to 'PURGE BINARY LOGS' until given binary log is reached
func PurgeBinaryLogsToCurrent(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}
	return PurgeBinaryLogsTo(instanceKey, instance.SelfBinlogCoordinates.LogFile)
}

func SetSemiSyncMaster(instanceKey *InstanceKey, enableMaster bool) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, err
	}
	if _, err := ExecInstance(instanceKey, "set @@global.rpl_semi_sync_master_enabled=?", enableMaster); err != nil {
		return instance, log.Errore(err)
	}
	return ReadTopologyInstance(instanceKey)
}

func SetSemiSyncReplica(instanceKey *InstanceKey, enableReplica bool) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, err
	}
	if instance.SemiSyncReplicaEnabled == enableReplica {
		log.Debugf("SetSemiSyncReplica: %+v slready in desired state", *instanceKey)
		return instance, nil
	}
	if _, err := ExecInstance(instanceKey, "set @@global.rpl_semi_sync_slave_enabled=?", enableReplica); err != nil {
		return instance, log.Errore(err)
	}
	if instance.Slave_IO_Running {
		// Need to apply change by stopping starting IO thread
		ExecInstance(instanceKey, "stop slave io_thread")
		if _, err := ExecInstance(instanceKey, "start slave io_thread"); err != nil {
			return instance, log.Errore(err)
		}
	}
	return ReadTopologyInstance(instanceKey)

}

// StopSlaveNicely stops a replica such that SQL_thread and IO_thread are aligned (i.e.
// SQL_thread consumes all relay log entries)
// It will actually START the sql_thread even if the replica is completely stopped.
func StopSlaveNicely(instanceKey *InstanceKey, timeout time.Duration) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if !instance.IsReplica() {
		return instance, fmt.Errorf("instance is not a replica: %+v", instanceKey)
	}

	// stop io_thread, start sql_thread but catch any errors
	for _, cmd := range []string{`stop slave io_thread`, `start slave sql_thread`} {
		if _, err = ExecInstance(instanceKey, cmd); err != nil {
			return nil, log.Errorf("%+v: StopSlaveNicely: %q failed: %+v", *instanceKey, cmd, err)
		}
	}

	if instance.SQLDelay == 0 {
		// Otherwise we don't bother.
		startTime := time.Now()
		for upToDate := false; !upToDate; {
			timeSinceStartTime := time.Since(startTime)
			if timeout > 0 && timeSinceStartTime >= timeout {
				// timeout
				return nil, log.Errorf("%+v: StopSlaveNicely timeout after %+v", *instanceKey, timeSinceStartTime)
			}
			instance, err = ReadTopologyInstance(instanceKey)
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
	_, err = ExecInstance(instanceKey, `stop slave`)
	if err != nil {
		// Patch; current MaxScale behavior for STOP SLAVE is to throw an error if replica already stopped.
		if instance.isMaxScale() && err.Error() == "Error 1199: Slave connection is not running" {
			err = nil
		}
	}
	if err != nil {
		return instance, log.Errore(err)
	}

	instance, err = ReadTopologyInstance(instanceKey)
	log.Infof("Stopped slave nicely on %+v, Self:%+v, Exec:%+v", *instanceKey, instance.SelfBinlogCoordinates, instance.ExecBinlogCoordinates)
	return instance, err
}

// StopSlaves will stop replication concurrently on given set of replicas.
// It will potentially do nothing, or attempt to stop _nicely_ or just stop normally, all according to stopReplicationMethod
func StopSlaves(replicas [](*Instance), stopReplicationMethod StopReplicationMethod, timeout time.Duration) [](*Instance) {
	if stopReplicationMethod == NoStopReplication {
		return replicas
	}
	refreshedReplicas := [](*Instance){}

	log.Debugf("Stopping %d replicas via %s", len(replicas), string(stopReplicationMethod))
	// use concurrency but wait for all to complete
	barrier := make(chan *Instance)
	for _, replica := range replicas {
		replica := replica
		go func() {
			updatedReplica := &replica
			// Signal completed replica
			defer func() { barrier <- *updatedReplica }()
			// Wait your turn to read a replica
			ExecuteOnTopology(func() {
				if stopReplicationMethod == StopReplicationNicely {
					StopSlaveNicely(&replica.Key, timeout)
				}
				replica, _ = StopSlave(&replica.Key)
				updatedReplica = &replica
			})
		}()
	}
	for range replicas {
		refreshedReplicas = append(refreshedReplicas, <-barrier)
	}
	return refreshedReplicas
}

// StopSlavesNicely will attemt to stop all given replicas nicely, up to timeout
func StopSlavesNicely(replicas [](*Instance), timeout time.Duration) [](*Instance) {
	return StopSlaves(replicas, StopReplicationNicely, timeout)
}

// StopSlave stops replication on a given instance
func StopSlave(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if !instance.IsReplica() {
		return instance, fmt.Errorf("instance is not a replica: %+v", instanceKey)
	}
	_, err = ExecInstance(instanceKey, `stop slave`)
	if err != nil {
		// Patch; current MaxScale behavior for STOP SLAVE is to throw an error if replica already stopped.
		if instance.isMaxScale() && err.Error() == "Error 1199: Slave connection is not running" {
			err = nil
		}
	}
	if err != nil {

		return instance, log.Errore(err)
	}
	instance, err = ReadTopologyInstance(instanceKey)

	log.Infof("Stopped slave on %+v, Self:%+v, Exec:%+v", *instanceKey, instance.SelfBinlogCoordinates, instance.ExecBinlogCoordinates)
	return instance, err
}

// sleep immediately after START SLAVE for a capped duration, until both replication threads are running.
// This is to give slack to IO thread to connect and begin streaming, and to SQL thread to start applying.
// Sleep is incremental with ongoing attempt to see whether replication is already up
func startSlavePostSleep(instanceKey *InstanceKey) error {
	waitDuration := time.Second
	waitInterval := 10 * time.Millisecond
	startTime := time.Now()

	for time.Since(startTime) < waitDuration {
		repliationRunning, err := AreReplicationThreadsRunning(instanceKey)
		if err != nil {
			return err
		}
		if repliationRunning {
			return nil
		}
		time.Sleep(waitInterval)
		if waitInterval < waitDuration {
			waitInterval = 2 * waitInterval
		}
	}
	return nil
}

// StartSlave starts replication on a given instance.
func StartSlave(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if !instance.IsReplica() {
		return instance, fmt.Errorf("instance is not a replica: %+v", instanceKey)
	}

	// If async fallback is disallowed, we'd better make sure to enable replicas to
	// send ACKs before START SLAVE. Replica ACKing is off at mysqld startup because
	// some replicas (those that must never be promoted) should never ACK.
	// Note: We assume that replicas use 'skip-slave-start' so they won't
	//       START SLAVE on their own upon restart.
	if instance.SemiSyncEnforced {
		// Send ACK only from promotable instances.
		sendACK := instance.PromotionRule != MustNotPromoteRule
		// Always disable master setting, in case we're converting a former master.
		if err := EnableSemiSync(instanceKey, false, sendACK); err != nil {
			return instance, log.Errore(err)
		}
	}

	_, err = ExecInstance(instanceKey, `start slave`)
	if err != nil {
		return instance, log.Errore(err)
	}
	log.Infof("Started slave on %+v", instanceKey)

	startSlavePostSleep(instanceKey)

	instance, err = ReadTopologyInstance(instanceKey)
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
func StartSlaves(replicas [](*Instance)) {
	// use concurrency but wait for all to complete
	log.Debugf("Starting %d replicas", len(replicas))
	barrier := make(chan InstanceKey)
	for _, instance := range replicas {
		instance := instance
		go func() {
			// Signal compelted replica
			defer func() { barrier <- instance.Key }()
			// Wait your turn to read a replica
			ExecuteOnTopology(func() { StartSlave(&instance.Key) })
		}()
	}
	for range replicas {
		<-barrier
	}
}

// StartSlaveUntilMasterCoordinates issuesa START SLAVE UNTIL... statement on given instance
func StartSlaveUntilMasterCoordinates(instanceKey *InstanceKey, masterCoordinates *BinlogCoordinates) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if !instance.IsReplica() {
		return instance, fmt.Errorf("instance is not a replica: %+v", instanceKey)
	}
	if instance.ReplicaRunning() {
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
	// This is the reason for ExecInstance
	_, err = ExecInstance(instanceKey, "start slave until master_log_file=?, master_log_pos=?",
		masterCoordinates.LogFile, masterCoordinates.LogPos)
	if err != nil {
		return instance, log.Errore(err)
	}

	for upToDate := false; !upToDate; {
		instance, err = ReadTopologyInstance(instanceKey)
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
	_, err := ExecInstance(instanceKey,
		`set global rpl_semi_sync_master_enabled = ?, global rpl_semi_sync_slave_enabled = ?`,
		master, slave)
	return err
}

// ChangeMasterCredentials issues a CHANGE MASTER TO... MASTER_USER=, MASTER_PASSWORD=...
func ChangeMasterCredentials(instanceKey *InstanceKey, masterUser string, masterPassword string) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}
	if masterUser == "" {
		return instance, log.Errorf("Empty user in ChangeMasterCredentials() for %+v", *instanceKey)
	}

	if instance.ReplicaRunning() {
		return instance, fmt.Errorf("ChangeMasterTo: Cannot change master on: %+v because slave is running", *instanceKey)
	}
	log.Debugf("ChangeMasterTo: will attempt changing master credentials on %+v", *instanceKey)

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting CHANGE MASTER TO operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}
	_, err = ExecInstance(instanceKey, "change master to master_user=?, master_password=?",
		masterUser, masterPassword)

	if err != nil {
		return instance, log.Errore(err)
	}

	log.Infof("ChangeMasterTo: Changed master credentials on %+v", *instanceKey)

	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}

// EnableMasterSSL issues CHANGE MASTER TO MASTER_SSL=1
func EnableMasterSSL(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if instance.ReplicaRunning() {
		return instance, fmt.Errorf("EnableMasterSSL: Cannot enable SSL replication on %+v because slave is running", *instanceKey)
	}
	log.Debugf("EnableMasterSSL: Will attempt enabling SSL replication on %+v", *instanceKey)

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting CHANGE MASTER TO MASTER_SSL=1 operation on %+v; signaling error but nothing went wrong.", *instanceKey)
	}
	_, err = ExecInstance(instanceKey, "change master to master_ssl=1")

	if err != nil {
		return instance, log.Errore(err)
	}

	log.Infof("EnableMasterSSL: Enabled SSL replication on %+v", *instanceKey)

	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}

// ChangeMasterTo changes the given instance's master according to given input.
func ChangeMasterTo(instanceKey *InstanceKey, masterKey *InstanceKey, masterBinlogCoordinates *BinlogCoordinates, skipUnresolve bool, gtidHint OperationGTIDHint) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if instance.ReplicaRunning() {
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
		// Keep on using GTID
		_, err = ExecInstance(instanceKey, "change master to master_host=?, master_port=?",
			changeToMasterKey.Hostname, changeToMasterKey.Port)
		changedViaGTID = true
	} else if instance.UsingMariaDBGTID && gtidHint == GTIDHintDeny {
		// Make sure to not use GTID
		_, err = ExecInstance(instanceKey, "change master to master_host=?, master_port=?, master_log_file=?, master_log_pos=?, master_use_gtid=no",
			changeToMasterKey.Hostname, changeToMasterKey.Port, masterBinlogCoordinates.LogFile, masterBinlogCoordinates.LogPos)
	} else if instance.IsMariaDB() && gtidHint == GTIDHintForce {
		// Is MariaDB; not using GTID, turn into GTID
		_, err = ExecInstance(instanceKey, "change master to master_host=?, master_port=?, master_use_gtid=slave_pos",
			changeToMasterKey.Hostname, changeToMasterKey.Port)
		changedViaGTID = true
	} else if instance.UsingOracleGTID && gtidHint != GTIDHintDeny {
		// Is Oracle; already uses GTID; keep using it.
		_, err = ExecInstance(instanceKey, "change master to master_host=?, master_port=?",
			changeToMasterKey.Hostname, changeToMasterKey.Port)
		changedViaGTID = true
	} else if instance.UsingOracleGTID && gtidHint == GTIDHintDeny {
		// Is Oracle; already uses GTID
		_, err = ExecInstance(instanceKey, "change master to master_host=?, master_port=?, master_log_file=?, master_log_pos=?, master_auto_position=0",
			changeToMasterKey.Hostname, changeToMasterKey.Port, masterBinlogCoordinates.LogFile, masterBinlogCoordinates.LogPos)
	} else if instance.SupportsOracleGTID && gtidHint == GTIDHintForce {
		// Is Oracle; not using GTID right now; turn into GTID
		_, err = ExecInstance(instanceKey, "change master to master_host=?, master_port=?, master_auto_position=1",
			changeToMasterKey.Hostname, changeToMasterKey.Port)
		changedViaGTID = true
	} else {
		// Normal binlog file:pos
		_, err = ExecInstance(instanceKey, "change master to master_host=?, master_port=?, master_log_file=?, master_log_pos=?",
			changeToMasterKey.Hostname, changeToMasterKey.Port, masterBinlogCoordinates.LogFile, masterBinlogCoordinates.LogPos)
	}
	if err != nil {
		return instance, log.Errore(err)
	}
	WriteMasterPositionEquivalence(&originalMasterKey, &originalExecBinlogCoordinates, changeToMasterKey, masterBinlogCoordinates)
	ResetInstanceRelaylogCoordinatesHistory(instanceKey)

	log.Infof("ChangeMasterTo: Changed master on %+v to: %+v, %+v. GTID: %+v", *instanceKey, masterKey, masterBinlogCoordinates, changedViaGTID)

	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}

// SkipToNextBinaryLog changes master position to beginning of next binlog
// USE WITH CARE!
// Use case is binlog servers where the master was gone & replaced by another.
func SkipToNextBinaryLog(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
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

// ResetSlave resets a replica, breaking the replication
func ResetSlave(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if instance.ReplicaRunning() {
		return instance, fmt.Errorf("Cannot reset slave on: %+v because slave is running", instanceKey)
	}

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting reset-slave operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}

	// MySQL's RESET SLAVE is done correctly; however SHOW SLAVE STATUS still returns old hostnames etc
	// and only resets till after next restart. This leads to orchestrator still thinking the instance replicates
	// from old host. We therefore forcibly modify the hostname.
	// RESET SLAVE ALL command solves this, but only as of 5.6.3
	_, err = ExecInstance(instanceKey, `change master to master_host='_'`)
	if err != nil {
		return instance, log.Errore(err)
	}
	_, err = ExecInstance(instanceKey, `reset slave /*!50603 all */`)
	if err != nil {
		return instance, log.Errore(err)
	}
	log.Infof("Reset slave %+v", instanceKey)

	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}

// ResetMaster issues a RESET MASTER statement on given instance. Use with extreme care!
func ResetMaster(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if instance.ReplicaRunning() {
		return instance, fmt.Errorf("Cannot reset master on: %+v because slave is running", instanceKey)
	}

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting reset-master operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}

	_, err = ExecInstance(instanceKey, `reset master`)
	if err != nil {
		return instance, log.Errore(err)
	}
	log.Infof("Reset master %+v", instanceKey)

	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}

// skipQueryClassic skips a query in normal binlog file:pos replication
func setGTIDPurged(instance *Instance, gtidPurged string) error {
	if *config.RuntimeCLIFlags.Noop {
		return fmt.Errorf("noop: aborting set-gtid-purged operation on %+v; signalling error but nothing went wrong.", instance.Key)
	}

	_, err := ExecInstance(&instance.Key, `set global gtid_purged := ?`, gtidPurged)
	return err
}

// skipQueryClassic skips a query in normal binlog file:pos replication
func skipQueryClassic(instance *Instance) error {
	_, err := ExecInstance(&instance.Key, `set global sql_slave_skip_counter := 1`)
	return err
}

// skipQueryOracleGtid skips a single query in an Oracle GTID replicating replica, by injecting an empty transaction
func skipQueryOracleGtid(instance *Instance) error {
	nextGtid, err := instance.NextGTID()
	if err != nil {
		return err
	}
	if nextGtid == "" {
		return fmt.Errorf("Empty NextGTID() in skipQueryGtid() for %+v", instance.Key)
	}
	if _, err := ExecInstance(&instance.Key, `SET GTID_NEXT=?`, nextGtid); err != nil {
		return err
	}
	if err := EmptyCommitInstance(&instance.Key); err != nil {
		return err
	}
	if _, err := ExecInstance(&instance.Key, `SET GTID_NEXT='AUTOMATIC'`); err != nil {
		return err
	}
	return nil
}

// SkipQuery skip a single query in a failed replication instance
func SkipQuery(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if !instance.IsReplica() {
		return instance, fmt.Errorf("instance is not a replica: %+v", instanceKey)
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

// DetachReplica detaches a replica from replication; forcibly corrupting the binlog coordinates (though in such way
// that is reversible)
func DetachReplica(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if instance.ReplicaRunning() {
		return instance, fmt.Errorf("Cannot detach slave on: %+v because slave is running", instanceKey)
	}

	isDetached, _ := instance.ExecBinlogCoordinates.ExtractDetachedCoordinates()

	if isDetached {
		return instance, fmt.Errorf("Cannot (need not) detach slave on: %+v because slave is already detached", instanceKey)
	}

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting detach-slave operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}

	detachedCoordinates := instance.ExecBinlogCoordinates.Detach()
	// Encode the current coordinates within the log file name, in such way that replication is broken, but info can still be resurrected
	_, err = ExecInstance(instanceKey, `change master to master_log_file=?, master_log_pos=?`, detachedCoordinates.LogFile, detachedCoordinates.LogPos)
	if err != nil {
		return instance, log.Errore(err)
	}

	log.Infof("Detach slave %+v", instanceKey)

	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}

// ReattachReplica restores a detached replica back into replication
func ReattachReplica(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if instance.ReplicaRunning() {
		return instance, fmt.Errorf("Cannot (need not) reattach slave on: %+v because slave is running", instanceKey)
	}

	isDetached, detachedCoordinates := instance.ExecBinlogCoordinates.ExtractDetachedCoordinates()

	if !isDetached {
		return instance, fmt.Errorf("Cannot reattach slave on: %+v because slave is not detached", instanceKey)
	}

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting reattach-slave operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}

	_, err = ExecInstance(instanceKey, `change master to master_log_file=?, master_log_pos=?`, detachedCoordinates.LogFile, detachedCoordinates.LogPos)
	if err != nil {
		return instance, log.Errore(err)
	}

	log.Infof("Reattach slave %+v", instanceKey)

	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}

// MasterPosWait issues a MASTER_POS_WAIT() an given instance according to given coordinates.
func MasterPosWait(instanceKey *InstanceKey, binlogCoordinates *BinlogCoordinates) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	_, err = ExecInstance(instanceKey, `select master_pos_wait(?, ?)`, binlogCoordinates.LogFile, binlogCoordinates.LogPos)
	if err != nil {
		return instance, log.Errore(err)
	}
	log.Infof("Instance %+v has reached coordinates: %+v", instanceKey, binlogCoordinates)

	instance, err = ReadTopologyInstance(instanceKey)
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
	instance, err := ReadTopologyInstance(instanceKey)
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

	if _, err := ExecInstance(instanceKey, "set global read_only = ?", readOnly); err != nil {
		return instance, log.Errore(err)
	}
	if config.Config.UseSuperReadOnly {
		if _, err := ExecInstance(instanceKey, "set global super_read_only = ?", readOnly); err != nil {
			// We don't bail out here. super_read_only is only available on
			// MySQL 5.7.8 and Percona Server 5.6.21-70
			// At this time orchestrator does not verify whether a server supports super_read_only or not.
			// It makes a best effort to set it.
			log.Errore(err)
		}
	}
	instance, err = ReadTopologyInstance(instanceKey)

	// If we just went read-only, it's safe to flip the master semi-sync switch
	// OFF, which is the default value so that replicas can make progress.
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
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting kill-query operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}

	_, err = ExecInstance(instanceKey, `kill query ?`, process)
	if err != nil {
		return instance, log.Errore(err)
	}

	instance, err = ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	log.Infof("Killed query on %+v", *instanceKey)
	AuditOperation("kill-query", instanceKey, fmt.Sprintf("Killed query %d", process))
	return instance, err
}

// injectPseudoGTID injects a Pseudo-GTID statement on a writable instance
func injectPseudoGTID(instance *Instance) (hint string, err error) {
	if *config.RuntimeCLIFlags.Noop {
		return hint, fmt.Errorf("noop: aborting inject-pseudo-gtid operation on %+v; signalling error but nothing went wrong.", instance.Key)
	}

	now := time.Now()
	randomHash := util.RandomHash()[0:16]
	hint = fmt.Sprintf("%.8x:%.8x:%s", now.Unix(), instance.ServerID, randomHash)
	query := fmt.Sprintf("drop view if exists `%s`.`_asc:%s`", config.PseudoGTIDSchema, hint)
	_, err = ExecInstance(&instance.Key, query)
	return hint, log.Errore(err)
}

// canInjectPseudoGTID checks orchestrator's grants to determine whether is has the
// privilege of auto-injecting pseudo-GTID
func canInjectPseudoGTID(instanceKey *InstanceKey) (canInject bool, err error) {
	if canInject, found := supportedAutoPseudoGTIDWriters.Get(instanceKey.StringCode()); found {
		return canInject.(bool), nil
	}
	db, err := db.OpenTopology(instanceKey.Hostname, instanceKey.Port)
	if err != nil {
		return canInject, err
	}

	foundAll := false
	foundDropOnAll := false
	foundAllOnSchema := false
	foundDropOnSchema := false

	err = sqlutils.QueryRowsMap(db, `show grants for current_user()`, func(m sqlutils.RowMap) error {
		for _, grantData := range m {
			grant := grantData.String
			if strings.Contains(grant, `GRANT ALL PRIVILEGES ON *.*`) {
				foundAll = true
			}
			if strings.Contains(grant, `DROP`) && strings.Contains(grant, ` ON *.*`) {
				foundDropOnAll = true
			}
			if strings.Contains(grant, fmt.Sprintf("GRANT ALL PRIVILEGES ON `%s`.*", config.PseudoGTIDSchema)) {
				foundAllOnSchema = true
			}
			if strings.Contains(grant, fmt.Sprintf(`GRANT ALL PRIVILEGES ON "%s".*`, config.PseudoGTIDSchema)) {
				foundAllOnSchema = true
			}
			if strings.Contains(grant, `DROP`) && strings.Contains(grant, fmt.Sprintf(" ON `%s`.*", config.PseudoGTIDSchema)) {
				foundDropOnSchema = true
			}
			if strings.Contains(grant, `DROP`) && strings.Contains(grant, fmt.Sprintf(` ON "%s".*`, config.PseudoGTIDSchema)) {
				foundDropOnSchema = true
			}
		}
		return nil
	})
	if err != nil {
		return canInject, err
	}

	canInject = foundAll || foundDropOnAll || foundAllOnSchema || foundDropOnSchema
	supportedAutoPseudoGTIDWriters.Set(instanceKey.StringCode(), canInject, cache.DefaultExpiration)
	return canInject, nil
}

// CheckAndInjectPseudoGTIDOnWriter checks whether pseudo-GTID can and
// should be injected on given instance, and if so, attempts to inject.
func CheckAndInjectPseudoGTIDOnWriter(instance *Instance) (injected bool, err error) {
	if !instance.IsWritableMaster() {
		return injected, nil
	}
	if !instance.IsLastCheckValid {
		return injected, nil
	}
	canInject, err := canInjectPseudoGTID(&instance.Key)
	if err != nil {
		return injected, log.Errore(err)
	}
	if !canInject {
		return injected, nil
	}
	if _, err := injectPseudoGTID(instance); err != nil {
		return injected, log.Errore(err)
	}
	injected = true
	if err := RegisterInjectedPseudoGTID(instance.ClusterName); err != nil {
		return injected, log.Errore(err)
	}
	return injected, nil
}
