/*
   Copyright 2017 GitHub Inc.

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

package remote

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/github/orchestrator/go/config"
	"github.com/github/orchestrator/go/inst"
	orcos "github.com/github/orchestrator/go/os"
	"github.com/github/orchestrator/go/process"
	"github.com/openark/golib/log"
)

// SyncRelaylogsChangeMasterIdentityFunc returns the identity and coorindates into which a replica would CHANGE MASTER TO
// having been applied with someone else's relaylogs.
// As examples, in testing we would point such replica back to the master (but with updated coordinates)
// In real master failover we would point to a candidate replica's self coordinates
type SyncRelaylogsChangeMasterIdentityFunc func(sourceReplica *inst.Instance, replicas ...*inst.Instance) (masterKey *inst.InstanceKey, coordinates *inst.BinlogCoordinates, err error)

// SyncRelaylogsChangeMasterToSharedMasterFunc points applied-to replica to the applied-from replica's master coordinates
var SyncRelaylogsChangeMasterToSharedMasterFunc SyncRelaylogsChangeMasterIdentityFunc = func(sourceReplica *inst.Instance, replicas ...*inst.Instance) (masterKey *inst.InstanceKey, coordinates *inst.BinlogCoordinates, err error) {
	return &sourceReplica.MasterKey, &sourceReplica.ExecBinlogCoordinates, nil
}

// SyncRelaylogsChangeMasterToSourceReplicaFunc points applied-to replica to the applied-from replica
var SyncRelaylogsChangeMasterToSourceReplicaFunc SyncRelaylogsChangeMasterIdentityFunc = func(sourceReplica *inst.Instance, replicas ...*inst.Instance) (masterKey *inst.InstanceKey, coordinates *inst.BinlogCoordinates, err error) {
	if !sourceReplica.LogBinEnabled {
		return nil, nil, fmt.Errorf("Cannot change master onto source replica %+v since it has no binary logs", sourceReplica.Key)
	}
	return &sourceReplica.Key, &sourceReplica.SelfBinlogCoordinates, nil
}

// SyncRelaylogsChangeMasterToTargetReplicaFunc points replica to the first target replica in given list
var SyncRelaylogsChangeMasterToTargetReplicaFunc SyncRelaylogsChangeMasterIdentityFunc = func(sourceReplica *inst.Instance, replicas ...*inst.Instance) (masterKey *inst.InstanceKey, coordinates *inst.BinlogCoordinates, err error) {
	if len(replicas) == 0 {
		return nil, nil, fmt.Errorf("No replicas given to SyncRelaylogsChangeMasterToTargetReplicaFunc")
	}
	replica := replicas[0]
	if !replica.LogBinEnabled {
		return nil, nil, fmt.Errorf("Cannot change master onto replica %+v since it has no binary logs", replica.Key)
	}
	return &replica.Key, &replica.SelfBinlogCoordinates, nil
}

func TestRemoteCommandOnInstance(instanceKey *inst.InstanceKey) error {
	sudoCommand := ""
	if config.Config.RemoteSSHCommandUseSudo {
		sudoCommand = "sudo -i"
	}

	tempFile, err := ioutil.TempFile("", "orchestrator-remote-test-")
	if err != nil {
		return err
	}
	defer os.Remove(tempFile.Name())

	randomToken := process.NewToken()

	command := config.Config.RemoteSSHCommand
	command = strings.Replace(command, "{hostname}", instanceKey.Hostname, -1)
	command = fmt.Sprintf("%s '%s echo %s' > %s", command, sudoCommand, randomToken.Hash, tempFile.Name())
	if err := orcos.CommandRun(command, orcos.EmptyEnv); err != nil {
		return err
	}
	bytes, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		return err
	}
	if content := strings.TrimSpace(string(bytes)); content != randomToken.Hash {
		return fmt.Errorf("TestRemoteCommandOnInstance: expected %s, got %s", randomToken.Hash, content)
	}
	return nil
}

func applySyncRelaylogsChangeMasterIdentityFunc(syncRelaylogsChangeMasterIdentityFunc SyncRelaylogsChangeMasterIdentityFunc, fromInstance, instance *inst.Instance) (updatedInstance *inst.Instance, err error) {
	if syncRelaylogsChangeMasterIdentityFunc == nil {
		return instance, nil
	}
	changeToKey, changeToCoordinates, err := syncRelaylogsChangeMasterIdentityFunc(fromInstance, instance)
	if err != nil {
		return instance, err
	}
	if changeToKey == nil {
		return instance, nil
	}
	if changeToCoordinates == nil {
		return instance, nil
	}
	if instance.Key.Equals(changeToKey) {
		return instance, nil
	}
	// We exempt the very instance from pointing onto itself

	instance, err = inst.ChangeMasterTo(&instance.Key, changeToKey, changeToCoordinates, false, inst.GTIDHintNeutral)

	return instance, nil
}

// SyncReplicaRelayLogs will align siblings by applying relaylogs from one to the other, via remote SSH
func SyncReplicaRelayLogs(instance, fromInstance *inst.Instance, syncRelaylogsChangeMasterIdentityFunc SyncRelaylogsChangeMasterIdentityFunc, startReplication bool) (*inst.Instance, error) {
	if config.Config.RemoteSSHCommand == "" {
		return instance, fmt.Errorf("RemoteSSHCommand not configured")
	}
	log.Debugf("Testing SSH on %+v", instance.Key)
	if err := TestRemoteCommandOnInstance(&instance.Key); err != nil {
		return instance, err
	}
	needToSync := true
	if instance.Key.Equals(&fromInstance.Key) {
		log.Debugf("SyncReplicaRelayLogs: will not pply from %+v onto itself.", instance.Key)
		needToSync = false
	}
	if needToSync && instance.ReplicaRunning() {
		return instance, log.Errorf("SyncReplicaRelayLogs: replication on %+v must not run", instance.Key)
	}
	if needToSync && fromInstance.ExecBinlogCoordinates.SmallerThan(&instance.ExecBinlogCoordinates) {
		return instance, log.Errorf("SyncReplicaRelayLogs: %+v is more up to date than %+v. Will not apply relaylogs from %+v", instance.Key, fromInstance.Key, fromInstance.Key)
	}
	if needToSync && fromInstance.ExecBinlogCoordinates.Equals(&instance.ExecBinlogCoordinates) {
		log.Debugf("SyncReplicaRelayLogs: %+v and %+v are at same coordinates. Not actually applying relay logs.", instance.Key, fromInstance.Key)
		needToSync = false
	}

	if needToSync {
		log.Debugf("SyncReplicaRelayLogs: correlating coordinates of %+v on %+v", instance.Key, fromInstance.Key)
		_, _, nextCoordinates, found, err := inst.CorrelateRelaylogCoordinates(instance, nil, fromInstance)
		if err != nil {
			return instance, err
		}
		if !found {
			return instance, err
		}
		log.Debugf("SyncReplicaRelayLogs: correlated next-coordinates of %+v on %+v are %+v", instance.Key, fromInstance.Key, *nextCoordinates)

		// We now have the correlation info needed to proceed with remote calls
		sudoCommand := ""
		if config.Config.RemoteSSHCommandUseSudo {
			sudoCommand = "sudo -i"
		}

		// Write get-relaylogs script locally
		getRelayLogContentsScriptFile, err := ioutil.TempFile("", "orchestrator-remote-get-relaylogs-content-")
		if err != nil {
			return instance, err
		}
		{
			defer os.Remove(getRelayLogContentsScriptFile.Name())
			script := GetRelayLogContentsScript
			script = strings.Replace(script, "$MAGIC_FIRST_RELAYLOG_FILE", nextCoordinates.LogFile, -1)
			script = strings.Replace(script, "$MAGIC_LAST_RELAYLOG_FILE", fromInstance.RelaylogCoordinates.LogFile, -1)
			script = strings.Replace(script, "$MAGIC_START_POSITION", fmt.Sprintf("%d", nextCoordinates.LogPos), -1)
			log.Debugf("fromInstance.RelaylogCoordinates = %+v", fromInstance.RelaylogCoordinates)
			script = strings.Replace(script, "$MAGIC_STOP_POSITION", fmt.Sprintf("%d", fromInstance.RelaylogCoordinates.LogPos), -1)
			if err := ioutil.WriteFile(getRelayLogContentsScriptFile.Name(), []byte(script), 0640); err != nil {
				return instance, err
			}
		}
		log.Debugf("getRelayLogContentsScriptFile: %+v", getRelayLogContentsScriptFile.Name())

		// Get relay log contents, save locally
		localRelayLogContentsFile, err := ioutil.TempFile("", "orchestrator-remote-relaylogs-content-")
		if err != nil {
			return instance, err
		}
		defer os.Remove(localRelayLogContentsFile.Name())
		localRelayLogContentsCopyFileName := fmt.Sprintf("%s.copy", localRelayLogContentsFile.Name())
		{
			command := config.Config.RemoteSSHCommand
			command = strings.Replace(command, "{hostname}", fromInstance.Key.Hostname, -1)
			command = fmt.Sprintf("cat %s | %s '%s' > %s", getRelayLogContentsScriptFile.Name(), command, sudoCommand, localRelayLogContentsFile.Name())
			if err := orcos.CommandRun(command, orcos.EmptyEnv); err != nil {
				return instance, err
			}
		}
		log.Debugf("Have fetched relay logs from %s, output file is %s", fromInstance.Key.Hostname, localRelayLogContentsFile.Name())
		// Copy local relay log contents to target host:
		{
			command := config.Config.RemoteSSHCommand
			command = strings.Replace(command, "{hostname}", instance.Key.Hostname, -1)
			command = fmt.Sprintf("cat %s | %s '%s cat - > %s'", localRelayLogContentsFile.Name(), command, sudoCommand, localRelayLogContentsCopyFileName)
			if err := orcos.CommandRun(command, orcos.EmptyEnv); err != nil {
				return instance, err
			}
		}
		log.Debugf("Have copied contents file to %s, output file is %s", instance.Key.Hostname, localRelayLogContentsFile.Name())

		// Generate the apply-relaylogs script, locally
		applyRelayLogContentsScriptFile, err := ioutil.TempFile("", "orchestrator-remote-apply-relaylogs-content-")
		if err != nil {
			return instance, err
		}
		{
			defer os.Remove(applyRelayLogContentsScriptFile.Name())
			script := ApplyRelayLogContentsScript
			script = strings.Replace(script, "$MAGIC_MYSQL_COMMAND", "", -1)
			script = strings.Replace(script, "$MAGIC_CONTENTS_FILE", localRelayLogContentsCopyFileName, -1)

			if err := ioutil.WriteFile(applyRelayLogContentsScriptFile.Name(), []byte(script), 0640); err != nil {
				return instance, err
			}
		}
		log.Debugf("applyRelayLogContentsScriptFile: %+v", applyRelayLogContentsScriptFile.Name())

		if *config.RuntimeCLIFlags.Noop {
			return instance, fmt.Errorf("noop: Not really applying scripts onto %+v; signalling error but nothing went wrong", instance.Key)
		}

		// apply relaylog contents on target host:
		{
			command := config.Config.RemoteSSHCommand
			command = strings.Replace(command, "{hostname}", instance.Key.Hostname, -1)
			command = fmt.Sprintf("cat %s | %s '%s'", applyRelayLogContentsScriptFile.Name(), command, sudoCommand)
			if err := orcos.CommandRun(command, orcos.EmptyEnv); err != nil {
				return instance, err
			}
		}
		log.Debugf("Have successfully applied relay logs on %s", instance.Key.Hostname)
	}

	var err error
	if instance, err = applySyncRelaylogsChangeMasterIdentityFunc(syncRelaylogsChangeMasterIdentityFunc, fromInstance, instance); err != nil {
		return instance, err
	}
	if startReplication {
		if instance, err = inst.StartSlave(&instance.Key); err != nil {
			return instance, err
		}
	}

	inst.AuditOperation("align-via-relaylogs-remote", &instance.Key, fmt.Sprintf("aligned %+v by relaylogs from %+v", instance.Key, fromInstance.Key))
	return instance, nil
}

func SyncReplicasRelayLogs(
	master *inst.Instance,
	syncRelaylogsChangeMasterIdentityFunc SyncRelaylogsChangeMasterIdentityFunc,
	avoidPostponingInstance *inst.Instance,
	startReplication bool,
	postponedFunctionsContainer *inst.PostponedFunctionsContainer,
) (
	synchedFromReplica *inst.Instance,
	syncedReplicas, failedReplicas, postponedReplicas [](*inst.Instance),
	err error,
) {
	postponedFunctionsContainer = nil
	var replicas [](*inst.Instance)
	if replicas, err = inst.GetSortedReplicas(&master.Key, inst.StopReplicationNormal); err != nil {
		return synchedFromReplica, syncedReplicas, replicas, postponedReplicas, err
	}

	synchedFromReplica = replicas[0]
	if len(replicas) <= 1 {
		// Nothing to be done
		return synchedFromReplica, syncedReplicas, replicas, postponedReplicas, err
	}
	applyToReplicas := replicas[1:]

	if err != nil {
		return synchedFromReplica, syncedReplicas, replicas, postponedReplicas, err
	}

	log.Debugf("SyncReplicasRelayLogs: Testing SSH on %+v", synchedFromReplica.Key)
	if err := TestRemoteCommandOnInstance(&synchedFromReplica.Key); err != nil {
		return synchedFromReplica, syncedReplicas, replicas, postponedReplicas, err
	}

	barrier := make(chan *inst.InstanceKey, len(applyToReplicas))
	allErrors := make(chan error, len(applyToReplicas))
	synchedReplicasChan := make(chan *inst.Instance, len(applyToReplicas))
	failedReplicasChan := make(chan *inst.Instance, len(applyToReplicas))

	applyToReplicaFunc := func(applyToReplica *inst.Instance) error {
		defer func() { barrier <- &applyToReplica.Key }()

		if _, err := SyncReplicaRelayLogs(applyToReplica, synchedFromReplica, syncRelaylogsChangeMasterIdentityFunc, startReplication); err == nil {
			synchedReplicasChan <- applyToReplica
		} else {
			err = fmt.Errorf("%+v: %+v", applyToReplica.Key, err.Error())
			failedReplicasChan <- applyToReplica
			allErrors <- err
		}
		return err
	}

	log.Debugf("SyncReplicasRelayLogs: Applying relay logs on %+v replicas", len(applyToReplicas))
	countImmediateApply := 0
	for _, applyToReplica := range applyToReplicas {
		applyToReplica := applyToReplica

		toBePostponed := false
		if postponedFunctionsContainer != nil {
			if config.Config.PostponeReplicaRecoveryOnLagMinutes > 0 &&
				applyToReplica.SQLDelay > config.Config.PostponeReplicaRecoveryOnLagMinutes*60 {
				toBePostponed = true
			}
			if applyToReplica.DataCenter != master.DataCenter {
				toBePostponed = true
			}
		}
		if avoidPostponingInstance != nil {
			if avoidPostponingInstance.Key.Equals(&applyToReplica.Key) {
				toBePostponed = false
			}
		}
		if toBePostponed {
			log.Debugf("SyncReplicasRelayLogs: postponing relaylog sync on %+v", applyToReplica.Key)
			postponedReplicas = append(postponedReplicas, applyToReplica)
			f := func() error { return applyToReplicaFunc(applyToReplica) }
			(*postponedFunctionsContainer).AddPostponedFunction(f, fmt.Sprintf("sync-replicas-relaylogs %+v", applyToReplica.Key))
		} else {
			log.Debugf("SyncReplicasRelayLogs: applying relaylog sync on %+v", applyToReplica.Key)
			countImmediateApply++
			go applyToReplicaFunc(applyToReplica)
		}
	}
	for i := 0; i < countImmediateApply; i++ {
		<-barrier
	}
	syncedReplicas = append(syncedReplicas, synchedFromReplica)
	for len(synchedReplicasChan) > 0 {
		syncedReplicas = append(syncedReplicas, <-synchedReplicasChan)
	}
	for len(failedReplicasChan) > 0 {
		failedReplicas = append(failedReplicas, <-failedReplicasChan)
	}
	countErrors := len(allErrors)
	for len(allErrors) > 0 {
		log.Errore(<-allErrors)
	}
	if synchedFromReplica, err = applySyncRelaylogsChangeMasterIdentityFunc(syncRelaylogsChangeMasterIdentityFunc, synchedFromReplica, synchedFromReplica); err != nil {
		return synchedFromReplica, syncedReplicas, failedReplicas, postponedReplicas, err
	}
	if startReplication {
		synchedFromReplica, err = inst.StartSlave(&synchedFromReplica.Key)
		log.Errore(err)
	}

	inst.AuditOperation("sync-replicas-relaylogs", &master.Key, fmt.Sprintf("aligned %+v replicas by relaylogs from %+v, got %+v errors", len(applyToReplicas), synchedFromReplica.Key, countErrors))
	return synchedFromReplica, syncedReplicas, failedReplicas, postponedReplicas, err
}

func SyncRelayLogsToSingleReplica(
	master *inst.Instance,
	syncRelaylogsChangeMasterIdentityFunc SyncRelaylogsChangeMasterIdentityFunc,
	startReplication bool,
) (synchedFromReplica, synchedToReplica *inst.Instance, err error) {
	var replicas [](*inst.Instance)
	if replicas, err = inst.GetSortedReplicas(&master.Key, inst.StopReplicationNormal); err != nil {
		return synchedFromReplica, synchedToReplica, err
	}

	synchedFromReplica = replicas[0]
	if len(replicas) <= 1 {
		// Nothing to be done
		return synchedFromReplica, synchedToReplica, err
	}
	applyToReplicas := replicas[1:]

	log.Debugf("SyncReplicasRelayLogs: Testing SSH on %+v", synchedFromReplica.Key)
	if err := TestRemoteCommandOnInstance(&synchedFromReplica.Key); err != nil {
		return synchedFromReplica, synchedToReplica, err
	}

	log.Debugf("SyncReplicasRelayLogs: searching for replica to apply relay logs to")
	for _, applyToReplica := range applyToReplicas {
		applyToReplica := applyToReplica

		if applyToReplica, err := SyncReplicaRelayLogs(applyToReplica, synchedFromReplica, syncRelaylogsChangeMasterIdentityFunc, startReplication); err == nil {
			log.Debugf("SyncReplicasRelayLogs: picked %+v", applyToReplica.Key)
			synchedToReplica = applyToReplica
			break
		}
	}
	if synchedFromReplica, err = applySyncRelaylogsChangeMasterIdentityFunc(syncRelaylogsChangeMasterIdentityFunc, synchedFromReplica, synchedFromReplica); err != nil {
		return synchedFromReplica, synchedToReplica, err
	}
	if startReplication {
		synchedFromReplica, err = inst.StartSlave(&synchedFromReplica.Key)
		log.Errore(err)
	}

	inst.AuditOperation("sync-replicas-relaylogs-to-single", &master.Key, fmt.Sprintf("applied relaylogs from %+v to %+v", synchedFromReplica.Key, synchedToReplica.Key))
	return synchedFromReplica, synchedToReplica, nil
}
