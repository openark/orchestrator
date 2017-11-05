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

package inst

import (
	"fmt"
	"strings"
)

type AnalysisCode string
type StructureAnalysisCode string

const (
	NoProblem                                             AnalysisCode = "NoProblem"
	DeadMasterWithoutSlaves                                            = "DeadMasterWithoutSlaves"
	DeadMaster                                                         = "DeadMaster"
	DeadMasterAndSlaves                                                = "DeadMasterAndSlaves"
	DeadMasterAndSomeSlaves                                            = "DeadMasterAndSomeSlaves"
	UnreachableMasterWithStaleSlaves                                   = "UnreachableMasterWithStaleSlaves"
	UnreachableMaster                                                  = "UnreachableMaster"
	MasterSingleSlaveNotReplicating                                    = "MasterSingleSlaveNotReplicating"
	MasterSingleSlaveDead                                              = "MasterSingleSlaveDead"
	AllMasterSlavesNotReplicating                                      = "AllMasterSlavesNotReplicating"
	AllMasterSlavesNotReplicatingOrDead                                = "AllMasterSlavesNotReplicatingOrDead"
	AllMasterSlavesStale                                               = "AllMasterSlavesStale"
	MasterWithoutSlaves                                                = "MasterWithoutSlaves"
	DeadCoMaster                                                       = "DeadCoMaster"
	DeadCoMasterAndSomeSlaves                                          = "DeadCoMasterAndSomeSlaves"
	UnreachableCoMaster                                                = "UnreachableCoMaster"
	AllCoMasterSlavesNotReplicating                                    = "AllCoMasterSlavesNotReplicating"
	DeadIntermediateMaster                                             = "DeadIntermediateMaster"
	DeadIntermediateMasterWithSingleSlave                              = "DeadIntermediateMasterWithSingleSlave"
	DeadIntermediateMasterWithSingleSlaveFailingToConnect              = "DeadIntermediateMasterWithSingleSlaveFailingToConnect"
	DeadIntermediateMasterAndSomeSlaves                                = "DeadIntermediateMasterAndSomeSlaves"
	DeadIntermediateMasterAndSlaves                                    = "DeadIntermediateMasterAndSlaves"
	UnreachableIntermediateMaster                                      = "UnreachableIntermediateMaster"
	AllIntermediateMasterSlavesFailingToConnectOrDead                  = "AllIntermediateMasterSlavesFailingToConnectOrDead"
	AllIntermediateMasterSlavesNotReplicating                          = "AllIntermediateMasterSlavesNotReplicating"
	FirstTierSlaveFailingToConnectToMaster                             = "FirstTierSlaveFailingToConnectToMaster"
	BinlogServerFailingToConnectToMaster                               = "BinlogServerFailingToConnectToMaster"
)

const (
	StatementAndMixedLoggingSlavesStructureWarning StructureAnalysisCode = "StatementAndMixedLoggingSlavesStructureWarning"
	StatementAndRowLoggingSlavesStructureWarning                         = "StatementAndRowLoggingSlavesStructureWarning"
	MixedAndRowLoggingSlavesStructureWarning                             = "MixedAndRowLoggingSlavesStructureWarning"
	MultipleMajorVersionsLoggingSlaves                                   = "MultipleMajorVersionsLoggingSlaves"
)

type InstanceAnalysis struct {
	key      *InstanceKey
	analysis AnalysisCode
}

func NewInstanceAnalysis(instanceKey *InstanceKey, analysis AnalysisCode) *InstanceAnalysis {
	return &InstanceAnalysis{
		key:      instanceKey,
		analysis: analysis,
	}
}

func (instanceAnalysis *InstanceAnalysis) String() string {
	return fmt.Sprintf("%s/%s", instanceAnalysis.key.StringCode(), string(instanceAnalysis.analysis))
}

// PeerAnalysisMap indicates the number of peers agreeing on an analysis.
// Key of this map is a InstanceAnalysis.String()
type PeerAnalysisMap map[string]int

// ReplicationAnalysis notes analysis on replication chain status, per instance
type ReplicationAnalysis struct {
	AnalyzedInstanceKey                       InstanceKey
	AnalyzedInstanceMasterKey                 InstanceKey
	ClusterDetails                            ClusterInfo
	IsMaster                                  bool
	IsCoMaster                                bool
	LastCheckValid                            bool
	CountReplicas                             uint
	CountValidReplicas                        uint
	CountValidReplicatingReplicas             uint
	CountReplicasFailingToConnectToMaster     uint
	CountStaleReplicas                        uint
	CountDowntimedReplicas                    uint
	ReplicationDepth                          uint
	SlaveHosts                                InstanceKeyMap
	IsFailingToConnectToMaster                bool
	Analysis                                  AnalysisCode
	Description                               string
	StructureAnalysis                         []StructureAnalysisCode
	IsDowntimed                               bool
	IsReplicasDowntimed                       bool // as good as downtimed because all replicas are downtimed AND analysis is all about the replicas (e.e. AllMasterSlavesNotReplicating)
	DowntimeEndTimestamp                      string
	DowntimeRemainingSeconds                  int
	IsBinlogServer                            bool
	PseudoGTIDImmediateTopology               bool
	OracleGTIDImmediateTopology               bool
	MariaDBGTIDImmediateTopology              bool
	BinlogServerImmediateTopology             bool
	CountStatementBasedLoggingReplicas        uint
	CountMixedBasedLoggingReplicas            uint
	CountRowBasedLoggingReplicas              uint
	CountDistinctMajorVersionsLoggingReplicas uint
	IsActionableRecovery                      bool
	ProcessingNodeHostname                    string
	ProcessingNodeToken                       string
	CountAdditionalAgreeingNodes              int
	StartActivePeriod                         string
}

type AnalysisMap map[string](*ReplicationAnalysis)

type ReplicationAnalysisChangelog struct {
	AnalyzedInstanceKey InstanceKey
	Changelog           []string
}

// ReadReplicaHostsFromString parses and reads replica keys from comma delimited string
func (this *ReplicationAnalysis) ReadReplicaHostsFromString(replicaHostsString string) error {
	this.SlaveHosts = *NewInstanceKeyMap()
	return this.SlaveHosts.ReadCommaDelimitedList(replicaHostsString)
}

// AnalysisString returns a human friendly description of all analysis issues
func (this *ReplicationAnalysis) AnalysisString() string {
	result := []string{}
	if this.Analysis != NoProblem {
		result = append(result, string(this.Analysis))
	}
	for _, structureAnalysis := range this.StructureAnalysis {
		result = append(result, string(structureAnalysis))
	}
	return strings.Join(result, ", ")
}

type ClusterAliasMaster struct {
	ClusterAlias string
	Master       InstanceKey
}

func NewClusterAliasMaster(clusterAlias string, master InstanceKey) *ClusterAliasMaster {
	return &ClusterAliasMaster{
		ClusterAlias: clusterAlias,
		Master:       master,
	}
}
