package orcraft

import (
	"encoding/json"
	"io"

	"github.com/openark/golib/log"

	"github.com/hashicorp/raft"
)

// fsm is a raft finite state machine
type fsm Store

// Apply applies a Raft log entry to the key-value store.
func (f *fsm) Apply(l *raft.Log) interface{} {
	var c storeCommand
	if err := json.Unmarshal(l.Data, &c); err != nil {
		log.Errorf("failed to unmarshal command: %s", err.Error())
	}

	return store.applier.ApplyCommand(c.Op, c.Value)
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
