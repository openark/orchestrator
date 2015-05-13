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

import ()

type AnalysisCode string

const (
	NoProblem                                 AnalysisCode = "NoProblem"
	DeadMasterWithoutSlaves                                = "DeadMasterWithoutSlaves"
	DeadMaster                                             = "DeadMaster"
	DeadMasterAndSlaves                                    = "DeadMasterAndSlaves"
	DeadMasterAndSomeSlaves                                = "DeadMasterAndSomeSlaves"
	UnreachableMaster                                      = "UnreachableMaster"
	AllMasterSlavesNotReplicating                          = "AllMasterSlavesNotReplicating"
	MasterWithoutSlaves                                    = "MasterWithoutSlaves"
	DeadCoMaster                                           = "DeadCoMaster"
	UnreachableCoMaster                                    = "UnreachableCoMaster"
	AllCoMasterSlavesNotReplicating                        = "AllCoMasterSlavesNotReplicating"
	DeadIntermediateMaster                                 = "DeadIntermediateMaster"
	DeadIntermediateMasterAndSomeSlaves                    = "DeadIntermediateMasterAndSomeSlaves"
	UnreachableIntermediateMaster                          = "UnreachableIntermediateMaster"
	AllIntermediateMasterSlavesNotReplicating              = "AllIntermediateMasterSlavesNotReplicating"
	FirstTierSlaveFailingToConnectToMaster                 = "FirstTierSlaveFailingToConnectToMaster"
)

// ReplicationAnalysis notes analysis on replication chain status, per instance
type ReplicationAnalysis struct {
	AnalyzedInstanceKey         InstanceKey
	AnalyzedInstanceMasterKey   InstanceKey
	ClusterName                 string
	IsMaster                    bool
	IsCoMaster                  bool
	LastCheckValid              bool
	CountSlaves                 uint
	CountValidSlaves            uint
	CountValidReplicatingSlaves uint
	ReplicationDepth            uint
	IsFailingToConnectToMaster  bool
	Analysis                    AnalysisCode
	Description                 string
}
