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

	"github.com/github/orchestrator/go/inst"

	"github.com/openark/golib/log"
)

type SnapshotData struct {
	Keys               []inst.InstanceKey
	DowntimedInstances []inst.Downtime
	Candidates         []inst.CandidateDatabaseInstance
	HostnameUnresolves []inst.HostnameUnresolve
	Pools              []inst.PoolInstancesSubmission
	RecoveryDisabled   bool
}

func NewSnapshotData() *SnapshotData {
	return &SnapshotData{}
}
func CreateSnapshotData() *SnapshotData {
	snapshotData := NewSnapshotData()

	// keys
	snapshotData.Keys, _ = inst.ReadAllInstanceKeys()
	// downtime
	snapshotData.DowntimedInstances, _ = inst.ReadDowntime()
	// candidates
	snapshotData.Candidates, _ = inst.BulkReadCandidateDatabaseInstance()
	// unresolve
	snapshotData.HostnameUnresolves, _ = inst.ReadAllHostnameUnresolves()
	// pool
	snapshotData.Pools, _ = inst.ReadAllPoolInstancesSubmissions()
	// recovery
	snapshotData.RecoveryDisabled, _ = IsRecoveryDisabled()

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
				log.Debugf("raft Restore(): forgetting %+v", key)
				inst.ForgetInstance(&key)
			}
		}
		for _, key := range instanceKeyMap.GetInstanceKeys() {
			log.Debugf("raft Restore(): discovering %+v", key)
			go func() {
				snapshotDiscoveryKeys <- key
			}()
		}
	}
	log.Debugf("raft restore applied")
	return nil
}
