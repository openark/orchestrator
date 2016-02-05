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
	"github.com/outbrain/orchestrator/go/process"
	"github.com/pmylund/go-cache"
	"github.com/rcrowley/go-metrics"
	"sort"
	"strings"
	"time"
)

// BlockedTopologyRecovery represents an entry in the blocked_topology_recovery table
type BlockedTopologyRecovery struct {
	FailedInstanceKey    inst.InstanceKey
	ClusterName          string
	Analysis             inst.AnalysisCode
	LastBlockedTimestamp string
	BlockingRecoveryId   int64
}

// TopologyRecovery represents an entry in the topology_recovery table
type TopologyRecovery struct {
	inst.PostponedFunctionsContainer

	Id                        int64
	AnalysisEntry             inst.ReplicationAnalysis
	SuccessorKey              *inst.InstanceKey
	IsActive                  bool
	IsSuccessful              bool
	LostSlaves                inst.InstanceKeyMap
	ParticipatingInstanceKeys inst.InstanceKeyMap
	AllErrors                 []string
	RecoveryStartTimestamp    string
	RecoveryEndTimestamp      string
	ProcessingNodeHostname    string
	ProcessingNodeToken       string
	PostponedFunctions        [](func() error)
	Acknowledged              bool
	AcknowledgedAt            string
	AcknowledgedBy            string
	AcknowledgedComment       string
	LastDetectionId           int64
	RelatedRecoveryId         int64
}

func NewTopologyRecovery(replicationAnalysis inst.ReplicationAnalysis) *TopologyRecovery {
	topologyRecovery := &TopologyRecovery{}
	topologyRecovery.AnalysisEntry = replicationAnalysis
	topologyRecovery.SuccessorKey = nil
	topologyRecovery.LostSlaves = *inst.NewInstanceKeyMap()
	topologyRecovery.ParticipatingInstanceKeys = *inst.NewInstanceKeyMap()
	topologyRecovery.AllErrors = []string{}
	topologyRecovery.PostponedFunctions = [](func() error){}
	return topologyRecovery
}

func (this *TopologyRecovery) AddError(err error) error {
	if err != nil {
		this.AllErrors = append(this.AllErrors, err.Error())
	}
	return err
}

func (this *TopologyRecovery) AddErrors(errs []error) {
	for _, err := range errs {
		this.AddError(err)
	}
}

type MasterRecoveryType string

const (
	MasterRecoveryGTID         MasterRecoveryType = "MasterRecoveryGTID"
	MasterRecoveryPseudoGTID                      = "MasterRecoveryPseudoGTID"
	MasterRecoveryBinlogServer                    = "MasterRecoveryBinlogServer"
)

var emptySlavesList [](*inst.Instance)

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

var recoverDeadMasterCounter = metrics.NewCounter()
var recoverDeadMasterSuccessCounter = metrics.NewCounter()
var recoverDeadMasterFailureCounter = metrics.NewCounter()
var recoverDeadIntermediateMasterCounter = metrics.NewCounter()
var recoverDeadIntermediateMasterSuccessCounter = metrics.NewCounter()
var recoverDeadIntermediateMasterFailureCounter = metrics.NewCounter()
var recoverDeadCoMasterCounter = metrics.NewCounter()
var recoverDeadCoMasterSuccessCounter = metrics.NewCounter()
var recoverDeadCoMasterFailureCounter = metrics.NewCounter()

func init() {
	metrics.Register("recover.dead_master.start", recoverDeadMasterCounter)
	metrics.Register("recover.dead_master.success", recoverDeadMasterSuccessCounter)
	metrics.Register("recover.dead_master.fail", recoverDeadMasterFailureCounter)
	metrics.Register("recover.dead_intermediate_master.start", recoverDeadIntermediateMasterCounter)
	metrics.Register("recover.dead_intermediate_master.success", recoverDeadIntermediateMasterSuccessCounter)
	metrics.Register("recover.dead_intermediate_master.fail", recoverDeadIntermediateMasterFailureCounter)
	metrics.Register("recover.dead_co_master.start", recoverDeadCoMasterCounter)
	metrics.Register("recover.dead_co_master.success", recoverDeadCoMasterSuccessCounter)
	metrics.Register("recover.dead_co_master.fail", recoverDeadCoMasterFailureCounter)
}

// replaceCommandPlaceholders replaxces agreed-upon placeholders with analysis data
func replaceCommandPlaceholders(command string, topologyRecovery *TopologyRecovery) string {
	analysisEntry := &topologyRecovery.AnalysisEntry
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
	command = strings.Replace(command, "{orchestratorHost}", process.ThisHostname, -1)

	command = strings.Replace(command, "{isSuccessful}", fmt.Sprint(topologyRecovery.SuccessorKey != nil), -1)
	if topologyRecovery.SuccessorKey != nil {
		command = strings.Replace(command, "{successorHost}", topologyRecovery.SuccessorKey.Hostname, -1)
		command = strings.Replace(command, "{successorPort}", fmt.Sprintf("%d", topologyRecovery.SuccessorKey.Port), -1)
	}

	command = strings.Replace(command, "{lostSlaves}", topologyRecovery.LostSlaves.ToCommaDelimitedList(), -1)
	command = strings.Replace(command, "{slaveHosts}", analysisEntry.SlaveHosts.ToCommaDelimitedList(), -1)

	return command
}

// executeProcesses executes a list of processes
func executeProcesses(processes []string, description string, topologyRecovery *TopologyRecovery, failOnError bool) error {
	var err error
	for _, command := range processes {
		command := replaceCommandPlaceholders(command, topologyRecovery)

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

func recoverDeadMasterInBinlogServerTopology(topologyRecovery *TopologyRecovery) (promotedSlave *inst.Instance, err error) {
	failedMasterKey := &topologyRecovery.AnalysisEntry.AnalyzedInstanceKey

	var promotedBinlogServer *inst.Instance

	_, promotedBinlogServer, err = inst.RegroupSlavesBinlogServers(failedMasterKey, true)
	if err != nil {
		return nil, log.Errore(err)
	}
	promotedBinlogServer, err = inst.StopSlave(&promotedBinlogServer.Key)
	if err != nil {
		return promotedSlave, log.Errore(err)
	}
	// Find candidate slave
	promotedSlave, err = inst.GetCandidateSlaveOfBinlogServerTopology(&promotedBinlogServer.Key)
	if err != nil {
		return promotedSlave, log.Errore(err)
	}
	// Align it with binlog server coordinates
	promotedSlave, err = inst.StopSlave(&promotedSlave.Key)
	if err != nil {
		return promotedSlave, log.Errore(err)
	}
	promotedSlave, err = inst.StartSlaveUntilMasterCoordinates(&promotedSlave.Key, &promotedBinlogServer.ExecBinlogCoordinates)
	if err != nil {
		return promotedSlave, log.Errore(err)
	}
	promotedSlave, err = inst.StopSlave(&promotedSlave.Key)
	if err != nil {
		return promotedSlave, log.Errore(err)
	}
	// Detach, flush binary logs forward
	promotedSlave, err = inst.ResetSlave(&promotedSlave.Key)
	if err != nil {
		return promotedSlave, log.Errore(err)
	}
	promotedSlave, err = inst.FlushBinaryLogsTo(&promotedSlave.Key, promotedBinlogServer.ExecBinlogCoordinates.LogFile)
	if err != nil {
		return promotedSlave, log.Errore(err)
	}
	promotedSlave, err = inst.FlushBinaryLogs(&promotedSlave.Key, 1)
	if err != nil {
		return promotedSlave, log.Errore(err)
	}
	promotedSlave, err = inst.PurgeBinaryLogsToCurrent(&promotedSlave.Key)
	if err != nil {
		return promotedSlave, log.Errore(err)
	}
	// Reconnect binlog servers to promoted slave (now master):
	promotedBinlogServer, err = inst.SkipToNextBinaryLog(&promotedBinlogServer.Key)
	if err != nil {
		return promotedSlave, log.Errore(err)
	}
	promotedBinlogServer, err = inst.Repoint(&promotedBinlogServer.Key, &promotedSlave.Key, inst.GTIDHintDeny)
	if err != nil {
		return nil, log.Errore(err)
	}

	func() {
		// Move binlog server slaves up to replicate from master.
		// This can only be done once a BLS has skipped to the next binlog
		// We postpone this operation. The master is already promoted and we're happy.
		binlogServerSlaves, err := inst.ReadBinlogServerSlaveInstances(&promotedBinlogServer.Key)
		if err != nil {
			return
		}
		maxBinlogServersToPromote := 3
		for i, binlogServerSlave := range binlogServerSlaves {
			binlogServerSlave := binlogServerSlave
			if i >= maxBinlogServersToPromote {
				return
			}
			postponedFunction := func() error {
				binlogServerSlave, err := inst.StopSlave(&binlogServerSlave.Key)
				if err != nil {
					return err
				}
				// Make sure the BLS has the "next binlog" -- the one the master flushed & purged to. Otherwise the BLS
				// will request a binlog the master does not have
				if binlogServerSlave.ExecBinlogCoordinates.SmallerThan(&promotedBinlogServer.ExecBinlogCoordinates) {
					binlogServerSlave, err = inst.StartSlaveUntilMasterCoordinates(&binlogServerSlave.Key, &promotedBinlogServer.ExecBinlogCoordinates)
					if err != nil {
						return err
					}
				}
				_, err = inst.Repoint(&binlogServerSlave.Key, &promotedSlave.Key, inst.GTIDHintDeny)
				return err
			}
			topologyRecovery.AddPostponedFunction(postponedFunction)
		}
	}()

	return promotedSlave, err
}

// RecoverDeadMaster recovers a dead master, complete logic inside
func RecoverDeadMaster(topologyRecovery *TopologyRecovery, skipProcesses bool) (promotedSlave *inst.Instance, lostSlaves [](*inst.Instance), err error) {
	analysisEntry := &topologyRecovery.AnalysisEntry
	failedInstanceKey := &analysisEntry.AnalyzedInstanceKey

	inst.AuditOperation("recover-dead-master", failedInstanceKey, "problem found; will recover")
	if !skipProcesses {
		if err := executeProcesses(config.Config.PreFailoverProcesses, "PreFailoverProcesses", topologyRecovery, true); err != nil {
			return nil, lostSlaves, topologyRecovery.AddError(err)
		}
	}

	log.Debugf("topology_recovery: RecoverDeadMaster: will recover %+v", *failedInstanceKey)

	var masterRecoveryType MasterRecoveryType = MasterRecoveryPseudoGTID
	if analysisEntry.OracleGTIDImmediateTopology || analysisEntry.MariaDBGTIDImmediateTopology {
		masterRecoveryType = MasterRecoveryGTID
	} else if analysisEntry.BinlogServerImmediateTopology {
		masterRecoveryType = MasterRecoveryBinlogServer
	}
	log.Debugf("topology_recovery: RecoverDeadMaster: masterRecoveryType=%+v", masterRecoveryType)

	switch masterRecoveryType {
	case MasterRecoveryGTID:
		{
			lostSlaves, _, promotedSlave, err = inst.RegroupSlavesGTID(failedInstanceKey, true, nil)
		}
	case MasterRecoveryPseudoGTID:
		{
			lostSlaves, _, _, promotedSlave, err = inst.RegroupSlavesPseudoGTIDIncludingSubSlavesOfBinlogServers(failedInstanceKey, true, nil, &topologyRecovery.PostponedFunctionsContainer)
		}
	case MasterRecoveryBinlogServer:
		{
			promotedSlave, err = recoverDeadMasterInBinlogServerTopology(topologyRecovery)
		}
	}
	topologyRecovery.AddError(err)

	if promotedSlave != nil && len(lostSlaves) > 0 && config.Config.DetachLostSlavesAfterMasterFailover {
		postponedFunction := func() error {
			log.Debugf("topology_recovery: - RecoverDeadMaster: lost %+v slaves during recovery process; detaching them", len(lostSlaves))
			for _, slave := range lostSlaves {
				slave := slave
				inst.DetachSlaveOperation(&slave.Key)
			}
			return nil
		}
		topologyRecovery.AddPostponedFunction(postponedFunction)
	}
	if config.Config.MasterFailoverLostInstancesDowntimeMinutes > 0 {
		postponedFunction := func() error {
			inst.BeginDowntime(failedInstanceKey, inst.GetMaintenanceOwner(), "RecoverDeadMaster indicates this instance is lost", config.Config.MasterFailoverLostInstancesDowntimeMinutes*60)
			for _, slave := range lostSlaves {
				slave := slave
				inst.BeginDowntime(&slave.Key, inst.GetMaintenanceOwner(), "RecoverDeadMaster indicates this instance is lost", config.Config.MasterFailoverLostInstancesDowntimeMinutes*60)
			}
			return nil
		}
		topologyRecovery.AddPostponedFunction(postponedFunction)
	}

	if promotedSlave == nil {
		inst.AuditOperation("recover-dead-master", failedInstanceKey, "Failure: no slave promoted.")
	} else {
		inst.AuditOperation("recover-dead-master", failedInstanceKey, fmt.Sprintf("promoted slave: %+v", promotedSlave.Key))
	}
	return promotedSlave, lostSlaves, err
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
		// Sanity. It IS the candidate, nothing to promote...
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
func checkAndRecoverDeadMaster(analysisEntry inst.ReplicationAnalysis, candidateInstanceKey *inst.InstanceKey, forceInstanceRecovery bool, skipProcesses bool) (bool, *TopologyRecovery, error) {
	if !(forceInstanceRecovery || analysisEntry.ClusterDetails.HasAutomatedMasterRecovery) {
		return false, nil, nil
	}
	topologyRecovery, err := AttemptRecoveryRegistration(&analysisEntry, !forceInstanceRecovery, !forceInstanceRecovery)
	if topologyRecovery == nil {
		log.Debugf("topology_recovery: found an active or recent recovery on %+v. Will not issue another RecoverDeadMaster.", analysisEntry.AnalyzedInstanceKey)
		return false, nil, err
	}

	// That's it! We must do recovery!
	log.Debugf("topology_recovery: will handle DeadMaster event on %+v", analysisEntry.ClusterDetails.ClusterName)
	recoverDeadMasterCounter.Inc(1)
	promotedSlave, lostSlaves, err := RecoverDeadMaster(topologyRecovery, skipProcesses)
	topologyRecovery.LostSlaves.AddInstances(lostSlaves)

	if promotedSlave != nil {
		promotedSlave, err = replacePromotedSlaveWithCandidate(&analysisEntry.AnalyzedInstanceKey, promotedSlave, candidateInstanceKey)
		topologyRecovery.AddError(err)
	}
	// And this is the end; whether successful or not, we're done.
	ResolveRecovery(topologyRecovery, promotedSlave)
	if promotedSlave != nil {
		// Success!
		recoverDeadMasterSuccessCounter.Inc(1)

		if config.Config.ApplyMySQLPromotionAfterMasterFailover {
			log.Debugf("topology_recovery: - RecoverDeadMaster: will apply MySQL changes to promoted master")
			inst.ResetSlaveOperation(&promotedSlave.Key)
			inst.SetReadOnly(&promotedSlave.Key, false)
		}
		if !skipProcesses {
			// Execute post master-failover processes
			executeProcesses(config.Config.PostMasterFailoverProcesses, "PostMasterFailoverProcesses", topologyRecovery, false)
		}
	} else {
		recoverDeadMasterFailureCounter.Inc(1)
	}

	return true, topologyRecovery, err
}

// isGeneralyValidAsCandidateSiblingOfIntermediateMaster sees that basic server configuration and state are valid
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

// isValidAsCandidateSiblingOfIntermediateMaster checks to see that the given sibling is capable to take over instance's slaves
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

// RecoverDeadIntermediateMaster performs intermediate master recovery; complete logic inside
func RecoverDeadIntermediateMaster(topologyRecovery *TopologyRecovery, skipProcesses bool) (successorInstance *inst.Instance, err error) {
	analysisEntry := &topologyRecovery.AnalysisEntry
	failedInstanceKey := &analysisEntry.AnalyzedInstanceKey
	recoveryResolved := false

	inst.AuditOperation("recover-dead-intermediate-master", failedInstanceKey, "problem found; will recover")
	if !skipProcesses {
		if err := executeProcesses(config.Config.PreFailoverProcesses, "PreFailoverProcesses", topologyRecovery, true); err != nil {
			return nil, topologyRecovery.AddError(err)
		}
	}

	intermediateMasterInstance, _, err := inst.ReadInstance(failedInstanceKey)
	if err != nil {
		return nil, topologyRecovery.AddError(err)
	}
	// Find possible candidate
	candidateSiblingOfIntermediateMaster, err := GetCandidateSiblingOfIntermediateMaster(intermediateMasterInstance)
	relocateSlavesToCandidateSibling := func() {
		if candidateSiblingOfIntermediateMaster == nil {
			return
		}
		// We have a candidate
		log.Debugf("topology_recovery: - RecoverDeadIntermediateMaster: will attempt a candidate intermediate master: %+v", candidateSiblingOfIntermediateMaster.Key)
		relocatedSlaves, candidateSibling, err, errs := inst.RelocateSlaves(failedInstanceKey, &candidateSiblingOfIntermediateMaster.Key, "")
		topologyRecovery.AddErrors(errs)
		topologyRecovery.ParticipatingInstanceKeys.AddKey(candidateSiblingOfIntermediateMaster.Key)

		if len(relocatedSlaves) == 0 {
			log.Debugf("topology_recovery: - RecoverDeadIntermediateMaster: failed to move any slave to candidate intermediate master (%+v)", candidateSibling.Key)
			return
		}
		if err != nil || len(errs) > 0 {
			log.Debugf("topology_recovery: - RecoverDeadIntermediateMaster: move to candidate intermediate master (%+v) did not complete: %+v", candidateSibling.Key, err)
			return
		}
		if err == nil {
			recoveryResolved = true
			successorInstance = candidateSibling

			inst.AuditOperation("recover-dead-intermediate-master", failedInstanceKey, fmt.Sprintf("Relocated %d slaves under candidate sibling: %+v; %d errors: %+v", len(relocatedSlaves), candidateSibling.Key, len(errs), errs))
		}
	}
	// Plan A: find a replacement intermediate master in same Data Center
	if candidateSiblingOfIntermediateMaster != nil && candidateSiblingOfIntermediateMaster.DataCenter == intermediateMasterInstance.DataCenter {
		relocateSlavesToCandidateSibling()
	}
	if !recoveryResolved {
		log.Debugf("topology_recovery: - RecoverDeadIntermediateMaster: will next attempt regrouping of slaves")
		// Plan B: regroup (we wish to reduce cross-DC replication streams)
		_, _, _, regroupPromotedSlave, err := inst.RegroupSlaves(failedInstanceKey, true, nil, nil)
		if err != nil {
			topologyRecovery.AddError(err)
			log.Debugf("topology_recovery: - RecoverDeadIntermediateMaster: regroup failed on: %+v", err)
		}
		if regroupPromotedSlave != nil {
			topologyRecovery.ParticipatingInstanceKeys.AddKey(regroupPromotedSlave.Key)
		}
		// Plan C: try replacement intermediate master in other DC...
		if candidateSiblingOfIntermediateMaster != nil && candidateSiblingOfIntermediateMaster.DataCenter != intermediateMasterInstance.DataCenter {
			log.Debugf("topology_recovery: - RecoverDeadIntermediateMaster: will next attempt relocating to another DC server")
			relocateSlavesToCandidateSibling()
		}
	}
	if !recoveryResolved {
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
		topologyRecovery.AddErrors(errs)
		topologyRecovery.ParticipatingInstanceKeys.AddKey(analysisEntry.AnalyzedInstanceMasterKey)

		if len(relocatedSlaves) > 0 {
			recoveryResolved = true
			inst.AuditOperation("recover-dead-intermediate-master", failedInstanceKey, fmt.Sprintf("Relocated slaves under: %+v %d errors: %+v", successorInstance.Key, len(errs), errs))
		} else {
			err = log.Errorf("topology_recovery: RecoverDeadIntermediateMaster failed to match up any slave from %+v", *failedInstanceKey)
			topologyRecovery.AddError(err)
		}
	}
	if !recoveryResolved {
		successorInstance = nil
	}
	ResolveRecovery(topologyRecovery, successorInstance)
	return successorInstance, err
}

// checkAndRecoverDeadIntermediateMaster checks a given analysis, decides whether to take action, and possibly takes action
// Returns true when action was taken.
func checkAndRecoverDeadIntermediateMaster(analysisEntry inst.ReplicationAnalysis, candidateInstanceKey *inst.InstanceKey, forceInstanceRecovery bool, skipProcesses bool) (bool, *TopologyRecovery, error) {
	if !(forceInstanceRecovery || analysisEntry.ClusterDetails.HasAutomatedIntermediateMasterRecovery) {
		return false, nil, nil
	}
	topologyRecovery, err := AttemptRecoveryRegistration(&analysisEntry, !forceInstanceRecovery, !forceInstanceRecovery)
	if topologyRecovery == nil {
		log.Debugf("topology_recovery: found an active or recent recovery on %+v. Will not issue another RecoverDeadIntermediateMaster.", analysisEntry.AnalyzedInstanceKey)
		return false, nil, err
	}

	// That's it! We must do recovery!
	recoverDeadIntermediateMasterCounter.Inc(1)
	promotedSlave, err := RecoverDeadIntermediateMaster(topologyRecovery, skipProcesses)
	if promotedSlave != nil {
		// success
		recoverDeadIntermediateMasterSuccessCounter.Inc(1)

		if !skipProcesses {
			// Execute post intermediate-master-failover processes
			topologyRecovery.SuccessorKey = &promotedSlave.Key
			executeProcesses(config.Config.PostIntermediateMasterFailoverProcesses, "PostIntermediateMasterFailoverProcesses", topologyRecovery, false)
		}
	} else {
		recoverDeadIntermediateMasterFailureCounter.Inc(1)
	}
	return true, topologyRecovery, err
}

// RecoverDeadCoMaster recovers a dead co-master, complete logic inside
func RecoverDeadCoMaster(topologyRecovery *TopologyRecovery, skipProcesses bool) (promotedSlave *inst.Instance, lostSlaves [](*inst.Instance), err error) {
	analysisEntry := &topologyRecovery.AnalysisEntry
	failedInstanceKey := &analysisEntry.AnalyzedInstanceKey
	otherCoMasterKey := &analysisEntry.AnalyzedInstanceMasterKey
	otherCoMaster, found, _ := inst.ReadInstance(otherCoMasterKey)
	if otherCoMaster == nil || !found {
		return nil, lostSlaves, topologyRecovery.AddError(log.Errorf("RecoverDeadCoMaster: could not read info for co-master %+v of %+v", *otherCoMasterKey, *failedInstanceKey))
	}
	inst.AuditOperation("recover-dead-co-master", failedInstanceKey, "problem found; will recover")
	if !skipProcesses {
		if err := executeProcesses(config.Config.PreFailoverProcesses, "PreFailoverProcesses", topologyRecovery, true); err != nil {
			return nil, lostSlaves, topologyRecovery.AddError(err)
		}
	}

	log.Debugf("topology_recovery: RecoverDeadCoMaster: will recover %+v", *failedInstanceKey)

	var coMasterRecoveryType MasterRecoveryType = MasterRecoveryPseudoGTID
	if analysisEntry.OracleGTIDImmediateTopology || analysisEntry.MariaDBGTIDImmediateTopology {
		coMasterRecoveryType = MasterRecoveryGTID
	}

	log.Debugf("topology_recovery: RecoverDeadCoMaster: coMasterRecoveryType=%+v", coMasterRecoveryType)

	switch coMasterRecoveryType {
	case MasterRecoveryGTID:
		{
			lostSlaves, _, promotedSlave, err = inst.RegroupSlavesGTID(failedInstanceKey, true, nil)
		}
	case MasterRecoveryPseudoGTID:
		{
			lostSlaves, _, _, promotedSlave, err = inst.RegroupSlavesPseudoGTIDIncludingSubSlavesOfBinlogServers(failedInstanceKey, true, nil, &topologyRecovery.PostponedFunctionsContainer)
		}
	}
	topologyRecovery.AddError(err)

	mustPromoteOtherCoMaster := config.Config.CoMasterRecoveryMustPromoteOtherCoMaster
	if !otherCoMaster.ReadOnly {
		log.Debugf("topology_recovery: RecoverDeadCoMaster: other co-master %+v is writeable hence has to be promoted", otherCoMaster.Key)
		mustPromoteOtherCoMaster = true
	}
	log.Debugf("topology_recovery: RecoverDeadCoMaster: mustPromoteOtherCoMaster? %+v", mustPromoteOtherCoMaster)

	if promotedSlave != nil {
		topologyRecovery.ParticipatingInstanceKeys.AddKey(promotedSlave.Key)
		if mustPromoteOtherCoMaster {
			log.Debugf("topology_recovery: mustPromoteOtherCoMaster. Verifying that %+v is/can be promoted", *otherCoMasterKey)
			promotedSlave, err = replacePromotedSlaveWithCandidate(failedInstanceKey, promotedSlave, otherCoMasterKey)
		} else {
			// We are allowed to promote any server
			promotedSlave, err = replacePromotedSlaveWithCandidate(failedInstanceKey, promotedSlave, nil)

			if promotedSlave.DataCenter == otherCoMaster.DataCenter &&
				promotedSlave.PhysicalEnvironment == otherCoMaster.PhysicalEnvironment && false {
				// and _still_ we prefer to promote the co-master! They're in same env & DC so no worries about geo issues!
				promotedSlave, err = replacePromotedSlaveWithCandidate(failedInstanceKey, promotedSlave, otherCoMasterKey)
			}
		}
		topologyRecovery.AddError(err)
	}
	if promotedSlave != nil {
		if mustPromoteOtherCoMaster && !promotedSlave.Key.Equals(otherCoMasterKey) {
			topologyRecovery.AddError(log.Errorf("RecoverDeadCoMaster: could not manage to promote other-co-master %+v; was only able to promote %+v; CoMasterRecoveryMustPromoteOtherCoMaster is true, therefore failing", *otherCoMasterKey, promotedSlave.Key))
			promotedSlave = nil
		}
	}
	if promotedSlave != nil {
		topologyRecovery.ParticipatingInstanceKeys.AddKey(promotedSlave.Key)
	}

	// OK, we may have someone promoted. Either this was the other co-master or another slave.
	// Noting down that we DO NOT attempt to set a new co-master topology. We are good with remaining with a single master.
	// I tried solving the "let's promote a slave and create a new co-master setup" but this turns so complex due to various factors.
	// I see this as risky and not worth the questionable benefit.
	// Maybe future me is a smarter person and finds a simple solution. Unlikely. I'm getting dumber.
	//
	// ...
	// Now that we're convinved, take a look at what we can be left with:
	// Say we started with M1<->M2<-S1, with M2 failing, and we promoted S1.
	// We now have M1->S1 (because S1 is promoted), S1->M2 (because that's what it remembers), M2->M1 (because that's what it remembers)
	// !! This is an evil 3-node circle that must be broken.
	// config.Config.ApplyMySQLPromotionAfterMasterFailover, if true, will cause it to break, because we would RESET SLAVE on S1
	// but we want to make sure the circle is broken no matter what.
	// So in the case we promoted not-the-other-co-master, we issue a detach-slave-master-host, which is a reversible operation
	if promotedSlave != nil && !promotedSlave.Key.Equals(otherCoMasterKey) {
		_, err = inst.DetachSlaveMasterHost(&promotedSlave.Key)
		topologyRecovery.AddError(log.Errore(err))
	}

	if promotedSlave != nil && len(lostSlaves) > 0 && config.Config.DetachLostSlavesAfterMasterFailover {
		postponedFunction := func() error {
			log.Debugf("topology_recovery: - RecoverDeadCoMaster: lost %+v slaves during recovery process; detaching them", len(lostSlaves))
			for _, slave := range lostSlaves {
				slave := slave
				inst.DetachSlaveOperation(&slave.Key)
			}
			return nil
		}
		topologyRecovery.AddPostponedFunction(postponedFunction)
	}
	if config.Config.MasterFailoverLostInstancesDowntimeMinutes > 0 {
		postponedFunction := func() error {
			inst.BeginDowntime(failedInstanceKey, inst.GetMaintenanceOwner(), "RecoverDeadCoMaster indicates this instance is lost", config.Config.MasterFailoverLostInstancesDowntimeMinutes*60)
			for _, slave := range lostSlaves {
				slave := slave
				inst.BeginDowntime(&slave.Key, inst.GetMaintenanceOwner(), "RecoverDeadCoMaster indicates this instance is lost", config.Config.MasterFailoverLostInstancesDowntimeMinutes*60)
			}
			return nil
		}
		topologyRecovery.AddPostponedFunction(postponedFunction)
	}

	return promotedSlave, lostSlaves, err
}

// checkAndRecoverDeadCoMaster checks a given analysis, decides whether to take action, and possibly takes action
// Returns true when action was taken.
func checkAndRecoverDeadCoMaster(analysisEntry inst.ReplicationAnalysis, candidateInstanceKey *inst.InstanceKey, forceInstanceRecovery bool, skipProcesses bool) (bool, *TopologyRecovery, error) {
	failedInstanceKey := &analysisEntry.AnalyzedInstanceKey
	if !(forceInstanceRecovery || analysisEntry.ClusterDetails.HasAutomatedMasterRecovery) {
		return false, nil, nil
	}
	topologyRecovery, err := AttemptRecoveryRegistration(&analysisEntry, !forceInstanceRecovery, !forceInstanceRecovery)
	if topologyRecovery == nil {
		log.Debugf("topology_recovery: found an active or recent recovery on %+v. Will not issue another RecoverDeadCoMaster.", analysisEntry.AnalyzedInstanceKey)
		return false, nil, err
	}

	// That's it! We must do recovery!
	recoverDeadCoMasterCounter.Inc(1)
	promotedSlave, lostSlaves, err := RecoverDeadCoMaster(topologyRecovery, skipProcesses)
	ResolveRecovery(topologyRecovery, promotedSlave)
	if promotedSlave == nil {
		inst.AuditOperation("recover-dead-co-master", failedInstanceKey, "Failure: no slave promoted.")
	} else {
		inst.AuditOperation("recover-dead-co-master", failedInstanceKey, fmt.Sprintf("promoted: %+v", promotedSlave.Key))
	}
	topologyRecovery.LostSlaves.AddInstances(lostSlaves)
	if promotedSlave != nil {
		// success
		recoverDeadCoMasterSuccessCounter.Inc(1)

		if config.Config.ApplyMySQLPromotionAfterMasterFailover {
			log.Debugf("topology_recovery: - RecoverDeadMaster: will apply MySQL changes to promoted master")
			inst.SetReadOnly(&promotedSlave.Key, false)
		}
		if !skipProcesses {
			// Execute post intermediate-master-failover processes
			topologyRecovery.SuccessorKey = &promotedSlave.Key
			executeProcesses(config.Config.PostMasterFailoverProcesses, "PostMasterFailoverProcesses", topologyRecovery, false)
		}
	} else {
		recoverDeadCoMasterFailureCounter.Inc(1)
	}
	return true, topologyRecovery, err
}

// checkAndRecoverGenericProblem is a general-purpose recovery function
func checkAndRecoverGenericProblem(analysisEntry inst.ReplicationAnalysis, candidateInstanceKey *inst.InstanceKey, forceInstanceRecovery bool, skipProcesses bool) (bool, *TopologyRecovery, error) {
	return false, nil, nil
}

// Force a re-read of a topology instance; this is done because we need to substantiate a suspicion that we may have a failover
// scenario. we want to speed up reading the complete picture.
func emergentlyReadTopologyInstance(instanceKey *inst.InstanceKey, analysisCode inst.AnalysisCode) {
	if existsInCacheError := emergencyReadTopologyInstanceMap.Add(instanceKey.DisplayString(), true, cache.DefaultExpiration); existsInCacheError != nil {
		// Just recently attempted
		return
	}
	go inst.ExecuteOnTopology(func() {
		inst.ReadTopologyInstance(instanceKey)
		inst.AuditOperation("emergently-read-topology-instance", instanceKey, string(analysisCode))
	})
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

// checkAndExecuteFailureDetectionProcesses tries to register for failure detection and potentially executes
// failure-detection processes.
func checkAndExecuteFailureDetectionProcesses(analysisEntry inst.ReplicationAnalysis, skipProcesses bool) (processesExecutionAttempted bool, err error) {
	if ok, _ := AttemptFailureDetectionRegistration(&analysisEntry); !ok {
		return false, nil
	}
	log.Debugf("topology_recovery: detected %+v failure on %+v", analysisEntry.Analysis, analysisEntry.AnalyzedInstanceKey)
	// Execute on-detection processes
	if skipProcesses {
		return false, nil
	}
	err = executeProcesses(config.Config.OnFailureDetectionProcesses, "OnFailureDetectionProcesses", NewTopologyRecovery(analysisEntry), true)
	return true, err
}

// executeCheckAndRecoverFunction will choose the correct check & recovery function based on analysis.
// It executes the function synchronuously
func executeCheckAndRecoverFunction(analysisEntry inst.ReplicationAnalysis, candidateInstanceKey *inst.InstanceKey, forceInstanceRecovery bool, skipProcesses bool) (recoveryAttempted bool, topologyRecovery *TopologyRecovery, err error) {
	var checkAndRecoverFunction func(analysisEntry inst.ReplicationAnalysis, candidateInstanceKey *inst.InstanceKey, forceInstanceRecovery bool, skipProcesses bool) (recoveryAttempted bool, topologyRecovery *TopologyRecovery, err error) = nil

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
		checkAndRecoverFunction = checkAndRecoverDeadCoMaster
	case inst.DeadCoMasterAndSomeSlaves:
		checkAndRecoverFunction = checkAndRecoverDeadCoMaster
	case inst.DeadMasterAndSlaves:
		go emergentlyReadTopologyInstance(&analysisEntry.AnalyzedInstanceMasterKey, analysisEntry.Analysis)
	case inst.UnreachableMaster:
		go emergentlyReadTopologyInstanceSlaves(&analysisEntry.AnalyzedInstanceKey, analysisEntry.Analysis)
	case inst.AllMasterSlavesNotReplicating:
		go emergentlyReadTopologyInstance(&analysisEntry.AnalyzedInstanceKey, analysisEntry.Analysis)
	case inst.FirstTierSlaveFailingToConnectToMaster:
		go emergentlyReadTopologyInstance(&analysisEntry.AnalyzedInstanceMasterKey, analysisEntry.Analysis)
	case inst.UnreachableMasterWithStaleSlaves:
		checkAndRecoverFunction = checkAndRecoverGenericProblem
	case inst.AllMasterSlavesStale:
		checkAndRecoverFunction = checkAndRecoverGenericProblem
	}

	if checkAndRecoverFunction == nil {
		// Unhandled problem type
		return false, nil, nil
	}
	// we have a recovery function; its execution still depends on filters if not disabled.
	log.Debugf("executeCheckAndRecoverFunction: proceeeding with %+v; skipProcesses: %+v", analysisEntry.AnalyzedInstanceKey, skipProcesses)

	if _, err := checkAndExecuteFailureDetectionProcesses(analysisEntry, skipProcesses); err != nil {
		return false, nil, err
	}

	recoveryAttempted, topologyRecovery, err = checkAndRecoverFunction(analysisEntry, candidateInstanceKey, forceInstanceRecovery, skipProcesses)
	if !recoveryAttempted {
		return recoveryAttempted, topologyRecovery, err
	}
	if topologyRecovery == nil {
		return recoveryAttempted, topologyRecovery, err
	}
	if !skipProcesses {
		if topologyRecovery.SuccessorKey == nil {
			// Execute general unsuccessful post failover processes
			executeProcesses(config.Config.PostUnsuccessfulFailoverProcesses, "PostUnsuccessfulFailoverProcesses", topologyRecovery, false)
		} else {
			// Execute general post failover processes
			executeProcesses(config.Config.PostFailoverProcesses, "PostFailoverProcesses", topologyRecovery, false)
		}
	}
	topologyRecovery.InvokePostponed()
	return recoveryAttempted, topologyRecovery, err
}

// CheckAndRecover is the main entry point for the recovery mechanism
func CheckAndRecover(specificInstance *inst.InstanceKey, candidateInstanceKey *inst.InstanceKey, skipProcesses bool) (recoveryAttempted bool, promotedSlaveKey *inst.InstanceKey, err error) {
	replicationAnalysis, err := inst.GetReplicationAnalysis("", true, true)
	if err != nil {
		return false, nil, log.Errore(err)
	}
	if *config.RuntimeCLIFlags.Noop {
		log.Debugf("--noop provided; will not execute processes")
		skipProcesses = true
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

		if specificInstance != nil {
			// force mode. Keep it synchronuous
			var topologyRecovery *TopologyRecovery
			recoveryAttempted, topologyRecovery, err = executeCheckAndRecoverFunction(analysisEntry, candidateInstanceKey, true, skipProcesses)
			if topologyRecovery != nil {
				promotedSlaveKey = topologyRecovery.SuccessorKey
			}
		} else {
			go executeCheckAndRecoverFunction(analysisEntry, candidateInstanceKey, false, skipProcesses)
		}
	}
	return recoveryAttempted, promotedSlaveKey, err
}
