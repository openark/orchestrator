package discovery

// Collect discovery metrics and manage their storage and retrieval for monitoring purposes.

import (
	"errors"
	"sync"
	"time"

	"github.com/github/orchestrator/go/inst"
)

// Metric holds a set of information of instance discovery metrics
type Metric struct {
	Timestamp       time.Time        // time the collection was taken
	InstanceKey     inst.InstanceKey // instance being monitored
	BackendLatency  time.Duration    // time taken talking to the backend
	InstanceLatency time.Duration    // time taken talking to the instance
	Err             error            // error (if applicable) doing the discovery process
}

// Equal compares if to Metrics are the same
func (m *Metric) Equal(m2 *Metric) bool {
	if m == nil && m2 == nil {
		return true // assume the same "empty" value
	}
	if m == nil || m2 == nil {
		return false // one or the other is "empty" so they must be different
	}
	return m.Timestamp == m2.Timestamp &&
		m.InstanceKey == m2.InstanceKey &&
		m.BackendLatency == m2.BackendLatency &&
		m.InstanceLatency == m2.InstanceLatency &&
		m.Err == m2.Err
}

// MetricsEqual compares two slices of Metrics to see if they are the same
func MetricsEqual(m1, m2 [](*Metric)) bool {
	if len(m1) != len(m2) {
		return false
	}
	for i := range m1 {
		if !m1[i].Equal(m2[i]) {
			return false
		}
	}
	return true
}

// MetricCollection contains a collection of Metrics
type MetricCollection struct {
	sync.Mutex             // for locking the structure
	collection [](*Metric) // may need impoving if the size of the collection grows too much
}

// NewMetricCollection returns the pointer to a new MetricCollection
func NewMetricCollection() *MetricCollection {
	mc := &MetricCollection{
		collection: nil,
	}
	return mc
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
