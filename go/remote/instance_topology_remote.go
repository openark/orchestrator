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

// AlignViaRelaylogCorrelation will align siblings by applying relaylogs from one to the other, via remote SSH
func AlignViaRelaylogCorrelation(instance, otherInstance *inst.Instance) (*inst.Instance, error) {
	if config.Config.RemoteSSHCommand == "" {
		return instance, fmt.Errorf("RemoteSSHCommand not configured")
	}
	log.Debugf("Testing SSH on %+v", instance.Key)
	if err := TestRemoteCommandOnInstance(&instance.Key); err != nil {
		return instance, err
	}
	log.Debugf("Testing SSH on %+v", otherInstance.Key)
	if err := TestRemoteCommandOnInstance(&otherInstance.Key); err != nil {
		return instance, err
	}

	log.Debugf("AlignViaRelaylogCorrelation: stopping replication")
	if instance.ReplicaRunning() {
		return instance, log.Errorf("AlignViaRelaylogCorrelation: replication on %+v must not run", instance.Key)
	}
	if otherInstance.ReplicaRunning() {
		return instance, log.Errorf("AlignViaRelaylogCorrelation: replication on %+v must not run", otherInstance.Key)
	}

	log.Debugf("AlignViaRelaylogCorrelation: correlating coordinates of %+v on %+v", instance.Key, otherInstance.Key)
	_, _, nextCoordinates, found, err := inst.CorrelateRelaylogCoordinates(instance, nil, otherInstance)
	if err != nil {
		return instance, err
	}
	if !found {
		return instance, err
	}
	log.Debugf("AlignViaRelaylogCorrelation: correlated next-coordinates are %+v", *nextCoordinates)

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
		script = strings.Replace(script, "$MAGIC_START_POSITION", fmt.Sprintf("%d", nextCoordinates.LogPos), -1)
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
		command = strings.Replace(command, "{hostname}", otherInstance.Key.Hostname, -1)
		command = fmt.Sprintf("cat %s | %s '%s' > %s", getRelayLogContentsScriptFile.Name(), command, sudoCommand, localRelayLogContentsFile.Name())
		if err := orcos.CommandRun(command, orcos.EmptyEnv); err != nil {
			return instance, err
		}
	}
	log.Debugf("Have executed on %s, output file is %s", otherInstance.Key.Hostname, localRelayLogContentsFile.Name())
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

	instance, err = inst.ChangeMasterTo(&instance.Key, &otherInstance.MasterKey, &otherInstance.ExecBinlogCoordinates, false, inst.GTIDHintNeutral)
	if err != nil {
		return instance, err
	}
	inst.AuditOperation("align-via-relaylogs-remote", &instance.Key, fmt.Sprintf("aligned %+v by relaylogs from %+v", instance.Key, otherInstance.Key))
	return instance, err
}
