package discovery

// Collect discovery metrics and manage their storage and retrieval for monitoring purposes.

import (
	"errors"
	"sync"
	"time"

	"github.com/outbrain/golib/log"

	"github.com/github/orchestrator/go/config"
)

// MC contains the last N discovery metrics which can then be accessed via an API call for monitoring.
// Currently this is accessed by ContinuousDiscovery() but also from http api calls.
// I may need to protect this better?
var MC *MetricCollection

// MetricCollection contains a collection of Metrics
type MetricCollection struct {
	sync.Mutex                 // for locking the structure
	collection   [](*Metric)   // may need impoving if the size of the collection grows too much
	done         chan struct{} // to indicate that we are finishing expiry
	expirePeriod time.Duration // time to keep the collection information for
}

// NewMetricCollection returns the pointer to a new MetricCollection
func NewMetricCollection(period time.Duration) *MetricCollection {
	mc := &MetricCollection{
		collection:   nil,
		done:         make(chan struct{}),
		expirePeriod: period,
	}
	go mc.autoExpire()

	return mc
}

// autoExpire is a private method which auto expires information
// periodically in the collection according to mc.expirePeriod.
// It will stop if it receives a message on channel mc.done.
func (mc *MetricCollection) autoExpire() {
	if mc == nil {
		return
	}
	ticker := time.NewTicker(time.Second) // hard-coded at every second

	for {
		select {
		case <-ticker.C: // do the periodic expiry
			mc.removeBefore(time.Now().Add(-mc.expirePeriod))
		case <-mc.done: // stop the ticker and return
			ticker.Stop()
			return
		}
	}
}

// Reload will check the retention period and adjust if needed.
func (mc *MetricCollection) Reload() {
	if mc == nil {
		return
	}
	mc.Lock()
	defer mc.Unlock()
	newExpirePeriod := time.Duration(config.Config.DiscoveryCollectionRetentionSeconds) * time.Second
	if mc.expirePeriod != newExpirePeriod {
		log.Infof("MetricCollection.Reload() Changing ExpirePeriod from %+v to %+v", mc.expirePeriod.String(), newExpirePeriod.String())
		mc.expirePeriod = time.Duration(config.Config.DiscoveryCollectionRetentionSeconds) * time.Second
	}
}

// Shutdown prepares to stop by terminating the expiry process
func (mc *MetricCollection) Shutdown() {
	log.Infof("MetricCollection.Shutdown: signalling mc.autoExpire() to stop")
	if mc == nil {
		return
	}
	mc.done <- struct{}{}
}

// Append a new Metric to the existing collection
func (mc *MetricCollection) Append(m *Metric) error {
	if mc == nil {
		return errors.New("MetricsCollection.Append: mc == nil")
	}
	mc.Lock()
	defer mc.Unlock()
	// we don't want to add nil metrics
	if m == nil {
		return errors.New("MetricsCollection.Append: m == nil")
	}
	// if no timestamp provided add one
	if m.Timestamp.IsZero() {
		m.Timestamp = time.Now()
	}
	mc.collection = append(mc.collection, m)

	return nil
}

// Since returns the Metrics on or after the given time. We assume
// the metrics are stored in ascending time.
// Iterate backwards until we reach the first value before the given time
// or the end of the array.
func (mc *MetricCollection) Since(t time.Time) ([](*Metric), error) {
	if mc == nil {
		return nil, errors.New("MetricsCollection.Since: mc == nil")
	}
	mc.Lock()
	defer mc.Unlock()
	if len(mc.collection) == 0 {
		return nil, nil // nothing to return
	}
	last := len(mc.collection)
	first := last - 1

	done := false
	for !done {
		if mc.collection[first].Timestamp.After(t) || mc.collection[first].Timestamp.Equal(t) {
			if first == 0 {
				break // as can't go lower
			}
			first--
		} else {
			if first != last {
				first++ // go back one (except if we're already at the end)
			}
			break
		}
	}

	return mc.collection[first:last], nil
}

// removeBefore is called by autoExpire and removes collection values
// from mc before the given time.
func (mc *MetricCollection) removeBefore(t time.Time) error {
	if mc == nil {
		return errors.New("MetricsCollection.removeBefore: mc == nil")
	}
	mc.Lock()
	defer mc.Unlock()

	cLen := len(mc.collection)
	if cLen == 0 {
		return nil // we have a collection but no data
	}
	// remove old data here.
	first := 0
	done := false
	for !done {
		if mc.collection[first].Timestamp.Before(t) {
			first++
			if first == cLen {
				break
			}
		} else {
			first--
			break
		}
	}

	// get the interval we need.
	if first == len(mc.collection) {
		mc.collection = nil // remove all entries
	} else if first != -1 {
		mc.collection = mc.collection[first:]
	}
	return nil // no errors
}
