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

package orchestrator

import (
	"fmt"
	"github.com/outbrain/golib/log"
	"github.com/outbrain/orchestrator/config"
	"github.com/outbrain/orchestrator/inst"
	"github.com/outbrain/orchestrator/os"
	"strings"
)

func RecoverDeadMaster(oldMaster *inst.InstanceKey, newMaster *inst.InstanceKey) error {

	var preProcessesFailures []error = []error{}
	barrier := make(chan error)
	for _, command := range config.Config.PreFailoverProcesses {
		command := command
		command = strings.Replace(command, "{failedHost}", oldMaster.Hostname, -1)
		command = strings.Replace(command, "{failedPort}", fmt.Sprintf("%d", oldMaster.Port), -1)
		command = strings.Replace(command, "{successorHost}", newMaster.Hostname, -1)
		command = strings.Replace(command, "{successorPort}", fmt.Sprintf("%d", newMaster.Port), -1)
		var cmdErr error = nil
		go func() {
			defer func() { barrier <- cmdErr }()
			cmdErr = os.CommandRun(command)
		}()
	}
	for _, _ = range config.Config.PreFailoverProcesses {
		if cmdErr := <-barrier; cmdErr != nil {
			preProcessesFailures = append(preProcessesFailures, cmdErr)
		}
	}
	if len(preProcessesFailures) > 0 {
		return preProcessesFailures[0]
	}

	return nil
}

func RecoverDeadIntermediateMaster(failedInstanceKey *inst.InstanceKey) (*inst.Instance, error) {
	log.Debugf("RecoverDeadIntermediateMaster: will recover %+v", *failedInstanceKey)
	candidateSibling, err := inst.GetCandidateSiblingOfIntermediateMaster(failedInstanceKey)
	if err == nil {
		log.Debugf("- RecoverDeadIntermediateMaster: will attempt a candidate intermediate master: %+v", candidateSibling.Key)
		// We have a candidate
		if _, belowInstance, err := inst.MultiMatchSlaves(failedInstanceKey, &candidateSibling.Key); err == nil {
			log.Debugf("- RecoverDeadIntermediateMaster: move to candidate intermediate master (%+v) went well", candidateSibling.Key)
			inst.AuditOperation("recover-dead-intermediate-master", failedInstanceKey, fmt.Sprintf("candidate sibling: %+v", candidateSibling.Key))
			return belowInstance, nil
		} else {
			log.Debugf("- RecoverDeadIntermediateMaster: move to candidate intermediate master (%+v) did not complete: %+v", candidateSibling.Key, err)
		}
	}
	// Either no candidate or only partial match of slaves. Match-up as plan B
	log.Debugf("- RecoverDeadIntermediateMaster: will next attempt a match up from %+v", *failedInstanceKey)
	_, successorInstance, err := inst.MatchUpSlaves(failedInstanceKey)
	log.Debugf("- RecoverDeadIntermediateMaster: matched up to %+v", successorInstance.Key)
	inst.AuditOperation("recover-dead-intermediate-master", failedInstanceKey, fmt.Sprintf("master: %+v", successorInstance.Key))

	return successorInstance, err
}
