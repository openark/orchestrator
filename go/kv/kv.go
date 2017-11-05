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
	"sync"
)

type KVStore interface {
	PutKeyValue(key string, value string) (err error)
	GetKeyValue(key string) (value string, err error)
}

var kvMutex sync.Mutex
var kvStores = []KVStore{}

func InitKVStores() {
	kvMutex.Lock()
	defer kvMutex.Unlock()

	kvStores = []KVStore{
		NewInternalKVStore(),
		NewConsulStore(),
		NewZkStore(),
	}
}

func getKVStores() (stores []KVStore) {
	kvMutex.Lock()
	defer kvMutex.Unlock()

	stores = kvStores
	return stores
}

func PutValue(key string, value string) (err error) {
	for _, store := range getKVStores() {
		if err := store.PutKeyValue(key, value); err != nil {
			return err
		}
	}
	return nil
}

func GetValue(key string) (value string, err error) {
	for _, store := range getKVStores() {
		return store.GetKeyValue(key)
	}
	return value, nil
}
