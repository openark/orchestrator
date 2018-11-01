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
	"sort"
	"time"

	"github.com/openark/golib/log"
	"github.com/rcrowley/go-metrics"

	"github.com/github/orchestrator/go/config"
)

var (
	iwbWriceCountMetric            = metrics.NewCounter()
	iwbErrorCountMetric            = metrics.NewCounter()
	iwbInstanceErrorsMetric        = metrics.NewCounter()
	iwbInstanceWritesMetric        = metrics.NewHistogram(metrics.NewUniformSample(16384))
	iwbWriteLatencyMetric          = metrics.NewTimer()
	iwbInstanceFlushIntervalMetric = metrics.NewFunctionalGauge(func() int64 {
		return int64(config.Config.InstanceFlushIntervalMilliseconds)
	})
	iwbInstanceWriteBufferSizeMetric = metrics.NewFunctionalGauge(func() int64 {
		return int64(config.Config.InstanceWriteBufferSize)
	})
)

type instanceUpdateObject struct {
	instance                 *Instance
	instanceWasActuallyFound bool
	lastError                error
}

type InstanceWriteBuffer struct {
	instanceWriteBuffer chan instanceUpdateObject
	forceFlushChannel   chan bool
}

func NewInstanceWriteBuffer() *InstanceWriteBuffer {
	metrics.Register("instance_write_buffer.write_count", iwbWriceCountMetric)
	metrics.Register("instance_write_buffer.error_count", iwbErrorCountMetric)
	metrics.Register("instance_write_buffer.instance_errors", iwbInstanceErrorsMetric)
	metrics.Register("instance_write_buffer.instance_writes", iwbInstanceWritesMetric)
	metrics.Register("instance_write_buffer.write_latency", iwbWriteLatencyMetric)
	metrics.Register("instance_write_buffer.instance_flush_interval", iwbInstanceFlushIntervalMetric)
	metrics.Register("instance_write_buffer.instance_write_buffer_size", iwbInstanceWriteBufferSizeMetric)

	return &InstanceWriteBuffer{
		instanceWriteBuffer: make(chan instanceUpdateObject, config.Config.InstanceWriteBufferSize),
		forceFlushChannel:   make(chan bool),
	}
}

func (iwb *InstanceWriteBuffer) EnqueueInstanceWrite(instance *Instance, instanceWasActuallyFound bool, lastError error) {
	if len(iwb.instanceWriteBuffer) == config.Config.InstanceWriteBufferSize {
		// Signal the "flushing" gorouting that there's work.
		// We prefer doing all bulk flushes from one goroutine.
		iwb.forceFlushChannel <- true
	}
	iwb.instanceWriteBuffer <- instanceUpdateObject{instance, instanceWasActuallyFound, lastError}
}

func (iwb *InstanceWriteBuffer) FlushBufferLoop(interval time.Duration) {
	flushTick := time.Tick(interval)
	for {
		// it is time to flush
		select {
		case <-flushTick:
			iwb.FlushInstanceWriteBuffer()
		case <-iwb.forceFlushChannel:
			iwb.FlushInstanceWriteBuffer()
		}
	}
}

func (iwb *InstanceWriteBuffer) FlushBufferAsync() {
	iwb.forceFlushChannel <- true
}

// FlushInstanceWriteBuffer saves enqueued instances to Orchestrator Db
func (iwb *InstanceWriteBuffer) FlushInstanceWriteBuffer() {
	var instances []*Instance
	var lastseen []*Instance // instances to update with last_seen field

	if len(iwb.instanceWriteBuffer) == 0 {
		return
	}

	iwbWriteLatencyMetric.Time(func() {
		for i := 0; i < len(iwb.instanceWriteBuffer); i++ {
			upd := <-iwb.instanceWriteBuffer
			if upd.instanceWasActuallyFound && upd.lastError == nil {
				lastseen = append(lastseen, upd.instance)
			} else {
				instances = append(instances, upd.instance)
				log.Debugf("flushInstanceWriteBuffer: will not update database_instance.last_seen due to error: %+v", upd.lastError)
			}
		}
		// sort instances by instanceKey (table pk) to make locking predictable
		sort.Sort(byInstanceKey(instances))
		sort.Sort(byInstanceKey(lastseen))

		writeFunc := func() error {
			err := writeManyInstances(instances, true, false)
			if err != nil {
				return log.Errorf("flushInstanceWriteBuffer writemany: %v", err)
			}
			writeInstanceCounter.Inc(int64(len(instances)))

			err = writeManyInstances(lastseen, true, true)
			if err != nil {
				return log.Errorf("flushInstanceWriteBuffer last_seen: %v", err)
			}

			writeInstanceCounter.Inc(int64(len(lastseen)))
			return nil
		}
		err := ExecDBWriteFunc(writeFunc)
		if err != nil {
			log.Errorf("flushInstanceWriteBuffer: %v", err)
		}

		iwbInstanceWritesMetric.Update(int64(len(instances) + len(lastseen)))
		iwbWriceCountMetric.Inc(1)

		if err != nil {
			iwbErrorCountMetric.Inc(1)
			iwbInstanceErrorsMetric.Inc(int64(len(instances) + len(lastseen)))
		}
	})
}
