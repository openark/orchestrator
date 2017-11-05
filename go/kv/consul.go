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
	"fmt"

	consulapi "github.com/armon/consul-api"
	"github.com/github/orchestrator/go/config"
	"github.com/openark/golib/log"
)

// A Consul store based on config's `ConsulAddress` and `ConsulKVPrefix`
type consulStore struct {
	client *consulapi.Client
}

// NewConsulStore creates a new consul store. It is possible that the client for this store is nil,
// which is the case if no consul config is provided.
func NewConsulStore() KVStore {
	store := &consulStore{}

	if config.Config.ConsulAddress != "" {
		consulConfig := consulapi.DefaultConfig()
		consulConfig.Address = config.Config.ConsulAddress
		if client, err := consulapi.NewClient(consulConfig); err != nil {
			log.Errore(err)
		} else {
			store.client = client
		}
	}
	return store
}

func (this *consulStore) fullKey(key string) string {
	if config.Config.ConsulKVPrefix != "" {
		return fmt.Sprintf("%s/%s", config.Config.ConsulKVPrefix, key)
	}
	return key
}

func (this *consulStore) PutKeyValue(key string, value string) (err error) {
	if this.client == nil {
		return nil
	}
	pair := &consulapi.KVPair{Key: this.fullKey(key), Value: []byte(value)}
	_, err = this.client.KV().Put(pair, nil)
	return err
}

func (this *consulStore) GetKeyValue(key string) (value string, err error) {
	if this.client == nil {
		return value, nil
	}
	pair, _, err := this.client.KV().Get(this.fullKey(key), nil)
	if err != nil {
		return value, err
	}
	return string(pair.Value), nil
}
