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
	"fmt"
	"github.com/outbrain/golib/log"
	"github.com/outbrain/orchestrator/go/config"
	"github.com/outbrain/orchestrator/go/inst"
	"github.com/outbrain/orchestrator/go/os"
	"github.com/pmylund/go-cache"
	"sort"
	"strings"
	"time"
)

// TopologyRecovery represents an entry in the topology_recovery table
type TopologyRecovery struct {
	Id                     int64
	AnalysisEntry          inst.ReplicationAnalysis
	SuccessorKey           inst.InstanceKey
	IsActive               bool
	RecoveryStartTimestamp string
	RecoveryEndTimestamp   string
	ProcessingNodeHostname string
	ProcessingNodeToken    string
}

type MasterRecoveryType string

const (
	MasterRecoveryGTID         MasterRecoveryType = "MasterRecoveryGTID"
	MasterRecoveryPseudoGTID                      = "MasterRecoveryPseudoGTID"
	MasterRecoveryBinlogServer                    = "MasterRecoveryBinlogServer"
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

// replaceCommandPlaceholders replaxces agreed-upon placeholders with analysis data
func replaceCommandPlaceholders(command string, analysisEntry inst.ReplicationAnalysis, successorInstance *inst.Instance) string {
	command = strings.Replace(command, "{failureType}", string(analysisEntry.Analysis), -1)
	command = strings.Replace(command, "{failureDescription}", analysisEntry.Description, -1)
	command = strings.Replace(command, "{failedHost}", analysisEntry.AnalyzedInstanceKey.Hostname, -1)
	command = strings.Replace(command, "{failedPort}", fmt.Sprintf("%d", analysisEntry.AnalyzedInstanceKey.Port), -1)
	command = strings.Replace(command, "{failureCluster}", analysisEntry.ClusterDetails.ClusterName, -1)
	command = strings.Replace(command, "{failureClusterAlias}", analysisEntry.ClusterDetails.ClusterAlias, -1)
	command = strings.Replace(command, "{countSlaves}", fmt.Sprintf("%d", analysisEntry.CountSlaves), -1)
	command = strings.Replace(command, "{isDowntimed}", fmt.Sprint(analysisEntry.IsDowntimed), -1)
	command = strings.Replace(command, "{autoMasterRecovery}", fmt.Sprint(analysisEntry.ClusterDetails.HasAutomatedMasterRecovery), -1)
	command = strings.Replace(command, "{autoIntermediateMasterRecovery}", fmt.Sprint(analysisEntry.ClusterDetails.HasAutomatedIntermediateMasterRecovery), -1)
	command = strings.Replace(command, "{orchestratorHost}", ThisHostname, -1)

	if successorInstance != nil {
		command = strings.Replace(command, "{successorHost}", successorInstance.Key.Hostname, -1)
		command = strings.Replace(command, "{successorPort}", fmt.Sprintf("%d", successorInstance.Key.Port), -1)
	}

	command = strings.Replace(command, "{slaveHosts}", analysisEntry.GetSlaveHostsAsString(), -1)

	return command
}

// executeProcesses executes a list of processes
func executeProcesses(processes []string, description string, analysisEntry inst.ReplicationAnalysis, successorInstance *inst.Instance, failOnError bool) error {
	var err error
	for _, command := range processes {
		command := replaceCommandPlaceholders(command, analysisEntry, successorInstance)

		if cmdErr := os.CommandRun(command); cmdErr == nil {
			log.Infof("Executed %s command: %s", description, command)
		} else {
			if err == nil {
				// Note first error
				err = cmdErr
			}
			log.Errorf("Failed to execute %s command: %s", description, command)
			if failOnError {
				return err
			}
		}
	}
	return err
}

func RecoverDeadMaster(analysisEntry inst.ReplicationAnalysis, skipProcesses bool) (success bool, promotedSlave *inst.Instance, err error) {
	failedInstanceKey := &analysisEntry.AnalyzedInstanceKey
	if ok, err := AttemptRecoveryRegistration(&analysisEntry); !ok {
		log.Debugf("topology_recovery: found an active or recent recovery on %+v. Will not issue another RecoverDeadMaster.", *failedInstanceKey)
		return false, nil, err
	}

	inst.AuditOperation("recover-dead-master", failedInstanceKey, "problem found; will recover")
	if !skipProcesses {
		if err := executeProcesses(config.Config.PreFailoverProcesses, "PreFailoverProcesses", analysisEntry, nil, true); err != nil {
			return false, nil, err
		}
	}

	log.Debugf("topology_recovery: RecoverDeadMaster: will recover %+v", *failedInstanceKey)

	var masterRecoveryType MasterRecoveryType = MasterRecoveryPseudoGTID
	if (analysisEntry.OracleGTIDImmediateTopology || analysisEntry.MariaDBGTIDImmediateTopology) && !analysisEntry.PseudoGTIDImmediateTopology {
		masterRecoveryType = MasterRecoveryGTID
	} else if analysisEntry.BinlogServerImmediateTopology {
		masterRecoveryType = MasterRecoveryBinlogServer
	}
	log.Debugf("topology_recovery: RecoverDeadMaster: masterRecoveryType=%+v", masterRecoveryType)
	switch masterRecoveryType {
	case MasterRecoveryGTID:
		{
			_, _, promotedSlave, err = inst.RegroupSlavesGTID(failedInstanceKey, true, nil)
		}
	case MasterRecoveryPseudoGTID:
		{
			_, _, _, promotedSlave, err = inst.RegroupSlavesIncludingSubSlavesOfBinlogServers(failedInstanceKey, true, nil)
		}
	}

	ResolveRecovery(failedInstanceKey, &promotedSlave.Key)

	log.Debugf("topology_recovery: - RecoverDeadMaster: candidate slave is %+v", promotedSlave.Key)
	inst.AuditOperation("recover-dead-master", failedInstanceKey, fmt.Sprintf("master: %+v", promotedSlave.Key))

	return true, promotedSlave, err
}

// replacePromotedSlaveWithCandidate is called after an intermediate master has died and been replaced by some promotedSlave.
// But, is there an even better slave to promote?
// if candidateInstanceKey is given, then it is forced to be promoted over the promotedSlave
// Otherwise, search for the best to promote!
func replacePromotedSlaveWithCandidate(deadInstanceKey *inst.InstanceKey, promotedSlave *inst.Instance, candidateInstanceKey *inst.InstanceKey) (*inst.Instance, error) {
	candidateSlaves, _ := inst.ReadClusterCandidateInstances(promotedSlave.ClusterName)
	// So we've already promoted a slave.
	// However, can we improve on our choice? Are there any slaves marked with "is_candidate"?
	// Maybe we actually promoted such a slave. Does that mean we should keep it?
	// The current logic is:
	// - 1. we prefer to promote a "is_candidate" which is in the same DC & env as the dead intermediate master (or do nothing if the promtoed slave is such one)
	// - 2. we prefer to promote a "is_candidate" which is in the same DC & env as the promoted slave (or do nothing if the promtoed slave is such one)
	// - 3. keep to current choice
	log.Infof("topology_recovery: checking if should replace promoted slave with a better candidate")
	if candidateInstanceKey == nil {
		if deadInstance, _, err := inst.ReadInstance(deadInstanceKey); err == nil && deadInstance != nil {
			for _, candidateSlave := range candidateSlaves {
				if promotedSlave.Key.Equals(&candidateSlave.Key) &&
					promotedSlave.DataCenter == deadInstance.DataCenter &&
					promotedSlave.PhysicalEnvironment == deadInstance.PhysicalEnvironment {
					// Seems like we promoted a candidate in the same DC & ENV as dead IM! Ideal! We're happy!
					log.Infof("topology_recovery: promoted slave %+v is the ideal candidate", promotedSlave.Key)
					return promotedSlave, nil
				}
			}
		}
	}
	// We didn't pick the ideal candidate; let's see if we can replace with a candidate from same DC and ENV
	if candidateInstanceKey == nil {
		// Try a candidate slave that is in same DC & env as the dead instance
		if deadInstance, _, err := inst.ReadInstance(deadInstanceKey); err == nil && deadInstance != nil {
			for _, candidateSlave := range candidateSlaves {
				if candidateSlave.DataCenter == deadInstance.DataCenter &&
					candidateSlave.PhysicalEnvironment == deadInstance.PhysicalEnvironment &&
					candidateSlave.MasterKey.Equals(&promotedSlave.Key) {
					// This would make a great candidate
					candidateInstanceKey = &candidateSlave.Key
					log.Debugf("topology_recovery: no candidate was offered for %+v but orchestrator picks %+v as candidate replacement, based on being in same DC & env as failed instance", promotedSlave.Key, candidateSlave.Key)
				}
			}
		}
	}
	if candidateInstanceKey == nil {
		// We cannot find a candidate in same DC and ENV as dead master
		for _, candidateSlave := range candidateSlaves {
			if promotedSlave.Key.Equals(&candidateSlave.Key) {
				// Seems like we promoted a candidate slave (though not in same DC and ENV as dead master). Good enough.
				// No further action required.
				log.Infof("topology_recovery: promoted slave %+v is a good candidate", promotedSlave.Key)
				return promotedSlave, nil
			}
		}
	}
	// Still nothing?
	if candidateInstanceKey == nil {
		// Try a candidate slave that is in same DC & env as the promoted slave (our promoted slave is not an "is_candidate")
		for _, candidateSlave := range candidateSlaves {
			if promotedSlave.DataCenter == candidateSlave.DataCenter &&
				promotedSlave.PhysicalEnvironment == candidateSlave.PhysicalEnvironment &&
				candidateSlave.MasterKey.Equals(&promotedSlave.Key) {
				// OK, better than nothing
				candidateInstanceKey = &candidateSlave.Key
				log.Debugf("topology_recovery: no candidate was offered for %+v but orchestrator picks %+v as candidate replacement, based on being in same DC & env as promoted instance", promotedSlave.Key, candidateSlave.Key)
			}
		}
	}

	// So do we have a candidate?
	if candidateInstanceKey == nil {
		// Found nothing. Stick with promoted slave
		return promotedSlave, nil
	}
	if promotedSlave.Key.Equals(candidateInstanceKey) {
		// Sanity. It IS the candidate
		return promotedSlave, nil
	}

	// Try and promote suggested candidate, if applicable and possible
	log.Debugf("topology_recovery: promoted instance %+v is not the suggested candidate %+v. Will see what can be done", promotedSlave.Key, *candidateInstanceKey)

	candidateInstance, _, err := inst.ReadInstance(candidateInstanceKey)
	if err != nil {
		return promotedSlave, log.Errore(err)
	}

	if candidateInstance.MasterKey.Equals(&promotedSlave.Key) {
		log.Debugf("topology_recovery: suggested candidate %+v is slave of promoted instance %+v. Will try and enslave its master", *candidateInstanceKey, promotedSlave.Key)
		candidateInstance, err = inst.EnslaveMaster(&candidateInstance.Key)
		if err != nil {
			return promotedSlave, log.Errore(err)
		}
		log.Debugf("topology_recovery: success promoting %+v over %+v", *candidateInstanceKey, promotedSlave.Key)
		return candidateInstance, nil
	}

	log.Debugf("topology_recovery: could not manage to promoted suggested candidate %+v", *candidateInstanceKey)
	return promotedSlave, nil
}

// checkAndRecoverDeadMaster checks a given analysis, decides whether to take action, and possibly takes action
// Returns true when action was taken.
func checkAndRecoverDeadMaster(analysisEntry inst.ReplicationAnalysis, candidateInstanceKey *inst.InstanceKey, skipFilters bool, skipProcesses bool) (bool, *inst.Instance, error) {
	if !(skipFilters || analysisEntry.ClusterDetails.HasAutomatedMasterRecovery) {
		return false, nil, nil
	}
	// Let's do dead master recovery!
	log.Debugf("topology_recovery: will handle DeadMaster event on %+v", analysisEntry.ClusterDetails.ClusterName)
	actionTaken, promotedSlave, err := RecoverDeadMaster(analysisEntry, skipProcesses)

	if actionTaken && promotedSlave != nil {
		promotedSlave, _ = replacePromotedSlaveWithCandidate(&analysisEntry.AnalyzedInstanceKey, promotedSlave, candidateInstanceKey)
	}
	if actionTaken && promotedSlave != nil {
		if !skipProcesses {
			// Execute post master-failover processes
			executeProcesses(config.Config.PostMasterFailoverProcesses, "PostMasterFailoverProcesses", analysisEntry, promotedSlave, false)
		}
	}

	return actionTaken, promotedSlave, err
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
	if sibling.HasReplicationFilters != intermediateMasterInstance.HasReplicationFilters {
		return false
	}
	if sibling.IsBinlogServer() != intermediateMasterInstance.IsBinlogServer() {
		// When both are binlog servers, failover is trivial.
		// When failed IM is binlog server, its sibling is still valid, but we catually prefer to just repoint the slave up -- simplest!
		return false
	}
	if sibling.ExecBinlogCoordinates.SmallerThan(&intermediateMasterInstance.ExecBinlogCoordinates) {
		return false
	}
	return true
}

// GetCandidateSiblingOfIntermediateMaster chooses the best sibling of a dead intermediate master
// to whom the IM's slaves can be moved.
func GetCandidateSiblingOfIntermediateMaster(intermediateMasterInstance *inst.Instance) (*inst.Instance, error) {

	siblings, err := inst.ReadSlaveInstances(&intermediateMasterInstance.MasterKey)
	if err != nil {
		return nil, err
	}
	if len(siblings) <= 1 {
		return nil, log.Errorf("topology_recovery: no siblings found for %+v", intermediateMasterInstance.Key)
	}

	sort.Sort(sort.Reverse(InstancesByCountSlaves(siblings)))

	// In the next series of steps we attempt to return a good replacement.
	// None of the below attempts is sure to pick a winning server. Perhaps picked server is not enough up-todate -- but
	// this has small likelihood in the general case, and, well, it's an attempt. It's a Plan A, but we have Plan B & C if this fails.

	// At first, we try to return an "is_candidate" server in same dc & env
	log.Infof("topology_recovery: searching for the best candidate sibling of dead intermediate master")
	for _, sibling := range siblings {
		sibling := sibling
		if isValidAsCandidateSiblingOfIntermediateMaster(intermediateMasterInstance, sibling) &&
			sibling.IsCandidate &&
			sibling.DataCenter == intermediateMasterInstance.DataCenter &&
			sibling.PhysicalEnvironment == intermediateMasterInstance.PhysicalEnvironment {
			log.Infof("topology_recovery: found %+v as the ideal candidate", sibling.Key)
			return sibling, nil
		}
	}
	// Go for something else in the same DC & ENV
	for _, sibling := range siblings {
		sibling := sibling
		if isValidAsCandidateSiblingOfIntermediateMaster(intermediateMasterInstance, sibling) &&
			sibling.DataCenter == intermediateMasterInstance.DataCenter &&
			sibling.PhysicalEnvironment == intermediateMasterInstance.PhysicalEnvironment {
			log.Infof("topology_recovery: found %+v as a replacement in same dc & environment", sibling.Key)
			return sibling, nil
		}
	}
	// Nothing in same DC & env, let's just go for some is_candidate
	for _, sibling := range siblings {
		sibling := sibling
		if isValidAsCandidateSiblingOfIntermediateMaster(intermediateMasterInstance, sibling) && sibling.IsCandidate {
			log.Infof("topology_recovery: found %+v as a good candidate", sibling.Key)
			return sibling, nil
		}
	}
	// Havent found an "is_candidate". Just whatever is valid.
	for _, sibling := range siblings {
		sibling := sibling
		if isValidAsCandidateSiblingOfIntermediateMaster(intermediateMasterInstance, sibling) {
			log.Infof("topology_recovery: found %+v as a replacement", sibling.Key)
			return sibling, nil
		}
	}
	return nil, log.Errorf("topology_recovery: cannot find candidate sibling of %+v", intermediateMasterInstance.Key)
}

func RecoverDeadIntermediateMaster(analysisEntry inst.ReplicationAnalysis, skipProcesses bool) (actionTaken bool, successorInstance *inst.Instance, err error) {
	failedInstanceKey := &analysisEntry.AnalyzedInstanceKey
	if ok, err := AttemptRecoveryRegistration(&analysisEntry); !ok {
		log.Debugf("topology_recovery: found an active or recent recovery on %+v. Will not issue another RecoverDeadIntermediateMaster.", *failedInstanceKey)
		return false, nil, err
	}

	inst.AuditOperation("recover-dead-intermediate-master", failedInstanceKey, "problem found; will recover")
	log.Debugf("topology_recovery: RecoverDeadIntermediateMaster: will recover %+v", *failedInstanceKey)
	if !skipProcesses {
		if err := executeProcesses(config.Config.PreFailoverProcesses, "PreFailoverProcesses", analysisEntry, nil, true); err != nil {
			return false, nil, err
		}
	}

	intermediateMasterInstance, _, err := inst.ReadInstance(failedInstanceKey)
	if err != nil {
		return false, nil, err
	}
	// Plan A: find a replacement intermediate master in same Data Center
	candidateSiblingOfIntermediateMaster, err := GetCandidateSiblingOfIntermediateMaster(intermediateMasterInstance)

	relocateSlavesToCandidateSibling := func() {
		if candidateSiblingOfIntermediateMaster == nil {
			return
		}
		log.Debugf("topology_recovery: - RecoverDeadIntermediateMaster: will attempt a candidate intermediate master: %+v", candidateSiblingOfIntermediateMaster.Key)
		// We have a candidate
		if relocatedSlaves, candidateSibling, err, errs := inst.RelocateSlaves(failedInstanceKey, &candidateSiblingOfIntermediateMaster.Key, ""); err == nil {
			ResolveRecovery(failedInstanceKey, &candidateSibling.Key)

			successorInstance = candidateSibling
			actionTaken = true

			log.Debugf("topology_recovery: - RecoverDeadIntermediateMaster: move to candidate intermediate master (%+v) went with %d errors", candidateSibling.Key, len(errs))
			inst.AuditOperation("recover-dead-intermediate-master", failedInstanceKey, fmt.Sprintf("Done. Relocated %d slaves under candidate sibling: %+v; %d errors: %+v", len(relocatedSlaves), candidateSibling.Key, len(errs), errs))
		} else {
			log.Debugf("topology_recovery: - RecoverDeadIntermediateMaster: move to candidate intermediate master (%+v) did not complete: %+v", candidateSibling.Key, err)
			inst.AuditOperation("recover-dead-intermediate-master", failedInstanceKey, fmt.Sprintf("Relocated %d slaves under candidate sibling: %+v; %d errors: %+v", len(relocatedSlaves), candidateSibling.Key, len(errs), errs))
		}
	}
	if candidateSiblingOfIntermediateMaster != nil && candidateSiblingOfIntermediateMaster.DataCenter == intermediateMasterInstance.DataCenter {
		relocateSlavesToCandidateSibling()
	}
	if !actionTaken {
		log.Debugf("topology_recovery: - RecoverDeadIntermediateMaster: will next attempt regrouping of slaves")
		// Plan B: regroup (we wish to reduce cross-DC replication streams)
		_, _, _, _, err = inst.RegroupSlaves(failedInstanceKey, true, nil)
		if err != nil {
			log.Debugf("topology_recovery: - RecoverDeadIntermediateMaster: regroup failed on: %+v", err)
		}
		// Plan C: try replacement intermediate master in other DC...
		if candidateSiblingOfIntermediateMaster != nil && candidateSiblingOfIntermediateMaster.DataCenter != intermediateMasterInstance.DataCenter {
			log.Debugf("topology_recovery: - RecoverDeadIntermediateMaster: will next attempt relocating to another DC server")
			relocateSlavesToCandidateSibling()
		}
	}
	if !actionTaken {
		// Do we still have leftovers? Some slaves couldn't move? Couldn't regroup? Only left with regroup's resulting leader?
		// nothing moved?
		// We don't care much if regroup made it or not. We prefer that it made it, in whcih case we only need to relocate up
		// one slave, but the operation is still valid if regroup partially/completely failed. We just promote anything
		// not regrouped.
		// So, match up all that's left, plan D
		log.Debugf("topology_recovery: - RecoverDeadIntermediateMaster: will next attempt to relocate up from %+v", *failedInstanceKey)

		var errs []error
		var relocatedSlaves [](*inst.Instance)
		relocatedSlaves, successorInstance, err, errs = inst.RelocateSlaves(failedInstanceKey, &analysisEntry.AnalyzedInstanceMasterKey, "")

		if len(relocatedSlaves) > 0 {
			actionTaken = true
			log.Debugf("topology_recovery: - RecoverDeadIntermediateMaster: relocated up to %+v", successorInstance.Key)
			inst.AuditOperation("recover-dead-intermediate-master", failedInstanceKey, fmt.Sprintf("Done. Relocated slaves under: %+v %d errors: %+v", successorInstance.Key, len(errs), errs))
		} else {
			err = log.Errorf("topology_recovery: RecoverDeadIntermediateMaster failed to match up any slave from %+v", *failedInstanceKey)
		}
	}
	if successorInstance != nil {
		ResolveRecovery(failedInstanceKey, &successorInstance.Key)
	} else {
		ResolveRecovery(failedInstanceKey, nil)
	}
	return actionTaken, successorInstance, err
}

// checkAndRecoverDeadIntermediateMaster checks a given analysis, decides whether to take action, and possibly takes action
// Returns true when action was taken.
func checkAndRecoverDeadIntermediateMaster(analysisEntry inst.ReplicationAnalysis, candidateInstanceKey *inst.InstanceKey, skipFilters bool, skipProcesses bool) (bool, *inst.Instance, error) {
	if !(skipFilters || analysisEntry.ClusterDetails.HasAutomatedIntermediateMasterRecovery) {
		return false, nil, nil
	}

	actionTaken, promotedSlave, err := RecoverDeadIntermediateMaster(analysisEntry, skipProcesses)
	if actionTaken {
		if !skipProcesses {
			// Execute post intermediate-master-failover processes
			executeProcesses(config.Config.PostIntermediateMasterFailoverProcesses, "PostIntermediateMasterFailoverProcesses", analysisEntry, promotedSlave, false)
		}
	}
	return actionTaken, promotedSlave, err
}

// checkAndRecoverGenericProblem is a general=purpose recovery function
func checkAndRecoverGenericProblem(analysisEntry inst.ReplicationAnalysis, candidateInstanceKey *inst.InstanceKey, skipFilters bool, skipProcesses bool) (bool, *inst.Instance, error) {
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
	slaves, err := inst.ReadSlaveInstancesIncludingBinlogServerSubSlaves(instanceKey)
	if err != nil {
		return
	}
	for _, slave := range slaves {
		go emergentlyReadTopologyInstance(&slave.Key, analysisCode)
	}
}

// executeCheckAndRecoverFunction will choose the correct check & recovery function based on analysis.
// It executes the function synchronuously
func executeCheckAndRecoverFunction(analysisEntry inst.ReplicationAnalysis, candidateInstanceKey *inst.InstanceKey, skipFilters bool, skipProcesses bool) (bool, *inst.Instance, error) {
	var checkAndRecoverFunction func(analysisEntry inst.ReplicationAnalysis, candidateInstanceKey *inst.InstanceKey, skipFilters bool, skipProcesses bool) (bool, *inst.Instance, error) = nil

	switch analysisEntry.Analysis {
	case inst.DeadMaster:
		checkAndRecoverFunction = checkAndRecoverDeadMaster
	case inst.DeadMasterAndSomeSlaves:
		checkAndRecoverFunction = checkAndRecoverDeadMaster
	case inst.DeadIntermediateMaster:
		checkAndRecoverFunction = checkAndRecoverDeadIntermediateMaster
	case inst.DeadIntermediateMasterAndSomeSlaves:
		checkAndRecoverFunction = checkAndRecoverDeadIntermediateMaster
	case inst.DeadIntermediateMasterWithSingleSlaveFailingToConnect:
		checkAndRecoverFunction = checkAndRecoverDeadIntermediateMaster
	case inst.AllIntermediateMasterSlavesFailingToConnectOrDead:
		checkAndRecoverFunction = checkAndRecoverDeadIntermediateMaster
	case inst.DeadCoMaster:
		checkAndRecoverFunction = checkAndRecoverDeadIntermediateMaster
	case inst.DeadMasterAndSlaves:
		go emergentlyReadTopologyInstance(&analysisEntry.AnalyzedInstanceMasterKey, analysisEntry.Analysis)
	case inst.UnreachableMaster:
		go emergentlyReadTopologyInstanceSlaves(&analysisEntry.AnalyzedInstanceKey, analysisEntry.Analysis)
	case inst.AllMasterSlavesNotReplicating:
		//checkAndRecoverFunction = checkAndRecoverGenericProblem
		go emergentlyReadTopologyInstance(&analysisEntry.AnalyzedInstanceKey, analysisEntry.Analysis)
	case inst.FirstTierSlaveFailingToConnectToMaster:
		go emergentlyReadTopologyInstance(&analysisEntry.AnalyzedInstanceMasterKey, analysisEntry.Analysis)
	}

	if checkAndRecoverFunction == nil {
		// Unhandled problem type
		return false, nil, nil
	}
	// we have a recovery function; its execution still depends on filters if not disabled.
	log.Debugf("executeCheckAndRecoverFunction: proceeeding with %+v; skipProcesses: %+v", analysisEntry.AnalyzedInstanceKey, skipProcesses)

	if ok, _ := AttemptFailureDetectionRegistration(&analysisEntry); ok {
		log.Debugf("topology_recovery: detected %+v failure on %+v", analysisEntry.Analysis, analysisEntry.AnalyzedInstanceKey)
		// Execute on-detection processes
		if !skipProcesses {
			if err := executeProcesses(config.Config.OnFailureDetectionProcesses, "OnFailureDetectionProcesses", analysisEntry, nil, true); err != nil {
				return false, nil, err
			}
		}
	}

	actionTaken, promotedSlave, err := checkAndRecoverFunction(analysisEntry, candidateInstanceKey, skipFilters, skipProcesses)
	if actionTaken {
		if !skipProcesses {
			// Execute post intermediate-master-failover processes
			executeProcesses(config.Config.PostFailoverProcesses, "PostFailoverProcesses", analysisEntry, promotedSlave, false)
		}
	}
	return actionTaken, promotedSlave, err
}

// CheckAndRecover is the main entry point for the recovery mechanism
func CheckAndRecover(specificInstance *inst.InstanceKey, candidateInstanceKey *inst.InstanceKey, skipFilters bool, skipProcesses bool) (actionTaken bool, instance *inst.Instance, err error) {
	replicationAnalysis, err := inst.GetReplicationAnalysis(true)
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
		if analysisEntry.IsDowntimed && specificInstance == nil {
			// Only recover a downtimed server if explicitly requested
			continue
		}

		if specificInstance != nil && skipFilters {
			// force mode. Keep it synchronuous
			actionTaken, instance, err = executeCheckAndRecoverFunction(analysisEntry, candidateInstanceKey, skipFilters, skipProcesses)
		} else {
			go executeCheckAndRecoverFunction(analysisEntry, candidateInstanceKey, skipFilters, skipProcesses)
		}
	}
	return actionTaken, instance, err
}
