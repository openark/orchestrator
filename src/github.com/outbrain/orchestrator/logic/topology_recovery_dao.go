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
	"github.com/outbrain/orchestrator/config"
	"github.com/outbrain/orchestrator/db"
	"github.com/outbrain/orchestrator/inst"
)

// AttemptRecoveryRegistration tries to add a recovery entry; if this fails that means recovery is already in place.
func AttemptRecoveryRegistration(failedKey *inst.InstanceKey) (bool, error) {

	db, err := db.OpenOrchestrator()
	if err != nil {
		return false, log.Errore(err)
	}

	sqlResult, err := sqlutils.Exec(db, `
			insert ignore 
				into topology_recovery (
					hostname, 
					port, 
					in_active_period, 
					start_active_period, 
					end_active_period_unixtime, 
					processing_node_hostname, 
					processcing_node_token
				) values (
					?,
					?,
					1,
					NOW(),
					0,
					?,
					?
				)
			`, failedKey.Hostname, failedKey.Port, ThisHostname, ProcessToken.Hash,
	)
	if err != nil {
		return false, log.Errore(err)
	}
	rows, err := sqlResult.RowsAffected()
	return (err == nil && rows > 0), err
}

// ClearActiveRecoveries clears the "in_active_period" flag for old-enough recoveries, thereby allowing for
// further recoveries on cleared instances.
func ClearActiveRecoveries() error {

	db, err := db.OpenOrchestrator()
	if err != nil {
		return log.Errore(err)
	}

	_, err = sqlutils.Exec(db, `
			update topology_recovery set 
				in_active_period = 0,
				end_active_period_unixtime = UNIX_TIMESTAMP()
			where
				in_active_period = 1
				AND start_active_period < NOW() - INTERVAL ? MINUTE
			`,
		config.Config.RecoveryPeriodBlockMinutes,
	)
	if err != nil {
		return log.Errore(err)
	}
	return nil
}

// ResolveRecovery is called on completion of a recovery process and updates the recovery status.
// It does not clear the "active period" as this still takes place in order to avoid flapping.
func ResolveRecovery(failedKey *inst.InstanceKey, successorKey *inst.InstanceKey) error {

	db, err := db.OpenOrchestrator()
	if err != nil {
		return log.Errore(err)
	}

	if successorKey == nil {
		successorKey = &inst.InstanceKey{}
	}

	_, err = sqlutils.Exec(db, `
			update topology_recovery set 
				successor_hostname = ?,
				successor_port = ?,
				end_recovery = NOW()
			where
				hostname = ?
				AND port = ?
				AND in_active_period = 1
				AND processing_node_hostname = ?
				AND processcing_node_token = ?
			`, successorKey.Hostname, successorKey.Port,
		failedKey.Hostname, failedKey.Port, ThisHostname, ProcessToken.Hash,
	)
	if err != nil {
		return log.Errore(err)
	}
	return nil
}
