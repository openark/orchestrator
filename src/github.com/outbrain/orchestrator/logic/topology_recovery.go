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
	"github.com/pmylund/go-cache"
	"regexp"
	"sort"
	"strings"
	"time"
)

var emergencyReadTopologyInstanceMap = cache.New(time.Duration(config.Config.DiscoveryPollSeconds)*time.Second, time.Duration(config.Config.DiscoveryPollSeconds)*time.Second)

// InstancesByCountSlaves sorts instances by umber of slaves, descending
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

func RecoverDeadMaster(failedInstanceKey *inst.InstanceKey) (bool, *inst.Instance, error) {
	if ok, err := AttemptRecoveryRegistration(failedInstanceKey); !ok {
		log.Debugf("Will not RecoverDeadMaster on %+v", *failedInstanceKey)
		return false, nil, err
	}

	inst.AuditOperation("recover-dead-master", failedInstanceKey, "problem found; will recover")

	log.Debugf("RecoverDeadMaster: will recover %+v", *failedInstanceKey)
	_, _, _, candidateSlave, err := inst.RegroupSlaves(failedInstanceKey, nil)

	ResolveRecovery(failedInstanceKey, &candidateSlave.Key)

	log.Debugf("- RecoverDeadIntermediateMaster: candidate slave is %+v", candidateSlave.Key)
	inst.AuditOperation("recover-dead-master", failedInstanceKey, fmt.Sprintf("master: %+v", candidateSlave.Key))

	return true, candidateSlave, err
}

// checkAndRecoverDeadMaster checks a given analysis, decides whether to take action, and possibly takes action
// Returns true when action was taken.
func checkAndRecoverDeadMaster(analysisEntry inst.ReplicationAnalysis, skipFilters bool) (bool, *inst.Instance, error) {
	filters := config.Config.RecoverMasterClusterFilters
	if skipFilters {
		// The following filter catches-all
		filters = append(filters, ".")
	}
	for _, filter := range filters {
		if matched, _ := regexp.MatchString(filter, analysisEntry.ClusterName); matched && filter != "" {
			log.Debugf("Will handle DeadMaster event on %+v", analysisEntry.ClusterName)
			actionTaken, promotedSlave, err := RecoverDeadMaster(&analysisEntry.AnalyzedInstanceKey)

			if actionTaken && promotedSlave != nil {
				// Execute post master-failover processes
				for _, command := range config.Config.PostMasterFailoverProcesses {
					command := command

					command = strings.Replace(command, "{failureType}", string(analysisEntry.Analysis), -1)
					command = strings.Replace(command, "{failureDescription}", analysisEntry.Description, -1)
					command = strings.Replace(command, "{failedHost}", analysisEntry.AnalyzedInstanceKey.Hostname, -1)
					command = strings.Replace(command, "{failedPort}", fmt.Sprintf("%d", analysisEntry.AnalyzedInstanceKey.Port), -1)
					command = strings.Replace(command, "{successorHost}", promotedSlave.Key.Hostname, -1)
					command = strings.Replace(command, "{successorPort}", fmt.Sprintf("%d", promotedSlave.Key.Port), -1)
					command = strings.Replace(command, "{failureCluster}", analysisEntry.ClusterName, -1)

					if cmdErr := os.CommandRun(command); cmdErr == nil {
						log.Infof("Executed post-master-failover command %s", command)
					} else {
						log.Errorf("Failed to execute post-master-failover command %s", command)
					}
				}
			}

			return actionTaken, promotedSlave, err
		}
	}
	return false, nil, nil
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
	if sibling.Key.Equals(&intermediateMasterInstance.Key) {
		// same instance
		return false
	}
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
	if sibling.IsMaxScale() || intermediateMasterInstance.IsMaxScale() {
		// With MaxScale the failover is different; we don't need this "move to uncle" logic.
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

func RecoverDeadIntermediateMaster(failedInstanceKey *inst.InstanceKey) (bool, *inst.Instance, error) {
	if ok, err := AttemptRecoveryRegistration(failedInstanceKey); !ok {
		log.Debugf("Will not RecoverDeadIntermediateMaster on %+v", *failedInstanceKey)
		return false, nil, err
	}

	inst.AuditOperation("recover-dead-intermediate-master", failedInstanceKey, "problem found; will recover")
	log.Debugf("RecoverDeadIntermediateMaster: will recover %+v", *failedInstanceKey)
	candidateSibling, err := GetCandidateSiblingOfIntermediateMaster(failedInstanceKey)
	if err == nil {
		log.Debugf("- RecoverDeadIntermediateMaster: will attempt a candidate intermediate master: %+v", candidateSibling.Key)
		// We have a candidate
		if matchedSlaves, belowInstance, err, errs := inst.MultiMatchSlaves(failedInstanceKey, &candidateSibling.Key, ""); err == nil {
			ResolveRecovery(failedInstanceKey, &candidateSibling.Key)
			log.Debugf("- RecoverDeadIntermediateMaster: move to candidate intermediate master (%+v) went with %d errors", candidateSibling.Key, len(errs))
			inst.AuditOperation("recover-dead-intermediate-master", failedInstanceKey, fmt.Sprintf("Done. Matched %d slaves under candidate sibling: %+v; %d errors: %+v", len(matchedSlaves), candidateSibling.Key, len(errs), errs))
			return true, belowInstance, nil
		} else {
			log.Debugf("- RecoverDeadIntermediateMaster: move to candidate intermediate master (%+v) did not complete: %+v", candidateSibling.Key, err)
			inst.AuditOperation("recover-dead-intermediate-master", failedInstanceKey, fmt.Sprintf("Matched %d slaves under candidate sibling: %+v; %d errors: %+v", len(matchedSlaves), candidateSibling.Key, len(errs), errs))
		}
	}
	// Either no candidate or only partial match of slaves. Regroup as plan B
	inst.RegroupSlaves(failedInstanceKey, nil)
	// And anyway, match up all that's left continue to match-up, plan C
	log.Debugf("- RecoverDeadIntermediateMaster: will next attempt a match up from %+v", *failedInstanceKey)
	_, successorInstance, err, errs := inst.MatchUpSlaves(failedInstanceKey, "")
	ResolveRecovery(failedInstanceKey, &successorInstance.Key)
	log.Debugf("- RecoverDeadIntermediateMaster: matched up to %+v", successorInstance.Key)
	inst.AuditOperation("recover-dead-intermediate-master", failedInstanceKey, fmt.Sprintf("Done. Matched slaves under: %+v %d errors: %+v", successorInstance.Key, len(errs), errs))

	return true, successorInstance, err
}

// checkAndRecoverDeadIntermediateMaster checks a given analysis, decides whether to take action, and possibly takes action
// Returns true when action was taken.
func checkAndRecoverDeadIntermediateMaster(analysisEntry inst.ReplicationAnalysis, skipFilters bool) (bool, *inst.Instance, error) {
	filters := config.Config.RecoverIntermediateMasterClusterFilters
	if skipFilters {
		// The following filter catches-all
		filters = append(filters, ".")
	}
	for _, filter := range filters {
		if matched, _ := regexp.MatchString(filter, analysisEntry.ClusterName); matched && filter != "" {
			log.Debugf("Will handle DeadIntermediateMaster event on %+v", analysisEntry.ClusterName)
			actionTaken, promotedSlave, err := RecoverDeadIntermediateMaster(&analysisEntry.AnalyzedInstanceKey)
			return actionTaken, promotedSlave, err
		}
	}
	return false, nil, nil
}

// Force a re-read of a topology instance; this is done because we need to substantiate a suspicion that we may have a failover
// scenario. we want to speed up rading the complete picture.
func emergentlyReadTopologyInstance(instanceKey *inst.InstanceKey, analysisCode inst.AnalysisCode) {
	if err := emergencyReadTopologyInstanceMap.Add(instanceKey.DisplayString(), true, 0); err == nil {
		emergencyReadTopologyInstanceMap.Set(instanceKey.DisplayString(), true, 0)
		go inst.ExecuteOnTopology(func() {
			inst.ReadTopologyInstance(instanceKey)
			inst.AuditOperation("emergently-read-topology-instance", instanceKey, string(analysisCode))
		})
	}
}

// Force reading of slaves of given instance. This is because we suspect the instance is dead, and want to speed up
// detection of replication failure from its slaves.
func emergentlyReadTopologyInstanceSlaves(instanceKey *inst.InstanceKey, analysisCode inst.AnalysisCode) {
	slaves, err := inst.ReadSlaveInstances(instanceKey)
	if err != nil {
		return
	}
	for _, slave := range slaves {
		go emergentlyReadTopologyInstance(&slave.Key, analysisCode)
	}
}

// executeCheckAndRecoverFunction will choose the correct check & recovery function based on analysis.
// It executes the function synchronuously
func executeCheckAndRecoverFunction(analysisEntry inst.ReplicationAnalysis, skipFilters bool) (bool, *inst.Instance, error) {
	var checkAndRecoverFunction func(analysisEntry inst.ReplicationAnalysis, skipFilters bool) (bool, *inst.Instance, error) = nil

	switch analysisEntry.Analysis {
	case inst.DeadMaster:
		checkAndRecoverFunction = checkAndRecoverDeadMaster
	case inst.DeadMasterAndSomeSlaves:
		checkAndRecoverFunction = checkAndRecoverDeadMaster
	case inst.DeadIntermediateMaster:
		checkAndRecoverFunction = checkAndRecoverDeadIntermediateMaster
	case inst.DeadIntermediateMasterAndSomeSlaves:
		checkAndRecoverFunction = checkAndRecoverDeadIntermediateMaster
	case inst.DeadCoMaster:
		checkAndRecoverFunction = checkAndRecoverDeadIntermediateMaster
	case inst.UnreachableMaster:
		go emergentlyReadTopologyInstanceSlaves(&analysisEntry.AnalyzedInstanceKey, analysisEntry.Analysis)
	case inst.AllMasterSlavesNotReplicating:
		go emergentlyReadTopologyInstance(&analysisEntry.AnalyzedInstanceKey, analysisEntry.Analysis)
	case inst.FirstTierSlaveFailingToConnectToMaster:
		go emergentlyReadTopologyInstance(&analysisEntry.AnalyzedInstanceMasterKey, analysisEntry.Analysis)
	}

	if checkAndRecoverFunction == nil {
		// Unhandled problem type
		return false, nil, nil
	}
	// we have a recovery function; its execution still depends on filters if not disabled.

	// Execute general pre-processes
	{
		var processesFailures []error = []error{}
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
				processesFailures = append(processesFailures, cmdErr)
			}
		}
		if len(processesFailures) > 0 {
			return false, nil, processesFailures[0]
		}
	}

	return checkAndRecoverFunction(analysisEntry, skipFilters)
}

// CheckAndRecover is the main entry point for the recovery mechanism
func CheckAndRecover(specificInstance *inst.InstanceKey, skipFilters bool) (actionTaken bool, instance *inst.Instance, err error) {
	replicationAnalysis, err := inst.GetReplicationAnalysis()
	if err != nil {
		return false, nil, log.Errore(err)
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
			actionTaken, instance, err = executeCheckAndRecoverFunction(analysisEntry, skipFilters)
		} else {
			go executeCheckAndRecoverFunction(analysisEntry, skipFilters)
		}
	}
	return actionTaken, instance, err
}
