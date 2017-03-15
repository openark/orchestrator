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
	"time"

	"github.com/github/orchestrator/go/config"
	"github.com/github/orchestrator/go/db"
	"github.com/openark/golib/log"
	"github.com/openark/golib/sqlutils"
)

const registrationPollSeconds = 10

type NodeHealth struct {
	Hostname        string
	Token           string
	AppVersion      string
	FirstSeenActive string
	LastSeenActive  string
}

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

var continuousRegistrationInitiated bool = false

// RegisterNode writes down this node in the node_health table
func RegisterNode(extraInfo string, command string, firstTime bool) (healthy bool, err error) {
	if firstTime {
		db.ExecOrchestrator(`
			insert ignore into node_health_history
				(hostname, token, first_seen_active, extra_info, command, app_version)
			values
				(?, ?, NOW(), ?, ?, ?)
			`,
			ThisHostname, ProcessToken.Hash, extraInfo, command,
			config.RuntimeCLIFlags.ConfiguredVersion,
		)
	}
	{
		sqlResult, err := db.ExecOrchestrator(`
			update node_health set
				last_seen_active = now(),
				extra_info = case when ? != '' then ? else extra_info end,
				app_version = ?
			where
				hostname = ?
				and token = ?
			`,
			extraInfo, extraInfo, config.RuntimeCLIFlags.ConfiguredVersion, ThisHostname, ProcessToken.Hash,
		)
		if err != nil {
			return false, log.Errore(err)
		}
		rows, err := sqlResult.RowsAffected()
		if err != nil {
			return false, log.Errore(err)
		}
		if rows > 0 {
			return true, nil
		}
	}
	{
		sqlResult, err := db.ExecOrchestrator(`
			insert ignore into node_health
				(hostname, token, first_seen_active, last_seen_active, extra_info, command, app_version)
			values
				(?, ?, now(), now(), ?, ?, ?)
			`,
			ThisHostname, ProcessToken.Hash, extraInfo, command,
			config.RuntimeCLIFlags.ConfiguredVersion,
		)
		if err != nil {
			return false, log.Errore(err)
		}
		rows, err := sqlResult.RowsAffected()
		if err != nil {
			return false, log.Errore(err)
		}
		if rows > 0 {
			return true, nil
		}
	}
	return false, nil
}

// HealthTest attempts to write to the backend database and get a result
func HealthTest() (*HealthStatus, error) {
	health := HealthStatus{Healthy: false, Hostname: ThisHostname, Token: ProcessToken.Hash}

	healthy, err := RegisterNode("", "", false)
	if err != nil {
		health.Error = err
		return &health, log.Errore(err)
	}
	health.Healthy = healthy
	health.ActiveNode, health.IsActiveNode, err = ElectedNode()
	if err != nil {
		health.Error = err
		return &health, log.Errore(err)
	}

	health.AvailableNodes, err = ReadAvailableNodes(true)

	return &health, nil
}

// ContinuousRegistration will continuously update the node_health
// table showing that the current process is still running.
func ContinuousRegistration(extraInfo string, command string) {
	if continuousRegistrationInitiated {
		// This is a simple mechanism to make sure this function is not being called multiple times in the lifespan of this process.
		// It is not concurrency-protected.
		// Original use case: multiple instances as in "-i instance1,instance2,instance3" flag
		return
	}
	continuousRegistrationInitiated = true

	tickOperation := func(firstTime bool) {
		if _, err := RegisterNode(extraInfo, command, firstTime); err != nil {
			log.Errorf("ContinuousRegistration: RegisterNode failed: %+v", err)
		}
	}
	// First one is synchronous
	tickOperation(true)
	go func() {
		registrationTick := time.Tick(time.Duration(registrationPollSeconds) * time.Second)
		for range registrationTick {
			// We already run inside a go-routine so
			// do not do this asynchronously.  If we
			// get stuck then we don't want to fill up
			// the backend pool with connections running
			// this maintenance operation.
			tickOperation(false)
		}
	}()
}

// ExpireAvailableNodes is an aggressive purging method to remove
// node entries who have skipped their keepalive for two times.
func ExpireAvailableNodes() {
	_, err := db.ExecOrchestrator(`
			delete
				from node_health
			where
				last_seen_active < now() - interval ? second
			`,
		registrationPollSeconds*2,
	)
	if err != nil {
		log.Errorf("ExpireAvailableNodes: failed to remove old entries: %+v", err)
	}
}

// ExpireNodesHistory cleans up the nodes history and is run by
// the orchestrator active node.
func ExpireNodesHistory() error {
	_, err := db.ExecOrchestrator(`
			delete
				from node_health_history
			where
				first_seen_active < now() - interval ? hour
			`,
		config.Config.UnseenInstanceForgetHours,
	)
	return log.Errore(err)
}

func ReadAvailableNodes(onlyHttpNodes bool) (nodes [](*NodeHealth), err error) {
	extraInfo := ""
	if onlyHttpNodes {
		extraInfo = string(OrchestratorExecutionHttpMode)
	}
	query := `
		select
			hostname, token, app_version, first_seen_active, last_seen_active
		from
			node_health
		where
			last_seen_active > now() - interval ? second
			and ? in (extra_info, '')
		order by
			hostname
		`

	err = db.QueryOrchestrator(query, sqlutils.Args(registrationPollSeconds*2, extraInfo), func(m sqlutils.RowMap) error {
		nodeHealth := &NodeHealth{
			Hostname:        m.GetString("hostname"),
			Token:           m.GetString("token"),
			AppVersion:      m.GetString("app_version"),
			FirstSeenActive: m.GetString("first_seen_active"),
			LastSeenActive:  m.GetString("last_seen_active"),
		}
		nodes = append(nodes, nodeHealth)
		return nil
	})
	return nodes, log.Errore(err)
}

func TokenBelongsToHealthyHttpService(token string) (result bool, err error) {
	extraInfo := string(OrchestratorExecutionHttpMode)

	query := `
		select
			token
		from
			node_health
		where
			and token = ?
			and extra_info = ?
		`

	err = db.QueryOrchestrator(query, sqlutils.Args(token, extraInfo), func(m sqlutils.RowMap) error {
		// Row exists? We're happy
		result = true
		return nil
	})
	return result, log.Errore(err)
}

// Just check to make sure we can connect to the database
func SimpleHealthTest() (*HealthStatus, error) {
	health := HealthStatus{Healthy: false, Hostname: ThisHostname, Token: ProcessToken.Hash}

	db, err := db.OpenOrchestrator()
	if err != nil {
		health.Error = err
		return &health, log.Errore(err)
	}

	if err = db.Ping(); err != nil {
		health.Error = err
		return &health, log.Errore(err)
	} else {
		health.Healthy = true
		return &health, nil
	}
}
