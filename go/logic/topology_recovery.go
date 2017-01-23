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
	goos "os"
	"sort"
	"strings"
	"time"

	"github.com/github/orchestrator/go/attributes"
	"github.com/github/orchestrator/go/config"
	"github.com/github/orchestrator/go/inst"
	"github.com/github/orchestrator/go/os"
	"github.com/github/orchestrator/go/process"
	"github.com/outbrain/golib/log"
	"github.com/patrickmn/go-cache"
	"github.com/rcrowley/go-metrics"
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
	SuccessorAlias            string
	IsActive                  bool
	IsSuccessful              bool
	LostReplicas              inst.InstanceKeyMap
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
	topologyRecovery.LostReplicas = *inst.NewInstanceKeyMap()
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

var emergencyReadTopologyInstanceMap = cache.New(time.Duration(config.Config.InstancePollSeconds)*time.Second, time.Second)

// InstancesByCountReplicas sorts instances by umber of replicas, descending
type InstancesByCountReplicas [](*inst.Instance)

func (this InstancesByCountReplicas) Len() int      { return len(this) }
func (this InstancesByCountReplicas) Swap(i, j int) { this[i], this[j] = this[j], this[i] }
func (this InstancesByCountReplicas) Less(i, j int) bool {
	if len(this[i].SlaveHosts) == len(this[j].SlaveHosts) {
		// Secondary sorting: prefer more advanced replicas
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
var recoverUnreachableMasterWithStaleSlavesCounter = metrics.NewCounter()
var recoverUnreachableMasterWithStaleSlavesSuccessCounter = metrics.NewCounter()
var recoverUnreachableMasterWithStaleSlavesFailureCounter = metrics.NewCounter()

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
	metrics.Register("recover.unreach_master_stale_slaves.start", recoverUnreachableMasterWithStaleSlavesCounter)
	metrics.Register("recover.unreach_master_stale_slaves.success", recoverUnreachableMasterWithStaleSlavesSuccessCounter)
	metrics.Register("recover.unreach_master_stale_slaves.fail", recoverUnreachableMasterWithStaleSlavesFailureCounter)
}

// replaceCommandPlaceholders replaces agreed-upon placeholders with analysis data
func replaceCommandPlaceholders(command string, topologyRecovery *TopologyRecovery) string {
	analysisEntry := &topologyRecovery.AnalysisEntry
	command = strings.Replace(command, "{failureType}", string(analysisEntry.Analysis), -1)
	command = strings.Replace(command, "{failureDescription}", analysisEntry.Description, -1)
	command = strings.Replace(command, "{failedHost}", analysisEntry.AnalyzedInstanceKey.Hostname, -1)
	command = strings.Replace(command, "{failedPort}", fmt.Sprintf("%d", analysisEntry.AnalyzedInstanceKey.Port), -1)
	command = strings.Replace(command, "{failureCluster}", analysisEntry.ClusterDetails.ClusterName, -1)
	command = strings.Replace(command, "{failureClusterAlias}", analysisEntry.ClusterDetails.ClusterAlias, -1)
	command = strings.Replace(command, "{failureClusterDomain}", analysisEntry.ClusterDetails.ClusterDomain, -1)
	command = strings.Replace(command, "{countSlaves}", fmt.Sprintf("%d", analysisEntry.CountReplicas), -1)
	command = strings.Replace(command, "{countReplicas}", fmt.Sprintf("%d", analysisEntry.CountReplicas), -1)
	command = strings.Replace(command, "{isDowntimed}", fmt.Sprint(analysisEntry.IsDowntimed), -1)
	command = strings.Replace(command, "{autoMasterRecovery}", fmt.Sprint(analysisEntry.ClusterDetails.HasAutomatedMasterRecovery), -1)
	command = strings.Replace(command, "{autoIntermediateMasterRecovery}", fmt.Sprint(analysisEntry.ClusterDetails.HasAutomatedIntermediateMasterRecovery), -1)
	command = strings.Replace(command, "{orchestratorHost}", process.ThisHostname, -1)

	command = strings.Replace(command, "{isSuccessful}", fmt.Sprint(topologyRecovery.SuccessorKey != nil), -1)
	if topologyRecovery.SuccessorKey != nil {
		command = strings.Replace(command, "{successorHost}", topologyRecovery.SuccessorKey.Hostname, -1)
		command = strings.Replace(command, "{successorPort}", fmt.Sprintf("%d", topologyRecovery.SuccessorKey.Port), -1)
		// As long as SucesssorKey != nil, we replace {successorAlias}.
		// If SucessorAlias is "", it's fine. We'll replace {successorAlias} with "".
		command = strings.Replace(command, "{successorAlias}", topologyRecovery.SuccessorAlias, -1)
	}

	command = strings.Replace(command, "{lostSlaves}", topologyRecovery.LostReplicas.ToCommaDelimitedList(), -1)
	command = strings.Replace(command, "{lostReplicas}", topologyRecovery.LostReplicas.ToCommaDelimitedList(), -1)
	command = strings.Replace(command, "{slaveHosts}", analysisEntry.SlaveHosts.ToCommaDelimitedList(), -1)
	command = strings.Replace(command, "{replicaHosts}", analysisEntry.SlaveHosts.ToCommaDelimitedList(), -1)

	return command
}

// replaceCommandPlaceholders replaces agreed-upon placeholders with analysis data
func applyEnvironmentVariables(topologyRecovery *TopologyRecovery) []string {
	analysisEntry := &topologyRecovery.AnalysisEntry
	env := goos.Environ()
	env = append(env, fmt.Sprintf("ORC_FAILURE_TYPE=%s", string(analysisEntry.Analysis)))
	env = append(env, fmt.Sprintf("ORC_FAILURE_DESCRIPTION=%s", analysisEntry.Description))
	env = append(env, fmt.Sprintf("ORC_FAILED_HOST=%s", analysisEntry.AnalyzedInstanceKey.Hostname))
	env = append(env, fmt.Sprintf("ORC_FAILED_PORT=%d", analysisEntry.AnalyzedInstanceKey.Port))
	env = append(env, fmt.Sprintf("ORC_FAILURE_CLUSTER=%s", analysisEntry.ClusterDetails.ClusterName))
	env = append(env, fmt.Sprintf("ORC_FAILURE_CLUSTER_ALIAS=%s", analysisEntry.ClusterDetails.ClusterAlias))
	env = append(env, fmt.Sprintf("ORC_FAILURE_CLUSTER_DOMAIN=%s", analysisEntry.ClusterDetails.ClusterDomain))
	env = append(env, fmt.Sprintf("ORC_COUNT_REPLICAS=%s", analysisEntry.CountReplicas))
	env = append(env, fmt.Sprintf("ORC_IS_DOWNTIMED=%s", analysisEntry.IsDowntimed))
	env = append(env, fmt.Sprintf("ORC_AUTO_MASTER_RECOVERY=%s", analysisEntry.ClusterDetails.HasAutomatedMasterRecovery))
	env = append(env, fmt.Sprintf("ORC_AUTO_INTERMEDIATE_MASTER_RECOVERY=%s", analysisEntry.ClusterDetails.HasAutomatedIntermediateMasterRecovery))
	env = append(env, fmt.Sprintf("ORC_ORCHESTRATOR_HOST=%s", process.ThisHostname))
	env = append(env, fmt.Sprintf("ORC_IS_SUCCESSFUL=%s", (topologyRecovery.SuccessorKey != nil)))
	env = append(env, fmt.Sprintf("ORC_LOST_REPLICAS=%s", topologyRecovery.LostReplicas.ToCommaDelimitedList()))
	env = append(env, fmt.Sprintf("ORC_REPLICA_HOSTS=%s", analysisEntry.SlaveHosts.ToCommaDelimitedList()))

	if topologyRecovery.SuccessorKey != nil {
		env = append(env, fmt.Sprintf("ORC_SUCCESSOR_HOST=%s", topologyRecovery.SuccessorKey.Hostname))
		env = append(env, fmt.Sprintf("ORC_SUCCESSOR_PORT=%d", topologyRecovery.SuccessorKey.Port))
		// As long as SucesssorKey != nil, we replace {successorAlias}.
		// If SucessorAlias is "", it's fine. We'll replace {successorAlias} with "".
		env = append(env, fmt.Sprintf("ORC_SUCCESSOR_ALIAS=%s", topologyRecovery.SuccessorAlias))
	}

	return env
}

// executeProcesses executes a list of processes
func executeProcesses(processes []string, description string, topologyRecovery *TopologyRecovery, failOnError bool) error {
	var err error
	for _, command := range processes {
		command := replaceCommandPlaceholders(command, topologyRecovery)
		env := applyEnvironmentVariables(topologyRecovery)

		if cmdErr := os.CommandRun(command, env); cmdErr == nil {
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

func recoverDeadMasterInBinlogServerTopology(topologyRecovery *TopologyRecovery) (promotedReplica *inst.Instance, err error) {
	failedMasterKey := &topologyRecovery.AnalysisEntry.AnalyzedInstanceKey

	var promotedBinlogServer *inst.Instance

	_, promotedBinlogServer, err = inst.RegroupReplicasBinlogServers(failedMasterKey, true)
	if err != nil {
		return nil, log.Errore(err)
	}
	promotedBinlogServer, err = inst.StopSlave(&promotedBinlogServer.Key)
	if err != nil {
		return promotedReplica, log.Errore(err)
	}
	// Find candidate replica
	promotedReplica, err = inst.GetCandidateReplicaOfBinlogServerTopology(&promotedBinlogServer.Key)
	if err != nil {
		return promotedReplica, log.Errore(err)
	}
	// Align it with binlog server coordinates
	promotedReplica, err = inst.StopSlave(&promotedReplica.Key)
	if err != nil {
		return promotedReplica, log.Errore(err)
	}
	promotedReplica, err = inst.StartSlaveUntilMasterCoordinates(&promotedReplica.Key, &promotedBinlogServer.ExecBinlogCoordinates)
	if err != nil {
		return promotedReplica, log.Errore(err)
	}
	promotedReplica, err = inst.StopSlave(&promotedReplica.Key)
	if err != nil {
		return promotedReplica, log.Errore(err)
	}
	// Detach, flush binary logs forward
	promotedReplica, err = inst.ResetSlave(&promotedReplica.Key)
	if err != nil {
		return promotedReplica, log.Errore(err)
	}
	promotedReplica, err = inst.FlushBinaryLogsTo(&promotedReplica.Key, promotedBinlogServer.ExecBinlogCoordinates.LogFile)
	if err != nil {
		return promotedReplica, log.Errore(err)
	}
	promotedReplica, err = inst.FlushBinaryLogs(&promotedReplica.Key, 1)
	if err != nil {
		return promotedReplica, log.Errore(err)
	}
	promotedReplica, err = inst.PurgeBinaryLogsToCurrent(&promotedReplica.Key)
	if err != nil {
		return promotedReplica, log.Errore(err)
	}
	// Reconnect binlog servers to promoted replica (now master):
	promotedBinlogServer, err = inst.SkipToNextBinaryLog(&promotedBinlogServer.Key)
	if err != nil {
		return promotedReplica, log.Errore(err)
	}
	promotedBinlogServer, err = inst.Repoint(&promotedBinlogServer.Key, &promotedReplica.Key, inst.GTIDHintDeny)
	if err != nil {
		return nil, log.Errore(err)
	}

	func() {
		// Move binlog server replicas up to replicate from master.
		// This can only be done once a BLS has skipped to the next binlog
		// We postpone this operation. The master is already promoted and we're happy.
		binlogServerReplicas, err := inst.ReadBinlogServerReplicaInstances(&promotedBinlogServer.Key)
		if err != nil {
			return
		}
		maxBinlogServersToPromote := 3
		for i, binlogServerReplica := range binlogServerReplicas {
			binlogServerReplica := binlogServerReplica
			if i >= maxBinlogServersToPromote {
				return
			}
			postponedFunction := func() error {
				binlogServerReplica, err := inst.StopSlave(&binlogServerReplica.Key)
				if err != nil {
					return err
				}
				// Make sure the BLS has the "next binlog" -- the one the master flushed & purged to. Otherwise the BLS
				// will request a binlog the master does not have
				if binlogServerReplica.ExecBinlogCoordinates.SmallerThan(&promotedBinlogServer.ExecBinlogCoordinates) {
					binlogServerReplica, err = inst.StartSlaveUntilMasterCoordinates(&binlogServerReplica.Key, &promotedBinlogServer.ExecBinlogCoordinates)
					if err != nil {
						return err
					}
				}
				_, err = inst.Repoint(&binlogServerReplica.Key, &promotedReplica.Key, inst.GTIDHintDeny)
				return err
			}
			topologyRecovery.AddPostponedFunction(postponedFunction)
		}
	}()

	return promotedReplica, err
}

// RecoverDeadMaster recovers a dead master, complete logic inside
func RecoverDeadMaster(topologyRecovery *TopologyRecovery, skipProcesses bool) (promotedReplica *inst.Instance, lostReplicas [](*inst.Instance), err error) {
	analysisEntry := &topologyRecovery.AnalysisEntry
	failedInstanceKey := &analysisEntry.AnalyzedInstanceKey
	var cannotReplicateReplicas [](*inst.Instance)

	inst.AuditOperation("recover-dead-master", failedInstanceKey, "problem found; will recover")
	if !skipProcesses {
		if err := executeProcesses(config.Config.PreFailoverProcesses, "PreFailoverProcesses", topologyRecovery, true); err != nil {
			return nil, lostReplicas, topologyRecovery.AddError(err)
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
			lostReplicas, _, cannotReplicateReplicas, promotedReplica, err = inst.RegroupReplicasGTID(failedInstanceKey, true, nil)
		}
	case MasterRecoveryPseudoGTID:
		{
			lostReplicas, _, _, cannotReplicateReplicas, promotedReplica, err = inst.RegroupReplicasPseudoGTIDIncludingSubReplicasOfBinlogServers(failedInstanceKey, true, nil, &topologyRecovery.PostponedFunctionsContainer)
		}
	case MasterRecoveryBinlogServer:
		{
			promotedReplica, err = recoverDeadMasterInBinlogServerTopology(topologyRecovery)
		}
	}
	topologyRecovery.AddError(err)
	lostReplicas = append(lostReplicas, cannotReplicateReplicas...)

	if promotedReplica != nil && len(lostReplicas) > 0 && config.Config.DetachLostReplicasAfterMasterFailover {
		postponedFunction := func() error {
			log.Debugf("topology_recovery: - RecoverDeadMaster: lost %+v replicas during recovery process; detaching them", len(lostReplicas))
			for _, replica := range lostReplicas {
				replica := replica
				inst.DetachReplicaOperation(&replica.Key)
			}
			return nil
		}
		topologyRecovery.AddPostponedFunction(postponedFunction)
	}
	if config.Config.MasterFailoverLostInstancesDowntimeMinutes > 0 {
		postponedFunction := func() error {
			inst.BeginDowntime(failedInstanceKey, inst.GetMaintenanceOwner(), inst.DowntimeLostInRecoveryMessage, config.Config.MasterFailoverLostInstancesDowntimeMinutes*60)
			for _, replica := range lostReplicas {
				replica := replica
				inst.BeginDowntime(&replica.Key, inst.GetMaintenanceOwner(), inst.DowntimeLostInRecoveryMessage, config.Config.MasterFailoverLostInstancesDowntimeMinutes*60)
			}
			return nil
		}
		topologyRecovery.AddPostponedFunction(postponedFunction)
	}

	if promotedReplica == nil {
		inst.AuditOperation("recover-dead-master", failedInstanceKey, "Failure: no replica promoted.")
	} else {
		inst.AuditOperation("recover-dead-master", failedInstanceKey, fmt.Sprintf("promoted replica: %+v", promotedReplica.Key))
	}
	return promotedReplica, lostReplicas, err
}

// replacePromotedReplicaWithCandidate is called after an intermediate master has died and been replaced by some promotedReplica.
// But, is there an even better replica to promote?
// if candidateInstanceKey is given, then it is forced to be promoted over the promotedReplica
// Otherwise, search for the best to promote!
func replacePromotedReplicaWithCandidate(deadInstanceKey *inst.InstanceKey, promotedReplica *inst.Instance, candidateInstanceKey *inst.InstanceKey) (*inst.Instance, error) {
	candidateReplicas, _ := inst.ReadClusterCandidateInstances(promotedReplica.ClusterName)
	// So we've already promoted a replica.
	// However, can we improve on our choice? Are there any replicas marked with "is_candidate"?
	// Maybe we actually promoted such a replica. Does that mean we should keep it?
	// The current logic is:
	// - 1. we prefer to promote a "is_candidate" which is in the same DC & env as the dead intermediate master (or do nothing if the promtoed replica is such one)
	// - 2. we prefer to promote a "is_candidate" which is in the same DC & env as the promoted replica (or do nothing if the promtoed replica is such one)
	// - 3. keep to current choice
	log.Infof("topology_recovery: checking if should replace promoted replica with a better candidate")
	if candidateInstanceKey == nil {
		if deadInstance, _, err := inst.ReadInstance(deadInstanceKey); err == nil && deadInstance != nil {
			for _, candidateReplica := range candidateReplicas {
				if promotedReplica.Key.Equals(&candidateReplica.Key) &&
					promotedReplica.DataCenter == deadInstance.DataCenter &&
					promotedReplica.PhysicalEnvironment == deadInstance.PhysicalEnvironment {
					// Seems like we promoted a candidate in the same DC & ENV as dead IM! Ideal! We're happy!
					log.Infof("topology_recovery: promoted replica %+v is the ideal candidate", promotedReplica.Key)
					return promotedReplica, nil
				}
			}
		}
	}
	// We didn't pick the ideal candidate; let's see if we can replace with a candidate from same DC and ENV
	if candidateInstanceKey == nil {
		// Try a candidate replica that is in same DC & env as the dead instance
		if deadInstance, _, err := inst.ReadInstance(deadInstanceKey); err == nil && deadInstance != nil {
			for _, candidateReplica := range candidateReplicas {
				if candidateReplica.DataCenter == deadInstance.DataCenter &&
					candidateReplica.PhysicalEnvironment == deadInstance.PhysicalEnvironment &&
					candidateReplica.MasterKey.Equals(&promotedReplica.Key) {
					// This would make a great candidate
					candidateInstanceKey = &candidateReplica.Key
					log.Debugf("topology_recovery: no candidate was offered for %+v but orchestrator picks %+v as candidate replacement, based on being in same DC & env as failed instance", promotedReplica.Key, candidateReplica.Key)
				}
			}
		}
	}
	if candidateInstanceKey == nil {
		// We cannot find a candidate in same DC and ENV as dead master
		for _, candidateReplica := range candidateReplicas {
			if promotedReplica.Key.Equals(&candidateReplica.Key) {
				// Seems like we promoted a candidate replica (though not in same DC and ENV as dead master). Good enough.
				// No further action required.
				log.Infof("topology_recovery: promoted replica %+v is a good candidate", promotedReplica.Key)
				return promotedReplica, nil
			}
		}
	}
	// Still nothing?
	if candidateInstanceKey == nil {
		// Try a candidate replica that is in same DC & env as the promoted replica (our promoted replica is not an "is_candidate")
		for _, candidateReplica := range candidateReplicas {
			if promotedReplica.DataCenter == candidateReplica.DataCenter &&
				promotedReplica.PhysicalEnvironment == candidateReplica.PhysicalEnvironment &&
				candidateReplica.MasterKey.Equals(&promotedReplica.Key) {
				// OK, better than nothing
				candidateInstanceKey = &candidateReplica.Key
				log.Debugf("topology_recovery: no candidate was offered for %+v but orchestrator picks %+v as candidate replacement, based on being in same DC & env as promoted instance", promotedReplica.Key, candidateReplica.Key)
			}
		}
	}

	// So do we have a candidate?
	if candidateInstanceKey == nil {
		// Found nothing. Stick with promoted replica
		return promotedReplica, nil
	}
	if promotedReplica.Key.Equals(candidateInstanceKey) {
		// Sanity. It IS the candidate, nothing to promote...
		return promotedReplica, nil
	}

	// Try and promote suggested candidate, if applicable and possible
	log.Debugf("topology_recovery: promoted instance %+v is not the suggested candidate %+v. Will see what can be done", promotedReplica.Key, *candidateInstanceKey)

	candidateInstance, _, err := inst.ReadInstance(candidateInstanceKey)
	if err != nil {
		return promotedReplica, log.Errore(err)
	}

	if candidateInstance.MasterKey.Equals(&promotedReplica.Key) {
		log.Debugf("topology_recovery: suggested candidate %+v is replica of promoted instance %+v. Will try and take its master", *candidateInstanceKey, promotedReplica.Key)
		candidateInstance, err = inst.TakeMaster(&candidateInstance.Key)
		if err != nil {
			return promotedReplica, log.Errore(err)
		}
		log.Debugf("topology_recovery: success promoting %+v over %+v", *candidateInstanceKey, promotedReplica.Key)
		return candidateInstance, nil
	}

	log.Debugf("topology_recovery: could not manage to promoted suggested candidate %+v", *candidateInstanceKey)
	return promotedReplica, nil
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
	promotedReplica, lostReplicas, err := RecoverDeadMaster(topologyRecovery, skipProcesses)
	topologyRecovery.LostReplicas.AddInstances(lostReplicas)

	if promotedReplica != nil {
		promotedReplica, err = replacePromotedReplicaWithCandidate(&analysisEntry.AnalyzedInstanceKey, promotedReplica, candidateInstanceKey)
		topologyRecovery.AddError(err)
	}
	// And this is the end; whether successful or not, we're done.
	ResolveRecovery(topologyRecovery, promotedReplica)
	if promotedReplica != nil {
		// Success!
		recoverDeadMasterSuccessCounter.Inc(1)

		if config.Config.ApplyMySQLPromotionAfterMasterFailover {
			log.Debugf("topology_recovery: - RecoverDeadMaster: will apply MySQL changes to promoted master")
			inst.ResetSlaveOperation(&promotedReplica.Key)
			inst.SetReadOnly(&promotedReplica.Key, false)
		}
		if !skipProcesses {
			// Execute post master-failover processes
			executeProcesses(config.Config.PostMasterFailoverProcesses, "PostMasterFailoverProcesses", topologyRecovery, false)
		}

		if config.Config.MasterFailoverDetachReplicaMasterHost {
			postponedFunction := func() error {
				log.Debugf("topology_recovery: - RecoverDeadMaster: detaching master host on promoted master")
				inst.DetachReplicaMasterHost(&promotedReplica.Key)
				return nil
			}
			topologyRecovery.AddPostponedFunction(postponedFunction)
		}
		postponedFunction := func() error {
			log.Debugf("topology_recovery: - RecoverDeadMaster: updating cluster_alias")
			inst.ReplaceAliasClusterName(analysisEntry.AnalyzedInstanceKey.StringCode(), promotedReplica.Key.StringCode())
			return nil
		}
		topologyRecovery.AddPostponedFunction(postponedFunction)

		attributes.SetGeneralAttribute(analysisEntry.ClusterDetails.ClusterDomain, promotedReplica.Key.StringCode())
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
	if !sibling.ReplicaRunning() {
		return false
	}
	if !sibling.IsLastCheckValid {
		return false
	}
	return true
}

// isValidAsCandidateSiblingOfIntermediateMaster checks to see that the given sibling is capable to take over instance's replicas
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
		// When failed IM is binlog server, its sibling is still valid, but we catually prefer to just repoint the replica up -- simplest!
		return false
	}
	if sibling.ExecBinlogCoordinates.SmallerThan(&intermediateMasterInstance.ExecBinlogCoordinates) {
		return false
	}
	return true
}

// GetCandidateSiblingOfIntermediateMaster chooses the best sibling of a dead intermediate master
// to whom the IM's replicas can be moved.
func GetCandidateSiblingOfIntermediateMaster(intermediateMasterInstance *inst.Instance) (*inst.Instance, error) {

	siblings, err := inst.ReadReplicaInstances(&intermediateMasterInstance.MasterKey)
	if err != nil {
		return nil, err
	}
	if len(siblings) <= 1 {
		return nil, log.Errorf("topology_recovery: no siblings found for %+v", intermediateMasterInstance.Key)
	}

	sort.Sort(sort.Reverse(InstancesByCountReplicas(siblings)))

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
	relocateReplicasToCandidateSibling := func() {
		if candidateSiblingOfIntermediateMaster == nil {
			return
		}
		// We have a candidate
		log.Debugf("topology_recovery: - RecoverDeadIntermediateMaster: will attempt a candidate intermediate master: %+v", candidateSiblingOfIntermediateMaster.Key)
		relocatedReplicas, candidateSibling, err, errs := inst.RelocateReplicas(failedInstanceKey, &candidateSiblingOfIntermediateMaster.Key, "")
		topologyRecovery.AddErrors(errs)
		topologyRecovery.ParticipatingInstanceKeys.AddKey(candidateSiblingOfIntermediateMaster.Key)

		if len(relocatedReplicas) == 0 {
			log.Debugf("topology_recovery: - RecoverDeadIntermediateMaster: failed to move any replica to candidate intermediate master (%+v)", candidateSibling.Key)
			return
		}
		if err != nil || len(errs) > 0 {
			log.Debugf("topology_recovery: - RecoverDeadIntermediateMaster: move to candidate intermediate master (%+v) did not complete: %+v", candidateSibling.Key, err)
			return
		}
		if err == nil {
			recoveryResolved = true
			successorInstance = candidateSibling

			inst.AuditOperation("recover-dead-intermediate-master", failedInstanceKey, fmt.Sprintf("Relocated %d replicas under candidate sibling: %+v; %d errors: %+v", len(relocatedReplicas), candidateSibling.Key, len(errs), errs))
		}
	}
	// Plan A: find a replacement intermediate master in same Data Center
	if candidateSiblingOfIntermediateMaster != nil && candidateSiblingOfIntermediateMaster.DataCenter == intermediateMasterInstance.DataCenter {
		relocateReplicasToCandidateSibling()
	}
	if !recoveryResolved {
		log.Debugf("topology_recovery: - RecoverDeadIntermediateMaster: will next attempt regrouping of replicas")
		// Plan B: regroup (we wish to reduce cross-DC replication streams)
		_, _, _, _, regroupPromotedReplica, err := inst.RegroupReplicas(failedInstanceKey, true, nil, nil)
		if err != nil {
			topologyRecovery.AddError(err)
			log.Debugf("topology_recovery: - RecoverDeadIntermediateMaster: regroup failed on: %+v", err)
		}
		if regroupPromotedReplica != nil {
			topologyRecovery.ParticipatingInstanceKeys.AddKey(regroupPromotedReplica.Key)
		}
		// Plan C: try replacement intermediate master in other DC...
		if candidateSiblingOfIntermediateMaster != nil && candidateSiblingOfIntermediateMaster.DataCenter != intermediateMasterInstance.DataCenter {
			log.Debugf("topology_recovery: - RecoverDeadIntermediateMaster: will next attempt relocating to another DC server")
			relocateReplicasToCandidateSibling()
		}
	}
	if !recoveryResolved {
		// Do we still have leftovers? some replicas couldn't move? Couldn't regroup? Only left with regroup's resulting leader?
		// nothing moved?
		// We don't care much if regroup made it or not. We prefer that it made it, in whcih case we only need to relocate up
		// one replica, but the operation is still valid if regroup partially/completely failed. We just promote anything
		// not regrouped.
		// So, match up all that's left, plan D
		log.Debugf("topology_recovery: - RecoverDeadIntermediateMaster: will next attempt to relocate up from %+v", *failedInstanceKey)

		var errs []error
		var relocatedReplicas [](*inst.Instance)
		relocatedReplicas, successorInstance, err, errs = inst.RelocateReplicas(failedInstanceKey, &analysisEntry.AnalyzedInstanceMasterKey, "")
		topologyRecovery.AddErrors(errs)
		topologyRecovery.ParticipatingInstanceKeys.AddKey(analysisEntry.AnalyzedInstanceMasterKey)

		if len(relocatedReplicas) > 0 {
			recoveryResolved = true
			inst.AuditOperation("recover-dead-intermediate-master", failedInstanceKey, fmt.Sprintf("Relocated replicas under: %+v %d errors: %+v", successorInstance.Key, len(errs), errs))
		} else {
			err = log.Errorf("topology_recovery: RecoverDeadIntermediateMaster failed to match up any replica from %+v", *failedInstanceKey)
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
	promotedReplica, err := RecoverDeadIntermediateMaster(topologyRecovery, skipProcesses)
	if promotedReplica != nil {
		// success
		recoverDeadIntermediateMasterSuccessCounter.Inc(1)

		if !skipProcesses {
			// Execute post intermediate-master-failover processes
			topologyRecovery.SuccessorKey = &promotedReplica.Key
			topologyRecovery.SuccessorAlias = promotedReplica.InstanceAlias
			executeProcesses(config.Config.PostIntermediateMasterFailoverProcesses, "PostIntermediateMasterFailoverProcesses", topologyRecovery, false)
		}
	} else {
		recoverDeadIntermediateMasterFailureCounter.Inc(1)
	}
	return true, topologyRecovery, err
}

// RecoverDeadCoMaster recovers a dead co-master, complete logic inside
func RecoverDeadCoMaster(topologyRecovery *TopologyRecovery, skipProcesses bool) (promotedReplica *inst.Instance, lostReplicas [](*inst.Instance), err error) {
	analysisEntry := &topologyRecovery.AnalysisEntry
	failedInstanceKey := &analysisEntry.AnalyzedInstanceKey
	otherCoMasterKey := &analysisEntry.AnalyzedInstanceMasterKey
	otherCoMaster, found, _ := inst.ReadInstance(otherCoMasterKey)
	if otherCoMaster == nil || !found {
		return nil, lostReplicas, topologyRecovery.AddError(log.Errorf("RecoverDeadCoMaster: could not read info for co-master %+v of %+v", *otherCoMasterKey, *failedInstanceKey))
	}
	inst.AuditOperation("recover-dead-co-master", failedInstanceKey, "problem found; will recover")
	if !skipProcesses {
		if err := executeProcesses(config.Config.PreFailoverProcesses, "PreFailoverProcesses", topologyRecovery, true); err != nil {
			return nil, lostReplicas, topologyRecovery.AddError(err)
		}
	}

	log.Debugf("topology_recovery: RecoverDeadCoMaster: will recover %+v", *failedInstanceKey)

	var coMasterRecoveryType MasterRecoveryType = MasterRecoveryPseudoGTID
	if analysisEntry.OracleGTIDImmediateTopology || analysisEntry.MariaDBGTIDImmediateTopology {
		coMasterRecoveryType = MasterRecoveryGTID
	}

	log.Debugf("topology_recovery: RecoverDeadCoMaster: coMasterRecoveryType=%+v", coMasterRecoveryType)

	var cannotReplicateReplicas [](*inst.Instance)
	switch coMasterRecoveryType {
	case MasterRecoveryGTID:
		{
			lostReplicas, _, cannotReplicateReplicas, promotedReplica, err = inst.RegroupReplicasGTID(failedInstanceKey, true, nil)
		}
	case MasterRecoveryPseudoGTID:
		{
			lostReplicas, _, _, cannotReplicateReplicas, promotedReplica, err = inst.RegroupReplicasPseudoGTIDIncludingSubReplicasOfBinlogServers(failedInstanceKey, true, nil, &topologyRecovery.PostponedFunctionsContainer)
		}
	}
	topologyRecovery.AddError(err)
	lostReplicas = append(lostReplicas, cannotReplicateReplicas...)

	mustPromoteOtherCoMaster := config.Config.CoMasterRecoveryMustPromoteOtherCoMaster
	if !otherCoMaster.ReadOnly {
		log.Debugf("topology_recovery: RecoverDeadCoMaster: other co-master %+v is writeable hence has to be promoted", otherCoMaster.Key)
		mustPromoteOtherCoMaster = true
	}
	log.Debugf("topology_recovery: RecoverDeadCoMaster: mustPromoteOtherCoMaster? %+v", mustPromoteOtherCoMaster)

	if promotedReplica != nil {
		topologyRecovery.ParticipatingInstanceKeys.AddKey(promotedReplica.Key)
		if mustPromoteOtherCoMaster {
			log.Debugf("topology_recovery: mustPromoteOtherCoMaster. Verifying that %+v is/can be promoted", *otherCoMasterKey)
			promotedReplica, err = replacePromotedReplicaWithCandidate(failedInstanceKey, promotedReplica, otherCoMasterKey)
		} else {
			// We are allowed to promote any server
			promotedReplica, err = replacePromotedReplicaWithCandidate(failedInstanceKey, promotedReplica, nil)

			if promotedReplica.DataCenter == otherCoMaster.DataCenter &&
				promotedReplica.PhysicalEnvironment == otherCoMaster.PhysicalEnvironment && false {
				// and _still_ we prefer to promote the co-master! They're in same env & DC so no worries about geo issues!
				promotedReplica, err = replacePromotedReplicaWithCandidate(failedInstanceKey, promotedReplica, otherCoMasterKey)
			}
		}
		topologyRecovery.AddError(err)
	}
	if promotedReplica != nil {
		if mustPromoteOtherCoMaster && !promotedReplica.Key.Equals(otherCoMasterKey) {
			topologyRecovery.AddError(log.Errorf("RecoverDeadCoMaster: could not manage to promote other-co-master %+v; was only able to promote %+v; CoMasterRecoveryMustPromoteOtherCoMaster is true, therefore failing", *otherCoMasterKey, promotedReplica.Key))
			promotedReplica = nil
		}
	}
	if promotedReplica != nil {
		topologyRecovery.ParticipatingInstanceKeys.AddKey(promotedReplica.Key)
	}

	// OK, we may have someone promoted. Either this was the other co-master or another replica.
	// Noting down that we DO NOT attempt to set a new co-master topology. We are good with remaining with a single master.
	// I tried solving the "let's promote a replica and create a new co-master setup" but this turns so complex due to various factors.
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
	// So in the case we promoted not-the-other-co-master, we issue a detach-replica-master-host, which is a reversible operation
	if promotedReplica != nil && !promotedReplica.Key.Equals(otherCoMasterKey) {
		_, err = inst.DetachReplicaMasterHost(&promotedReplica.Key)
		topologyRecovery.AddError(log.Errore(err))
	}

	if promotedReplica != nil && len(lostReplicas) > 0 && config.Config.DetachLostReplicasAfterMasterFailover {
		postponedFunction := func() error {
			log.Debugf("topology_recovery: - RecoverDeadCoMaster: lost %+v replicas during recovery process; detaching them", len(lostReplicas))
			for _, replica := range lostReplicas {
				replica := replica
				inst.DetachReplicaOperation(&replica.Key)
			}
			return nil
		}
		topologyRecovery.AddPostponedFunction(postponedFunction)
	}
	if config.Config.MasterFailoverLostInstancesDowntimeMinutes > 0 {
		postponedFunction := func() error {
			inst.BeginDowntime(failedInstanceKey, inst.GetMaintenanceOwner(), inst.DowntimeLostInRecoveryMessage, config.Config.MasterFailoverLostInstancesDowntimeMinutes*60)
			for _, replica := range lostReplicas {
				replica := replica
				inst.BeginDowntime(&replica.Key, inst.GetMaintenanceOwner(), inst.DowntimeLostInRecoveryMessage, config.Config.MasterFailoverLostInstancesDowntimeMinutes*60)
			}
			return nil
		}
		topologyRecovery.AddPostponedFunction(postponedFunction)
	}

	return promotedReplica, lostReplicas, err
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
	promotedReplica, lostReplicas, err := RecoverDeadCoMaster(topologyRecovery, skipProcesses)
	ResolveRecovery(topologyRecovery, promotedReplica)
	if promotedReplica == nil {
		inst.AuditOperation("recover-dead-co-master", failedInstanceKey, "Failure: no replica promoted.")
	} else {
		inst.AuditOperation("recover-dead-co-master", failedInstanceKey, fmt.Sprintf("promoted: %+v", promotedReplica.Key))
	}
	topologyRecovery.LostReplicas.AddInstances(lostReplicas)
	if promotedReplica != nil {
		// success
		recoverDeadCoMasterSuccessCounter.Inc(1)

		if config.Config.ApplyMySQLPromotionAfterMasterFailover {
			log.Debugf("topology_recovery: - RecoverDeadMaster: will apply MySQL changes to promoted master")
			inst.SetReadOnly(&promotedReplica.Key, false)
		}
		if !skipProcesses {
			// Execute post intermediate-master-failover processes
			topologyRecovery.SuccessorKey = &promotedReplica.Key
			topologyRecovery.SuccessorAlias = promotedReplica.InstanceAlias
			executeProcesses(config.Config.PostMasterFailoverProcesses, "PostMasterFailoverProcesses", topologyRecovery, false)
		}
	} else {
		recoverDeadCoMasterFailureCounter.Inc(1)
	}
	return true, topologyRecovery, err
}

// checkAndRecoverUnreachableMasterWithStaleSlaves executes an external process. No other action is taken.
// Returns false.
func checkAndRecoverUnreachableMasterWithStaleSlaves(analysisEntry inst.ReplicationAnalysis, candidateInstanceKey *inst.InstanceKey, forceInstanceRecovery bool, skipProcesses bool) (bool, *TopologyRecovery, error) {
	topologyRecovery, err := AttemptRecoveryRegistration(&analysisEntry, !forceInstanceRecovery, !forceInstanceRecovery)
	if topologyRecovery == nil {
		log.Debugf("topology_recovery: found an active or recent recovery on %+v. Will not issue another UnreachableMasterWithStaleSlaves.", analysisEntry.AnalyzedInstanceKey)
	} else {
		recoverUnreachableMasterWithStaleSlavesCounter.Inc(1)
		if !skipProcesses {
			err := executeProcesses(config.Config.UnreachableMasterWithStaleSlavesProcesses, "UnreachableMasterWithStaleSlavesProcesses", topologyRecovery, false)
			if err != nil {
				recoverUnreachableMasterWithStaleSlavesFailureCounter.Inc(1)
			} else {
				recoverUnreachableMasterWithStaleSlavesSuccessCounter.Inc(1)
			}
		}
	}
	return false, nil, err
}

// checkAndRecoverGenericProblem is a general-purpose recovery function
func checkAndRecoverGenericProblem(analysisEntry inst.ReplicationAnalysis, candidateInstanceKey *inst.InstanceKey, forceInstanceRecovery bool, skipProcesses bool) (bool, *TopologyRecovery, error) {
	return false, nil, nil
}

// Force a re-read of a topology instance; this is done because we need to substantiate a suspicion
// that we may have a failover scenario. we want to speed up reading the complete picture.
func emergentlyReadTopologyInstance(instanceKey *inst.InstanceKey, analysisCode inst.AnalysisCode) {
	if existsInCacheError := emergencyReadTopologyInstanceMap.Add(instanceKey.StringCode(), true, cache.DefaultExpiration); existsInCacheError != nil {
		// Just recently attempted
		return
	}
	go inst.ExecuteOnTopology(func() {
		inst.ReadTopologyInstance(instanceKey)
		inst.AuditOperation("emergently-read-topology-instance", instanceKey, string(analysisCode))
	})
}

// Force reading of replicas of given instance. This is because we suspect the instance is dead, and want to speed up
// detection of replication failure from its replicas.
func emergentlyReadTopologyInstanceReplicas(instanceKey *inst.InstanceKey, analysisCode inst.AnalysisCode) {
	replicas, err := inst.ReadReplicaInstancesIncludingBinlogServerSubReplicas(instanceKey)
	if err != nil {
		return
	}
	for _, replica := range replicas {
		go emergentlyReadTopologyInstance(&replica.Key, analysisCode)
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
		go emergentlyReadTopologyInstanceReplicas(&analysisEntry.AnalyzedInstanceKey, analysisEntry.Analysis)
	case inst.AllMasterSlavesNotReplicating:
		go emergentlyReadTopologyInstance(&analysisEntry.AnalyzedInstanceKey, analysisEntry.Analysis)
	case inst.FirstTierSlaveFailingToConnectToMaster:
		go emergentlyReadTopologyInstance(&analysisEntry.AnalyzedInstanceMasterKey, analysisEntry.Analysis)
	case inst.UnreachableMasterWithStaleSlaves:
		checkAndRecoverFunction = checkAndRecoverUnreachableMasterWithStaleSlaves
	}
	// Right now this is mostly causing noise with no clear action.
	// Will revisit this in the future.
	// case inst.AllMasterSlavesStale:
	// 	checkAndRecoverFunction = checkAndRecoverGenericProblem

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
			inst.EndDowntime(topologyRecovery.SuccessorKey)
			executeProcesses(config.Config.PostFailoverProcesses, "PostFailoverProcesses", topologyRecovery, false)
		}
	}
	topologyRecovery.InvokePostponed()
	return recoveryAttempted, topologyRecovery, err
}

// CheckAndRecover is the main entry point for the recovery mechanism
func CheckAndRecover(specificInstance *inst.InstanceKey, candidateInstanceKey *inst.InstanceKey, skipProcesses bool) (recoveryAttempted bool, promotedReplicaKey *inst.InstanceKey, err error) {
	// Allow the analysis to run evern if we don't want to recover
	replicationAnalysis, err := inst.GetReplicationAnalysis("", true, true)
	if err != nil {
		return false, nil, log.Errore(err)
	}
	// Check for recovery being disabled globally
	recoveryDisabledGlobally, err := IsRecoveryDisabled()
	if err != nil {
		log.Warningf("Unable to determine if recovery is disabled globally: %v", err)
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
				promotedReplicaKey = topologyRecovery.SuccessorKey
			}
		} else if recoveryDisabledGlobally {
			log.Infof("CheckAndRecover: InstanceKey: %+v, candidateInstanceKey: %+v, "+
				"skipProcesses: %v: NOT Recovering host (disabled globally)",
				analysisEntry.AnalyzedInstanceKey, candidateInstanceKey, skipProcesses)
		} else {
			go executeCheckAndRecoverFunction(analysisEntry, candidateInstanceKey, false, skipProcesses)
		}
	}
	return recoveryAttempted, promotedReplicaKey, err
}

// ForceExecuteRecovery can be called to issue a recovery process even if analysis says there is no recovery case.
// The caller of this function injects the type of analysis it wishes the function to assume.
// By calling this function one takes responsibility for one's actions.
func ForceExecuteRecovery(clusterName string, analysisCode inst.AnalysisCode, failedInstanceKey *inst.InstanceKey, candidateInstanceKey *inst.InstanceKey, skipProcesses bool) (recoveryAttempted bool, topologyRecovery *TopologyRecovery, err error) {
	clusterInfo, err := inst.ReadClusterInfo(clusterName)
	if err != nil {
		return recoveryAttempted, topologyRecovery, err
	}

	analysisEntry := inst.ReplicationAnalysis{
		Analysis:            analysisCode,
		ClusterDetails:      *clusterInfo,
		AnalyzedInstanceKey: *failedInstanceKey,
	}
	return executeCheckAndRecoverFunction(analysisEntry, candidateInstanceKey, true, skipProcesses)
}

// ForceMasterTakeover *trusts* master of given cluster is dead and fails over to designated instance,
// which has to be its direct child.
func ForceMasterTakeover(clusterName string, destination *inst.Instance) (topologyRecovery *TopologyRecovery, err error) {
	clusterMasters, err := inst.ReadClusterWriteableMaster(clusterName)
	if err != nil {
		return nil, fmt.Errorf("Cannot deduce cluster master for %+v", clusterName)
	}
	if len(clusterMasters) != 1 {
		return nil, fmt.Errorf("Cannot deduce cluster master for %+v", clusterName)
	}
	clusterMaster := clusterMasters[0]

	if !destination.MasterKey.Equals(&clusterMaster.Key) {
		return nil, fmt.Errorf("You may only promote a direct child of the master %+v. The master of %+v is %+v.", clusterMaster.Key, destination.Key, destination.MasterKey)
	}
	log.Debugf("Will demote %+v and promote %+v instead", clusterMaster.Key, destination.Key)

	recoveryAttempted, topologyRecovery, err := ForceExecuteRecovery(clusterName, inst.DeadMaster, &clusterMaster.Key, &destination.Key, false)
	if err != nil {
		return nil, err
	}
	if !recoveryAttempted {
		return nil, fmt.Errorf("Unexpected error: recovery not attempted. This should not happen")
	}
	if topologyRecovery == nil {
		return nil, fmt.Errorf("Recovery attempted but with no results. This should not happen")
	}
	if topologyRecovery.SuccessorKey == nil {
		return nil, fmt.Errorf("Recovery attempted yet no replica promoted")
	}
	return topologyRecovery, nil
}

// GracefulMasterTakeover will demote master of existing topology and promote its
// direct replica instead.
// It expects that replica to have no siblings.
// This function is graceful in that it will first lock down the master, then wait
// for the designated replica to catch up with last position.
// It will point old master at the newly promoted master at the correct coordinates, but will not start replication.
func GracefulMasterTakeover(clusterName string) (topologyRecovery *TopologyRecovery, promotedMasterCoordinates *inst.BinlogCoordinates, err error) {
	clusterMasters, err := inst.ReadClusterWriteableMaster(clusterName)
	if err != nil {
		return nil, nil, fmt.Errorf("Cannot deduce cluster master for %+v", clusterName)
	}
	if len(clusterMasters) != 1 {
		return nil, nil, fmt.Errorf("Cannot deduce cluster master for %+v. Found %+v potential masters", clusterName, len(clusterMasters))
	}
	clusterMaster := clusterMasters[0]
	if len(clusterMaster.SlaveHosts) == 0 {
		return nil, nil, fmt.Errorf("Master %+v doesn't seem to have replicas", clusterMaster.Key)
	}
	if len(clusterMaster.SlaveHosts) > 1 {
		return nil, nil, fmt.Errorf("GracefulMasterTakeover: master %+v should only have one replica (making the takeover safe and simple), but has %+v. Aborting", clusterMaster.Key, len(clusterMaster.SlaveHosts))
	}

	designatedInstanceKey := &(clusterMaster.SlaveHosts.GetInstanceKeys()[0])
	designatedInstance, err := inst.ReadTopologyInstance(designatedInstanceKey)
	if err != nil {
		return nil, nil, err
	}
	masterOfDesigntaedInstance, err := inst.GetInstanceMaster(designatedInstance)
	if err != nil {
		return nil, nil, err
	}
	if !masterOfDesigntaedInstance.Key.Equals(&clusterMaster.Key) {
		return nil, nil, fmt.Errorf("Sanity check failure. It seems like the designated instance %+v does not replicate from the master %+v (designated instance's master key is %+v). This error is strange. Panicking", designatedInstance.Key, clusterMaster.Key, designatedInstance.MasterKey)
	}
	if !designatedInstance.HasReasonableMaintenanceReplicationLag() {
		return nil, nil, fmt.Errorf("Desginated instance %+v seems to be lagging to much for thie operation. Aborting.", designatedInstance.Key)
	}
	log.Debugf("Will demote %+v and promote %+v instead", clusterMaster.Key, designatedInstance.Key)

	if designatedInstance, err = inst.StopSlave(&designatedInstance.Key); err != nil {
		return nil, nil, err
	}
	log.Debugf("Will set %+v as read_only", clusterMaster.Key)
	if clusterMaster, err = inst.SetReadOnly(&clusterMaster.Key, true); err != nil {
		return nil, nil, err
	}

	log.Debugf("Will advance %+v to master coordinates %+v", designatedInstance.Key, clusterMaster.SelfBinlogCoordinates)
	if designatedInstance, err = inst.StartSlaveUntilMasterCoordinates(&designatedInstance.Key, &clusterMaster.SelfBinlogCoordinates); err != nil {
		return nil, nil, err
	}
	promotedMasterCoordinates = &designatedInstance.SelfBinlogCoordinates

	recoveryAttempted, topologyRecovery, err := ForceExecuteRecovery(clusterName, inst.DeadMaster, &clusterMaster.Key, &designatedInstance.Key, false)
	if err != nil {
		return nil, nil, err
	}
	if !recoveryAttempted {
		return nil, nil, fmt.Errorf("Unexpected error: recovery not attempted. This should not happen")
	}
	if topologyRecovery == nil {
		return nil, nil, fmt.Errorf("Recovery attempted but with no results. This should not happen")
	}
	if topologyRecovery.SuccessorKey == nil {
		return nil, nil, fmt.Errorf("Recovery attempted yet no replica promoted")
	}

	clusterMaster, err = inst.ChangeMasterTo(&clusterMaster.Key, &designatedInstance.Key, promotedMasterCoordinates, false, inst.GTIDHintNeutral)

	return topologyRecovery, promotedMasterCoordinates, err
}
