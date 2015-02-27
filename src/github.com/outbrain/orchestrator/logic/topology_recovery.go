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
	"regexp"
	"sort"
	"strings"
)

type InstancesByCountSlaves [](*inst.Instance)

func (this InstancesByCountSlaves) Len() int      { return len(this) }
func (this InstancesByCountSlaves) Swap(i, j int) { this[i], this[j] = this[j], this[i] }
func (this InstancesByCountSlaves) Less(i, j int) bool {
	if len(this[i].SlaveHosts) == len(this[j].SlaveHosts) {
		// Secondary sorting: prefer more advanced slaves
		return !this[i].ExecBinlogCoordinates.SmallerThan(&this[j].ExecBinlogCoordinates)
	}
	return len(this[i].SlaveHosts) < len(this[j].SlaveHosts)
}

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

func isValidAsCandidateSiblingOfIntermediateMaster(sibling *inst.Instance) bool {
	if !sibling.LogBinEnabled {
		return false
	}
	if !sibling.LogSlaveUpdatesEnabled {
		return false
	}
	if !sibling.SlaveRunning() {
		return false
	}
	if !sibling.IsLastCheckValid {
		return false
	}
	return true
}

// GetCandidateSlave chooses the best slave to promote given a (possibly dead) master
func GetCandidateSiblingOfIntermediateMaster(intermediateMasterKey *inst.InstanceKey) (*inst.Instance, error) {
	intermediateMasterInstance, _, err := inst.ReadInstance(intermediateMasterKey)
	if err != nil {
		return nil, err
	}

	siblings, err := inst.ReadSlaveInstances(&intermediateMasterInstance.MasterKey)
	if err != nil {
		return nil, err
	}
	if len(siblings) <= 1 {
		return nil, log.Errorf("No siblings found for %+v", *intermediateMasterKey)
	}

	sort.Sort(sort.Reverse(InstancesByCountSlaves(siblings)))

	for _, sibling := range siblings {
		sibling := sibling
		if isValidAsCandidateSiblingOfIntermediateMaster(sibling) {
			if sibling.DataCenter == intermediateMasterInstance.DataCenter && sibling.PhysicalEnvironment == intermediateMasterInstance.PhysicalEnvironment {
				if !sibling.Key.Equals(intermediateMasterKey) && intermediateMasterInstance.ExecBinlogCoordinates.SmallerThan(&sibling.ExecBinlogCoordinates) {
					// this is *assumed* to be a good choice.
					// We don't know for sure:
					// - the dead intermediate master's position may have been more advanced then last recorded
					// - and the candidate's position may have been stalled in the past seconds
					// But it's an attempt...
					return sibling, nil
				}
			}
		}
	}
	return nil, log.Errorf("Cannot find candidate sibling of %+v", *intermediateMasterKey)
}

func RecoverDeadIntermediateMaster(failedInstanceKey *inst.InstanceKey) (*inst.Instance, error) {
	log.Debugf("RecoverDeadIntermediateMaster: will recover %+v", *failedInstanceKey)
	candidateSibling, err := GetCandidateSiblingOfIntermediateMaster(failedInstanceKey)
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
	// Either no candidate or only partial match of slaves. Regroup as plan B
	inst.RegroupSlaves(failedInstanceKey, nil)
	// Either no candidate or only partial match of slaves, and/or regroup: continue to match-up, plan C
	log.Debugf("- RecoverDeadIntermediateMaster: will next attempt a match up from %+v", *failedInstanceKey)
	_, successorInstance, err := inst.MatchUpSlaves(failedInstanceKey)
	log.Debugf("- RecoverDeadIntermediateMaster: matched up to %+v", successorInstance.Key)
	inst.AuditOperation("recover-dead-intermediate-master", failedInstanceKey, fmt.Sprintf("master: %+v", successorInstance.Key))

	return successorInstance, err
}

// checkAndRecoverDeadIntermediateMaster checks a given analysis, decides whether to take action, and possibly takes action
// Returns true when action was taken.
func checkAndRecoverDeadIntermediateMaster(analysisEntry inst.ReplicationAnalysis, skipFilters bool) (bool, error) {
	for _, filter := range config.Config.RecoverIntermediateMasterClusterFilters {
		if matched, _ := regexp.MatchString(filter, analysisEntry.ClusterName); matched || skipFilters {
			log.Debugf("Will handle DeadIntermediateMaster event on %+v", analysisEntry.ClusterName)
			_, err := RecoverDeadIntermediateMaster(&analysisEntry.AnalyzedInstanceKey)
			return true, err
		}
	}
	return false, nil
}

func CheckAndRecover(specificInstance *inst.InstanceKey, skipFilters bool) (bool, error) {
	actionTaken := false
	replicationAnalysis, err := inst.GetReplicationAnalysis()
	if err != nil {
		return false, log.Errore(err)
	}
	for _, analysisEntry := range replicationAnalysis {
		if specificInstance != nil {
			// We are looking for a specific instance; if this is not the one, skip!
			if !specificInstance.Equals(&analysisEntry.AnalyzedInstanceKey) {
				continue
			}
		}
		if analysisEntry.Analysis == inst.DeadIntermediateMaster {
			actionTaken, err = checkAndRecoverDeadIntermediateMaster(analysisEntry, skipFilters)
		}
	}
	return actionTaken, err
}
