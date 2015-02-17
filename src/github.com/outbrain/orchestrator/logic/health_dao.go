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

package orchestrator

import (
	"github.com/outbrain/golib/log"
	"github.com/outbrain/golib/sqlutils"
	"github.com/outbrain/orchestrator/db"
)

type HealthStatus struct {
	Healthy      bool
	IsActiveNode bool
	ActiveNode   string
	Error        error
}

// HealthTest attempts to write to the backend database and get a result
func HealthTest() (*HealthStatus, error) {
	health := HealthStatus{Healthy: false}

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
	health.ActiveNode, err = ElectedNode()
	health.IsActiveNode, err = IsElected()
	if err != nil {
		health.Error = err
		return &health, log.Errore(err)
	}

	return &health, nil
}
