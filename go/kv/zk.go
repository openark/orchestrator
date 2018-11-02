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

package kv

// Internal key-value store, based on relational backend
type zkStore struct {
}

// TODO: use config.Config.ZkAddress to put/get k/v in ZooKeeper. See
// - https://github.com/outbrain/zookeepercli
// - https://github.com/samuel/go-zookeeper/zk

func NewZkStore() KVStore {
	return &zkStore{}
}

func (this *zkStore) PutKeyValue(key string, value string) (err error) {
	return
}

func (this *zkStore) GetKeyValue(key string) (value string, err error) {
	return
}
