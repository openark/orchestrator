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
	"github.com/openark/golib/log"
	"github.com/openark/golib/sqlutils"
	"github.com/openark/orchestrator/go/db"
)

// Internal key-value store, based on relational backend
type internalKVStore struct {
}

func NewInternalKVStore() KVStore {
	return &internalKVStore{}
}

func (this *internalKVStore) PutKeyValue(key string, value string) (err error) {
	_, err = db.ExecOrchestrator(`
		replace
			into kv_store (
        store_key, store_value, last_updated
			) values (
				?, ?, now()
			)
		`, key, value,
	)
	return log.Errore(err)
}

func (this *internalKVStore) GetKeyValue(key string) (value string, found bool, err error) {
	query := `
		select
			store_value
		from
			kv_store
		where
      store_key = ?
		`

	err = db.QueryOrchestrator(query, sqlutils.Args(key), func(m sqlutils.RowMap) error {
		value = m.GetString("store_value")
		found = true
		return nil
	})

	return value, found, log.Errore(err)
}

func (this *internalKVStore) PutKVPairs(kvPairs []*KVPair) (err error) {
	for _, pair := range kvPairs {
		if err := this.PutKeyValue(pair.Key, pair.Value); err != nil {
			return err
		}
	}
	return nil
}

func (this *internalKVStore) DistributePairs(kvPairs [](*KVPair)) (err error) {
	return nil
}
