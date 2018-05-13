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

import (
	"math/rand"
	"strings"
	"time"

	"github.com/github/orchestrator/go/config"
	"github.com/outbrain/zookeepercli/zk"
)

// Internal key-value store, based on relational backend
type zkStore struct {
	servers []string
}

func NewZkStore() KVStore {
	rand.Seed(time.Now().UnixNano())
	serversArray := strings.Split(config.Config.ZkAddress, ",")

	return &zkStore{
		servers: serversArray,
	}
}

func (this *zkStore) PutKeyValue(key string, value string) (err error) {
	aclstr := ""
	_, err = zk.Create(key, []byte(value), aclstr, true)
	return err
}

func (this *zkStore) GetKeyValue(key string) (value string, err error) {
	result, err := zk.Get(key)
	if err != nil {
		return value, err
	}
	return string(result), nil
}
