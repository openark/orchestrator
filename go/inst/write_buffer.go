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

	"github.com/rcrowley/go-metrics"

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

var (
	writeCount            = metrics.NewCounter()
	errorCount            = metrics.NewCounter()
	instancesErrors       = metrics.NewCounter()
	instancesWritten      = metrics.NewHistogram(metrics.NewUniformSample(16384))
	writeLatency          = metrics.NewTimer()
	instanceFlushInterval = metrics.NewFunctionalGauge(func() int64 {
		return int64(config.Config.InstanceFlushIntervalMilliseconds)
	})
	instanceWriteBufferSize = metrics.NewFunctionalGauge(func() int64 {
		return int64(config.Config.InstanceWriteBufferSize)
	})
)

func InitWriteBufferMetrics() {
	metrics.Register("write_buffer.write_count", writeCount)
	metrics.Register("write_buffer.error_count", errorCount)
	metrics.Register("write_buffer.instance_errors", instancesErrors)
	metrics.Register("write_buffer.instance_writes", instancesWritten)
	metrics.Register("write_buffer.write_latency", writeLatency)
	metrics.Register("constant.write_buffer.instance_flush_interval", instanceFlushInterval)
	metrics.Register("constant.write_buffer.instance_write_buffer_size", instanceWriteBufferSize)
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
	InstanceFlushIntervalMilliseconds int64 // config setting
	InstanceWriteBufferSize           int64 // config setting
	WriteCount                        int64 // number of writes done
	ErrorCount                        int64 // number of writes with errors
	SumInstancesWritten               int64 // total number of rows written
	SumInstancesErrors                int64 // total number of rows with errors
	MaxInstancesWritten               int64 // metrics for instances written in each call
	MeanInstancesWritten              float64
	MedianInstancesWritten            float64
	P75InstancesWritten               float64
	P95InstancesWritten               float64
	P99InstancesWritten               float64
	MaxWriteLatencySeconds            int64 // metrics for time to write each set of instances
	MeanWriteLatencySeconds           float64
	MedianWriteLatencySeconds         float64
	P75WriteLatencySeconds            float64
	P95WriteLatencySeconds            float64
	P99WriteLatencySeconds            float64
}

// WriteBufferAggregatedSince returns the aggregated for metrics in the collection since the specified time.
func WriteBufferAggregatedSince(c *collection.Collection, t time.Time) WriteBufferAggregate {

	// Retrieve values since the time specified
	values, err := c.Since(t)
	wba := WriteBufferAggregate{
		InstanceFlushIntervalMilliseconds: instanceFlushInterval.Value(),
		InstanceWriteBufferSize:           instanceWriteBufferSize.Value(),
	}

	if err != nil {
		return wba // empty data
	}

	// generate the metrics
	for _, v := range values {
		writeLatency.Update(v.(*WriteBufferMetric).WriteLatency)
		instancesWritten.Update(int64(v.(*WriteBufferMetric).Items))
		if v.(*WriteBufferMetric).Err != nil {
			errorCount.Inc(1)
			instancesErrors.Inc(int64(v.(*WriteBufferMetric).Items))
		}
	}

	wba.WriteCount = writeLatency.Count()
	wba.ErrorCount = errorCount.Count()
	wba.SumInstancesErrors = instancesErrors.Count()
	wba.SumInstancesWritten = instancesWritten.Sum()
	wba.MaxInstancesWritten = instancesWritten.Max()
	wba.MeanInstancesWritten = instancesWritten.Mean()
	wba.MedianInstancesWritten = instancesWritten.Percentile(50)
	wba.P75InstancesWritten = instancesWritten.Percentile(75)
	wba.P95InstancesWritten = instancesWritten.Percentile(95)
	wba.P99InstancesWritten = instancesWritten.Percentile(99)
	wba.MaxWriteLatencySeconds = writeLatency.Max()
	wba.MeanWriteLatencySeconds = writeLatency.Mean()
	wba.MedianWriteLatencySeconds = writeLatency.Percentile(50)
	wba.P75WriteLatencySeconds = writeLatency.Percentile(75)
	wba.P95WriteLatencySeconds = writeLatency.Percentile(95)
	wba.P99WriteLatencySeconds = writeLatency.Percentile(99)

	return wba
}
