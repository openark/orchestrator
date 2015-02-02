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
	"fmt"
	"github.com/outbrain/golib/log"
	"github.com/outbrain/golib/sqlutils"
	"github.com/outbrain/orchestrator/config"
	"github.com/outbrain/orchestrator/db"
)

// WriteResolvedHostname stores a hostname and the resolved hostname to backend database
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

// ReadResolvedHostname returns the resolved hostname given a hostname, or empty if not exists
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
