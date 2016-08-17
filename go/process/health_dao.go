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
	"database/sql"
	"fmt"
	"time"

	"github.com/outbrain/golib/log"
	"github.com/outbrain/golib/sqlutils"
	"github.com/outbrain/orchestrator/go/config"
	"github.com/outbrain/orchestrator/go/db"
)

const registrationPollSeconds = 10

type HealthStatus struct {
	Healthy        bool
	Hostname       string
	Token          string
	IsActiveNode   bool
	ActiveNode     string
	Error          error
	AvailableNodes []string
}

type OrchestratorExecutionMode string

const (
	OrchestratorExecutionCliMode  OrchestratorExecutionMode = "CLIMode"
	OrchestratorExecutionHttpMode                           = "HttpMode"
)

var continuousRegistrationInitiated bool = false

// RegisterNode writes down this node in the node_health table
func RegisterNode(extraInfo string, command string, firstTime bool) (sql.Result, error) {
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
	return db.ExecOrchestrator(`
			insert into node_health
				(hostname, token, last_seen_active, extra_info, command, app_version)
			values
				(?, ?, NOW(), ?, ?, ?)
			on duplicate key update
				token=values(token),
				last_seen_active=values(last_seen_active),
				extra_info=if(values(extra_info) != '', values(extra_info), extra_info),
				app_version=values(app_version)
			`,
		ThisHostname, ProcessToken.Hash, extraInfo, command,
		config.RuntimeCLIFlags.ConfiguredVersion,
	)
}

// HealthTest attempts to write to the backend database and get a result
func HealthTest() (*HealthStatus, error) {
	health := HealthStatus{Healthy: false, Hostname: ThisHostname, Token: ProcessToken.Hash}

	sqlResult, err := RegisterNode("", "", false)
	if err != nil {
		health.Error = err
		return &health, log.Errore(err)
	}
	rows, err := sqlResult.RowsAffected()
	if err != nil {
		health.Error = err
		return &health, log.Errore(err)
	}
	health.Healthy = (rows > 0)
	activeHostname, activeToken, isActive, err := ElectedNode()
	if err != nil {
		health.Error = err
		return &health, log.Errore(err)
	}
	health.ActiveNode = fmt.Sprintf("%s;%s", activeHostname, activeToken)
	health.IsActiveNode = isActive

	health.AvailableNodes, err = ReadAvailableNodes(true)

	return &health, nil
}

func ContinuousRegistration(extraInfo string, command string) {
	if continuousRegistrationInitiated {
		// This is a simple mechanism to make sure this function is not being called multiple times in the lifespan of this process.
		// It is not concurrency-protected.
		// Original use case: multiple instances as in "-i instance1,instance2,instance3" flag
		return
	}
	continuousRegistrationInitiated = true

	tickOperation := func(firstTime bool) {
		RegisterNode(extraInfo, command, firstTime)
		expireAvailableNodes()
	}
	// First one is synchronous
	tickOperation(true)
	go func() {
		registrationTick := time.Tick(time.Duration(registrationPollSeconds) * time.Second)
		for range registrationTick {
			go tickOperation(false)
		}
	}()
}

// expireAvailableNodes is an aggressive purging method to remove node entries who have skipped
// their keepalive for two times
func expireAvailableNodes() error {
	if !config.Config.NodeHealthExpiry {
		return nil
	}
	_, err := db.ExecOrchestrator(`
			delete
				from node_health
			where
				last_seen_active < now() - interval ? second
			`,
		registrationPollSeconds*2,
	)
	return log.Errore(err)
}

// ExpireNodesHistory cleans up the nodes history
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

func ReadAvailableNodes(onlyHttpNodes bool) ([]string, error) {
	res := []string{}
	extraInfo := ""
	if onlyHttpNodes {
		extraInfo = string(OrchestratorExecutionHttpMode)
	}
	query := `
		select
			concat(hostname, ';', token, ';', app_version) as node
		from
			node_health
		where
			last_seen_active > now() - interval ? second
			and ? in (extra_info, '')
		order by
			hostname
		`

	err := db.QueryOrchestrator(query, sqlutils.Args(registrationPollSeconds*2, extraInfo), func(m sqlutils.RowMap) error {
		res = append(res, m.GetString("node"))
		return nil
	})
	if err != nil {
		log.Errore(err)
	}
	return res, err
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
