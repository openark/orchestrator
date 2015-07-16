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
	"strings"
)

type AnalysisCode string

const (
	NoProblem                                             AnalysisCode = "NoProblem"
	DeadMasterWithoutSlaves                                            = "DeadMasterWithoutSlaves"
	DeadMaster                                                         = "DeadMaster"
	DeadMasterAndSlaves                                                = "DeadMasterAndSlaves"
	DeadMasterAndSomeSlaves                                            = "DeadMasterAndSomeSlaves"
	UnreachableMaster                                                  = "UnreachableMaster"
	MasterSingleSlaveNotReplicating                                    = "MasterSingleSlaveNotReplicating"
	MasterSingleSlaveDead                                              = "MasterSingleSlaveDead"
	AllMasterSlavesNotReplicating                                      = "AllMasterSlavesNotReplicating"
	AllMasterSlavesNotReplicatingOrDead                                = "AllMasterSlavesNotReplicatingOrDead"
	MasterWithoutSlaves                                                = "MasterWithoutSlaves"
	DeadCoMaster                                                       = "DeadCoMaster"
	UnreachableCoMaster                                                = "UnreachableCoMaster"
	AllCoMasterSlavesNotReplicating                                    = "AllCoMasterSlavesNotReplicating"
	DeadIntermediateMaster                                             = "DeadIntermediateMaster"
	DeadIntermediateMasterWithSingleSlave                              = "DeadIntermediateMasterWithSingleSlave"
	DeadIntermediateMasterWithSingleSlaveFailingToConnect              = "DeadIntermediateMasterWithSingleSlaveFailingToConnect"
	DeadIntermediateMasterAndSomeSlaves                                = "DeadIntermediateMasterAndSomeSlaves"
	UnreachableIntermediateMaster                                      = "UnreachableIntermediateMaster"
	AllIntermediateMasterSlavesFailingToConnectOrDead                  = "AllIntermediateMasterSlavesFailingToConnectOrDead"
	AllIntermediateMasterSlavesNotReplicating                          = "AllIntermediateMasterSlavesNotReplicating"
	FirstTierSlaveFailingToConnectToMaster                             = "FirstTierSlaveFailingToConnectToMaster"
	MaxscaleFailingToConnectToMaster                                   = "MaxscaleFailingToConnectToMaster"
)

// ReplicationAnalysis notes analysis on replication chain status, per instance
type ReplicationAnalysis struct {
	AnalyzedInstanceKey                 InstanceKey
	AnalyzedInstanceMasterKey           InstanceKey
	ClusterDetails                      ClusterInfo
	IsMaster                            bool
	IsCoMaster                          bool
	LastCheckValid                      bool
	CountSlaves                         uint
	CountValidSlaves                    uint
	CountValidReplicatingSlaves         uint
	CountSlavesFailingToConnectToMaster uint
	ReplicationDepth                    uint
	SlaveHosts                          InstanceKeyMap
	IsFailingToConnectToMaster          bool
	Analysis                            AnalysisCode
	Description                         string
	IsDowntimed                         bool
	DowntimeEndTimestamp                string
	DowntimeRemainingSeconds            int
	IsMaxscale                          bool
}

// GetSlaveHostsAsString serializes all slave keys as a single comma delimited string
func (this *ReplicationAnalysis) GetSlaveHostsAsString() string {
	slaveHostsStrings := []string{}
	for slaveKey := range this.SlaveHosts {
		slaveHostsStrings = append(slaveHostsStrings, slaveKey.DisplayString())
	}
	return strings.Join(slaveHostsStrings, ",")
}

// ReadSlaveHostsFromString parses and reads slave keys from comma delimited string
func (this *ReplicationAnalysis) ReadSlaveHostsFromString(slaveHostsString string) error {
	this.SlaveHosts = make(map[InstanceKey]bool)

	slaveHostsStrings := strings.Split(slaveHostsString, ",")
	for _, slaveKeyString := range slaveHostsStrings {
		slaveKey, err := ParseInstanceKey(slaveKeyString)
		if err != nil {
			return err
		}
		this.SlaveHosts[*slaveKey] = true
	}
	return nil
}
