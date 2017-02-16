/*
   Copyright 2017 Simon J Mudd

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

package discovery

// Collect discovery metrics and manage their storage and retrieval for monitoring purposes.

import (
	"reflect"
	"time"

	"github.com/github/orchestrator/go/inst"
)

// Metric holds a set of information of instance discovery metrics
type Metric struct {
	Timestamp       time.Time        // time the collection was taken
	InstanceKey     inst.InstanceKey // instance being monitored
	BackendLatency  time.Duration    // time taken talking to the backend
	InstanceLatency time.Duration    // time taken talking to the instance
	TotalLatency    time.Duration    // total time taken doing the discovery
	Err             error            // error (if applicable) doing the discovery process
}

// When did the metric happen
func (m Metric) When() time.Time {
	return m.Timestamp
}

// MetricsEqual compares two slices of Metrics to see if they are the same
func MetricsEqual(m1, m2 [](*Metric)) bool {
	if len(m1) != len(m2) {
		return false
	}
	for i := range m1 {
		if !reflect.DeepEqual(m1[i],m2[i]) {
			return false
		}
	}
	return true
}
