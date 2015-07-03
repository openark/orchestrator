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

// writePoolInstances will write (and override) a single cluster name mapping
func writePoolInstances(pool string, instanceKeys [](*InstanceKey)) error {
	writeFunc := func() error {
		db, err := db.OpenOrchestrator()
		if err != nil {
			return log.Errore(err)
		}

		tx, err := db.Begin()
		stmt, err := tx.Prepare(`delete from database_instance_pool where pool = ?`)
		_, err = stmt.Exec(pool)
		if err != nil {
			tx.Rollback()
			return log.Errore(err)
		}
		stmt, err = tx.Prepare(`insert into database_instance_pool values (?, ?, ?)`)
		for _, instanceKey := range instanceKeys {
			_, err := stmt.Exec(instanceKey.Hostname, instanceKey.Port, pool)
			if err != nil {
				tx.Rollback()
				return log.Errore(err)
			}
		}
		if err != nil {
			tx.Rollback()
			return log.Errore(err)
		}
		tx.Commit()

		return nil
	}
	return ExecDBWriteFunc(writeFunc)
}

func ReadClusterPoolInstances(clusterName string) (*PoolInstancesMap, error) {
	var poolInstancesMap = make(PoolInstancesMap)

	query := fmt.Sprintf(`
		select 
			database_instance_pool.*
		from 
			database_instance
			join database_instance_pool using (hostname, port)
		where
			database_instance.cluster_name = '%s'
		`, clusterName)
	db, err := db.OpenOrchestrator()
	if err != nil {
		goto Cleanup
	}

	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		pool := m.GetString("pool")
		hostname := m.GetString("hostname")
		port := m.GetInt("port")
		if _, ok := poolInstancesMap[pool]; !ok {
			poolInstancesMap[pool] = [](*InstanceKey){}
		}
		poolInstancesMap[pool] = append(poolInstancesMap[pool], &InstanceKey{Hostname: hostname, Port: port})
		return nil
	})
Cleanup:

	if err != nil {
		return nil, err
	}

	return &poolInstancesMap, nil

}

func ReadAllClusterPoolInstances() ([](*ClusterPoolInstance), error) {
	var result [](*ClusterPoolInstance) = [](*ClusterPoolInstance){}
	query := fmt.Sprintf(`
		select 
			cluster_name,
			ifnull(alias, cluster_name) as alias,
			database_instance_pool.*
		from 
			database_instance
			join database_instance_pool using (hostname, port)
			left join cluster_alias using (cluster_name)
		`)
	db, err := db.OpenOrchestrator()
	if err != nil {
		goto Cleanup
	}

	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		clusterPoolInstance := ClusterPoolInstance{
			ClusterName:  m.GetString("cluster_name"),
			ClusterAlias: m.GetString("alias"),
			Pool:         m.GetString("pool"),
			Hostname:     m.GetString("hostname"),
			Port:         m.GetInt("port"),
		}
		result = append(result, &clusterPoolInstance)
		return nil
	})
Cleanup:

	if err != nil {
		return nil, err
	}

	return result, nil

}
