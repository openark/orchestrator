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
	"math"
	"testing"
	"time"
)

// helper function convert seconds into a duration - it's just shorter!
func seconds(s float64) time.Duration {
	return time.Duration(int64(s * 1e9))
}

var first, middle, last time.Time
var input, inputPercentiles [](*Metric)

func init() {
	first = time.Now()
	middle = first.Add(time.Hour * 24)
	last = middle.Add(time.Hour * 24)

	input = [](*Metric){
		&Metric{Timestamp: last, TotalLatency: seconds(7.0), BackendLatency: seconds(6), InstanceLatency: seconds(4)},
		&Metric{Timestamp: middle, TotalLatency: seconds(7.2), BackendLatency: seconds(2), InstanceLatency: seconds(5)},
		&Metric{Timestamp: first, TotalLatency: seconds(8.0), BackendLatency: seconds(1), InstanceLatency: seconds(9)},
	}
	for i := 0; i <= 100; i++ {
		inputPercentiles = append(inputPercentiles, &Metric{Timestamp: first.Add(time.Hour * 24 * time.Duration(i)), TotalLatency: seconds(float64(i))})
	}
}

func TestFirstAndLastSeen(t *testing.T) {
	input := [](*Metric){
		&Metric{Timestamp: last},
		&Metric{Timestamp: middle},
		&Metric{Timestamp: first},
	}

	output := aggregate(input)

	if !output.FirstSeen.Equal(first) {
		t.Errorf("TestFirstAndLast: FirstSeen not as expected. Got: %v, Expected: %v", output.FirstSeen, first)
	}
	if !output.LastSeen.Equal(last) {
		t.Errorf("TestFirstAndLast: LastSeen not as expected. Got: %v, Expected: %v", output.LastSeen, last)
	}
}

func TestMax(t *testing.T) {
	output := aggregate(input)

	if output.MaxTotalSeconds != 8.0 {
		t.Errorf("TestMax: MaxTotalSeconds not as expected. Got: %.3f, Expected: %.3f", output.MaxTotalSeconds, 8.0)
	}
	if output.MaxBackendSeconds != 6.0 {
		t.Errorf("TestMax: MaxBackendSeconds not as expected. Got: %.3f, Expected: %.3f", output.MaxBackendSeconds, 6.0)
	}
	if output.MaxInstanceSeconds != 9.0 {
		t.Errorf("TestMax: MaxInstanceSeconds not as expected. Got: %.3f, Expected: %.3f", output.MaxInstanceSeconds, 9.0)
	}
}

func TestMean(t *testing.T) {
	output := aggregate(input)

	if math.Abs(output.MeanTotalSeconds-7.4) >= 0.0005 { // groan rounding differences
		t.Errorf("TestMean: MeanTotalSeconds not as expected. Got: %.3f, Expected: %.3f", output.MeanTotalSeconds, 7.4)
	}
	if output.MeanBackendSeconds != 3.0 {
		t.Errorf("TestMean: MeanBackendSeconds not as expected. Got: %.3f, Expected: %.3f", output.MeanBackendSeconds, 3.0)
	}
	if output.MeanInstanceSeconds != 6.0 {
		t.Errorf("TestMean: MeanInstanceSeconds not as expected. Got: %.3f, Expected: %.3f", output.MeanInstanceSeconds, 6.0)
	}
}

func TestMedian(t *testing.T) {
	output := aggregate(input)

	if output.MedianTotalSeconds != 7.2 {
		t.Errorf("TestMedian: MedianTotalSeconds not as expected. Got: %.3f, Expected: %.3f", output.MedianTotalSeconds, 7.4)
	}
	if output.MedianBackendSeconds != 2 {
		t.Errorf("TestMedian: MedianBackendSeconds not as expected. Got: %.3f, Expected: %.3f", output.MedianBackendSeconds, 2)
	}
	if output.MedianInstanceSeconds != 5 {
		t.Errorf("TestMedian: MedianInstanceSeconds not as expected. Got: %.3f, Expected: %.3f", output.MedianInstanceSeconds, 5)
	}
}

func TestPercentile(t *testing.T) {
	output := aggregate(inputPercentiles)

	if output.P95TotalSeconds != 95 {
		t.Errorf("TestPercentile: P95TotalSeconds not as expected. Got: %.3f, Expected: %.3f", output.P95TotalSeconds, 95.0)
	}
}
