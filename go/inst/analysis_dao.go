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
	"regexp"
)

// GetReplicationAnalysis will check for replication problems (dead master; unreachable master; etc)
func GetReplicationAnalysis(includeDowntimed bool) ([]ReplicationAnalysis, error) {
	result := []ReplicationAnalysis{}

	query := fmt.Sprintf(`
		    SELECT 
		        master_instance.hostname,
		        master_instance.port,
		        MIN(master_instance.master_host) AS master_host,
		        MIN(master_instance.master_port) AS master_port,
		        MIN(master_instance.cluster_name) AS cluster_name,
		        MIN(IFNULL(cluster_alias.alias, master_instance.cluster_name)) AS cluster_alias,
		        MIN(
		        		master_instance.last_checked <= master_instance.last_seen
		        		AND master_instance.last_attempted_check <= master_instance.last_seen + INTERVAL (2 * %d) SECOND
		        	) IS TRUE AS is_last_check_valid,
		        MIN(master_instance.master_host IN ('' , '_')
		            OR master_instance.master_port = 0) AS is_master,
		        MIN(master_instance.is_co_master) AS is_co_master,
		        MIN(CONCAT(master_instance.hostname,
		                ':',
		                master_instance.port) = master_instance.cluster_name) AS is_cluster_master,
		        COUNT(slave_instance.server_id) AS count_slaves,
		        IFNULL(SUM(slave_instance.last_checked <= slave_instance.last_seen),
		                0) AS count_valid_slaves,
		        IFNULL(SUM(slave_instance.last_checked <= slave_instance.last_seen
		                    AND slave_instance.slave_io_running != 0
		                    AND slave_instance.slave_sql_running != 0),
		                0) AS count_valid_replicating_slaves,
		        IFNULL(SUM(slave_instance.last_checked <= slave_instance.last_seen
		                    AND slave_instance.slave_io_running = 0
		                    AND slave_instance.last_io_error RLIKE 'error (connecting|reconnecting) to master'
		                    AND slave_instance.slave_sql_running = 1),
		                0) AS count_slaves_failing_to_connect_to_master,
		        MIN(master_instance.replication_depth) AS replication_depth,
		        MIN(master_instance.slave_hosts) AS slave_hosts,
		        MIN(
		            master_instance.slave_sql_running = 1
		            AND master_instance.slave_io_running = 0
		            AND master_instance.last_io_error RLIKE 'error (connecting|reconnecting) to master'
		          ) AS is_failing_to_connect_to_master,
		        MIN(
		    		database_instance_downtime.downtime_active IS NULL
		    		OR database_instance_downtime.end_timestamp < NOW()
		    	) IS FALSE AS is_downtimed,
		    	MIN(
		    		IFNULL(database_instance_downtime.end_timestamp, '')
		    	) AS downtime_end_timestamp,
		    	MIN(
		    		IFNULL(TIMESTAMPDIFF(SECOND, NOW(), database_instance_downtime.end_timestamp), 0)
		    	) AS downtime_remaining_seconds,
		    	MIN(
		    		master_instance.version like '%%maxscale%%'
		    	) AS is_maxscale
		    FROM
		        database_instance master_instance
		            LEFT JOIN
		        hostname_resolve ON (master_instance.hostname = hostname_resolve.hostname)
		            LEFT JOIN
		        database_instance slave_instance ON (COALESCE(hostname_resolve.resolved_hostname,
		                master_instance.hostname) = slave_instance.master_host
		            	AND master_instance.port = slave_instance.master_port)
		            LEFT JOIN
		        database_instance_maintenance ON (master_instance.hostname = database_instance_maintenance.hostname
		        		AND master_instance.port = database_instance_maintenance.port
		        		AND database_instance_maintenance.maintenance_active = 1)
		            LEFT JOIN
		        database_instance_downtime ON (master_instance.hostname = database_instance_downtime.hostname
		        		AND master_instance.port = database_instance_downtime.port
		        		AND database_instance_downtime.downtime_active = 1)
		        	LEFT JOIN 
		        cluster_alias ON (cluster_alias.cluster_name = master_instance.cluster_name)
		    WHERE
		    	database_instance_maintenance.database_instance_maintenance_id IS NULL
		    GROUP BY 
			    master_instance.hostname, 
			    master_instance.port
		    ORDER BY 
			    is_master DESC , 
			    is_cluster_master DESC, 
			    count_slaves DESC
	`, config.Config.InstancePollSeconds)

	db, err := db.OpenOrchestrator()
	if err != nil {
		goto Cleanup
	}

	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		a := ReplicationAnalysis{Analysis: NoProblem}

		a.IsMaster = m.GetBool("is_master")
		a.IsCoMaster = m.GetBool("is_co_master")
		a.AnalyzedInstanceKey = InstanceKey{Hostname: m.GetString("hostname"), Port: m.GetInt("port")}
		a.AnalyzedInstanceMasterKey = InstanceKey{Hostname: m.GetString("master_host"), Port: m.GetInt("master_port")}
		a.ClusterDetails.ClusterName = m.GetString("cluster_name")
		a.ClusterDetails.ClusterAlias = m.GetString("cluster_alias")
		a.LastCheckValid = m.GetBool("is_last_check_valid")
		a.CountSlaves = m.GetUint("count_slaves")
		a.CountValidSlaves = m.GetUint("count_valid_slaves")
		a.CountValidReplicatingSlaves = m.GetUint("count_valid_replicating_slaves")
		a.CountSlavesFailingToConnectToMaster = m.GetUint("count_slaves_failing_to_connect_to_master")
		a.ReplicationDepth = m.GetUint("replication_depth")
		a.IsFailingToConnectToMaster = m.GetBool("is_failing_to_connect_to_master")
		a.IsDowntimed = m.GetBool("is_downtimed")
		a.DowntimeEndTimestamp = m.GetString("downtime_end_timestamp")
		a.DowntimeRemainingSeconds = m.GetInt("downtime_remaining_seconds")
		a.IsMaxscale = m.GetBool("is_maxscale")
		a.ClusterDetails.ReadRecoveryInfo()

		instance := &Instance{}
		instance.ReadSlaveHostsFromJson(m.GetString("slave_hosts"))
		a.SlaveHosts = instance.SlaveHosts

		if a.IsMaster && !a.LastCheckValid && a.CountSlaves == 0 {
			a.Analysis = DeadMasterWithoutSlaves
			a.Description = "Master cannot be reached by orchestrator and has no slave"
			//
		} else if a.IsMaster && !a.LastCheckValid && a.CountValidSlaves == a.CountSlaves && a.CountValidReplicatingSlaves == 0 {
			a.Analysis = DeadMaster
			a.Description = "Master cannot be reached by orchestrator and none of its slaves is replicating"
			//
		} else if a.IsMaster && !a.LastCheckValid && a.CountSlaves > 0 && a.CountValidSlaves == 0 && a.CountValidReplicatingSlaves == 0 {
			a.Analysis = DeadMasterAndSlaves
			a.Description = "Master cannot be reached by orchestrator and none of its slaves is replicating"
			//
		} else if a.IsMaster && !a.LastCheckValid && a.CountValidSlaves < a.CountSlaves && a.CountValidSlaves > 0 && a.CountValidReplicatingSlaves == 0 {
			a.Analysis = DeadMasterAndSomeSlaves
			a.Description = "Master cannot be reached by orchestrator; some of its slaves are unreachable and none of its reachable slaves is replicating"
			//
		} else if a.IsMaster && !a.LastCheckValid && a.CountValidSlaves > 0 && a.CountValidReplicatingSlaves > 0 {
			a.Analysis = UnreachableMaster
			a.Description = "Master cannot be reached by orchestrator but it has replicating slaves; possibly a network/host issue"
			//
		} else if a.IsMaster && a.LastCheckValid && a.CountSlaves == 1 && a.CountValidSlaves == a.CountSlaves && a.CountValidReplicatingSlaves == 0 {
			a.Analysis = MasterSingleSlaveNotReplicating
			a.Description = "Master is reachable but its single slave is not replicating"
			//
		} else if a.IsMaster && a.LastCheckValid && a.CountSlaves == 1 && a.CountValidSlaves == 0 {
			a.Analysis = MasterSingleSlaveDead
			a.Description = "Master is reachable but its single slave is dead"
			//
		} else if a.IsMaster && a.LastCheckValid && a.CountSlaves > 1 && a.CountValidSlaves == a.CountSlaves && a.CountValidReplicatingSlaves == 0 {
			a.Analysis = AllMasterSlavesNotReplicating
			a.Description = "Master is reachable but none of its slaves is replicating"
			//
		} else if a.IsMaster && a.LastCheckValid && a.CountSlaves > 1 && a.CountValidSlaves < a.CountSlaves && a.CountValidSlaves > 0 && a.CountValidReplicatingSlaves == 0 {
			a.Analysis = AllMasterSlavesNotReplicatingOrDead
			a.Description = "Master is reachable but none of its slaves is replicating"
			//
		} else /* co-master */ if a.IsCoMaster && !a.LastCheckValid && a.CountSlaves > 0 && a.CountValidSlaves == a.CountSlaves && a.CountValidReplicatingSlaves == 0 {
			a.Analysis = DeadCoMaster
			a.Description = "Co-master cannot be reached by orchestrator and none of its slaves is replicating"
			//
		} else if a.IsCoMaster && !a.LastCheckValid && a.CountValidSlaves > 0 && a.CountValidReplicatingSlaves > 0 {
			a.Analysis = UnreachableCoMaster
			a.Description = "Co-master cannot be reached by orchestrator but it has replicating slaves; possibly a network/host issue"
			//
		} else if a.IsCoMaster && a.LastCheckValid && a.CountSlaves > 0 && a.CountValidReplicatingSlaves == 0 {
			a.Analysis = AllCoMasterSlavesNotReplicating
			a.Description = "Co-master is reachable but none of its slaves is replicating"
			//
		} else /* intermediate-master */ if !a.IsMaster && !a.LastCheckValid && a.CountSlaves == 1 && a.CountValidSlaves == a.CountSlaves && a.CountSlavesFailingToConnectToMaster == a.CountSlaves && a.CountValidReplicatingSlaves == 0 {
			a.Analysis = DeadIntermediateMasterWithSingleSlaveFailingToConnect
			a.Description = "Intermediate master cannot be reached by orchestrator and its (single) slave is failing to connect"
			//
		} else /* intermediate-master */ if !a.IsMaster && !a.LastCheckValid && a.CountSlaves == 1 && a.CountValidSlaves == a.CountSlaves && a.CountValidReplicatingSlaves == 0 {
			a.Analysis = DeadIntermediateMasterWithSingleSlave
			a.Description = "Intermediate master cannot be reached by orchestrator and its (single) slave is not replicating"
			//
		} else /* intermediate-master */ if !a.IsMaster && !a.LastCheckValid && a.CountSlaves > 1 && a.CountValidSlaves == a.CountSlaves && a.CountValidReplicatingSlaves == 0 {
			a.Analysis = DeadIntermediateMaster
			a.Description = "Intermediate master cannot be reached by orchestrator and none of its slaves is replicating"
			//
		} else if !a.IsMaster && !a.LastCheckValid && a.CountValidSlaves < a.CountSlaves && a.CountValidSlaves > 0 && a.CountValidReplicatingSlaves == 0 {
			a.Analysis = DeadIntermediateMasterAndSomeSlaves
			a.Description = "Intermediate master cannot be reached by orchestrator; some of its slaves are unreachable and none of its reachable slaves is replicating"
			//
		} else if !a.IsMaster && !a.LastCheckValid && a.CountValidSlaves > 0 && a.CountValidReplicatingSlaves > 0 {
			a.Analysis = UnreachableIntermediateMaster
			a.Description = "Intermediate master cannot be reached by orchestrator but it has replicating slaves; possibly a network/host issue"
			//
		} else if !a.IsMaster && a.LastCheckValid && a.CountSlaves > 1 && a.CountValidReplicatingSlaves == 0 &&
			a.CountSlavesFailingToConnectToMaster > 0 && a.CountSlavesFailingToConnectToMaster == a.CountValidSlaves {
			// All slaves are either failing to connect to master (and at least one of these have to exist)
			// or completely dead.
			// Must have at least two slaves to reach such conclusion -- do note that the intermediate master is still
			// reachable to orchestrator, so we base our conclusion on slaves only at this point.
			a.Analysis = AllIntermediateMasterSlavesFailingToConnectOrDead
			a.Description = "Intermediate master is reachable but all of its slaves are failing to connect"
			//
		} else if !a.IsMaster && a.LastCheckValid && a.CountSlaves > 0 && a.CountValidReplicatingSlaves == 0 {
			a.Analysis = AllIntermediateMasterSlavesNotReplicating
			a.Description = "Intermediate master is reachable but none of its slaves is replicating"
			//
		} else if a.IsMaxscale && a.IsFailingToConnectToMaster {
			a.Analysis = MaxscaleFailingToConnectToMaster
			a.Description = "Maxscale is unable to connect to its master"
			//
		} else if a.ReplicationDepth == 1 && a.IsFailingToConnectToMaster {
			a.Analysis = FirstTierSlaveFailingToConnectToMaster
			a.Description = "1st tier slave (directly replicating from topology master) is unable to connect to the master"
			//
		}
		//		 else if a.IsMaster && a.CountSlaves == 0 {
		//			a.Analysis = MasterWithoutSlaves
		//			a.Description = "Master has no slaves"
		//		}

		if a.Analysis != NoProblem {
			skipThisHost := false
			for _, filter := range config.Config.RecoveryIgnoreHostnameFilters {
				if matched, _ := regexp.MatchString(filter, a.AnalyzedInstanceKey.Hostname); matched {
					skipThisHost = true
				}
			}
			if a.IsDowntimed && !includeDowntimed {
				skipThisHost = true
			}
			if !skipThisHost {
				result = append(result, a)
			}
		}
		return nil
	})
Cleanup:

	if err != nil {
		log.Errore(err)
	}
	return result, err

}
