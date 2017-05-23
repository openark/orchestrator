//
// Provide distributed consensus services.
// Underlying implementation is Raft, via https://godoc.org/github.com/hashicorp/raft
//
// This file provides generic access functions to setup & check group communication.
//

package orcraft

import (
	"fmt"
	"strings"
	"time"

	"github.com/github/orchestrator/go/config"
	"github.com/openark/golib/log"

	"github.com/hashicorp/raft"
)

const RaftDBFile = "freno-raft.db"

var store *Store

func IsRaftEnabled() bool {
	return store != nil
}

// Setup creates the entire raft shananga. Creates the store, associates with the throttler,
// contacts peer nodes, and subscribes to leader changes to export them.
func Setup() error {
	log.Debugf("Setting up raft")
	store = NewStore(config.Config.RaftDataDir, normalizeRaftNode(config.Config.RaftBind))
	peerNodes := []string{}
	for _, raftNode := range config.Config.RaftNodes {
		peerNodes = append(peerNodes, normalizeRaftNode(raftNode))
	}
	if err := store.Open(peerNodes); err != nil {
		return log.Errorf("failed to open raft store: %s", err.Error())
	}

	return nil
}

// getRaft is a convenience method
func getRaft() *raft.Raft {
	return store.raft
}

// normalizeRaftNode attempts to make sure there's a port to the given node.
// It consults the DefaultRaftPort when there isn't
func normalizeRaftNode(node string) string {
	if strings.Contains(node, ":") {
		return node
	}
	if config.Config.DefaultRaftPort == 0 {
		return node
	}
	node = fmt.Sprintf("%s:%d", node, config.Config.DefaultRaftPort)
	return node
}

// IsLeader tells if this node is the current raft leader
func IsLeader() bool {
	future := getRaft().VerifyLeader()
	err := future.Error()
	return err == nil
}

// GetLeader returns identity of raft leader
func GetLeader() string {
	return getRaft().Leader()
}

// GetState returns current raft state
func GetState() raft.RaftState {
	return getRaft().State()
}

// Monitor is a utility function to routinely observe leadership state.
// It doesn't actually do much; merely takes notes.
func Monitor() {
	t := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-t.C:
			leaderHint := GetLeader()

			if IsLeader() {
				leaderHint = fmt.Sprintf("%s (this host)", leaderHint)
			}
			log.Debugf("raft leader is %s; state: %s", leaderHint, GetState().String())
		}
	}
}
