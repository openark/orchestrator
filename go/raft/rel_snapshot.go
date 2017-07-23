package orcraft

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	golog "log"
	"os"
	"time"

	"github.com/github/orchestrator/go/db"

	"github.com/hashicorp/raft"
	"github.com/openark/golib/log"
	"github.com/openark/golib/sqlutils"
)

type dummyReadCloser struct {
}

func (c *dummyReadCloser) Close() error {
	return nil
}

func (c *dummyReadCloser) Read(p []byte) (n int, err error) {
	// return 0, io.EOF
	return 0, fmt.Errorf("==== dummyReadCloser refuses to read snapshot, bwahahahahahahaha")
}

// RelSnapshotStore implements the SnapshotStore interface and allows
// snapshots to be made on the local disk.
type RelSnapshotStore struct {
	retain int
	logger *golog.Logger
}

// RelSnapshotSink implements SnapshotSink with a file.
type RelSnapshotSink struct {
	store  *RelSnapshotStore
	logger *golog.Logger
	meta   relSnapshotMeta

	closed bool
}

// relSnapshotMeta is stored on disk. We also put a CRC
// on disk so that we can verify the snapshot.
type relSnapshotMeta struct {
	raft.SnapshotMeta
	CRC []byte
}

// bufferedFile is returned when we open a snapshot. This way
// reads are buffered and the file still gets closed.
type bufferedFile struct {
	bh *bufio.Reader
	fh *os.File
}

func (b *bufferedFile) Read(p []byte) (n int, err error) {
	log.Debugf("===== bufferedFile.Read")
	return b.bh.Read(p)
}

func (b *bufferedFile) Close() error {
	return b.fh.Close()
}

// NewRelSnapshotStoreWithLogger creates a new RelSnapshotStore based
// on a base directory. The `retain` parameter controls how many
// snapshots are retained. Must be at least 1.
func NewRelSnapshotStoreWithLogger(retain int, logger *golog.Logger) (*RelSnapshotStore, error) {
	if retain < 1 {
		return nil, log.Errorf("must retain at least one snapshot")
	}
	if logger == nil {
		logger = golog.New(os.Stderr, "", golog.LstdFlags)
	}

	// Setup the store
	store := &RelSnapshotStore{
		retain: retain,
		logger: logger,
	}

	return store, nil
}

// NewRelSnapshotStore creates a new RelSnapshotStore based
// on a base directory. The `retain` parameter controls how many
// snapshots are retained. Must be at least 1.
func NewRelSnapshotStore(retain int, logOutput io.Writer) (*RelSnapshotStore, error) {
	if logOutput == nil {
		logOutput = os.Stderr
	}
	return NewRelSnapshotStoreWithLogger(retain, golog.New(logOutput, "", golog.LstdFlags))
}

// snapshotName generates a name for the snapshot.
func snapshotName(term, index uint64) string {
	now := time.Now()
	msec := now.UnixNano() / int64(time.Millisecond)
	return fmt.Sprintf("%d-%d-%d", term, index, msec)
}

// Create is used to start a new snapshot
func (f *RelSnapshotStore) Create(index, term uint64, peers []byte) (raft.SnapshotSink, error) {
	// Create a new path
	name := snapshotName(term, index)

	// Create the sink
	sink := &RelSnapshotSink{
		store:  f,
		logger: f.logger,
		meta: relSnapshotMeta{
			SnapshotMeta: raft.SnapshotMeta{
				ID:    name,
				Index: index,
				Term:  term,
				Peers: peers,
			},
			CRC: nil,
		},
	}

	// Write out the meta data
	if err := sink.writeMeta(); err != nil {
		f.logger.Printf("[ERR] snapshot: Failed to write metadata: %v", err)
		return nil, log.Errore(err)
	}

	// Done
	return sink, nil
}

// List returns available snapshots in the store.
func (f *RelSnapshotStore) List() ([]*raft.SnapshotMeta, error) {
	log.Debugf("===== RelSnapshotStore.List")
	// Get the eligible snapshots
	snapshots, err := f.getSnapshots()
	if err != nil {
		f.logger.Printf("[ERR] snapshot: Failed to get snapshots: %v", err)
		return nil, log.Errore(err)
	}

	var snapMeta []*raft.SnapshotMeta
	for _, meta := range snapshots {
		snapMeta = append(snapMeta, &meta.SnapshotMeta)
		if len(snapMeta) == f.retain {
			break
		}
	}
	return snapMeta, nil
}

// readSnapshots reads snapshots by query
func (f *RelSnapshotStore) readSnapshots(query string, args []interface{}) (snapMeta []*relSnapshotMeta, err error) {
	log.Debugf("===== RelSnapshotStore.readSnapshots; query=%+v", query)
	err = db.QueryOrchestrator(query, args, func(m sqlutils.RowMap) error {
		snapshotMetaText := m.GetString("snapshot_meta")

		meta := &relSnapshotMeta{}
		if err := json.Unmarshal([]byte(snapshotMetaText), &meta); err != nil {
			return err
		}

		snapMeta = append(snapMeta, meta)
		return nil
	})
	return snapMeta, log.Errore(err)
}

func (f *RelSnapshotStore) readMeta(name string) (*relSnapshotMeta, error) {
	log.Debugf("===== RelSnapshotStore.readMeta; name=%+v", name)
	query := `select snapshot_meta from raft_snapshot where snapshot_name=?`
	snapshots, err := f.readSnapshots(query, sqlutils.Args(name))
	if err != nil {
		return nil, log.Errore(err)
	}
	if len(snapshots) == 1 {
		return snapshots[0], nil
	}
	return nil, log.Errorf("Found %+v snapshots for %s", len(snapshots), name)
}

// getSnapshots returns all the known snapshots.
func (f *RelSnapshotStore) getSnapshots() (snapMeta []*relSnapshotMeta, err error) {
	query := `select snapshot_meta from raft_snapshot order by snapshot_id desc`
	return f.readSnapshots(query, sqlutils.Args())
}

// Open takes a snapshot ID and returns a ReadCloser for that snapshot.
func (f *RelSnapshotStore) Open(id string) (*raft.SnapshotMeta, io.ReadCloser, error) {
	log.Debugf("===== RelSnapshotStore.Open; id=%+v", id)
	// Get the metadata
	meta, err := f.readMeta(id)
	if err != nil {
		f.logger.Printf("[ERR] snapshot: Failed to get meta data to open snapshot: %v", err)
		return nil, nil, log.Errore(err)
	}
	return &meta.SnapshotMeta, &dummyReadCloser{}, nil
}

// ReapSnapshots reaps any snapshots beyond the retain count.
func (f *RelSnapshotStore) ReapSnapshots() error {
	snapshots, err := f.getSnapshots()
	if err != nil {
		f.logger.Printf("[ERR] snapshot: Failed to get snapshots: %v", err)
		return log.Errore(err)
	}

	for i := f.retain; i < len(snapshots); i++ {
		if _, err := db.ExecOrchestrator(`delete from raft_snapshot where snapshot_name=?`, snapshots[i].ID); err != nil {
			f.logger.Printf("[ERR] snapshot: Failed to reap snapshot %v: %v", snapshots[i].ID, err)
			return log.Errore(err)
		}
	}
	return nil
}

// ID returns the ID of the snapshot, can be used with Open()
// after the snapshot is finalized.
func (s *RelSnapshotSink) ID() string {
	return s.meta.ID
}

// Write is used to append to the state file. We write to the
// buffered IO object to reduce the amount of context switches.
func (s *RelSnapshotSink) Write(b []byte) (int, error) {
	return 0, nil
}

// Close is used to indicate a successful end.
func (s *RelSnapshotSink) Close() error {
	// Make sure close is idempotent
	if s.closed {
		return nil
	}
	s.closed = true
	// Write out the meta data
	if err := s.writeMeta(); err != nil {
		s.logger.Printf("[ERR] snapshot: Failed to write metadata: %v", err)
		return log.Errore(err)
	}

	// Reap any old snapshots
	if err := s.store.ReapSnapshots(); err != nil {
		return log.Errore(err)
	}

	return nil
}

// Cancel is used to indicate an unsuccessful end.
func (s *RelSnapshotSink) Cancel() error {
	// Make sure close is idempotent
	if s.closed {
		return nil
	}
	s.closed = true

	return nil
}

// writeMeta is used to write out the metadata we have.
func (s *RelSnapshotSink) writeMeta() error {
	b, err := json.Marshal(&s.meta)
	if err != nil {
		return log.Errore(err)
	}
	snapshotMetaText := string(b)
	_, err = db.ExecOrchestrator(`
		replace into raft_snapshot
			(snapshot_id, snapshot_name, snapshot_meta, created_at)
		values
			(null, ?, ?, now())
			`, s.meta.ID, snapshotMetaText)
	return log.Errore(err)
}
