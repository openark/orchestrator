package orcraft

import (
	"io"

	"github.com/hashicorp/raft"
)

// fsm is a raft finite state machine, that is freno aware. It applies events/commands
// onto the freno throttler.
type fsm Store

// Apply applies a Raft log entry to the key-value store.
func (f *fsm) Apply(l *raft.Log) interface{} {
	return nil
}

// Snapshot returns a snapshot object of freno's state
func (f *fsm) Snapshot() (raft.FSMSnapshot, error) {
	snapshot := newFsmSnapshot()
	return snapshot, nil
}

// Restore restores freno state
func (f *fsm) Restore(rc io.ReadCloser) error {
	defer rc.Close()
	return nil
}
