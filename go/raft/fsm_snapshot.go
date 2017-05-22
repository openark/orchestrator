package orcraft

import (
	"github.com/hashicorp/raft"
)

// fsmSnapshot handles raft persisting of snapshots
type fsmSnapshot struct {
}

func newFsmSnapshot() *fsmSnapshot {
	return &fsmSnapshot{}
}

// Persist
func (f *fsmSnapshot) Persist(sink raft.SnapshotSink) error {
	return sink.Close()
}

// Release
func (f *fsmSnapshot) Release() {
}
