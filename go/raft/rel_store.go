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
	"encoding/binary"

	"github.com/github/orchestrator/go/db"

	"github.com/hashicorp/raft"
)

// RelationalStoreimplements:
// - hashicorp/raft.StableStore
// - hashicorp/log.LogStore
type RelationalStore struct {
}

func NewRelationalStore() *RelationalStore {
	return &RelationalStore{}
}

func (relStore *RelationalStore) Set(key []byte, val []byte) error {
	db, err := db.OpenOrchestrator()
	if err != nil {
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(`delete from raft_store where store_key = ?`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(key)
	if err != nil {
		tx.Rollback()
		return err
	}
	stmt, err = tx.Prepare(`insert into raft_store (store_key, store_value) values (?, ?)`)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(key, val)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()

	return err
}

// Get returns the value for key, or an empty byte slice if key was not found.
func (relStore *RelationalStore) Get(key []byte) (val []byte, err error) {
	db, err := db.OpenOrchestrator()
	if err != nil {
		return val, err
	}
	err = db.QueryRow("select min(store_value) from raft_store where store_key = ?", key).Scan(&val)
	return val, err
}

func (relStore *RelationalStore) SetUint64(key []byte, val uint64) error {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, val)

	return relStore.Set(key, b)
}

// GetUint64 returns the uint64 value for key, or 0 if key was not found.
func (relStore *RelationalStore) GetUint64(key []byte) (uint64, error) {
	b, err := relStore.Get(key)
	if err != nil {
		return 0, err
	}
	if len(b) == 0 {
		// Not found
		return 0, nil
	}
	i := binary.LittleEndian.Uint64(b)
	return i, nil
}

func (relStore *RelationalStore) FirstIndex() (idx uint64, err error) {
	db, err := db.OpenOrchestrator()
	if err != nil {
		return idx, err
	}
	err = db.QueryRow("select ifnull(min(log_index), 0) from raft_log").Scan(&idx)
	return idx, err
}

// LastIndex returns the last index written. 0 for no entries.
func (relStore *RelationalStore) LastIndex() (idx uint64, err error) {
	db, err := db.OpenOrchestrator()
	if err != nil {
		return idx, err
	}
	err = db.QueryRow("select ifnull(max(log_index), 0) from raft_log").Scan(&idx)
	return idx, err
}

// GetLog gets a log entry at a given index.
func (relStore *RelationalStore) GetLog(index uint64, log *raft.Log) error {
	db, err := db.OpenOrchestrator()
	if err != nil {
		return err
	}
	err = db.QueryRow(`
    select log_index, term, log_type, data
      from raft_log
      where log_index = ?
    `, index).Scan(&log.Index, &log.Term, &log.Type, &log.Data)
	return err
}

// StoreLog stores a log entry.
func (relStore *RelationalStore) StoreLog(log *raft.Log) error {
	return relStore.StoreLogs([]*raft.Log{log})
}

// StoreLogs stores multiple log entries.
func (relStore *RelationalStore) StoreLogs(logs []*raft.Log) error {
	db, err := db.OpenOrchestrator()
	if err != nil {
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(`
    replace into raft_log (
      log_index, term, log_type, data
    ) values (
      ?, ?, ?, ?
    )`)
	if err != nil {
		return err
	}
	for _, raftLog := range logs {
		_, err = stmt.Exec(raftLog.Index, raftLog.Term, int(raftLog.Type), raftLog.Data)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

// DeleteRange deletes a range of log entries. The range is inclusive.
func (relStore *RelationalStore) DeleteRange(min, max uint64) error {
	_, err := db.ExecOrchestrator("delete from raft_log where log_index >= ? and log_index <= ?", min, max)
	return err
}
