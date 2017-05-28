/*
   Copyright 2015 Shlomi Noach, courtesy Booking.com

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

package process

import (
	"sync"
	"time"

	"github.com/github/orchestrator/go/config"

	"github.com/github/orchestrator/go/raft"
	"github.com/openark/golib/log"
)

const registrationPollSeconds = 5

type NodeHealth struct {
	Hostname        string
	Token           string
	AppVersion      string
	FirstSeenActive string
	LastSeenActive  string

	LastReported time.Time
	once         sync.Once
}

func NewNodeHealth() *NodeHealth {
	return &NodeHealth{
		Hostname:   ThisHostname,
		Token:      ProcessToken.Hash,
		AppVersion: config.RuntimeCLIFlags.ConfiguredVersion,
	}
}

var ThisNodeHealth = NewNodeHealth()

type HealthStatus struct {
	Healthy        bool
	Hostname       string
	Token          string
	IsActiveNode   bool
	ActiveNode     NodeHealth
	Error          error
	AvailableNodes [](*NodeHealth)
}

type OrchestratorExecutionMode string

const (
	OrchestratorExecutionCliMode  OrchestratorExecutionMode = "CLIMode"
	OrchestratorExecutionHttpMode                           = "HttpMode"
)

var continuousRegistrationOnce sync.Once

// HealthTest attempts to write to the backend database and get a result
func HealthTest() (*HealthStatus, error) {
	health := HealthStatus{Healthy: false, Hostname: ThisHostname, Token: ProcessToken.Hash}

	healthy, err := RegisterNode(ThisNodeHealth, "", "")
	if err != nil {
		health.Error = err
		return &health, log.Errore(err)
	}
	health.Healthy = healthy
	if orcraft.IsRaftEnabled() {
		health.ActiveNode.Hostname = orcraft.GetLeader()
		health.IsActiveNode = orcraft.IsLeader()
	} else {
		health.ActiveNode, health.IsActiveNode, err = ElectedNode()
		if err != nil {
			health.Error = err
			return &health, log.Errore(err)
		}
	}

	health.AvailableNodes, err = ReadAvailableNodes(true)

	return &health, nil
}

// ContinuousRegistration will continuously update the node_health
// table showing that the current process is still running.
func ContinuousRegistration(extraInfo string, command string) {
	continuousRegistrationOnce.Do(func() {
		tickOperation := func() {
			if _, err := RegisterNode(ThisNodeHealth, extraInfo, command); err != nil {
				log.Errorf("ContinuousRegistration: RegisterNode failed: %+v", err)
			}
		}
		// First one is synchronous
		tickOperation()
		go func() {
			registrationTick := time.Tick(time.Duration(registrationPollSeconds) * time.Second)
			for range registrationTick {
				// We already run inside a go-routine so
				// do not do this asynchronously.  If we
				// get stuck then we don't want to fill up
				// the backend pool with connections running
				// this maintenance operation.
				tickOperation()
			}
		}()
	})
}
