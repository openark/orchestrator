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

import (
	"testing"
	"time"

	"github.com/github/orchestrator/go/inst"
)

var (
	dm  [](*Metric)
	dmc *MetricCollection
	ts  time.Time
)

func TestAppend(t *testing.T) {
	// some random base timestamp
	ts = time.Date(2016, 12, 27, 13, 36, 40, 0, time.Local)

	dm = make([](*Metric), 0)

	dm = append(dm, &Metric{
		Timestamp:       ts,
		InstanceKey:     inst.InstanceKey{Hostname: "host1", Port: 3306},
		BackendLatency:  time.Second,
		InstanceLatency: time.Second * 2,
	})
	dm = append(dm, &Metric{
		Timestamp:       ts.Add(time.Second),
		InstanceKey:     inst.InstanceKey{Hostname: "host2", Port: 3306},
		BackendLatency:  time.Second * 3,
		InstanceLatency: time.Second * 4,
	})
	dm = append(dm, &Metric{
		Timestamp:       ts.Add(time.Second * 2),
		InstanceKey:     inst.InstanceKey{Hostname: "host3", Port: 3306},
		BackendLatency:  time.Second * 5,
		InstanceLatency: time.Second * 6,
	})
	dmc = NewMetricCollection(time.Second)

	dmc.Append(dm[0])
	dmc.Append(dm[1])
	dmc.Append(dm[2])

	//	t.Logf("dmc: %+v", dmc)
	//	for i := range dmc.collection {
	//		t.Logf("  i: %d, %+v", i, *dmc.collection[i])
	//	}

	// retrieve metrics from ts + i Seconds
	for i := 0; i <= 2; i++ {
		refTime := ts.Add(time.Duration(i) * time.Second)
		if since, err := dmc.Since(refTime); err != nil {
			t.Errorf("TestAppend: Got Error trying to retrieve data since %+v", refTime)
		} else {
			// check expected values we get back
			if len(since) != (3 - i) {
				t.Errorf("TestAppend: len(since) not as expected. Got %d, expected: %d", len(since), 3-i)
				for j := range since {
					t.Logf("TestAppend: i: %d, %+v", i, *since[j])
				}
			}
			// check we get back the values we expect
			if !MetricsEqual(since, dm[i:]) {
				t.Errorf("TestAppend: since[%d] != dm[%d:%d]", i, i, len(dm))
				for j := range since {
					t.Logf("TestAppend: j: %d, Got %+v, expected: %+v", *since[j], dm[i+j])
				}
			}
		}
	}
}

// Add some values and then check that as we remove them we get back what we expect.
// This test is expected to run AFTER TestAppend
func TestremoveBefore(t *testing.T) {
	if len(dm) != 3 {
		t.Errorf("TestremoveBefore: dm is not the expected length. Got: %d, Expected: %d", len(dm), 3)
	}

	// we expect to have the 3 value in the array so remove them one by one.
	// the first iteration should remove nothing.
	for i := -1; i <= 3; i++ {
		refTime := ts.Add(time.Duration(i) * time.Second)
		// t.Logf("TestremoveBefore: i: %d, refTime: %+v", i, refTime)

		if err := dmc.removeBefore(refTime); err != nil {
			t.Errorf("TestremoveBefore: failed. i: %d, refTime: %+v, err: %+v", i, refTime, err)
		}
		since, err := dmc.Since(refTime)
		if err != nil {
			t.Errorf("TestremoveBefore: Got Error trying to retrieve data since %+v: %+v", refTime, err)
		} else {
			// check expected size we get back
			var expecting int
			if i < 0 {
				expecting = 3
			} else if i > 3 {
				expecting = 0
			} else {
				expecting = 3 - i
			}
			if len(since) != expecting {
				t.Errorf("TestremoveBefore(c): i: %d. len(since) wrong. Got: %d, Expecting: %d", i, len(since), expecting)
			}
		}
	}
}
