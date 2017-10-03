/*
   Copyright 2017 Shlomi Noach, GitHub Inc.

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
	"encoding/json"
	"io"

	"github.com/github/orchestrator/go/db"
	"github.com/github/orchestrator/go/inst"

	"github.com/openark/golib/log"
	"github.com/openark/golib/sqlutils"
)

type SnapshotData struct {
	Keys             []inst.InstanceKey
	RecoveryDisabled bool

	ClusterAlias,
	ClusterAliasOverride,
	ClusterDomainName,
	HostAttributes,
	AccessToken,
	PoolInstances,
	HostnameResolves,
	HostnameUnresolves,
	DowntimedInstances,
	Candidates,
	Detections,
	Recovery,
	RecoverySteps sqlutils.NamedResultData
}

func NewSnapshotData() *SnapshotData {
	return &SnapshotData{}
}

func readTableData(tableName string, data *sqlutils.NamedResultData) error {
	orcdb, err := db.OpenOrchestrator()
	if err != nil {
		return log.Errore(err)
	}
	*data, err = sqlutils.ScanTable(orcdb, tableName)
	return log.Errore(err)
}

func writeTableData(tableName string, data *sqlutils.NamedResultData) error {
	orcdb, err := db.OpenOrchestrator()
	if err != nil {
		return log.Errore(err)
	}
	err = sqlutils.WriteTable(orcdb, tableName, *data)
	return log.Errore(err)
}

func CreateSnapshotData() *SnapshotData {
	snapshotData := NewSnapshotData()

	// keys
	snapshotData.Keys, _ = inst.ReadAllInstanceKeys()
	snapshotData.RecoveryDisabled, _ = IsRecoveryDisabled()

	readTableData("cluster_alias", &snapshotData.ClusterAlias)
	readTableData("cluster_alias_override", &snapshotData.ClusterAliasOverride)
	readTableData("cluster_domain_name", &snapshotData.ClusterDomainName)
	readTableData("access_token", &snapshotData.AccessToken)
	readTableData("host_attributes", &snapshotData.HostAttributes)
	readTableData("database_instance_pool", &snapshotData.PoolInstances)
	readTableData("hostname_resolve", &snapshotData.HostnameResolves)
	readTableData("hostname_unresolve", &snapshotData.HostnameUnresolves)
	readTableData("database_instance_downtime", &snapshotData.DowntimedInstances)
	readTableData("candidate_database_instance", &snapshotData.Candidates)
	readTableData("topology_failure_detection", &snapshotData.Detections)
	readTableData("topology_recovery", &snapshotData.Recovery)
	readTableData("topology_recovery_steps", &snapshotData.RecoverySteps)

	log.Debugf("raft snapshot data created")
	return snapshotData
}

type SnapshotDataCreatorApplier struct {
}

func NewSnapshotDataCreatorApplier() *SnapshotDataCreatorApplier {
	generator := &SnapshotDataCreatorApplier{}
	return generator
}

func (this *SnapshotDataCreatorApplier) GetData() (data []byte, err error) {
	snapshotData := CreateSnapshotData()
	return json.Marshal(snapshotData)
}

func (this *SnapshotDataCreatorApplier) Restore(rc io.ReadCloser) error {
	snapshotData := NewSnapshotData()
	// if err := json.Unmarshal(data, &snapshotData); err != nil {
	if err := json.NewDecoder(rc).Decode(&snapshotData); err != nil {
		return err
	}

	// keys
	{
		instanceKeyMap := inst.NewInstanceKeyMap()
		instanceKeyMap.AddKeys(snapshotData.Keys)

		allKeys, _ := inst.ReadAllInstanceKeys()
		for _, key := range allKeys {
			if !instanceKeyMap.HasKey(key) {
				inst.ForgetInstance(&key)
			}
		}
		for _, key := range instanceKeyMap.GetInstanceKeys() {
			go func() {
				snapshotDiscoveryKeys <- key
			}()
		}
	}
	writeTableData("cluster_alias", &snapshotData.ClusterAlias)
	writeTableData("cluster_alias_override", &snapshotData.ClusterAliasOverride)
	writeTableData("cluster_domain_name", &snapshotData.ClusterDomainName)
	writeTableData("access_token", &snapshotData.AccessToken)
	writeTableData("host_attributes", &snapshotData.HostAttributes)
	writeTableData("database_instance_pool", &snapshotData.PoolInstances)
	writeTableData("hostname_resolve", &snapshotData.HostnameResolves)
	writeTableData("hostname_unresolve", &snapshotData.HostnameUnresolves)
	writeTableData("database_instance_downtime", &snapshotData.DowntimedInstances)
	writeTableData("candidate_database_instance", &snapshotData.Candidates)
	writeTableData("topology_recovery", &snapshotData.Recovery)
	writeTableData("topology_failure_detection", &snapshotData.Detections)
	writeTableData("topology_recovery_steps", &snapshotData.RecoverySteps)

	// recovery disable
	{
		SetRecoveryDisabled(snapshotData.RecoveryDisabled)
	}
	log.Debugf("raft restore applied")
	return nil
}
