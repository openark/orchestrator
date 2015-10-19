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
	"github.com/outbrain/golib/log"
	"github.com/outbrain/golib/sqlutils"
	"github.com/outbrain/orchestrator/go/config"
	"github.com/outbrain/orchestrator/go/db"
)

// AttemptElection tries to grab leadership (become active node)
func AttemptElection() (bool, error) {
	sqlResult, err := db.ExecOrchestrator(`
			insert ignore into active_node (
					anchor, hostname, token, last_seen_active
				) values (
					1, ?, ?, now()
				) on duplicate key update
					hostname = if(last_seen_active < now() - interval ? second, values(hostname), hostname),
					token    = if(last_seen_active < now() - interval ? second, values(token),    token),
					last_seen_active = if(hostname = values(hostname) and token = values(token), values(last_seen_active), last_seen_active)					
			`,
		ThisHostname, ProcessToken.Hash, config.Config.ActiveNodeExpireSeconds, config.Config.ActiveNodeExpireSeconds,
	)
	if err != nil {
		return false, log.Errore(err)
	}
	rows, err := sqlResult.RowsAffected()
	return (rows > 0), log.Errore(err)
}

// GrabElection forcibly grabs leadership. Use with care!!
func GrabElection() error {
	_, err := db.ExecOrchestrator(`
			replace into active_node (
					anchor, hostname, token, last_seen_active
				) values (
					1, ?, ?, NOW()
				)
			`,
		ThisHostname, ProcessToken.Hash,
	)
	return log.Errore(err)
}

// Reelect clears the way for re-elections. Active node is immediately demoted.
func Reelect() error {
	_, err := db.ExecOrchestrator(`delete from active_node where anchor = 1`)
	return log.Errore(err)
}

// ElectedNode returns the details of the elected node, as well as answering the question "is this process the elected one"?
func ElectedNode() (hostname string, token string, isElected bool, err error) {
	query := `
		select 
			hostname,
			token
		from 
			active_node
		where
			anchor = 1
		`
	err = db.QueryOrchestratorRowsMap(query, func(m sqlutils.RowMap) error {
		hostname = m.GetString("hostname")
		token = m.GetString("token")
		return nil
	})

	if err != nil {
		log.Errore(err)
	}
	isElected = (hostname == ThisHostname && token == ProcessToken.Hash)
	return hostname, token, isElected, err
}
