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
	"github.com/outbrain/orchestrator/go/config"
	"github.com/outbrain/orchestrator/go/db"
)

func CreateElectionAnchor(force bool) error {
	db, err := db.OpenOrchestrator()
	if err != nil {
		return log.Errore(err)
	}

	statement := `insert ignore`
	if force {
		statement = `replace`
	}

	_, err = sqlutils.Exec(db, statement+`
			 into active_node
				(anchor, last_seen_active, hostname)
			values
				(1, now() - interval (? + 1) second, '')
		`, config.Config.ActiveNodeExpireSeconds)
	if err != nil {
		return log.Errore(err)
	}
	return nil

}

// AttemptElection tries to grab leadership (become active node)
func AttemptElection() (bool, error) {

	db, err := db.OpenOrchestrator()
	if err != nil {
		return false, log.Errore(err)
	}

	sqlResult, err := sqlutils.Exec(db, `
			update active_node set 
				hostname = ?,
				token = ?,
				last_seen_active = now()
			where
				anchor = 1
				and (
					last_seen_active < now() - interval ? second
					or hostname = ''
					or (hostname = ? and token = ?)
				)					
			`,
		ThisHostname, ProcessToken.Hash, config.Config.ActiveNodeExpireSeconds, ThisHostname, ProcessToken.Hash,
	)
	if err != nil {
		return false, log.Errore(err)
	}
	rows, err := sqlResult.RowsAffected()
	return (rows > 0), err
}

// GrabElection forcibly grabs leadership. Use with care!!
func GrabElection() (bool, error) {

	db, err := db.OpenOrchestrator()
	if err != nil {
		return false, log.Errore(err)
	}

	sqlResult, err := sqlutils.Exec(db, `
			update active_node set 
				hostname = ?,
				token = ?,
				last_seen_active = now()
			where
				anchor = 1
			`,
		ThisHostname, ProcessToken.Hash,
	)
	if err != nil {
		return false, log.Errore(err)
	}
	rows, err := sqlResult.RowsAffected()
	return (rows > 0), err
}

// Reelect clears the way for re-elections. Active node is immediately demoted.
func Reelect() error {
	return CreateElectionAnchor(true)
}

// IsElected checks whether this node is the elected active node
func IsElected() (bool, error) {
	isElected := false
	query := fmt.Sprintf(`
		select 
			count(*) as is_elected
		from 
			active_node
		where
			anchor = 1
			and hostname = '%s'
			and token = '%s'
		`,
		ThisHostname, ProcessToken.Hash)
	db, err := db.OpenOrchestrator()
	if err != nil {
		goto Cleanup
	}

	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		isElected = m.GetBool("is_elected")
		return nil
	})
Cleanup:

	if err != nil {
		log.Errore(err)
	}
	return isElected, err
}

// ElectedNode returns the hostname of the elected node
func ElectedNode() (string, string, bool, error) {
	hostname := ""
	token := ""
	isElected := false
	query := fmt.Sprintf(`
		select 
			ifnull(max(hostname), '') as hostname,
			ifnull(max(token), '') as token,
			(ifnull(max(hostname), '') = '%s') and (ifnull(max(token), '') = '%s') as is_elected
		from 
			active_node
		where
			anchor = 1
		`,
		ThisHostname, ProcessToken.Hash)
	db, err := db.OpenOrchestrator()
	if err != nil {
		goto Cleanup
	}

	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		hostname = m.GetString("hostname")
		token = m.GetString("token")
		isElected = m.GetBool("is_elected")
		return nil
	})
Cleanup:

	if err != nil {
		log.Errore(err)
	}
	return hostname, token, isElected, err
}
