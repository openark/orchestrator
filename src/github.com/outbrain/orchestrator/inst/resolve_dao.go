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
	"fmt"
	"github.com/outbrain/golib/log"
	"github.com/outbrain/golib/sqlutils"
	"github.com/outbrain/orchestrator/config"
	"github.com/outbrain/orchestrator/db"
)

// WriteResolvedHostname stores a hostname and the resolved hostname to backend database
func WriteResolvedHostname(hostname string, resolvedHostname string) error {
	writeFunc := func() error {
		db, err := db.OpenOrchestrator()
		if err != nil {
			return log.Errore(err)
		}

		_, err = sqlutils.Exec(db, `
			insert into  
					hostname_resolve (hostname, resolved_hostname, resolved_timestamp)
				values
					(?, ?, NOW())
				on duplicate key update
					resolved_hostname = VALUES(resolved_hostname), 
					resolved_timestamp = VALUES(resolved_timestamp)
			`,
			hostname,
			resolvedHostname)
		if err != nil {
			return log.Errore(err)
		}
		log.Debugf("WriteResolvedHostname: resolved %s to %s", hostname, resolvedHostname)

		return nil
	}
	return execDBWriteFunc(writeFunc)
}

// ReadResolvedHostname returns the resolved hostname given a hostname, or empty if not exists
func ReadResolvedHostname(hostname string) (string, error) {
	var resolvedHostname string = ""

	query := fmt.Sprintf(`
		select 
			resolved_hostname
		from 
			hostname_resolve
		where
			hostname = '%s'
		`, hostname)
	db, err := db.OpenOrchestrator()
	if err != nil {
		goto Cleanup
	}

	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		resolvedHostname = m.GetString("resolved_hostname")
		return nil
	})
Cleanup:

	if err != nil {
		log.Errore(err)
	}
	return resolvedHostname, err
}

func readAllHostnameResolves() ([]HostnameResolve, error) {
	res := []HostnameResolve{}
	query := fmt.Sprintf(`
		select 
			hostname, 
			resolved_hostname  
		from 
			hostname_resolve
		`)
	db, err := db.OpenOrchestrator()
	if err != nil {
		goto Cleanup
	}

	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		hostnameResolve := HostnameResolve{hostname: m.GetString("hostname"), resolvedHostname: m.GetString("resolved_hostname")}

		res = append(res, hostnameResolve)
		return nil
	})
Cleanup:

	if err != nil {
		log.Errore(err)
	}
	return res, err
}

// ForgetLongUnseenInstances will remove entries of all instances that have long since been last seen.
func ForgetExpiredHostnameResolves() error {
	db, err := db.OpenOrchestrator()
	if err != nil {
		return log.Errore(err)
	}

	_, err = sqlutils.Exec(db, `
			delete 
				from hostname_resolve 
			where 
				resolved_timestamp < NOW() - interval (? * 2) minute`,
		config.Config.ExpiryHostnameResolvesMinutes,
	)
	return err
}

// deleteHostnameResolves compeltely erases the database cache
func deleteHostnameResolves() error {
	db, err := db.OpenOrchestrator()
	if err != nil {
		return log.Errore(err)
	}

	_, err = sqlutils.Exec(db, `
			delete 
				from hostname_resolve`,
	)
	return err
}
