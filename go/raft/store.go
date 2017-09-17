package orcraft

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/openark/golib/log"

	"github.com/hashicorp/raft"
)

const (
	retainSnapshotCount = 2880
	snapshotInterval    = 10 * time.Second
	raftTimeout         = 10 * time.Second
)

type Store struct {
	raftDir  string
	raftBind string

	raft      *raft.Raft // The consensus mechanism
	peerStore raft.PeerStore

	applier                CommandApplier
	snapshotCreatorApplier SnapshotCreatorApplier
}

type storeCommand struct {
	Op    string `json:"op,omitempty"`
	Value []byte `json:"value,omitempty"`
}

// NewStore inits and returns a new store
func NewStore(raftDir string, raftBind string, applier CommandApplier, snapshotCreatorApplier SnapshotCreatorApplier) *Store {
	return &Store{
		raftDir:                raftDir,
		raftBind:               raftBind,
		applier:                applier,
		snapshotCreatorApplier: snapshotCreatorApplier,
	}
}

// Open opens the store. If enableSingle is set, and there are no existing peers,
// then this node becomes the first node, and therefore leader, of the cluster.
func (store *Store) Open(peerNodes []string) error {
	// Setup Raft configuration.
	config := raft.DefaultConfig()
	config.SnapshotThreshold = 1
	config.SnapshotInterval = snapshotInterval

	// Setup Raft communication.
	addr, err := net.ResolveTCPAddr("tcp", store.raftBind)
	if err != nil {
		return err
	}
	log.Debugf("raft: addr=%+v", addr)
	transport, err := raft.NewTCPTransport(store.raftBind, addr, 3, 10*time.Second, os.Stderr)
	if err != nil {
		return err
	}
	log.Debugf("raft: transport=%+v", transport)

	peers := make([]string, 0, 10)
	for _, peerNode := range peerNodes {
		peerNode = strings.TrimSpace(peerNode)
		peers = raft.AddUniquePeer(peers, peerNode)
	}
	log.Debugf("raft: peers=%+v", peers)

	// Create peer storage.
	peerStore := &raft.StaticPeers{}
	if err := peerStore.SetPeers(peers); err != nil {
		return err
	}

	// Allow the node to enter single-mode, potentially electing itself, if
	// explicitly enabled and there is only 1 node in the cluster already.
	if len(peerNodes) == 0 && len(peers) <= 1 {
		log.Infof("enabling single-node mode")
		config.EnableSingleNode = true
		config.DisableBootstrapAfterElect = false
	}

	// Create the snapshot store. This allows the Raft to truncate the log.
	snapshots, err := raft.NewFileSnapshotStore(store.raftDir, retainSnapshotCount, os.Stderr)
	//snapshots, err := NewRelSnapshotStore(retainSnapshotCount, os.Stderr)
	if err != nil {
		return log.Errorf("file snapshot store: %s", err)
	}
	//	snapshots := raft.NewDiscardSnapshotStore()

	// Create the log store and stable store.
	logStore := NewRelationalStore()
	log.Debugf("raft: logStore=%+v", logStore)

	// Instantiate the Raft systems.
	if store.raft, err = raft.NewRaft(config, (*fsm)(store), logStore, logStore, snapshots, peerStore, transport); err != nil {
		return fmt.Errorf("error creating new raft: %s", err)
	}
	store.peerStore = peerStore
	log.Infof("new raft created")

	return nil
}

// Join joins a node, located at addr, to this store. The node must be ready to
// respond to Raft communications at that address.
func (store *Store) Join(addr string) error {
	log.Infof("received join request for remote node as %s", addr)

	f := store.raft.AddPeer(addr)
	if f.Error() != nil {
		return f.Error()
	}
	log.Infof("node at %s joined successfully", addr)
	return nil
}

// genericCommand requests consensus for applying a single command.
// This is an internal orchestrator implementation
func (store *Store) genericCommand(op string, bytes []byte) (response interface{}, err error) {
	if store.raft.State() != raft.Leader {
		return nil, fmt.Errorf("not leader")
	}

	b, err := json.Marshal(&storeCommand{Op: op, Value: bytes})
	if err != nil {
		return nil, err
	}

	f := store.raft.Apply(b, raftTimeout)
	if err = f.Error(); err != nil {
		return nil, err
	}
	r := f.Response()
	return r, nil
}
