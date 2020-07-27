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
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/openark/golib/log"
	"github.com/openark/golib/sqlutils"
	"github.com/openark/orchestrator/go/config"
	"github.com/openark/orchestrator/go/db"
	"github.com/openark/orchestrator/go/util"
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

const (
	Error1201CouldnotInitializeMainInfoStructure = "Error 1201:"
)

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

// GetReplicationRestartPreserveStatements returns a sequence of statements that make sure a replica is stopped
// and then returned to the same state. For example, if the replica was fully running, this will issue
// a STOP on both io_thread and sql_thread, followed by START on both. If one of them is not running
// at the time this function is called, said thread will be neither stopped nor started.
// The caller may provide an injected statememt, to be executed while the replica is stopped.
// This is useful for CHANGE MASTER TO commands, that unfortunately must take place while the replica
// is completely stopped.
func GetReplicationRestartPreserveStatements(instanceKey *InstanceKey, injectedStatement string) (statements []string, err error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return statements, err
	}
	if instance.ReplicationIOThreadRuning {
		statements = append(statements, SemicolonTerminated(`stop subordinate io_thread`))
	}
	if instance.ReplicationSQLThreadRuning {
		statements = append(statements, SemicolonTerminated(`stop subordinate sql_thread`))
	}
	if injectedStatement != "" {
		statements = append(statements, SemicolonTerminated(injectedStatement))
	}
	if instance.ReplicationSQLThreadRuning {
		statements = append(statements, SemicolonTerminated(`start subordinate sql_thread`))
	}
	if instance.ReplicationIOThreadRuning {
		statements = append(statements, SemicolonTerminated(`start subordinate io_thread`))
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

// purgeBinaryLogsTo attempts to 'PURGE BINARY LOGS' until given binary log is reached
func purgeBinaryLogsTo(instanceKey *InstanceKey, logFile string) (*Instance, error) {
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

func SetSemiSyncMain(instanceKey *InstanceKey, enableMain bool) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, err
	}
	if _, err := ExecInstance(instanceKey, "set @@global.rpl_semi_sync_main_enabled=?", enableMain); err != nil {
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
		return instance, nil
	}
	if _, err := ExecInstance(instanceKey, "set @@global.rpl_semi_sync_subordinate_enabled=?", enableReplica); err != nil {
		return instance, log.Errore(err)
	}
	if instance.ReplicationIOThreadRuning {
		// Need to apply change by stopping starting IO thread
		ExecInstance(instanceKey, "stop subordinate io_thread")
		if _, err := ExecInstance(instanceKey, "start subordinate io_thread"); err != nil {
			return instance, log.Errore(err)
		}
	}
	return ReadTopologyInstance(instanceKey)
}

func RestartReplicationQuick(instanceKey *InstanceKey) error {
	for _, cmd := range []string{`stop subordinate sql_thread`, `stop subordinate io_thread`, `start subordinate io_thread`, `start subordinate sql_thread`} {
		if _, err := ExecInstance(instanceKey, cmd); err != nil {
			return log.Errorf("%+v: RestartReplicationQuick: '%q' failed: %+v", *instanceKey, cmd, err)
		} else {
			log.Infof("%s on %+v as part of RestartReplicationQuick", cmd, *instanceKey)
		}
	}
	return nil
}

// StopReplicationNicely stops a replica such that SQL_thread and IO_thread are aligned (i.e.
// SQL_thread consumes all relay log entries)
// It will actually START the sql_thread even if the replica is completely stopped.
func StopReplicationNicely(instanceKey *InstanceKey, timeout time.Duration) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if !instance.ReplicationThreadsExist() {
		return instance, fmt.Errorf("instance is not a replica: %+v", instanceKey)
	}

	// stop io_thread, start sql_thread but catch any errors
	for _, cmd := range []string{`stop subordinate io_thread`, `start subordinate sql_thread`} {
		if _, err := ExecInstance(instanceKey, cmd); err != nil {
			return nil, log.Errorf("%+v: StopReplicationNicely: '%q' failed: %+v", *instanceKey, cmd, err)
		}
	}

	if instance.SQLDelay == 0 {
		// Otherwise we don't bother.
		if instance, err = WaitForSQLThreadUpToDate(instanceKey, timeout, 0); err != nil {
			return instance, err
		}
	}

	_, err = ExecInstance(instanceKey, `stop subordinate`)
	if err != nil {
		// Patch; current MaxScale behavior for STOP SLAVE is to throw an error if replica already stopped.
		if instance.isMaxScale() && err.Error() == "Error 1199: Subordinate connection is not running" {
			err = nil
		}
	}
	if err != nil {
		return instance, log.Errore(err)
	}

	instance, err = ReadTopologyInstance(instanceKey)
	log.Infof("Stopped replication nicely on %+v, Self:%+v, Exec:%+v", *instanceKey, instance.SelfBinlogCoordinates, instance.ExecBinlogCoordinates)
	return instance, err
}

func WaitForSQLThreadUpToDate(instanceKey *InstanceKey, overallTimeout time.Duration, staleCoordinatesTimeout time.Duration) (instance *Instance, err error) {
	// Otherwise we don't bother.
	var lastExecBinlogCoordinates BinlogCoordinates

	if overallTimeout == 0 {
		overallTimeout = 24 * time.Hour
	}
	if staleCoordinatesTimeout == 0 {
		staleCoordinatesTimeout = time.Duration(config.Config.ReasonableReplicationLagSeconds) * time.Second
	}
	generalTimer := time.NewTimer(overallTimeout)
	staleTimer := time.NewTimer(staleCoordinatesTimeout)
	for {
		instance, err := RetryInstanceFunction(func() (*Instance, error) {
			return ReadTopologyInstance(instanceKey)
		})
		if err != nil {
			return instance, log.Errore(err)
		}

		if instance.SQLThreadUpToDate() {
			// Woohoo
			return instance, nil
		}
		if instance.SQLDelay != 0 {
			return instance, log.Errorf("WaitForSQLThreadUpToDate: instance %+v has SQL Delay %+v. Operation is irrelevant", *instanceKey, instance.SQLDelay)
		}

		if !instance.ExecBinlogCoordinates.Equals(&lastExecBinlogCoordinates) {
			// means we managed to apply binlog events. We made progress...
			// so we reset the "staleness" timer
			if !staleTimer.Stop() {
				<-staleTimer.C
			}
			staleTimer.Reset(staleCoordinatesTimeout)
		}
		lastExecBinlogCoordinates = instance.ExecBinlogCoordinates

		select {
		case <-generalTimer.C:
			return instance, log.Errorf("WaitForSQLThreadUpToDate timeout on %+v after duration %+v", *instanceKey, overallTimeout)
		case <-staleTimer.C:
			return instance, log.Errorf("WaitForSQLThreadUpToDate stale coordinates timeout on %+v after duration %+v", *instanceKey, staleCoordinatesTimeout)
		default:
			log.Debugf("WaitForSQLThreadUpToDate waiting on %+v", *instanceKey)
			time.Sleep(retryInterval)
		}
	}
}

// StopReplicas will stop replication concurrently on given set of replicas.
// It will potentially do nothing, or attempt to stop _nicely_ or just stop normally, all according to stopReplicationMethod
func StopReplicas(replicas [](*Instance), stopReplicationMethod StopReplicationMethod, timeout time.Duration) [](*Instance) {
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
				if stopReplicationMethod == StopReplicationNice {
					StopReplicationNicely(&replica.Key, timeout)
				}
				replica, _ = StopReplication(&replica.Key)
				updatedReplica = &replica
			})
		}()
	}
	for range replicas {
		refreshedReplicas = append(refreshedReplicas, <-barrier)
	}
	return refreshedReplicas
}

// StopReplicasNicely will attemt to stop all given replicas nicely, up to timeout
func StopReplicasNicely(replicas [](*Instance), timeout time.Duration) [](*Instance) {
	return StopReplicas(replicas, StopReplicationNice, timeout)
}

// StopReplication stops replication on a given instance
func StopReplication(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if !instance.IsReplica() {
		return instance, fmt.Errorf("instance is not a replica: %+v", instanceKey)
	}
	_, err = ExecInstance(instanceKey, `stop subordinate`)
	if err != nil {
		// Patch; current MaxScale behavior for STOP SLAVE is to throw an error if replica already stopped.
		if instance.isMaxScale() && err.Error() == "Error 1199: Subordinate connection is not running" {
			err = nil
		}
	}
	if err != nil {
		return instance, log.Errore(err)
	}
	instance, err = ReadTopologyInstance(instanceKey)

	log.Infof("Stopped replication on %+v, Self:%+v, Exec:%+v", *instanceKey, instance.SelfBinlogCoordinates, instance.ExecBinlogCoordinates)
	return instance, err
}

// waitForReplicationState waits for both replication threads to be either running or not running, together.
// This is useful post- `start subordinate` operation, ensuring both threads are actually running,
// or post `stop subordinate` operation, ensuring both threads are not running.
func waitForReplicationState(instanceKey *InstanceKey, expectedState ReplicationThreadState) (expectationMet bool, err error) {
	waitDuration := time.Second
	waitInterval := 10 * time.Millisecond
	startTime := time.Now()

	for {
		// Since this is an incremental aggressive polling, it's OK if an occasional
		// error is observed. We don't bail out on a single error.
		if expectationMet, _ := expectReplicationThreadsState(instanceKey, expectedState); expectationMet {
			return true, nil
		}
		if time.Since(startTime)+waitInterval > waitDuration {
			break
		}
		time.Sleep(waitInterval)
		waitInterval = 2 * waitInterval
	}
	return false, nil
}

// StartReplication starts replication on a given instance.
func StartReplication(instanceKey *InstanceKey) (*Instance, error) {
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
	// Note: We assume that replicas use 'skip-subordinate-start' so they won't
	//       START SLAVE on their own upon restart.
	if instance.SemiSyncEnforced {
		// Send ACK only from promotable instances.
		sendACK := instance.PromotionRule != MustNotPromoteRule
		// Always disable main setting, in case we're converting a former main.
		if err := EnableSemiSync(instanceKey, false, sendACK); err != nil {
			return instance, log.Errore(err)
		}
	}

	_, err = ExecInstance(instanceKey, `start subordinate`)
	if err != nil {
		return instance, log.Errore(err)
	}
	log.Infof("Started replication on %+v", instanceKey)

	waitForReplicationState(instanceKey, ReplicationThreadStateRunning)

	instance, err = ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}
	if !instance.ReplicaRunning() {
		return instance, ReplicationNotRunningError
	}
	return instance, nil
}

// RestartReplication stops & starts replication on a given instance
func RestartReplication(instanceKey *InstanceKey) (instance *Instance, err error) {
	instance, err = StopReplication(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}
	instance, err = StartReplication(instanceKey)
	return instance, log.Errore(err)
}

// StartReplicas will do concurrent start-replica
func StartReplicas(replicas [](*Instance)) {
	// use concurrency but wait for all to complete
	log.Debugf("Starting %d replicas", len(replicas))
	barrier := make(chan InstanceKey)
	for _, instance := range replicas {
		instance := instance
		go func() {
			// Signal compelted replica
			defer func() { barrier <- instance.Key }()
			// Wait your turn to read a replica
			ExecuteOnTopology(func() { StartReplication(&instance.Key) })
		}()
	}
	for range replicas {
		<-barrier
	}
}

func WaitForExecBinlogCoordinatesToReach(instanceKey *InstanceKey, coordinates *BinlogCoordinates, maxWait time.Duration) (instance *Instance, exactMatch bool, err error) {
	startTime := time.Now()
	for {
		if maxWait != 0 && time.Since(startTime) > maxWait {
			return nil, exactMatch, fmt.Errorf("WaitForExecBinlogCoordinatesToReach: reached maxWait %+v on %+v", maxWait, *instanceKey)
		}
		instance, err = ReadTopologyInstance(instanceKey)
		if err != nil {
			return instance, exactMatch, log.Errore(err)
		}

		switch {
		case instance.ExecBinlogCoordinates.SmallerThan(coordinates):
			time.Sleep(retryInterval)
		case instance.ExecBinlogCoordinates.Equals(coordinates):
			return instance, true, nil
		case coordinates.SmallerThan(&instance.ExecBinlogCoordinates):
			return instance, false, nil
		}
	}
	return instance, exactMatch, err
}

// StartReplicationUntilMainCoordinates issuesa START SLAVE UNTIL... statement on given instance
func StartReplicationUntilMainCoordinates(instanceKey *InstanceKey, mainCoordinates *BinlogCoordinates) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if !instance.IsReplica() {
		return instance, fmt.Errorf("instance is not a replica: %+v", instanceKey)
	}
	if !instance.ReplicationThreadsStopped() {
		return instance, fmt.Errorf("replication threads are not stopped: %+v", instanceKey)
	}

	log.Infof("Will start replication on %+v until coordinates: %+v", instanceKey, mainCoordinates)

	if instance.SemiSyncEnforced {
		// Send ACK only from promotable instances.
		sendACK := instance.PromotionRule != MustNotPromoteRule
		// Always disable main setting, in case we're converting a former main.
		if err := EnableSemiSync(instanceKey, false, sendACK); err != nil {
			return instance, log.Errore(err)
		}
	}

	// MariaDB has a bug: a CHANGE MASTER TO statement does not work properly with prepared statement... :P
	// See https://mariadb.atlassian.net/browse/MDEV-7640
	// This is the reason for ExecInstance
	_, err = ExecInstance(instanceKey, "start subordinate until main_log_file=?, main_log_pos=?",
		mainCoordinates.LogFile, mainCoordinates.LogPos)
	if err != nil {
		return instance, log.Errore(err)
	}

	instance, exactMatch, err := WaitForExecBinlogCoordinatesToReach(instanceKey, mainCoordinates, 0)
	if err != nil {
		return instance, log.Errore(err)
	}
	if !exactMatch {
		return instance, fmt.Errorf("Start SLAVE UNTIL is past coordinates: %+v", instanceKey)
	}

	instance, err = StopReplication(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	return instance, err
}

// EnableSemiSync sets the rpl_semi_sync_(main|replica)_enabled variables
// on a given instance.
func EnableSemiSync(instanceKey *InstanceKey, main, replica bool) error {
	log.Infof("instance %+v rpl_semi_sync_main_enabled: %t, rpl_semi_sync_subordinate_enabled: %t", instanceKey, main, replica)
	_, err := ExecInstance(instanceKey,
		`set global rpl_semi_sync_main_enabled = ?, global rpl_semi_sync_subordinate_enabled = ?`,
		main, replica)
	return err
}

// ChangeMainCredentials issues a CHANGE MASTER TO... MASTER_USER=, MASTER_PASSWORD=...
func ChangeMainCredentials(instanceKey *InstanceKey, mainUser string, mainPassword string) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}
	if mainUser == "" {
		return instance, log.Errorf("Empty user in ChangeMainCredentials() for %+v", *instanceKey)
	}

	if instance.ReplicationThreadsExist() && !instance.ReplicationThreadsStopped() {
		return instance, fmt.Errorf("ChangeMainTo: Cannot change main on: %+v because replication is running", *instanceKey)
	}
	log.Debugf("ChangeMainTo: will attempt changing main credentials on %+v", *instanceKey)

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting CHANGE MASTER TO operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}
	_, err = ExecInstance(instanceKey, "change main to main_user=?, main_password=?",
		mainUser, mainPassword)

	if err != nil {
		return instance, log.Errore(err)
	}

	log.Infof("ChangeMainTo: Changed main credentials on %+v", *instanceKey)

	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}

// EnableMainSSL issues CHANGE MASTER TO MASTER_SSL=1
func EnableMainSSL(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if instance.ReplicationThreadsExist() && !instance.ReplicationThreadsStopped() {
		return instance, fmt.Errorf("EnableMainSSL: Cannot enable SSL replication on %+v because replication threads are not stopped", *instanceKey)
	}
	log.Debugf("EnableMainSSL: Will attempt enabling SSL replication on %+v", *instanceKey)

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting CHANGE MASTER TO MASTER_SSL=1 operation on %+v; signaling error but nothing went wrong.", *instanceKey)
	}
	_, err = ExecInstance(instanceKey, "change main to main_ssl=1")

	if err != nil {
		return instance, log.Errore(err)
	}

	log.Infof("EnableMainSSL: Enabled SSL replication on %+v", *instanceKey)

	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}

// See https://bugs.mysql.com/bug.php?id=83713
func workaroundBug83713(instanceKey *InstanceKey) {
	log.Debugf("workaroundBug83713: %+v", *instanceKey)
	queries := []string{
		`reset subordinate`,
		`start subordinate IO_THREAD`,
		`stop subordinate IO_THREAD`,
		`reset subordinate`,
	}
	for _, query := range queries {
		if _, err := ExecInstance(instanceKey, query); err != nil {
			log.Debugf("workaroundBug83713: error on %s: %+v", query, err)
		}
	}
}

// ChangeMainTo changes the given instance's main according to given input.
func ChangeMainTo(instanceKey *InstanceKey, mainKey *InstanceKey, mainBinlogCoordinates *BinlogCoordinates, skipUnresolve bool, gtidHint OperationGTIDHint) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if instance.ReplicationThreadsExist() && !instance.ReplicationThreadsStopped() {
		return instance, fmt.Errorf("ChangeMainTo: Cannot change main on: %+v because replication threads are not stopped", *instanceKey)
	}
	log.Debugf("ChangeMainTo: will attempt changing main on %+v to %+v, %+v", *instanceKey, *mainKey, *mainBinlogCoordinates)
	changeToMainKey := mainKey
	if !skipUnresolve {
		unresolvedMainKey, nameUnresolved, err := UnresolveHostname(mainKey)
		if err != nil {
			log.Debugf("ChangeMainTo: aborting operation on %+v due to resolving error on %+v: %+v", *instanceKey, *mainKey, err)
			return instance, err
		}
		if nameUnresolved {
			log.Debugf("ChangeMainTo: Unresolved %+v into %+v", *mainKey, unresolvedMainKey)
		}
		changeToMainKey = &unresolvedMainKey
	}

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting CHANGE MASTER TO operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}

	originalMainKey := instance.MainKey
	originalExecBinlogCoordinates := instance.ExecBinlogCoordinates

	var changeMainFunc func() error
	changedViaGTID := false
	if instance.UsingMariaDBGTID && gtidHint != GTIDHintDeny {
		// Keep on using GTID
		changeMainFunc = func() error {
			_, err := ExecInstance(instanceKey, "change main to main_host=?, main_port=?",
				changeToMainKey.Hostname, changeToMainKey.Port)
			return err
		}
		changedViaGTID = true
	} else if instance.UsingMariaDBGTID && gtidHint == GTIDHintDeny {
		// Make sure to not use GTID
		changeMainFunc = func() error {
			_, err = ExecInstance(instanceKey, "change main to main_host=?, main_port=?, main_log_file=?, main_log_pos=?, main_use_gtid=no",
				changeToMainKey.Hostname, changeToMainKey.Port, mainBinlogCoordinates.LogFile, mainBinlogCoordinates.LogPos)
			return err
		}
	} else if instance.IsMariaDB() && gtidHint == GTIDHintForce {
		// Is MariaDB; not using GTID, turn into GTID
		mariadbGTIDHint := "subordinate_pos"
		if !instance.ReplicationThreadsExist() {
			// This instance is currently a main. As per https://mariadb.com/kb/en/change-master-to/#master_use_gtid
			// we should be using current_pos.
			// See also:
			// - https://github.com/openark/orchestrator/issues/1146
			// - https://dba.stackexchange.com/a/234323
			mariadbGTIDHint = "current_pos"
		}
		changeMainFunc = func() error {
			_, err = ExecInstance(instanceKey, fmt.Sprintf("change main to main_host=?, main_port=?, main_use_gtid=%s", mariadbGTIDHint),
				changeToMainKey.Hostname, changeToMainKey.Port)
			return err
		}
		changedViaGTID = true
	} else if instance.UsingOracleGTID && gtidHint != GTIDHintDeny {
		// Is Oracle; already uses GTID; keep using it.
		changeMainFunc = func() error {
			_, err = ExecInstance(instanceKey, "change main to main_host=?, main_port=?",
				changeToMainKey.Hostname, changeToMainKey.Port)
			return err
		}
		changedViaGTID = true
	} else if instance.UsingOracleGTID && gtidHint == GTIDHintDeny {
		// Is Oracle; already uses GTID
		changeMainFunc = func() error {
			_, err = ExecInstance(instanceKey, "change main to main_host=?, main_port=?, main_log_file=?, main_log_pos=?, main_auto_position=0",
				changeToMainKey.Hostname, changeToMainKey.Port, mainBinlogCoordinates.LogFile, mainBinlogCoordinates.LogPos)
			return err
		}
	} else if instance.SupportsOracleGTID && gtidHint == GTIDHintForce {
		// Is Oracle; not using GTID right now; turn into GTID
		changeMainFunc = func() error {
			_, err = ExecInstance(instanceKey, "change main to main_host=?, main_port=?, main_auto_position=1",
				changeToMainKey.Hostname, changeToMainKey.Port)
			return err
		}
		changedViaGTID = true
	} else {
		// Normal binlog file:pos
		changeMainFunc = func() error {
			_, err = ExecInstance(instanceKey, "change main to main_host=?, main_port=?, main_log_file=?, main_log_pos=?",
				changeToMainKey.Hostname, changeToMainKey.Port, mainBinlogCoordinates.LogFile, mainBinlogCoordinates.LogPos)
			return err
		}
	}
	err = changeMainFunc()
	if err != nil && instance.UsingOracleGTID && strings.Contains(err.Error(), Error1201CouldnotInitializeMainInfoStructure) {
		log.Debugf("ChangeMainTo: got %+v", err)
		workaroundBug83713(instanceKey)
		err = changeMainFunc()
	}
	if err != nil {
		return instance, log.Errore(err)
	}
	WriteMainPositionEquivalence(&originalMainKey, &originalExecBinlogCoordinates, changeToMainKey, mainBinlogCoordinates)
	ResetInstanceRelaylogCoordinatesHistory(instanceKey)

	log.Infof("ChangeMainTo: Changed main on %+v to: %+v, %+v. GTID: %+v", *instanceKey, mainKey, mainBinlogCoordinates, changedViaGTID)

	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}

// SkipToNextBinaryLog changes main position to beginning of next binlog
// USE WITH CARE!
// Use case is binlog servers where the main was gone & replaced by another.
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

	instance, err = ChangeMainTo(&instance.Key, &instance.MainKey, &nextFileCoordinates, false, GTIDHintNeutral)
	if err != nil {
		return instance, log.Errore(err)
	}
	AuditOperation("skip-binlog", instanceKey, fmt.Sprintf("Skipped replication to next binary log: %+v", nextFileCoordinates.LogFile))
	return StartReplication(instanceKey)
}

// ResetReplication resets a replica, breaking the replication
func ResetReplication(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if instance.ReplicationThreadsExist() && !instance.ReplicationThreadsStopped() {
		return instance, fmt.Errorf("Cannot reset replication on: %+v because replication threads are not stopped", instanceKey)
	}

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting reset-replication operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}

	// MySQL's RESET SLAVE is done correctly; however SHOW SLAVE STATUS still returns old hostnames etc
	// and only resets till after next restart. This leads to orchestrator still thinking the instance replicates
	// from old host. We therefore forcibly modify the hostname.
	// RESET SLAVE ALL command solves this, but only as of 5.6.3
	_, err = ExecInstance(instanceKey, `change main to main_host='_'`)
	if err != nil {
		return instance, log.Errore(err)
	}
	_, err = ExecInstance(instanceKey, `reset subordinate /*!50603 all */`)
	if err != nil && strings.Contains(err.Error(), Error1201CouldnotInitializeMainInfoStructure) {
		log.Debugf("ResetReplication: got %+v", err)
		workaroundBug83713(instanceKey)
		_, err = ExecInstance(instanceKey, `reset subordinate /*!50603 all */`)
	}
	if err != nil {
		return instance, log.Errore(err)
	}
	log.Infof("Reset replication %+v", instanceKey)

	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}

// ResetMain issues a RESET MASTER statement on given instance. Use with extreme care!
func ResetMain(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if instance.ReplicationThreadsExist() && !instance.ReplicationThreadsStopped() {
		return instance, fmt.Errorf("Cannot reset main on: %+v because replication threads are not stopped", instanceKey)
	}

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting reset-main operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}

	_, err = ExecInstance(instanceKey, `reset main`)
	if err != nil {
		return instance, log.Errore(err)
	}
	log.Infof("Reset main %+v", instanceKey)

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

// injectEmptyGTIDTransaction
func injectEmptyGTIDTransaction(instanceKey *InstanceKey, gtidEntry *OracleGtidSetEntry) error {
	db, err := db.OpenTopology(instanceKey.Hostname, instanceKey.Port)
	if err != nil {
		return err
	}
	ctx := context.Background()
	conn, err := db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	if _, err := conn.ExecContext(ctx, fmt.Sprintf(`SET GTID_NEXT="%s"`, gtidEntry.String())); err != nil {
		return err
	}
	tx, err := conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	if _, err := conn.ExecContext(ctx, `SET GTID_NEXT="AUTOMATIC"`); err != nil {
		return err
	}
	return nil
}

// skipQueryClassic skips a query in normal binlog file:pos replication
func skipQueryClassic(instance *Instance) error {
	_, err := ExecInstance(&instance.Key, `set global sql_subordinate_skip_counter := 1`)
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
	if instance.ReplicationSQLThreadRuning {
		return instance, fmt.Errorf("Replication SQL thread is running on %+v", instanceKey)
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
	return StartReplication(instanceKey)
}

// MainPosWait issues a MASTER_POS_WAIT() an given instance according to given coordinates.
func MainPosWait(instanceKey *InstanceKey, binlogCoordinates *BinlogCoordinates) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	_, err = ExecInstance(instanceKey, `select main_pos_wait(?, ?)`, binlogCoordinates.LogFile, binlogCoordinates.LogPos)
	if err != nil {
		return instance, log.Errore(err)
	}
	log.Infof("Instance %+v has reached coordinates: %+v", instanceKey, binlogCoordinates)

	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}

// Attempt to read and return replication credentials from the mysql.subordinate_main_info system table
func ReadReplicationCredentials(instanceKey *InstanceKey) (replicationUser string, replicationPassword string, err error) {
	if config.Config.ReplicationCredentialsQuery != "" {
		err = ScanInstanceRow(instanceKey, config.Config.ReplicationCredentialsQuery, &replicationUser, &replicationPassword)
		if err == nil && replicationUser == "" {
			err = fmt.Errorf("Empty username retrieved by ReplicationCredentialsQuery")
		}
		if err == nil {
			return replicationUser, replicationPassword, nil
		}
		log.Errore(err)
	}
	// Didn't get credentials from ReplicationCredentialsQuery, or ReplicationCredentialsQuery doesn't exist in the first place?
	// We brute force our way through mysql.subordinate_main_info
	{
		query := `
			select
				ifnull(max(User_name), '') as user,
				ifnull(max(User_password), '') as password
			from
				mysql.subordinate_main_info
		`
		err = ScanInstanceRow(instanceKey, query, &replicationUser, &replicationPassword)
		if err == nil && replicationUser == "" {
			err = fmt.Errorf("Empty username found in mysql.subordinate_main_info")
		}
	}
	return replicationUser, replicationPassword, log.Errore(err)
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

	// If async fallback is disallowed, we're responsible for flipping the main
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

	// If we just went read-only, it's safe to flip the main semi-sync switch
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
	if instance == nil {
		return injected, log.Errorf("CheckAndInjectPseudoGTIDOnWriter: instance is nil")
	}
	if instance.ReadOnly {
		return injected, log.Errorf("CheckAndInjectPseudoGTIDOnWriter: instance is read-only: %+v", instance.Key)
	}
	if !instance.IsLastCheckValid {
		return injected, nil
	}
	canInject, err := canInjectPseudoGTID(&instance.Key)
	if err != nil {
		return injected, log.Errore(err)
	}
	if !canInject {
		if util.ClearToLog("CheckAndInjectPseudoGTIDOnWriter", instance.Key.StringCode()) {
			log.Warningf("AutoPseudoGTID enabled, but orchestrator has no priviliges on %+v to inject pseudo-gtid", instance.Key)
		}

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

func GTIDSubtract(instanceKey *InstanceKey, gtidSet string, gtidSubset string) (gtidSubtract string, err error) {
	db, err := db.OpenTopology(instanceKey.Hostname, instanceKey.Port)
	if err != nil {
		return gtidSubtract, err
	}
	err = db.QueryRow("select gtid_subtract(?, ?)", gtidSet, gtidSubset).Scan(&gtidSubtract)
	return gtidSubtract, err
}

func ShowMainStatus(instanceKey *InstanceKey) (mainStatusFound bool, executedGtidSet string, err error) {
	db, err := db.OpenTopology(instanceKey.Hostname, instanceKey.Port)
	if err != nil {
		return mainStatusFound, executedGtidSet, err
	}
	err = sqlutils.QueryRowsMap(db, "show main status", func(m sqlutils.RowMap) error {
		mainStatusFound = true
		executedGtidSet = m.GetStringD("Executed_Gtid_Set", "")
		return nil
	})
	return mainStatusFound, executedGtidSet, err
}

func ShowBinaryLogs(instanceKey *InstanceKey) (binlogs []string, err error) {
	db, err := db.OpenTopology(instanceKey.Hostname, instanceKey.Port)
	if err != nil {
		return binlogs, err
	}
	err = sqlutils.QueryRowsMap(db, "show binary logs", func(m sqlutils.RowMap) error {
		binlogs = append(binlogs, m.GetString("Log_name"))
		return nil
	})
	return binlogs, err
}
