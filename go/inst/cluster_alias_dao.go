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
	query := `
		select 
			cluster_name,
			alias
		from 
			cluster_alias
		`
	err := db.QueryOrchestratorRowsMap(query, func(m sqlutils.RowMap) error {
		updatedMap[m.GetString("cluster_name")] = m.GetString("alias")
		return nil
	})
	if err != nil {
		log.Errore(err)
	}
	clusterAliasMapMutex.Lock()
	defer clusterAliasMapMutex.Unlock()
	clusterAliasMap = updatedMap
	return err

}

// ReadClusterByAlias
func ReadClusterByAlias(alias string) (string, error) {
	clusterName := ""
	query := `
		select 
			cluster_name
		from 
			cluster_alias
		where
			alias = ?
		`
	err := db.QueryOrchestrator(query, sqlutils.Args(alias), func(m sqlutils.RowMap) error {
		clusterName = m.GetString("cluster_name")
		return nil
	})
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
		_, err := db.ExecOrchestrator(`
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

//
func UpdateClusterAliases() error {
	writeFunc := func() error {
		_, err := db.ExecOrchestrator(`
			replace into  
					cluster_alias (alias, cluster_name, last_registered)
				select 
				    suggested_cluster_alias, 
				    substring_index(group_concat(cluster_name order by cluster_name), ',', 1) as cluster_name,
				    NOW()
				  from 
				    database_instance 
				    left join database_instance_downtime using (hostname, port)
				  where 
				    suggested_cluster_alias!='' 
				    and not (
				      (hostname, port) in (select hostname, port from topology_recovery where start_active_period >= now() - interval 11111 day) 
				      and (
				        database_instance_downtime.downtime_active IS NULL
				        or database_instance_downtime.end_timestamp < NOW()
					  ) is false
				    )
				  group by 
				    suggested_cluster_alias
			`)
		if err == nil {
			err = ReadClusterAliases()
		}
		return log.Errore(err)
	}
	return ExecDBWriteFunc(writeFunc)
}
