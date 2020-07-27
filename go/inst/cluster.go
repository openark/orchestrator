/*
   Copyright 2014 Outbrain Inc.

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

package inst

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/openark/orchestrator/go/config"
	"github.com/openark/orchestrator/go/kv"
)

func GetClusterMainKVKey(clusterAlias string) string {
	return fmt.Sprintf("%s%s", config.Config.KVClusterMainPrefix, clusterAlias)
}

func getClusterMainKVPair(clusterAlias string, mainKey *InstanceKey) *kv.KVPair {
	if clusterAlias == "" {
		return nil
	}
	if mainKey == nil {
		return nil
	}
	return kv.NewKVPair(GetClusterMainKVKey(clusterAlias), mainKey.StringCode())
}

// GetClusterMainKVPairs returns all KV pairs associated with a main. This includes the
// full identity of the main as well as a breakdown by hostname, port, ipv4, ipv6
func GetClusterMainKVPairs(clusterAlias string, mainKey *InstanceKey) (kvPairs [](*kv.KVPair)) {
	mainKVPair := getClusterMainKVPair(clusterAlias, mainKey)
	if mainKVPair == nil {
		return kvPairs
	}
	kvPairs = append(kvPairs, mainKVPair)

	addPair := func(keySuffix, value string) {
		key := fmt.Sprintf("%s/%s", mainKVPair.Key, keySuffix)
		kvPairs = append(kvPairs, kv.NewKVPair(key, value))
	}

	addPair("hostname", mainKey.Hostname)
	addPair("port", fmt.Sprintf("%d", mainKey.Port))
	if ipv4, ipv6, err := readHostnameIPs(mainKey.Hostname); err == nil {
		addPair("ipv4", ipv4)
		addPair("ipv6", ipv6)
	}
	return kvPairs
}

// mappedClusterNameToAlias attempts to match a cluster with an alias based on
// configured ClusterNameToAlias map
func mappedClusterNameToAlias(clusterName string) string {
	for pattern, alias := range config.Config.ClusterNameToAlias {
		if pattern == "" {
			// sanity
			continue
		}
		if matched, _ := regexp.MatchString(pattern, clusterName); matched {
			return alias
		}
	}
	return ""
}

// ClusterInfo makes for a cluster status/info summary
type ClusterInfo struct {
	ClusterName                            string
	ClusterAlias                           string // Human friendly alias
	ClusterDomain                          string // CNAME/VIP/A-record/whatever of the main of this cluster
	CountInstances                         uint
	HeuristicLag                           int64
	HasAutomatedMainRecovery             bool
	HasAutomatedIntermediateMainRecovery bool
}

// ReadRecoveryInfo
func (this *ClusterInfo) ReadRecoveryInfo() {
	this.HasAutomatedMainRecovery = this.filtersMatchCluster(config.Config.RecoverMainClusterFilters)
	this.HasAutomatedIntermediateMainRecovery = this.filtersMatchCluster(config.Config.RecoverIntermediateMainClusterFilters)
}

// filtersMatchCluster will see whether the given filters match the given cluster details
func (this *ClusterInfo) filtersMatchCluster(filters []string) bool {
	for _, filter := range filters {
		if filter == this.ClusterName {
			return true
		}
		if filter == this.ClusterAlias {
			return true
		}
		if strings.HasPrefix(filter, "alias=") {
			// Match by exact cluster alias name
			alias := strings.SplitN(filter, "=", 2)[1]
			if alias == this.ClusterAlias {
				return true
			}
		} else if strings.HasPrefix(filter, "alias~=") {
			// Match by cluster alias regex
			aliasPattern := strings.SplitN(filter, "~=", 2)[1]
			if matched, _ := regexp.MatchString(aliasPattern, this.ClusterAlias); matched {
				return true
			}
		} else if filter == "*" {
			return true
		} else if matched, _ := regexp.MatchString(filter, this.ClusterName); matched && filter != "" {
			return true
		}
	}
	return false
}

// ApplyClusterAlias updates the given clusterInfo's ClusterAlias property
func (this *ClusterInfo) ApplyClusterAlias() {
	if this.ClusterAlias != "" && this.ClusterAlias != this.ClusterName {
		// Already has an alias; abort
		return
	}
	if alias := mappedClusterNameToAlias(this.ClusterName); alias != "" {
		this.ClusterAlias = alias
	}
}
