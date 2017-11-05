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

package consul

import (
	"fmt"
	"sync"

	"github.com/github/orchestrator/go/config"

	consulapi "github.com/hashicorp/consul/api"
)

var consulClient *consulapi.Client
var clientMutex sync.Mutex

func getConsulClient() *consulapi.Client {
	clientMutex.Lock()
	defer clientMutex.Unlock()

	if consulClient == nil && config.Config.ConsulAddress != "" {
		consulConfig := consulapi.DefaultConfig()
		consulConfig.Address = config.Config.ConsulAddress
		if client, err := consulapi.NewClient(consulConfig); err == nil {
			consulClient = client
		}
	}
	return consulClient
}

func fullKey(key string) string {
	if config.Config.ConsulKVPrefix != "" {
		return fmt.Sprintf("%s/%s", config.Config.ConsulKVPrefix, key)
	}
	return key
}

func PutKeyValue(key string, value string) (err error) {
	client := getConsulClient()
	if client == nil {
		return nil
	}
	p := &consulapi.KVPair{Key: fullKey(key), Value: []byte(value)}
	_, err = client.KV().Put(p, nil)
	return err
}

func GetKeyValue(key string) (value string, err error) {
	client := getConsulClient()
	if client == nil {
		return value, nil
	}
	pair, _, err := client.KV().Get(fullKey(key), nil)
	if err != nil {
		return value, err
	}
	return string(pair.Value), nil
}
