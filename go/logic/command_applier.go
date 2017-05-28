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

package logic

import (
	"encoding/json"

	"github.com/github/orchestrator/go/inst"
	"github.com/github/orchestrator/go/process"

	"github.com/openark/golib/log"
)

// AsyncRequest represents an entry in the async_request table
type CommandApplier struct {
}

func NewCommandApplier() *CommandApplier {
	applier := &CommandApplier{}
	return applier
}

func (applier *CommandApplier) ApplyCommand(op string, value []byte) interface{} {
	log.Debugf("orchestrator/raft: applying command: %s", op)
	switch op {
	case "register-node":
		return applier.registerNode(value)
	case "discover":
		return applier.discover(value)
	case "forget":
		return applier.forget(value)
	}
	return log.Errorf("Unknown command op: %s", op)
}

func (applier *CommandApplier) registerNode(value []byte) interface{} {
	nodeHealth := process.NodeHealth{}
	if err := json.Unmarshal(value, &nodeHealth); err != nil {
		return log.Errore(err)
	}
	log.Debugf("====== nodeHealth: %+v", nodeHealth)
	if _, err := process.WriteRegisterNode(&nodeHealth); err != nil {
		return log.Errore(err)
	}

	return &nodeHealth
}

func (applier *CommandApplier) discover(value []byte) interface{} {
	instanceKey := inst.InstanceKey{}
	if err := json.Unmarshal(value, &instanceKey); err != nil {
		return log.Errore(err)
	}
	if instance, err := inst.ReadTopologyInstance(&instanceKey); err != nil {
		return log.Errore(err)
	} else {
		return instance
	}
}

func (applier *CommandApplier) forget(value []byte) interface{} {
	instanceKey := inst.InstanceKey{}
	if err := json.Unmarshal(value, &instanceKey); err != nil {
		return log.Errore(err)
	}
	err := inst.ForgetInstance(&instanceKey)
	return err
}
