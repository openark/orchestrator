/*
   Copyright 2014 Outbrain Inc.

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
	"github.com/github/orchestrator/go/db"
	"github.com/openark/golib/log"
	"github.com/openark/golib/sqlutils"
)

// WriteLongRunningProcesses rewrites current state of long running processes for given instance
func WriteLongRunningProcesses(instanceKey *InstanceKey, processes []Process) error {
	writeFunc := func() error {
		_, err := db.ExecOrchestrator(`
			delete from
					database_instance_long_running_queries
				where
					hostname = ?
					and port = ?
			`,
			instanceKey.Hostname,
			instanceKey.Port)
		if err != nil {
			return log.Errore(err)
		}

		for _, process := range processes {
			_, merr := db.ExecOrchestrator(`
      	insert ignore into database_instance_long_running_queries (
      		hostname,
      		port,
      		process_id,
      		process_started_at,
					process_user,
					process_host,
					process_db,
					process_command,
					process_time_seconds,
					process_state,
					process_info
				) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
				instanceKey.Hostname,
				instanceKey.Port,
				process.Id,
				process.StartedAt,
				process.User,
				process.Host,
				process.Db,
				process.Command,
				process.Time,
				process.State,
				process.Info,
			)
			if merr != nil {
				err = merr
			}
		}
		if err != nil {
			return log.Errore(err)
		}

		return nil
	}
	return ExecDBWriteFunc(writeFunc)
}

// ReadLongRunningProcesses returns the list of current known long running processes of all instances
func ReadLongRunningProcesses(filter string) ([]Process, error) {
	longRunningProcesses := []Process{}

	if filter != "" {
		filter = "%" + filter + "%"
	} else {
		filter = "%"
	}
	query := `
		select
			hostname,
			port,
			process_id,
			process_started_at,
			process_user,
			process_host,
			process_db,
			process_command,
			process_time_seconds,
			process_state,
			process_info
		from
			database_instance_long_running_queries
		where
			hostname like ?
			or process_user like ?
			or process_host like ?
			or process_db like ?
			or process_command like ?
			or process_state like ?
			or process_info like ?
		order by
			process_time_seconds desc
		`
	args := sqlutils.Args(filter, filter, filter, filter, filter, filter, filter)
	err := db.QueryOrchestrator(query, args, func(m sqlutils.RowMap) error {
		process := Process{}
		process.InstanceHostname = m.GetString("hostname")
		process.InstancePort = m.GetInt("port")
		process.Id = m.GetInt64("process_id")
		process.User = m.GetString("process_user")
		process.Host = m.GetString("process_host")
		process.Db = m.GetString("process_db")
		process.Command = m.GetString("process_command")
		process.Time = m.GetInt64("process_time_seconds")
		process.State = m.GetString("process_state")
		process.Info = m.GetString("process_info")
		process.StartedAt = m.GetString("process_started_at")

		longRunningProcesses = append(longRunningProcesses, process)
		return nil
	})

	if err != nil {
		log.Errore(err)
	}
	return longRunningProcesses, err

}
