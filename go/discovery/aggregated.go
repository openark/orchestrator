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
	"time"

	"github.com/montanaflynn/stats"
)

// this contains routines for cooked DiscoveryMetric values
// which can be called from api/discovery-metrics-aggregated/:seconds

type AggregatedDiscoveryMetrics struct {
	FirstSeen                       time.Time // timestamp of the first data seen
	LastSeen                        time.Time // timestamp of the last data seen
	DistinctInstanceKeys            int       // number of distinct Instances seen
	DistinctFailedInstanceKeys      int       // number of distinct Instances which failed
	FailedDiscoveries               uint64    // number of failed discoveries
	SuccessfulDiscoveries           uint64    // number of successful discoveries
	PollAvgTotalSeconds             float64   // Average Total Poll Time
	PollAvgBackendSeconds           float64
	PollAvgInstanceSeconds          float64
	PollFailedAvgTotalSeconds       float64 // Average Total Poll Time for failed discoveries
	PollFailedAvgBackendSeconds     float64
	PollFailedAvgInstanceSeconds    float64
	PollMaxTotalSeconds             float64
	PollMaxBackendSeconds           float64
	PollMaxInstanceSeconds          float64
	PollFailedMaxTotalSeconds       float64
	PollFailedMaxBackendSeconds     float64
	PollFailedMaxInstanceSeconds    float64
	PollMedianTotalSeconds          float64
	PollMedianBackendSeconds        float64
	PollMedianInstanceSeconds       float64
	PollFailedMedianTotalSeconds    float64
	PollFailedMedianBackendSeconds  float64
	PollFailedMedianInstanceSeconds float64
	PollQ95TotalSeconds             float64
	PollQ95BackendSeconds           float64
	PollQ95InstanceSeconds          float64
	PollFailedQ95TotalSeconds       float64
	PollFailedQ95BackendSeconds     float64
	PollFailedQ95InstanceSeconds    float64
}

func sum(values stats.Float64Data) float64 {
	var sum float64

	for _, v := range values {
		sum += v
	}
	return sum
}

func avg(values stats.Float64Data) float64 {
	length := len(values)

	if length == 0 {
		return 0
	}

	return sum(values) / float64(length)
}

// return the requested percentile value given.
func percentile(values stats.Float64Data, percent float64) float64 {
	s, err := stats.Percentile(values, percent)
	if err != nil {
		return 0
	}
	return s
}

// internal routine to return the maximum value or 0
func max(values stats.Float64Data) float64 {
	s, err := stats.Max(values)
	if err != nil {
		return 0
	}
	return s
}

// internal routine to return the median or 0
func median(values stats.Float64Data) float64 {
	s, err := stats.Median(values)
	if err != nil {
		return 0
	}
	return s
}

// AggregatedSince returns a large number of aggregated metrics
// based on the raw metrics collected since the given time.
func (mc *MetricCollection) AggregatedSince(t time.Time) AggregatedDiscoveryMetrics {
	results, err := mc.Since(t)
	if err != nil {
		return AggregatedDiscoveryMetrics{}
	}

	return aggregate(results)
}

func aggregate(results [](*Metric)) AggregatedDiscoveryMetrics {
	if len(results) == 0 {
		return AggregatedDiscoveryMetrics{}
	}

	var (
		first    time.Time
		last     time.Time
		counters map[string]uint64              // map of string based counters
		names    map[string](map[string]int)    // map of string based names (using a map)
		timings  map[string](stats.Float64Data) // map of string based float64 values
	)

	// initialise counters
	for _, v := range []string{"FailedDiscoveries", "Discoveries"} {
		counters[v] = 0
	}
	// initialise names
	for _, v := range []string{"InstanceKeys", "FailedInstanceKeys"} {
		names[v] = make(map[string]int)
	}
	// initialise timers
	for _, v := range []string{"TotalSeconds", "BackendSeconds", "InstanceSeconds", "FailedTotalSeconds", "FailedBackendSeconds", "FailedInstanceSeconds"} {
		timings[v] = nil
	}

	// iterate over results storing required values
	for _, v := range results {
		// first and last
		if first.IsZero() || first.After(v.Timestamp) {
			first = v.Timestamp
		}
		if last.Before(v.Timestamp) {
			last = v.Timestamp
		}
		// failed / successful names
		if v.Err == nil {
			x := names["InstanceKeys"]
			x[v.InstanceKey.String()] = 1 // Value doesn't matter
			names["InstanceKeys"] = x
		} else {
			x := names["FailedInstanceKeys"]
			x[v.InstanceKey.String()] = 1 // Value doesn't matter
			names["FailedInstanceKeys"] = x
		}

		counters["Discoveries"]++
		if v.Err == nil {
			counters["FailedDiscoveries"]++
		}

		timings["TotalSeconds"] = append(timings["TotalSeconds"], v.TotalLatency.Seconds())
		timings["BackendSeconds"] = append(timings["BackendSeconds"], v.BackendLatency.Seconds())
		timings["InstanceSeconds"] = append(timings["InstanceSeconds"], v.InstanceLatency.Seconds())

		if v.Err != nil {
			timings["FailedTotalSeconds"] = append(timings["FailedTotalSeconds"], v.TotalLatency.Seconds())
			timings["FailedBackendSeconds"] = append(timings["FailedBackendSeconds"], v.BackendLatency.Seconds())
			timings["FailedInstanceSeconds"] = append(timings["FailedInstanceSeconds"], v.InstanceLatency.Seconds())
		}

	}

	return AggregatedDiscoveryMetrics{
		FirstSeen:                       first,
		LastSeen:                        last,
		DistinctInstanceKeys:            len(names["InstanceKeys"]),
		DistinctFailedInstanceKeys:      len(names["FailedInstanceKeys"]),
		FailedDiscoveries:               counters["FailedDiscoveries"],
		SuccessfulDiscoveries:           counters["Discoveries"],
		PollAvgTotalSeconds:             avg(timings["TotalSeconds"]),
		PollAvgBackendSeconds:           avg(timings["BackendSeconds"]),
		PollAvgInstanceSeconds:          avg(timings["InstanceSeconds"]),
		PollFailedAvgTotalSeconds:       avg(timings["FailedTotalSeconds"]),
		PollFailedAvgBackendSeconds:     avg(timings["FailedBackendSeconds"]),
		PollFailedAvgInstanceSeconds:    avg(timings["FailedInstanceSeconds"]),
		PollMaxTotalSeconds:             max(timings["TotalSeconds"]),
		PollMaxBackendSeconds:           max(timings["BackendSeconds"]),
		PollMaxInstanceSeconds:          max(timings["InstanceSeconds"]),
		PollFailedMaxTotalSeconds:       max(timings["FailedTotalSeconds"]),
		PollFailedMaxBackendSeconds:     max(timings["FailedBackendSeconds"]),
		PollFailedMaxInstanceSeconds:    max(timings["FailedInstanceSeconds"]),
		PollMedianTotalSeconds:          median(timings["TotalSeconds"]),
		PollMedianBackendSeconds:        median(timings["BackendSeconds"]),
		PollMedianInstanceSeconds:       median(timings["InstanceSeconds"]),
		PollFailedMedianTotalSeconds:    median(timings["FailedTotalSeconds"]),
		PollFailedMedianBackendSeconds:  median(timings["FailedBackendSeconds"]),
		PollFailedMedianInstanceSeconds: median(timings["FailedInstanceSeconds"]),
		PollQ95TotalSeconds:             percentile(timings["TotalSeconds"], 95),
		PollQ95BackendSeconds:           percentile(timings["BackendSeconds"], 95),
		PollQ95InstanceSeconds:          percentile(timings["InstanceSeconds"], 95),
		PollFailedQ95TotalSeconds:       percentile(timings["FailedTotalSeconds"], 95),
		PollFailedQ95BackendSeconds:     percentile(timings["FailedBackendSeconds"], 95),
		PollFailedQ95InstanceSeconds:    percentile(timings["FailedInstanceSeconds"], 95),
	}
}
