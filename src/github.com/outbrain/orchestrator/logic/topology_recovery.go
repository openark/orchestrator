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
	candidateSibling, err := inst.GetCandidateSiblingOfIntermediateMaster(failedInstanceKey)
	if err == nil {
		// We have a candidate
		if _, belowInstance, err := inst.MultiMatchSlaves(failedInstanceKey, &candidateSibling.Key); err == nil {
			return belowInstance, nil
		}
	}
	// Either no candidate or only partial match of slaves. Match-up as plan B
	_, successorInstance, err := inst.MatchUpSlaves(failedInstanceKey)

	return successorInstance, err
}
