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
	"github.com/outbrain/golib/log"
	"github.com/outbrain/golib/sqlutils"
	"github.com/outbrain/orchestrator/go/db"
	"time"
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

// RegisterNode writes down this node in the node_health table
func RegisterNode(extraInfo string) (sql.Result, error) {
	return db.ExecOrchestrator(`
			insert into node_health 
				(hostname, token, last_seen_active, extra_info)
			values
				(?, ?, NOW(), ?)
			on duplicate key update
				token=values(token),
				last_seen_active=values(last_seen_active),
				extra_info=if(values(extra_info) != '', values(extra_info), extra_info)
			`,
		ThisHostname, ProcessToken.Hash, extraInfo,
	)
}

// HealthTest attempts to write to the backend database and get a result
func HealthTest() (*HealthStatus, error) {
	health := HealthStatus{Healthy: false, Hostname: ThisHostname, Token: ProcessToken.Hash}

	sqlResult, err := RegisterNode("")
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

	health.AvailableNodes, err = readAvailableNodes(true)

	return &health, nil
}

func ContinuousRegistration(extraInfo string) {
	registrationTick := time.Tick(time.Duration(registrationPollSeconds) * time.Second)
	tickOperation := func() {
		RegisterNode(extraInfo)
		expireAvailableNodes()
	}
	go func() {
		go tickOperation()
		for range registrationTick {
			go tickOperation()
		}
	}()
}

// expireAvailableNodes is an aggressive puring method to remove node entries who have skipped
// their keepalive for two times
func expireAvailableNodes() error {
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

func readAvailableNodes(onlyHttpNodes bool) ([]string, error) {
	res := []string{}
	extraConditions := ``
	if onlyHttpNodes {
		extraConditions = fmt.Sprintf(`%s and extra_info = '%s'`, extraConditions, string(OrchestratorExecutionHttpMode))
	}
	query := fmt.Sprintf(`
		select 
			concat(hostname, ';', token) as node
		from 
			node_health
		where
			last_seen_active > now() - interval %d second
			%s
		order by
			hostname
		`, registrationPollSeconds*2, extraConditions)

	err := db.QueryOrchestratorRowsMap(query, func(m sqlutils.RowMap) error {
		res = append(res, m.GetString("node"))
		return nil
	})
	if err != nil {
		log.Errore(err)
	}
	return res, err
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
