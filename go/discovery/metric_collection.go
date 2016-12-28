package discovery

// Collect discovery metrics and manage their storage and retrieval for monitoring purposes.

import (
	"errors"
	"sync"
	"time"

	"github.com/outbrain/golib/log"

	"github.com/github/orchestrator/go/config"
)

// MetricCollection contains a collection of Metrics
type MetricCollection struct {
	sync.Mutex                 // for locking the structure
	collection   [](*Metric)   // may need impoving if the size of the collection grows too much
	done         chan struct{} // to indicate that we are finishing expiry
	ticker       *time.Ticker  // expiry ticker
        expirePeriod time.Duration // time to keep the collection information for
}

// NewMetricCollection returns the pointer to a new MetricCollection
func NewMetricCollection(period time.Duration) *MetricCollection {
	mc := &MetricCollection{
		collection: nil,
		done: make(chan struct{}),
		expirePeriod: period,
	}
	go mc.Expire()

	return mc
}

// Expire removes old values periodically given mc.expirePeriod
func (mc *MetricCollection) Expire() {
	if mc == nil {
		return
	}
	log.Infof("MetricCollection: Expiring values every second and keeping values for last %+v", mc.expirePeriod.String())
	mc.ticker = time.NewTicker(time.Second) // hard-coded at every second

	for {
		select {
		case <-mc.ticker.C: // do the periodic expiry
			mc.RemoveBefore(time.Now().Add(-mc.expirePeriod))
		case <-mc.done: // stop the ticker and return
			mc.ticker.Stop()
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
	if mc.collection == nil || len(mc.collection) == 0 {
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

// RemoveBefore collection values from mc before the given time.
func (mc *MetricCollection) RemoveBefore(t time.Time) error {
	if mc == nil {
		return errors.New("MetricsCollection.RemoveBefore: mc == nil")
	}
	mc.Lock()
	defer mc.Unlock()
	if mc.collection == nil {
		return nil // no data so nothing to do
	}
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
