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
		// ConsulAclToken defaults to ""
		consulConfig.Token = config.Config.ConsulAclToken
		if client, err := consulapi.NewClient(consulConfig); err != nil {
			log.Errore(err)
		} else {
			store.client = client
		}
	}
	return store
}

func (this *consulStore) PutKeyValue(key string, value string) (err error) {
	if this.client == nil {
		return nil
	}
	pair := &consulapi.KVPair{Key: key, Value: []byte(value)}
	_, err = this.client.KV().Put(pair, nil)
	return err
}

func (this *consulStore) GetKeyValue(key string) (value string, found bool, err error) {
	if this.client == nil {
		return value, found, nil
	}
	pair, _, err := this.client.KV().Get(key, nil)
	if err != nil {
		return value, found, err
	}
	return string(pair.Value), (pair != nil), nil
}

func (this *consulStore) DistributePairs(kvPairs [](*KVPair)) (failedDistributions []string, err error) {
	log.Debugf("consulStore.DistributePairs(): %d pairs", len(kvPairs))
	if !config.Config.ConsulCrossDataCenterDistribution {
		log.Debugf("consulStore.DistributePairs(): !config.Config.ConsulCrossDataCenterDistribution")
		return failedDistributions, nil
	}
	datacenters, err := this.client.Catalog().Datacenters()
	if err != nil {
		return failedDistributions, err
	}
	log.Debugf("consulStore.DistributePairs(): %d datacenters", len(datacenters))
	consulPairs := [](*consulapi.KVPair){}
	for _, kvPair := range kvPairs {
		log.Debugf("consulStore.DistributePairs():kvpair: %+v", kvPair)
		consulPairs = append(consulPairs, &consulapi.KVPair{Key: kvPair.Key, Value: []byte(kvPair.Value)})
	}
	for _, datacenter := range datacenters {
		log.Debugf("consulStore.DistributePairs():datacenter: %+v", datacenter)
		writeOptions := &consulapi.WriteOptions{Datacenter: datacenter}
		for _, consulPair := range consulPairs {
			if _, e := this.client.KV().Put(consulPair, writeOptions); e != nil {
				log.Errorf("consulStore.DistributePairs(): failed writing %s", datacenter)
				failedDistributions = append(failedDistributions, datacenter)
				err = e
			}
		}
	}
	return failedDistributions, err
}
