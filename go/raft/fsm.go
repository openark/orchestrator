/*
   Copyright 2017 Shlomi Noach, GitHub Inc.

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

	if c.Op == "yield" {
		toPeer := string(c.Value)
		isThisPeer, err := IsPeer(toPeer)
		if err != nil {
			return log.Errorf("failed to unmarshal command: %s", err.Error())
		}
		if isThisPeer {
			log.Debugf("Will not yield to myself")
			return nil
		}
		log.Debugf("Yielding to %s", toPeer)
		return yield()
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
