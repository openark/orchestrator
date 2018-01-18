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

import (
	"time"

	"github.com/montanaflynn/stats"

	"github.com/github/orchestrator/go/collection"
	"github.com/github/orchestrator/go/config"
)

// WriteBufferMetric holds the result of a single buffered write attempt to the backend
type WriteBufferMetric struct {
	Timestamp    time.Time
	WriteLatency time.Duration // time that we had to wait before starting query execution
	Items        int           // number of rows written
	Err          error         // any error resulting from the query execution
}

// NewWriteBufferMetric returns a new metric with timestamp starting from now
func NewWriteBufferMetric() *WriteBufferMetric {
	wbm := &WriteBufferMetric{
		Timestamp: time.Now(),
	}

	return wbm
}

// Update the metrics based on the given values
func (wbm *WriteBufferMetric) Update(items int, err error) {
	wbm.WriteLatency = time.Since(wbm.Timestamp)
	wbm.Items = items
	wbm.Err = err
}

// When records the timestamp of the start of the recording
func (wbm WriteBufferMetric) When() time.Time {
	return wbm.Timestamp
}

// WriteBufferAggregate holds the results of a number of buffered write attempts to the backend
type WriteBufferAggregate struct {
	InstanceFlushIntervalMilliseconds int     // config setting
	InstanceWriteBufferSize           int     // config setting
	WriteCount                        int     // number of writes done
	ErrorCount                        int     // number of writes with errors
	SumInstancesWritten               int     // total number of rows written
	SumInstancesErrors                int     // total number of rows with errors
	MaxInstancesWritten               float64 // metrics for instances written in each call
	MeanInstancesWritten              float64
	MedianInstancesWritten            float64
	P75InstancesWritten               float64
	P95InstancesWritten               float64
	P99InstancesWritten               float64
	MaxWriteLatencySeconds            float64 // metrics for time to write each set of instances
	MeanWriteLatencySeconds           float64
	MedianWriteLatencySeconds         float64
	P75WriteLatencySeconds            float64
	P95WriteLatencySeconds            float64
	P99WriteLatencySeconds            float64
}

// WriteBufferAggregatedSince returns the aggregated for metrics in the collection since the specified time.
func WriteBufferAggregatedSince(c *collection.Collection, t time.Time) WriteBufferAggregate {
	var (
		writeTimings []float64
		writeItems   []float64
	)

	// Retrieve values since the time specified
	values, err := c.Since(t)
	wba := WriteBufferAggregate{
		InstanceFlushIntervalMilliseconds: config.Config.InstanceFlushIntervalMilliseconds,
		InstanceWriteBufferSize:           config.Config.InstanceWriteBufferSize,
	}

	if err != nil {
		return wba // empty data
	}

	// generate the metrics
	for _, v := range values {
		writeTimings = append(writeTimings, v.(*WriteBufferMetric).WriteLatency.Seconds())
		writeItems = append(writeItems, float64(v.(*WriteBufferMetric).Items))
		wba.SumInstancesWritten += v.(*WriteBufferMetric).Items
		if v.(*WriteBufferMetric).Err != nil {
			wba.ErrorCount++
			wba.SumInstancesErrors += v.(*WriteBufferMetric).Items
		}
	}

	wba.WriteCount = len(writeTimings)

	// generate aggregate timing metrics
	if s, err := stats.Max(stats.Float64Data(writeTimings)); err == nil {
		wba.MaxWriteLatencySeconds = s
	}
	if s, err := stats.Mean(stats.Float64Data(writeTimings)); err == nil {
		wba.MeanWriteLatencySeconds = s
	}
	if s, err := stats.Median(stats.Float64Data(writeTimings)); err == nil {
		wba.MedianWriteLatencySeconds = s
	}
	if s, err := stats.Percentile(stats.Float64Data(writeTimings), 75); err == nil {
		wba.P75WriteLatencySeconds = s
	}
	if s, err := stats.Percentile(stats.Float64Data(writeTimings), 95); err == nil {
		wba.P95WriteLatencySeconds = s
	}
	if s, err := stats.Percentile(stats.Float64Data(writeTimings), 99); err == nil {
		wba.P99WriteLatencySeconds = s
	}
	// generate aggregate item size metrics
	if s, err := stats.Max(stats.Float64Data(writeItems)); err == nil {
		wba.MaxInstancesWritten = s
	}
	if s, err := stats.Mean(stats.Float64Data(writeItems)); err == nil {
		wba.MeanInstancesWritten = s
	}
	if s, err := stats.Median(stats.Float64Data(writeItems)); err == nil {
		wba.MedianInstancesWritten = s
	}
	if s, err := stats.Percentile(stats.Float64Data(writeItems), 75); err == nil {
		wba.P75InstancesWritten = s
	}
	if s, err := stats.Percentile(stats.Float64Data(writeItems), 95); err == nil {
		wba.P95InstancesWritten = s
	}
	if s, err := stats.Percentile(stats.Float64Data(writeItems), 99); err == nil {
		wba.P99InstancesWritten = s
	}

	return wba
}
