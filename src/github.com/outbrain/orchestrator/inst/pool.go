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
	"github.com/outbrain/golib/log"
	"strings"
)

// PoolInstancesMap lists instance keys per pool name
type PoolInstancesMap map[string]([]*InstanceKey)

// ClusterPoolInstance is an instance mapping a cluster, pool & instance
type ClusterPoolInstance struct {
	ClusterName  string
	ClusterAlias string
	Pool         string
	Hostname     string
	Port         int
}

func ApplyPoolInstances(pool string, instancesList string) error {
	var instanceKeys [](*InstanceKey)
	if instancesList != "" {
		instancesStrings := strings.Split(instancesList, ",")
		for _, instanceString := range instancesStrings {

			instanceKey, err := ParseInstanceKeyLoose(instanceString)
			log.Debugf("%+v", instanceKey)
			if err != nil {
				return log.Errore(err)
			}

			instanceKeys = append(instanceKeys, instanceKey)
		}
	}
	writePoolInstances(pool, instanceKeys)
	return nil
}
