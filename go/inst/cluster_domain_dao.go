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

// ReadClusterDomainName reads the domain name associated with a cluster, if any
func ReadClusterDomainName(clusterName string) (string, error) {
	domainName := ""
	query := fmt.Sprintf(`
		select 
			domain_name
		from 
			cluster_domain_name
		where
			cluster_name = '%s'
		`, clusterName)
	db, err := db.OpenOrchestrator()
	if err != nil {
		goto Cleanup
	}

	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		domainName = m.GetString("domain_name")
		return nil
	})
Cleanup:

	if err != nil {
		return "", err
	}
	if domainName == "" {
		err = fmt.Errorf("No domain name found for cluster %s", clusterName)
	}
	return domainName, err

}

// WriteClusterDomainName will write (and override) the domain name of a cluster
func WriteClusterDomainName(clusterName string, domainName string) error {
	writeFunc := func() error {
		db, err := db.OpenOrchestrator()
		if err != nil {
			return log.Errore(err)
		}

		_, err = sqlutils.Exec(db, `
			insert into  
					cluster_domain_name (cluster_name, domain_name, last_registered)
				values
					(?, ?, NOW())
				on duplicate key update
					domain_name=values(domain_name),
					last_registered=values(last_registered)
			`,
			clusterName,
			domainName)
		if err != nil {
			return log.Errore(err)
		}

		return nil
	}
	return ExecDBWriteFunc(writeFunc)
}

// ExpireClusterDomainName expires cluster_domain_name entries that haven't been updated recently.
func ExpireClusterDomainName() error {
	writeFunc := func() error {
		_, err := db.ExecOrchestrator(`
        	delete from cluster_domain_name 
				where last_registered < NOW() - INTERVAL ? MINUTE
				`, config.Config.ExpiryHostnameResolvesMinutes,
		)
		return log.Errore(err)
	}
	return ExecDBWriteFunc(writeFunc)
}
