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

func RecoverDeadMaster(failedInstanceKey *inst.InstanceKey) (*inst.Instance, error) {

	inst.AuditOperation("recover-dead-master", failedInstanceKey, "problem found; will recover")

	log.Debugf("RecoverDeadMaster: will recover %+v", *failedInstanceKey)
	_, _, _, candidateSlave, err := inst.RegroupSlaves(failedInstanceKey, nil)
	log.Debugf("- RecoverDeadIntermediateMaster: candidate slave is %+v", candidateSlave.Key)
	inst.AuditOperation("recover-dead-master", failedInstanceKey, fmt.Sprintf("master: %+v", candidateSlave.Key))

	return candidateSlave, err
}

// checkAndRecoverDeadMaster checks a given analysis, decides whether to take action, and possibly takes action
// Returns true when action was taken.
func checkAndRecoverDeadMaster(analysisEntry inst.ReplicationAnalysis, skipFilters bool) (bool, error) {
	filters := config.Config.RecoverMasterClusterFilters
	if skipFilters {
		// The following filter catches-all
		filters = append(filters, ".")
	}
	for _, filter := range filters {
		if matched, _ := regexp.MatchString(filter, analysisEntry.ClusterName); matched && filter != "" {
			log.Debugf("Will handle DeadMaster event on %+v", analysisEntry.ClusterName)
			_, err := RecoverDeadMaster(&analysisEntry.AnalyzedInstanceKey)
			return true, err
		}
	}
	return false, nil
}

func isGeneralyValidAsCandidateSiblingOfIntermediateMaster(sibling *inst.Instance) bool {
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

func isValidAsCandidateSiblingOfIntermediateMaster(intermediateMasterInstance *inst.Instance, sibling *inst.Instance) bool {
	if !isGeneralyValidAsCandidateSiblingOfIntermediateMaster(sibling) {
		return false
	}
	if sibling.DataCenter != intermediateMasterInstance.DataCenter {
		return false
	}
	if sibling.PhysicalEnvironment != intermediateMasterInstance.PhysicalEnvironment {
		return false
	}
	if sibling.HasReplicationFilters != intermediateMasterInstance.HasReplicationFilters {
		return false
	}
	if sibling.Key.Equals(&intermediateMasterInstance.Key) {
		return false
	}
	if sibling.ExecBinlogCoordinates.SmallerThan(&intermediateMasterInstance.ExecBinlogCoordinates) {
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
		if isValidAsCandidateSiblingOfIntermediateMaster(intermediateMasterInstance, sibling) {
			// this is *assumed* to be a good choice.
			// We don't know for sure:
			// - the dead intermediate master's position may have been more advanced then last recorded
			// - and the candidate's position may have been stalled in the past seconds
			// But it's an attempt...
			return sibling, nil
		}
	}
	return nil, log.Errorf("Cannot find candidate sibling of %+v", *intermediateMasterKey)
}

func RecoverDeadIntermediateMaster(failedInstanceKey *inst.InstanceKey) (*inst.Instance, error) {
	if ok, err := AttemptRecoveryRegistration(failedInstanceKey); !ok {
		log.Debugf("Will not RecoverDeadIntermediateMaster on %+v", *failedInstanceKey)
		return nil, err
	}

	inst.AuditOperation("recover-dead-intermediate-master", failedInstanceKey, "problem found; will recover")
	log.Debugf("RecoverDeadIntermediateMaster: will recover %+v", *failedInstanceKey)
	candidateSibling, err := GetCandidateSiblingOfIntermediateMaster(failedInstanceKey)
	if err == nil {
		log.Debugf("- RecoverDeadIntermediateMaster: will attempt a candidate intermediate master: %+v", candidateSibling.Key)
		// We have a candidate
		if matchedSlaves, belowInstance, err := inst.MultiMatchSlaves(failedInstanceKey, &candidateSibling.Key); err == nil {
			ResolveRecovery(failedInstanceKey, &candidateSibling.Key)
			log.Debugf("- RecoverDeadIntermediateMaster: move to candidate intermediate master (%+v) went well", candidateSibling.Key)
			inst.AuditOperation("recover-dead-intermediate-master", failedInstanceKey, fmt.Sprintf("Done. Matched %d slaves under candidate sibling: %+v", len(matchedSlaves), candidateSibling.Key))
			return belowInstance, nil
		} else {
			inst.AuditOperation("recover-dead-intermediate-master", failedInstanceKey, fmt.Sprintf("Matched %d slaves under candidate sibling: %+v", len(matchedSlaves), candidateSibling.Key))
			log.Debugf("- RecoverDeadIntermediateMaster: move to candidate intermediate master (%+v) did not complete: %+v", candidateSibling.Key, err)
		}
	}
	// Either no candidate or only partial match of slaves. Regroup as plan B
	inst.RegroupSlaves(failedInstanceKey, nil)
	// And anyway, match up all that's left continue to match-up, plan C
	log.Debugf("- RecoverDeadIntermediateMaster: will next attempt a match up from %+v", *failedInstanceKey)
	_, successorInstance, err := inst.MatchUpSlaves(failedInstanceKey)
	ResolveRecovery(failedInstanceKey, &successorInstance.Key)
	log.Debugf("- RecoverDeadIntermediateMaster: matched up to %+v", successorInstance.Key)
	inst.AuditOperation("recover-dead-intermediate-master", failedInstanceKey, fmt.Sprintf("Done. Matched slaves under: %+v", successorInstance.Key))

	return successorInstance, err
}

// checkAndRecoverDeadIntermediateMaster checks a given analysis, decides whether to take action, and possibly takes action
// Returns true when action was taken.
func checkAndRecoverDeadIntermediateMaster(analysisEntry inst.ReplicationAnalysis, skipFilters bool) (bool, error) {
	filters := config.Config.RecoverIntermediateMasterClusterFilters
	if skipFilters {
		// The following filter catches-all
		filters = append(filters, ".")
	}
	for _, filter := range filters {
		if matched, _ := regexp.MatchString(filter, analysisEntry.ClusterName); matched && filter != "" {
			log.Debugf("Will handle DeadIntermediateMaster event on %+v", analysisEntry.ClusterName)
			_, err := RecoverDeadIntermediateMaster(&analysisEntry.AnalyzedInstanceKey)
			return true, err
		}
	}
	return false, nil
}

// executeCheckAndRecoverFunction will choose the correct check & recovery function based on analysis.
// It executes the function synchronuously
func executeCheckAndRecoverFunction(analysisEntry inst.ReplicationAnalysis, skipFilters bool) (bool, error) {
	var checkAndRecoverFunction func(analysisEntry inst.ReplicationAnalysis, skipFilters bool) (bool, error) = nil
	if analysisEntry.Analysis == inst.DeadMaster {
		checkAndRecoverFunction = checkAndRecoverDeadMaster
	} else if analysisEntry.Analysis == inst.DeadIntermediateMaster || analysisEntry.Analysis == inst.DeadIntermediateMasterAndSomeSlaves {
		checkAndRecoverFunction = checkAndRecoverDeadIntermediateMaster
	}
	if checkAndRecoverFunction == nil {
		// Unhandled problem type
		return false, nil
	}
	// we have a recovery function; its execution still depends on filters if not disabled.

	// Execute general pre-processes
	{
		var preProcessesFailures []error = []error{}
		barrier := make(chan error)
		for _, command := range config.Config.PreFailoverProcesses {
			command := command

			command = strings.Replace(command, "{failureType}", string(analysisEntry.Analysis), -1)
			command = strings.Replace(command, "{failureDescription}", analysisEntry.Description, -1)
			command = strings.Replace(command, "{failedHost}", analysisEntry.AnalyzedInstanceKey.Hostname, -1)
			command = strings.Replace(command, "{failedPort}", fmt.Sprintf("%d", analysisEntry.AnalyzedInstanceKey.Port), -1)
			command = strings.Replace(command, "{failureCluster}", analysisEntry.ClusterName, -1)

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
			return false, preProcessesFailures[0]
		}
	}

	return checkAndRecoverFunction(analysisEntry, skipFilters)
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

		if specificInstance != nil && skipFilters {
			// force mode. Keep it synchronuous
			actionTaken, err = executeCheckAndRecoverFunction(analysisEntry, skipFilters)
		} else {
			go executeCheckAndRecoverFunction(analysisEntry, skipFilters)
		}
	}
	return actionTaken, err
}
