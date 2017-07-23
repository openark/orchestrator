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
	"fmt"
	"strings"
	"time"

	"github.com/github/orchestrator/go/config"
	"github.com/openark/golib/log"

	"github.com/hashicorp/raft"
)

const (
	YieldCommand     = "yield"
	YieldHintCommand = "yield-hint"
)

var RaftNotRunning = fmt.Errorf("raft is not configured/running")
var store *Store
var ThisHostname string

var fatalRaftErrorChan = make(chan error)

func IsRaftEnabled() bool {
	return store != nil
}

func FatalRaftError(err error) error {
	if err != nil {
		go func() { fatalRaftErrorChan <- err }()
	}
	return err
}

// Setup creates the entire raft shananga. Creates the store, associates with the throttler,
// contacts peer nodes, and subscribes to leader changes to export them.
func Setup(applier CommandApplier, thisHostname string) error {
	log.Debugf("Setting up raft")
	ThisHostname = thisHostname
	store = NewStore(config.Config.RaftDataDir, normalizeRaftNode(config.Config.RaftBind), applier)
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
	return GetState() == raft.Leader
}

// GetLeader returns identity of raft leader
func GetLeader() string {
	return getRaft().Leader()
}

func QuorumSize() (int, error) {
	peers, err := GetPeers()
	if err != nil {
		return 0, err
	}
	return len(peers)/2 + 1, nil
}

// GetState returns current raft state
func GetState() raft.RaftState {
	return getRaft().State()
}

func StepDown() {
	getRaft().StepDown()
}

func yield() error {
	if !IsRaftEnabled() {
		return RaftNotRunning
	}
	return getRaft().Yield()
}

func GetPeers() ([]string, error) {
	if !IsRaftEnabled() {
		return []string{}, RaftNotRunning
	}
	return store.peerStore.Peers()
}

func IsPeer(peer string) (bool, error) {
	if !IsRaftEnabled() {
		return false, RaftNotRunning
	}
	return (store.raftBind == peer), nil
}

// PublishCommand will distribute a command across the group
func PublishCommand(op string, value interface{}) (response interface{}, err error) {
	if !IsRaftEnabled() {
		return nil, RaftNotRunning
	}
	b, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	return store.genericCommand(op, b)
}

func PublishYield(toPeer string) (response interface{}, err error) {
	toPeer = normalizeRaftNode(toPeer)
	return store.genericCommand(YieldCommand, []byte(toPeer))
}

func PublishYieldHostnameHint(hostnameHint string) (response interface{}, err error) {
	return store.genericCommand(YieldHintCommand, []byte(hostnameHint))
}

// Monitor is a utility function to routinely observe leadership state.
// It doesn't actually do much; merely takes notes.
func Monitor() {
	t := time.Tick(5 * time.Second)
	heartbeat := time.Tick(1 * time.Minute)
	for {
		select {
		case <-t:
			leaderHint := GetLeader()

			if IsLeader() {
				leaderHint = fmt.Sprintf("%s (this host)", leaderHint)
			}
			log.Debugf("raft leader is %s; state: %s", leaderHint, GetState().String())

		case <-heartbeat:
			if IsLeader() {
				go PublishCommand("heartbeat", "")
			}
		case err := <-fatalRaftErrorChan:
			log.Fatale(err)
		}
	}
}
