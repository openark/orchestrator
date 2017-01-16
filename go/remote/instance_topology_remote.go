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
	"github.com/outbrain/golib/log"
)

type SyncRelaylogsChangeMasterMethod string

const (
	SyncRelaylogsNoChangeMaster              SyncRelaylogsChangeMasterMethod = "SyncRelaylogsNoChangeMaster"
	SyncRelaylogsChangeMasterToSharedMaster                                  = "SyncRelaylogsChangeMasterToSharedMaster"
	SyncRelaylogsChangeMasterToSourceReplica                                 = "SyncRelaylogsChangeMasterToSourceReplica"
)

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

// SyncReplicaRelayLogs will align siblings by applying relaylogs from one to the other, via remote SSH
func SyncReplicaRelayLogs(instance, fromInstance *inst.Instance, syncRelaylogsChangeMasterMethod SyncRelaylogsChangeMasterMethod) (*inst.Instance, error) {
	if config.Config.RemoteSSHCommand == "" {
		return instance, fmt.Errorf("RemoteSSHCommand not configured")
	}
	log.Debugf("Testing SSH on %+v", instance.Key)
	if err := TestRemoteCommandOnInstance(&instance.Key); err != nil {
		return instance, err
	}

	if syncRelaylogsChangeMasterMethod == SyncRelaylogsChangeMasterToSourceReplica && !fromInstance.LogBinEnabled {
		return instance, log.Errorf("SyncReplicaRelayLogs: cannot change master to source replica %+v since it has no binary logs", fromInstance.Key)
	}
	if instance.ReplicaRunning() {
		return instance, log.Errorf("SyncReplicaRelayLogs: replication on %+v must not run", instance.Key)
	}
	if fromInstance.ReplicaRunning() {
		return instance, log.Errorf("SyncReplicaRelayLogs: replication on %+v must not run", fromInstance.Key)
	}
	if fromInstance.ExecBinlogCoordinates.Equals(&instance.ExecBinlogCoordinates) {
		log.Debugf("SyncReplicaRelayLogs: %+v and %+v are at same coordinates. Nothing to apply.", instance.Key, fromInstance.Key)
		return instance, nil
	}
	if fromInstance.ExecBinlogCoordinates.SmallerThan(&instance.ExecBinlogCoordinates) {
		return instance, log.Errorf("SyncReplicaRelayLogs: %+v is more up to date than %+v. Will not apply relaylogs from %+v", instance.Key, fromInstance.Key, fromInstance.Key)
	}

	log.Debugf("SyncReplicaRelayLogs: correlating coordinates of %+v on %+v", instance.Key, fromInstance.Key)
	_, _, nextCoordinates, found, err := inst.CorrelateRelaylogCoordinates(instance, nil, fromInstance)
	if err != nil {
		return instance, err
	}
	if !found {
		return instance, err
	}
	log.Debugf("SyncReplicaRelayLogs: correlated next-coordinates on %+v are %+v", fromInstance.Key, *nextCoordinates)

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

	switch syncRelaylogsChangeMasterMethod {
	case SyncRelaylogsChangeMasterToSharedMaster:
		{
			instance, err = inst.ChangeMasterTo(&instance.Key, &fromInstance.MasterKey, &fromInstance.ExecBinlogCoordinates, false, inst.GTIDHintNeutral)
			if err != nil {
				return instance, err
			}
		}
	case SyncRelaylogsChangeMasterToSourceReplica:
		{
			instance, err = inst.ChangeMasterTo(&instance.Key, &fromInstance.Key, &fromInstance.SelfBinlogCoordinates, false, inst.GTIDHintNeutral)
			if err != nil {
				return instance, err
			}
		}
	}
	inst.AuditOperation("align-via-relaylogs-remote", &instance.Key, fmt.Sprintf("aligned %+v by relaylogs from %+v", instance.Key, fromInstance.Key))
	return instance, err
}

func SyncReplicasRelayLogs(masterKey *inst.InstanceKey, syncRelaylogsChangeMasterMethod SyncRelaylogsChangeMasterMethod, postponedFunctionsContainer *inst.PostponedFunctionsContainer) (
	synchedFromReplica *inst.Instance, syncedReplicas, failedReplicas, postponedReplicas [](*inst.Instance), err error,
) {
	var replicas [](*inst.Instance)
	if replicas, err = inst.GetSortedReplicas(masterKey, inst.StopReplicationNormal); err != nil {
		return synchedFromReplica, syncedReplicas, replicas, postponedReplicas, err
	}
	synchedFromReplica = replicas[0]
	if len(replicas) <= 1 {
		// Nothing to be done
		return synchedFromReplica, syncedReplicas, replicas, postponedReplicas, err
	}
	applyToReplicas := replicas[1:]

	log.Debugf("Testing SSH on %+v", synchedFromReplica.Key)
	if err := TestRemoteCommandOnInstance(&synchedFromReplica.Key); err != nil {
		return synchedFromReplica, syncedReplicas, replicas, postponedReplicas, err
	}

	barrier := make(chan *inst.InstanceKey, len(applyToReplicas))
	allErrors := make(chan error, len(applyToReplicas))
	synchedReplicasChan := make(chan *inst.Instance, len(applyToReplicas))
	failedReplicasChan := make(chan *inst.Instance, len(applyToReplicas))

	applyToReplicaFunc := func(applyToReplica *inst.Instance) error {
		defer func() { barrier <- &applyToReplica.Key }()

		if _, err := SyncReplicaRelayLogs(applyToReplica, synchedFromReplica, syncRelaylogsChangeMasterMethod); err == nil {
			synchedReplicasChan <- applyToReplica
		} else {
			err = fmt.Errorf("%+v: %+v", applyToReplica.Key, err.Error())
			failedReplicasChan <- applyToReplica
			allErrors <- err
		}
		return err
	}

	log.Debugf("Applying relay logs on %+v replicas", len(applyToReplicas))
	countImmediateApply := 0
	for _, applyToReplica := range applyToReplicas {
		applyToReplica := applyToReplica

		if postponedFunctionsContainer != nil &&
			config.Config.PostponeReplicaRecoveryOnLagMinutes > 0 &&
			applyToReplica.SQLDelay > config.Config.PostponeReplicaRecoveryOnLagMinutes*60 {
			postponedReplicas = append(postponedReplicas, applyToReplica)
			(*postponedFunctionsContainer).AddPostponedFunction(func() error { return applyToReplicaFunc(applyToReplica) })
		} else {
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
	inst.AuditOperation("sync-replicas-relaylogs", masterKey, fmt.Sprintf("aligned %+v replicas by relaylogs from %+v, got %+v errors", len(applyToReplicas), synchedFromReplica.Key, countErrors))
	return synchedFromReplica, syncedReplicas, failedReplicas, postponedReplicas, err
}
