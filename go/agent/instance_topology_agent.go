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

package agent

import (
	"encoding/json"
	"fmt"

	"github.com/github/orchestrator/go/inst"
	"github.com/openark/golib/log"
)

func SyncReplicaRelayLogs(instance, otherInstance *inst.Instance) (*inst.Instance, error) {
	var err error
	var found bool
	var nextCoordinates *inst.BinlogCoordinates
	var content string
	onResponse := func(contentBytes []byte) {
		json.Unmarshal(contentBytes, &content)
	}
	log.Debugf("SyncReplicaRelayLogs: stopping replication")

	if !instance.ReplicationThreadsStopped() {
		return instance, log.Errorf("SyncReplicaRelayLogs: replication on %+v must not run", instance.Key)
	}
	if !otherInstance.ReplicationThreadsStopped() {
		return instance, log.Errorf("SyncReplicaRelayLogs: replication on %+v must not run", otherInstance.Key)
	}

	agent, err := ReadAgent(instance.Key.Hostname)
	if err != nil {
		return instance, log.Errore(err)
	}

	otherAgent, err := ReadAgent(otherInstance.Key.Hostname)
	if err != nil {
		return otherInstance, log.Errore(err)
	}

	log.Debugf("SyncReplicaRelayLogs: correlating coordinates of %+v on %+v", instance.Key, otherInstance.Key)
	_, _, nextCoordinates, found, err = inst.CorrelateRelaylogCoordinates(instance, nil, otherInstance)
	if err != nil {
		goto Cleanup
	}
	if !found {
		goto Cleanup
	}
	log.Debugf("SyncReplicaRelayLogs: correlated next-coordinates are %+v", *nextCoordinates)

	InitHttpClient()

	if err := otherAgent.relaylogContentsTail(nextCoordinates, &onResponse); err != nil {
		goto Cleanup
	}
	log.Debugf("SyncReplicaRelayLogs: got content (%d bytes)", len(content))

	if err := agent.applyRelaylogContents(content); err != nil {
		goto Cleanup
	}
	log.Debugf("SyncReplicaRelayLogs: applied content (%d bytes)", len(content))

	instance, err = inst.ChangeMasterTo(&instance.Key, &otherInstance.MasterKey, &otherInstance.ExecBinlogCoordinates, false, inst.GTIDHintNeutral)
	if err != nil {
		goto Cleanup
	}

Cleanup:
	if err != nil {
		return instance, log.Errore(err)
	}
	// and we're done (pending deferred functions)
	inst.AuditOperation("align-via-relaylogs", &instance.Key, fmt.Sprintf("aligned %+v by relaylogs from %+v", instance.Key, otherInstance.Key))

	return instance, err
}
