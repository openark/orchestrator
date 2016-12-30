/*
   Copyright 2016 Simon J Mudd

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

// Equal compares if to Metrics are the same
func (m *Metric) Equal(m2 *Metric) bool {
	if m == nil && m2 == nil {
		return true // assume the same "empty" value
	}
	if m == nil || m2 == nil {
		return false // one or the other is "empty" so they must be different
	}
	return m.Timestamp == m2.Timestamp &&
		m.InstanceKey == m2.InstanceKey &&
		m.BackendLatency == m2.BackendLatency &&
		m.InstanceLatency == m2.InstanceLatency &&
		m.TotalLatency == m2.TotalLatency &&
		m.Err == m2.Err
}

// MetricsEqual compares two slices of Metrics to see if they are the same
func MetricsEqual(m1, m2 [](*Metric)) bool {
	if len(m1) != len(m2) {
		return false
	}
	for i := range m1 {
		if !m1[i].Equal(m2[i]) {
			return false
		}
	}
	return true
}
