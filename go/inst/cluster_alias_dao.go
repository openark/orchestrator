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
	"github.com/outbrain/orchestrator/go/db"
)

// ReadClusterAliases reads the entrie cluster name aliases mapping
func ReadClusterAliases() error {
	updatedMap := make(map[string]string)
	query := fmt.Sprintf(`
		select 
			cluster_name,
			alias
		from 
			cluster_alias
		`)
	db, err := db.OpenOrchestrator()
	if err != nil {
		goto Cleanup
	}

	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		updatedMap[m.GetString("cluster_name")] = m.GetString("alias")
		return err
	})
Cleanup:

	if err != nil {
		log.Errore(err)
	}
	clusterAliasMapMutex.Lock()
	clusterAliasMap = updatedMap
	clusterAliasMapMutex.Unlock()
	return err

}

// ReadClusterAliases reads the entrie cluster name aliases mapping
func ReadClusterByAlias(alias string) (string, error) {
	clusterName := ""
	query := fmt.Sprintf(`
		select 
			cluster_name
		from 
			cluster_alias
		where
			alias = '%s'
		`, alias)
	db, err := db.OpenOrchestrator()
	if err != nil {
		goto Cleanup
	}

	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		clusterName = m.GetString("cluster_name")
		return nil
	})
Cleanup:

	if err != nil {
		return "", err
	}
	if clusterName == "" {
		err = fmt.Errorf("No cluster found for alias %s", alias)
	}
	return clusterName, err

}

// WriteClusterAlias will write (and override) a single cluster name mapping
func WriteClusterAlias(clusterName string, alias string) error {
	writeFunc := func() error {
		db, err := db.OpenOrchestrator()
		if err != nil {
			return log.Errore(err)
		}

		_, err = sqlutils.Exec(db, `
			replace into  
					cluster_alias (cluster_name, alias)
				values
					(?, ?)
			`,
			clusterName,
			alias)
		if err != nil {
			return log.Errore(err)
		}

		return nil
	}
	return ExecDBWriteFunc(writeFunc)
}
