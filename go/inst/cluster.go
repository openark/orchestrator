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
	"github.com/outbrain/orchestrator/go/config"
	"regexp"
	"strings"
)

// ClusterInfo makes for a cluster status/info summary
type ClusterInfo struct {
	ClusterName                            string
	ClusterAlias                           string // Human friendly alias
	CountInstances                         uint
	HeuristicLag                           int64
	HasAutomatedMasterRecovery             bool
	HasAutomatedIntermediateMasterRecovery bool
}

// ReadRecoveryInfo
func (this *ClusterInfo) ReadRecoveryInfo() {
	this.HasAutomatedMasterRecovery = this.filtersMatchCluster(config.Config.RecoverMasterClusterFilters)
	this.HasAutomatedIntermediateMasterRecovery = this.filtersMatchCluster(config.Config.RecoverIntermediateMasterClusterFilters)
}

// filtersMatchCluster will see whether the given filters match the given cluster details
func (this *ClusterInfo) filtersMatchCluster(filters []string) bool {
	for _, filter := range filters {
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
		} else if matched, _ := regexp.MatchString(filter, this.ClusterName); matched && filter != "" {
			return true
		}
	}
	return false
}
