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

package inst

/*
  query holds information about query metrics and records the time taken
  waiting before doing the query plus the time taken executing the query.
*/
import (
	"time"

	"github.com/github/orchestrator/go/collection"
	"github.com/montanaflynn/stats"
)

// Metric records query metrics of backend writes that go through
// a sized channel.  It allows us to compare the time waiting to
// execute the query against the time needed to run it and in a
// "sized channel" the wait time may be significant and is good to
// measure.
type WriteBufferMetric struct {
	Timestamp         time.Time     // time the metric was started
	WaitLatency       time.Duration // time that we had to wait before flushing the buffer
	EnqueuedInstances int
}

// NewMetric returns a new metric with timestamp starting from now
func NewWriteBufferMetric() *WriteBufferMetric {
	wbm := &WriteBufferMetric{
		Timestamp: time.Now(),
	}

	return wbm
}

// Update the metrics based on the given values
func (wbm *WriteBufferMetric) Update(instances int) {
	wbm.WaitLatency = time.Since(wbm.Timestamp)
	wbm.EnqueuedInstances = instances
}

// When records the timestamp of the start of the recording
func (m WriteBufferMetric) When() time.Time {
	return m.Timestamp
}

type AggregatedWriteBufferMetric struct {
	InstanceFlushIntervalMilliseconds int // config setting
	InstanceWriteBufferSize           int // config setting
	EnqueuedInstancesCount            int
	MaxEnqueuedInstances              float64
	MeanEnqueuedInstances             float64
	MedianEnqueuedInstances           float64
	P95EnqueuedInstances              float64
	MaxWaitSeconds                    float64
	MeanWaitSeconds                   float64
	MedianWaitSeconds                 float64
	P95WaitSeconds                    float64
}

// AggregatedSince returns the aggregated query metrics for the period
// given from the values provided.
func AggregatedSince(c *collection.Collection, t time.Time) AggregatedWriteBufferMetric {

	// Initialise metrics
	var enquedInstances []float64
	var waitTimings []float64

	// Retrieve values since the time specified
	values, err := c.Since(t)
	a := AggregatedWriteBufferMetric{}
	if err != nil {
		return a // empty data
	}

	// generate the metrics
	for _, v := range values {
		enquedInstances = append(enquedInstances, float64(v.(*WriteBufferMetric).EnqueuedInstances))
		waitTimings = append(waitTimings, v.(*WriteBufferMetric).WaitLatency.Seconds())
	}

	a.EnqueuedInstancesCount = len(waitTimings)

	// generate aggregate values
	if s, err := stats.Max(stats.Float64Data(enquedInstances)); err == nil {
		a.MaxEnqueuedInstances = s
	}
	if s, err := stats.Mean(stats.Float64Data(enquedInstances)); err == nil {
		a.MeanEnqueuedInstances = s
	}
	if s, err := stats.Median(stats.Float64Data(enquedInstances)); err == nil {
		a.MedianEnqueuedInstances = s
	}
	if s, err := stats.Percentile(stats.Float64Data(enquedInstances), 95); err == nil {
		a.P95EnqueuedInstances = s
	}
	if s, err := stats.Max(stats.Float64Data(waitTimings)); err == nil {
		a.MaxWaitSeconds = s
	}
	if s, err := stats.Mean(stats.Float64Data(waitTimings)); err == nil {
		a.MeanWaitSeconds = s
	}
	if s, err := stats.Median(stats.Float64Data(waitTimings)); err == nil {
		a.MedianWaitSeconds = s
	}
	if s, err := stats.Percentile(stats.Float64Data(waitTimings), 95); err == nil {
		a.P95WaitSeconds = s
	}

	return a
}
