/*
   Copyright 2020 Shlomi Noach, GitHub Inc.

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
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/openark/orchestrator/go/config"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/patrickmn/go-cache"

	"github.com/openark/golib/log"
)

// A Consul store based on config's `ConsulAddress`, `ConsulScheme`, and `ConsulKVPrefix`
type consulTxnStore struct {
	client                        *consulapi.Client
	kvCache                       *cache.Cache
	pairsDistributionSuccessMutex sync.Mutex
	distributionReentry           int64
}

// NewConsulTxnStore creates a new consul store that uses Consul Transactions to read/write multiple KVPairs.
// It is possible that the client for this store is nil, which is the case if no consul config is provided
func NewConsulTxnStore() KVStore {
	store := &consulTxnStore{
		kvCache: cache.New(cache.NoExpiration, cache.DefaultExpiration),
	}

	if config.Config.ConsulAddress != "" {
		consulConfig := consulapi.DefaultConfig()
		consulConfig.Address = config.Config.ConsulAddress
		consulConfig.Scheme = config.Config.ConsulScheme
		if config.Config.ConsulScheme == "https" {
			consulConfig.HttpClient = &http.Client{
				Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
			}
		}
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

// doWriteTxn performs one or many of write operations using a Consul Transaction and handles any client/server
// or transaction-level errors. Updates are all-or-nothing - all operations are rolled-back on any txn error
func (this *consulTxnStore) doWriteTxn(txnOps consulapi.TxnOps, queryOptions *consulapi.QueryOptions) (err error) {
	ok, resp, _, err := this.client.Txn().Txn(txnOps, queryOptions)
	if err != nil {
		log.Errorf("consulTxnStore.doWriteTxn(): failed %v", err)
		return err
	} else if !ok {
		for _, terr := range resp.Errors {
			txnOp := txnOps[terr.OpIndex]
			log.Errorf("consulTxnStore.doWriteTxn(): transaction error %q for KV %s on %s", terr.What, txnOp.KV.Verb, txnOp.KV.Key)
			err = fmt.Errorf("%v", terr.What)
		}
	}
	return err
}

func (this *consulTxnStore) updateDatacenterKVPairs(wg *sync.WaitGroup, dc string, kvPairs []*consulapi.KVPair) (skipped, existing, written, failed int, err error) {
	defer wg.Done()

	queryOptions := &consulapi.QueryOptions{Datacenter: dc}
	kcCacheKeys := make([]string, 0)

	// get the current key-values in a single transaction
	var terr error
	var getTxnOps consulapi.TxnOps
	var getTxnResp *consulapi.TxnResponse
	var possibleSetKVPairs []*consulapi.KVPair
	for _, kvPair := range kvPairs {
		val := string(kvPair.Value)
		kcCacheKey := getConsulKVCacheKey(dc, kvPair.Key)
		kcCacheKeys = append(kcCacheKeys, kcCacheKey)
		if value, found := this.kvCache.Get(kcCacheKey); found && val == value {
			skipped++
			continue
		}
		getTxnOps = append(getTxnOps, &consulapi.TxnOp{
			KV: &consulapi.KVTxnOp{
				Verb: consulapi.KVGet,
				Key:  kvPair.Key,
			},
		})
		possibleSetKVPairs = append(possibleSetKVPairs, kvPair)
	}
	if len(getTxnOps) > 0 {
		_, getTxnResp, _, terr = this.client.Txn().Txn(getTxnOps, queryOptions)
		if terr != nil {
			err = terr
		}
	}

	// find key-value pairs that need updating, add pairs that need updating to set transaction
	var setTxnOps consulapi.TxnOps
	for _, pair := range possibleSetKVPairs {
		var kvExistsAndEqual bool
		if len(getTxnResp.Errors) > 0 {
			for _, terr := range getTxnResp.Errors {
				if !strings.HasSuffix(terr.What, "doesn't exist") {
					log.Errorf("consulTxnStore.DistributePairs(): get transaction error %v", terr.What)
				}
			}
		} else if getTxnResp != nil {
			for _, result := range getTxnResp.Results {
				if pair.Key == result.KV.Key && string(pair.Value) == string(result.KV.Value) {
					this.kvCache.SetDefault(getConsulKVCacheKey(dc, pair.Key), string(pair.Value))
					kvExistsAndEqual = true
					existing++
					break
				}
			}
		}
		if !kvExistsAndEqual {
			setTxnOps = append(setTxnOps, &consulapi.TxnOp{
				KV: &consulapi.KVTxnOp{
					Verb:  consulapi.KVSet,
					Key:   pair.Key,
					Value: pair.Value,
				},
			})
		}
	}

	// update key-value pairs in a single Consul Transaction
	if len(setTxnOps) > 0 {
		if err = this.doWriteTxn(setTxnOps, queryOptions); err != nil {
			log.Errorf("consulTxnStore.DistributePairs(): failed %v", kcCacheKeys)
			failed = len(setTxnOps)
		} else {
			for _, txnOp := range setTxnOps {
				this.kvCache.SetDefault(getConsulKVCacheKey(dc, txnOp.KV.Key), string(txnOp.KV.Value))
				written++
			}
		}
	}

	return skipped, existing, written, failed, err
}

// GetKeyValue returns the value of a Consul KV if it exists
func (this *consulTxnStore) GetKeyValue(key string) (value string, found bool, err error) {
	if this.client == nil {
		return value, found, nil
	}
	pair, _, err := this.client.KV().Get(key, nil)
	if err != nil {
		return value, found, err
	}
	return string(pair.Value), (pair != nil), nil
}

// PutKeyValue performs a Consul KV put operation for a key/value
func (this *consulTxnStore) PutKeyValue(key string, value string) (err error) {
	if this.client == nil {
		return nil
	}
	pair := &consulapi.KVPair{Key: key, Value: []byte(value)}
	_, err = this.client.KV().Put(pair, nil)
	return err
}

// PutKVPairs updates one or more KV pairs in a single, atomic Consul operation.
// If a single KV pair is provided PutKeyValue is used to update the pair
func (this *consulTxnStore) PutKVPairs(kvPairs []*KVPair) (err error) {
	if this.client == nil {
		return nil
	}
	// use .PutKeyValue for single KVPair puts
	if len(kvPairs) == 1 {
		return this.PutKeyValue(kvPairs[0].Key, kvPairs[0].Value)
	}
	var txnOps consulapi.TxnOps
	for _, pair := range kvPairs {
		txnOps = append(txnOps, &consulapi.TxnOp{
			KV: &consulapi.KVTxnOp{
				Verb:  consulapi.KVSet,
				Key:   pair.Key,
				Value: []byte(pair.Value),
			},
		})
	}
	return this.doWriteTxn(txnOps, nil)
}

// DistributePairs updates all known Consul Datacenters with one or more KV pairs
func (this *consulTxnStore) DistributePairs(kvPairs [](*KVPair)) (err error) {
	// This function is non re-entrant (it can only be running once at any point in time)
	if atomic.CompareAndSwapInt64(&this.distributionReentry, 0, 1) {
		defer atomic.StoreInt64(&this.distributionReentry, 0)
	} else {
		return
	}

	if !config.Config.ConsulCrossDataCenterDistribution {
		return nil
	}

	datacenters, err := this.client.Catalog().Datacenters()
	if err != nil {
		return err
	}
	log.Debugf("consulTxnStore.DistributePairs(): distributing %d pairs to %d datacenters", len(kvPairs), len(datacenters))
	consulPairs := []*consulapi.KVPair{}
	for _, kvPair := range kvPairs {
		consulPairs = append(consulPairs, &consulapi.KVPair{Key: kvPair.Key, Value: []byte(kvPair.Value)})
	}
	var wg sync.WaitGroup
	for _, datacenter := range datacenters {
		var skipped, existing, written, failed int
		datacenter := datacenter

		wg.Add(1)
		skipped, existing, written, failed, err = this.updateDatacenterKVPairs(&wg, datacenter, consulPairs)
		log.Debugf("consulTxnStore.DistributePairs(): datacenter: %s; skipped: %d, existing: %d, written: %d, failed: %d", datacenter, skipped, existing, written, failed)
	}
	wg.Wait()
	return err
}
