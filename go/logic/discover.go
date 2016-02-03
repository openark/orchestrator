package logic

import (
	"sync"

	"github.com/outbrain/golib/log"
	"github.com/outbrain/orchestrator/go/inst"
)

// record a current and maxUsage count and if we need to alert
type info struct {
	count int  // usage count of this instance
	max   int  // max usage count of this instnce (should be 1)
	alert bool // do we need to alert when we get to 0?
}

// data records how many of each instance are being looked at so
// we can see if we really have concurrency issues.
var (
	data map[inst.InstanceKey]info
	m    sync.Mutex
)

func init() {
	data = make(map[inst.InstanceKey]info)
}

// DiscoverStart records how many copies we are disconvering of each instance, and
// alerts if this number is > 1.
func DiscoverStart(instance inst.InstanceKey) {
	m.Lock()

	old, found := data[instance]
	if !found {
		data[instance] = info{count: 1, max: 1, alert: false}
	} else {
		old.count++
		log.Debugf("DiscoverStart(%+v) count: %d", instance, old.count)
		if old.count > old.max {
			old.max = old.count
			old.alert = true
		}
	}
	m.Unlock()
}

// DiscoverEnd reduces the usage count of each instance and alerts if the values
// were not as expected (i.e. max > 1)
func DiscoverEnd(instance inst.InstanceKey) {
	m.Lock()

	old, found := data[instance]
	if !found {
		log.Debugf("DiscoverEnd(%+v) not found", instance) // should not happen
	} else {
		old.count--
		if old.count == 0 {
			if old.alert {
				log.Debugf("DiscoverEnd(%+v) no longer used, max concurrent usage: %d", instance, old.max)
			}
			delete(data, instance)
		} else {
			log.Debugf("DiscoverEnd(%+v) count: %d", instance, old.count)
			data[instance] = old
		}
	}
	m.Unlock()
}
