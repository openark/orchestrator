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

package logic

import (
	"fmt"
	"github.com/outbrain/golib/log"
	"github.com/outbrain/golib/sqlutils"
	"github.com/outbrain/orchestrator/go/db"
)

type HealthStatus struct {
	Healthy        bool
	Hostname       string
	Token          string
	IsActiveNode   bool
	ActiveNode     string
	Error          error
	AvailableNodes []string
}

// HealthTest attempts to write to the backend database and get a result
func HealthTest() (*HealthStatus, error) {
	health := HealthStatus{Healthy: false, Hostname: ThisHostname, Token: ProcessToken.Hash}

	db, err := db.OpenOrchestrator()
	if err != nil {
		health.Error = err
		return &health, log.Errore(err)
	}

	sqlResult, err := sqlutils.Exec(db, `
			insert into node_health 
				(hostname, token, last_seen_active)
			values
				(?, ?, NOW())
			on duplicate key update
				token=values(token),
				last_seen_active=values(last_seen_active)
			`,
		ThisHostname, ProcessToken.Hash,
	)
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

	health.AvailableNodes, err = readAvailableNodes()

	return &health, nil
}

func readAvailableNodes() ([]string, error) {
	res := []string{}
	query := fmt.Sprintf(`
		select 
			concat(hostname, ';', token) as node
		from 
			node_health
		where
			last_seen_active > now() - interval 5 minute
		order by
			hostname
		`)
	db, err := db.OpenOrchestrator()
	if err != nil {
		goto Cleanup
	}

	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		res = append(res, m.GetString("node"))
		return nil
	})
Cleanup:

	if err != nil {
		log.Errore(err)
	}
	return res, err
}
