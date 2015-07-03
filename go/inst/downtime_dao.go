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

package inst

import (
	"fmt"
	"github.com/outbrain/golib/log"
	"github.com/outbrain/golib/sqlutils"
	"github.com/outbrain/orchestrator/go/config"
	"github.com/outbrain/orchestrator/go/db"
)

// BeginDowntime will make mark an instance as downtimed (or override existing downtime period)
func BeginDowntime(instanceKey *InstanceKey, owner string, reason string, durationSeconds uint) error {
	db, err := db.OpenOrchestrator()
	if err != nil {
		return log.Errore(err)
	}

	if durationSeconds == 0 {
		durationSeconds = config.Config.MaintenanceExpireMinutes * 60
	}
	_, err = sqlutils.Exec(db, `
			insert 
				into database_instance_downtime (
					hostname, port, downtime_active, begin_timestamp, end_timestamp, owner, reason
				) VALUES (
					?, ?, 1, NOW(), NOW() + INTERVAL ? SECOND, ?, ?
				)
				on duplicate key update
					downtime_active=values(downtime_active),
					begin_timestamp=values(begin_timestamp),
					end_timestamp=values(end_timestamp),
					owner=values(owner),
					reason=values(reason)
			`,
		instanceKey.Hostname,
		instanceKey.Port,
		durationSeconds,
		owner,
		reason,
	)
	if err != nil {
		return log.Errore(err)
	}

	AuditOperation("begin-downtime", instanceKey, fmt.Sprintf("owner: %s, reason: %s", owner, reason))

	return nil
}

// EndDowntime will remove downtime flag from an instance
func EndDowntime(instanceKey *InstanceKey) error {
	db, err := db.OpenOrchestrator()
	if err != nil {
		return log.Errore(err)
	}

	res, err := sqlutils.Exec(db, `
			update
				database_instance_downtime
			set  
				downtime_active = NULL,
				end_timestamp = NOW()
			where
				hostname = ? 
				and port = ?
				and downtime_active = 1
			`,
		instanceKey.Hostname,
		instanceKey.Port,
	)
	if err != nil {
		return log.Errore(err)
	}

	if affected, _ := res.RowsAffected(); affected == 0 {
		err = fmt.Errorf("Instance is not in downtime mode: %+v", instanceKey)
	} else {
		// success
		AuditOperation("end-downtime", instanceKey, "")
	}
	return err
}

// ExpireDowntime will remove the maintenance flag on old downtimes
func ExpireDowntime() error {
	db, err := db.OpenOrchestrator()
	if err != nil {
		return log.Errore(err)
	}

	{
		res, err := sqlutils.Exec(db, `
			delete from
				database_instance_downtime
			where
				downtime_active is null
				and end_timestamp < NOW() - INTERVAL ? DAY 
			`,
			config.Config.MaintenancePurgeDays,
		)
		if err != nil {
			return log.Errore(err)
		}
		if rowsAffected, _ := res.RowsAffected(); rowsAffected > 0 {
			AuditOperation("expire-downtime", nil, fmt.Sprintf("Purged %d historical entries", rowsAffected))
		}
	}
	{
		res, err := sqlutils.Exec(db, `
			update
				database_instance_downtime
			set  
				downtime_active = NULL				
			where
				downtime_active = 1
				and end_timestamp < NOW() 
			`,
		)
		if err != nil {
			return log.Errore(err)
		}
		if rowsAffected, _ := res.RowsAffected(); rowsAffected > 0 {
			AuditOperation("expire-downtime", nil, fmt.Sprintf("Expired %d entries", rowsAffected))
		}
	}

	return err
}
