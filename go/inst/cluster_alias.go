/*
   Copyright 2015 Shlomi Noach, courtesy Booking.com

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
	"github.com/outbrain/orchestrator/go/config"
	"regexp"
	"sync"
)

// clusterAlias maps a cluster name to an alias
var clusterAliasMap = make(map[string]string)
var clusterAliasMapMutex = &sync.Mutex{}

func ApplyClusterAlias(clusterInfo *ClusterInfo) {
	for pattern, _ := range config.Config.ClusterNameToAlias {
		if matched, _ := regexp.MatchString(pattern, clusterInfo.ClusterName); matched {
			clusterInfo.ClusterAlias = config.Config.ClusterNameToAlias[pattern]
		}
	}
	clusterAliasMapMutex.Lock()
	if alias, ok := clusterAliasMap[clusterInfo.ClusterName]; ok {
		clusterInfo.ClusterAlias = alias
	}
	clusterAliasMapMutex.Unlock()
}

// SetClusterAlias will write (and override) a single cluster name mapping
func SetClusterAlias(clusterName string, alias string) error {
	err := WriteClusterAlias(clusterName, alias)
	if err != nil {
		return err
	}
	clusterAliasMapMutex.Lock()
	clusterAliasMap[clusterName] = alias
	clusterAliasMapMutex.Unlock()
	return nil
}

// GetClusterByAlias returns the cluster name associated with given alias.
// The function returns with error when:
// - No cluster is associated with the alias
// - More than one cluster is associated with the alias
func GetClusterByAlias(alias string) (string, error) {
	clusterName := ""
	clusterAliasMapMutex.Lock()
	defer clusterAliasMapMutex.Unlock()

	for mappedName, mappedAlias := range clusterAliasMap {
		if mappedAlias == alias {
			if clusterName == "" {
				clusterName = mappedName
			} else {
				return "", fmt.Errorf("GetClusterByAlias: multiple clusters for alias %s", alias)
			}
		}
	}
	if clusterName == "" {
		return "", fmt.Errorf("GetClusterByAlias: no cluster found for alias %s", alias)
	}
	return clusterName, nil
}
