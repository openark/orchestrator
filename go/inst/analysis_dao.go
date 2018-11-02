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
	"regexp"
	"time"

	"github.com/github/orchestrator/go/config"
	"github.com/github/orchestrator/go/db"
	"github.com/github/orchestrator/go/process"
	"github.com/github/orchestrator/go/raft"

	"github.com/openark/golib/log"
	"github.com/openark/golib/sqlutils"
	"github.com/patrickmn/go-cache"
	"github.com/rcrowley/go-metrics"
)

var analysisChangeWriteAttemptCounter = metrics.NewCounter()
var analysisChangeWriteCounter = metrics.NewCounter()

var recentInstantAnalysis *cache.Cache

func init() {
	metrics.Register("analysis.change.write.attempt", analysisChangeWriteAttemptCounter)
	metrics.Register("analysis.change.write", analysisChangeWriteCounter)

	go initializeAnalysisDaoPostConfiguration()
}

func initializeAnalysisDaoPostConfiguration() {
	config.WaitForConfigurationToBeLoaded()

	recentInstantAnalysis = cache.New(time.Duration(config.RecoveryPollSeconds*2)*time.Second, time.Second)
}

// GetReplicationAnalysis will check for replication problems (dead master; unreachable master; etc)
func GetReplicationAnalysis(clusterName string, includeDowntimed bool, auditAnalysis bool) ([]ReplicationAnalysis, error) {
	result := []ReplicationAnalysis{}

	args := sqlutils.Args(ValidSecondsFromSeenToLastAttemptedCheck(), clusterName)
	analysisQueryReductionClause := ``
	if config.Config.ReduceReplicationAnalysisCount {
		analysisQueryReductionClause = `
			HAVING
				(MIN(
					master_instance.last_checked <= master_instance.last_seen
					and master_instance.last_attempted_check <= master_instance.last_seen + interval ? second
       	 ) = 1 /* AS is_last_check_valid */) = 0
				OR (IFNULL(SUM(slave_instance.last_checked <= slave_instance.last_seen
		                    AND slave_instance.slave_io_running = 0
		                    AND slave_instance.last_io_error like '%error %connecting to master%'
		                    AND slave_instance.slave_sql_running = 1),
		                0) /* AS count_slaves_failing_to_connect_to_master */ > 0)
				OR (IFNULL(SUM(slave_instance.last_checked <= slave_instance.last_seen),
		                0) /* AS count_valid_slaves */ < COUNT(slave_instance.server_id) /* AS count_slaves */)
				OR (IFNULL(SUM(slave_instance.last_checked <= slave_instance.last_seen
		                    AND slave_instance.slave_io_running != 0
		                    AND slave_instance.slave_sql_running != 0),
		                0) /* AS count_valid_replicating_slaves */ < COUNT(slave_instance.server_id) /* AS count_slaves */)
				OR (MIN(
		            master_instance.slave_sql_running = 1
		            AND master_instance.slave_io_running = 0
		            AND master_instance.last_io_error like '%error %connecting to master%'
		          ) /* AS is_failing_to_connect_to_master */)
				OR (COUNT(slave_instance.server_id) /* AS count_slaves */ > 0)
			`
		args = append(args, ValidSecondsFromSeenToLastAttemptedCheck())
	}
	// "OR count_slaves > 0" above is a recent addition, which, granted, makes some previous conditions redundant.
	// It gives more output, and more "NoProblem" messages that I am now interested in for purpose of auditing in database_instance_analysis_changelog
	query := fmt.Sprintf(`
		    SELECT
		        master_instance.hostname,
		        master_instance.port,
						master_instance.data_center,
						master_instance.physical_environment,
		        MIN(master_instance.master_host) AS master_host,
		        MIN(master_instance.master_port) AS master_port,
		        MIN(master_instance.cluster_name) AS cluster_name,
		        MIN(IFNULL(cluster_alias.alias, master_instance.cluster_name)) AS cluster_alias,
		        MIN(
							master_instance.last_checked <= master_instance.last_seen
							and master_instance.last_attempted_check <= master_instance.last_seen + interval ? second
		        	) = 1 AS is_last_check_valid,
		        MIN(master_instance.master_host IN ('' , '_')
		            OR master_instance.master_port = 0
								OR substr(master_instance.master_host, 1, 2) = '//') AS is_master,
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
		                    AND slave_instance.last_io_error like '%%error %%connecting to master%%'
		                    AND slave_instance.slave_sql_running = 1),
		                0) AS count_slaves_failing_to_connect_to_master,
		        MIN(master_instance.replication_depth) AS replication_depth,
		        GROUP_CONCAT(concat(slave_instance.Hostname, ':', slave_instance.Port)) as slave_hosts,
		        MIN(
		            master_instance.slave_sql_running = 1
		            AND master_instance.slave_io_running = 0
		            AND master_instance.last_io_error like '%%error %%connecting to master%%'
		          ) AS is_failing_to_connect_to_master,
						MIN(
								master_downtime.downtime_active is not null
								and ifnull(master_downtime.end_timestamp, now()) > now()
							) AS is_downtimed,
			    	MIN(
				    		IFNULL(master_downtime.end_timestamp, '')
				    	) AS downtime_end_timestamp,
			    	MIN(
				    		IFNULL(unix_timestamp() - unix_timestamp(master_downtime.end_timestamp), 0)
				    	) AS downtime_remaining_seconds,
			    	MIN(
				    		master_instance.binlog_server
				    	) AS is_binlog_server,
			    	MIN(
				    		master_instance.pseudo_gtid
				    	) AS is_pseudo_gtid,
			    	MIN(
				    		master_instance.supports_oracle_gtid
				    	) AS supports_oracle_gtid,
			    	SUM(
				    		slave_instance.oracle_gtid
				    	) AS count_oracle_gtid_slaves,
			      IFNULL(SUM(slave_instance.last_checked <= slave_instance.last_seen
	              AND slave_instance.oracle_gtid != 0),
              0) AS count_valid_oracle_gtid_slaves,
			    	SUM(
				    		slave_instance.binlog_server
				    	) AS count_binlog_server_slaves,
		        IFNULL(SUM(slave_instance.last_checked <= slave_instance.last_seen
                  AND slave_instance.binlog_server != 0),
              0) AS count_valid_binlog_server_slaves,
			    	MIN(
				    		master_instance.mariadb_gtid
				    	) AS is_mariadb_gtid,
			    	SUM(
				    		slave_instance.mariadb_gtid
				    	) AS count_mariadb_gtid_slaves,
		        IFNULL(SUM(slave_instance.last_checked <= slave_instance.last_seen
                  AND slave_instance.mariadb_gtid != 0),
              0) AS count_valid_mariadb_gtid_slaves,
						IFNULL(SUM(slave_instance.log_bin
							  AND slave_instance.log_slave_updates
								AND slave_instance.binlog_format = 'STATEMENT'),
              0) AS count_statement_based_loggin_slaves,
						IFNULL(SUM(slave_instance.log_bin
								AND slave_instance.log_slave_updates
								AND slave_instance.binlog_format = 'MIXED'),
              0) AS count_mixed_based_loggin_slaves,
						IFNULL(SUM(slave_instance.log_bin
								AND slave_instance.log_slave_updates
								AND slave_instance.binlog_format = 'ROW'),
              0) AS count_row_based_loggin_slaves,
						IFNULL(SUM(
								replica_downtime.downtime_active is not null
								and ifnull(replica_downtime.end_timestamp, now()) > now()),
              0) AS count_downtimed_replicas,
						COUNT(DISTINCT case
								when slave_instance.log_bin AND slave_instance.log_slave_updates
								then slave_instance.major_version
								else NULL
							end
						) AS count_distinct_logging_major_versions
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
		        database_instance_downtime as master_downtime ON (master_instance.hostname = master_downtime.hostname
		        		AND master_instance.port = master_downtime.port
		        		AND master_downtime.downtime_active = 1)
					LEFT JOIN
		        database_instance_downtime as replica_downtime ON (slave_instance.hostname = replica_downtime.hostname
		        		AND slave_instance.port = replica_downtime.port
		        		AND replica_downtime.downtime_active = 1)
        	LEFT JOIN
		        cluster_alias ON (cluster_alias.cluster_name = master_instance.cluster_name)
				  LEFT JOIN
						database_instance_recent_relaylog_history ON (
								slave_instance.hostname = database_instance_recent_relaylog_history.hostname
		        		AND slave_instance.port = database_instance_recent_relaylog_history.port)
		    WHERE
		    	database_instance_maintenance.database_instance_maintenance_id IS NULL
		    	AND ? IN ('', master_instance.cluster_name)
		    GROUP BY
			    master_instance.hostname,
			    master_instance.port
			%s
		    ORDER BY
			    is_master DESC ,
			    is_cluster_master DESC,
			    count_slaves DESC
	`, analysisQueryReductionClause)
	err := db.QueryOrchestrator(query, args, func(m sqlutils.RowMap) error {
		a := ReplicationAnalysis{
			Analysis:               NoProblem,
			ProcessingNodeHostname: process.ThisHostname,
			ProcessingNodeToken:    process.ProcessToken.Hash,
		}

		a.IsMaster = m.GetBool("is_master")
		a.IsCoMaster = m.GetBool("is_co_master")
		a.AnalyzedInstanceKey = InstanceKey{Hostname: m.GetString("hostname"), Port: m.GetInt("port")}
		a.AnalyzedInstanceMasterKey = InstanceKey{Hostname: m.GetString("master_host"), Port: m.GetInt("master_port")}
		a.AnalyzedInstanceDataCenter = m.GetString("data_center")
		a.AnalyzedInstancePhysicalEnvironment = m.GetString("physical_environment")
		a.ClusterDetails.ClusterName = m.GetString("cluster_name")
		a.ClusterDetails.ClusterAlias = m.GetString("cluster_alias")
		a.LastCheckValid = m.GetBool("is_last_check_valid")
		a.CountReplicas = m.GetUint("count_slaves")
		a.CountValidReplicas = m.GetUint("count_valid_slaves")
		a.CountValidReplicatingReplicas = m.GetUint("count_valid_replicating_slaves")
		a.CountReplicasFailingToConnectToMaster = m.GetUint("count_slaves_failing_to_connect_to_master")
		a.CountDowntimedReplicas = m.GetUint("count_downtimed_replicas")
		a.CountStaleReplicas = 0
		a.ReplicationDepth = m.GetUint("replication_depth")
		a.IsFailingToConnectToMaster = m.GetBool("is_failing_to_connect_to_master")
		a.IsDowntimed = m.GetBool("is_downtimed")
		a.DowntimeEndTimestamp = m.GetString("downtime_end_timestamp")
		a.DowntimeRemainingSeconds = m.GetInt("downtime_remaining_seconds")
		a.IsBinlogServer = m.GetBool("is_binlog_server")
		a.ClusterDetails.ReadRecoveryInfo()

		a.SlaveHosts = *NewInstanceKeyMap()
		a.SlaveHosts.ReadCommaDelimitedList(m.GetString("slave_hosts"))

		countValidOracleGTIDSlaves := m.GetUint("count_valid_oracle_gtid_slaves")
		a.OracleGTIDImmediateTopology = countValidOracleGTIDSlaves == a.CountValidReplicas && a.CountValidReplicas > 0
		countValidMariaDBGTIDSlaves := m.GetUint("count_valid_mariadb_gtid_slaves")
		a.MariaDBGTIDImmediateTopology = countValidMariaDBGTIDSlaves == a.CountValidReplicas && a.CountValidReplicas > 0
		countValidBinlogServerSlaves := m.GetUint("count_valid_binlog_server_slaves")
		a.BinlogServerImmediateTopology = countValidBinlogServerSlaves == a.CountValidReplicas && a.CountValidReplicas > 0
		a.PseudoGTIDImmediateTopology = m.GetBool("is_pseudo_gtid")

		a.CountStatementBasedLoggingReplicas = m.GetUint("count_statement_based_loggin_slaves")
		a.CountMixedBasedLoggingReplicas = m.GetUint("count_mixed_based_loggin_slaves")
		a.CountRowBasedLoggingReplicas = m.GetUint("count_row_based_loggin_slaves")
		a.CountDistinctMajorVersionsLoggingReplicas = m.GetUint("count_distinct_logging_major_versions")

		if a.IsMaster && !a.LastCheckValid && a.CountReplicas == 0 {
			a.Analysis = DeadMasterWithoutSlaves
			a.Description = "Master cannot be reached by orchestrator and has no slave"
			//
		} else if a.IsMaster && !a.LastCheckValid && a.CountValidReplicas == a.CountReplicas && a.CountValidReplicatingReplicas == 0 {
			a.Analysis = DeadMaster
			a.Description = "Master cannot be reached by orchestrator and none of its replicas is replicating"
			//
		} else if a.IsMaster && !a.LastCheckValid && a.CountReplicas > 0 && a.CountValidReplicas == 0 && a.CountValidReplicatingReplicas == 0 {
			a.Analysis = DeadMasterAndSlaves
			a.Description = "Master cannot be reached by orchestrator and none of its replicas is replicating"
			//
		} else if a.IsMaster && !a.LastCheckValid && a.CountValidReplicas < a.CountReplicas && a.CountValidReplicas > 0 && a.CountValidReplicatingReplicas == 0 {
			a.Analysis = DeadMasterAndSomeSlaves
			a.Description = "Master cannot be reached by orchestrator; some of its replicas are unreachable and none of its reachable replicas is replicating"
			//
		} else if a.IsMaster && !a.LastCheckValid && a.CountStaleReplicas == a.CountReplicas && a.CountValidReplicatingReplicas > 0 {
			a.Analysis = UnreachableMasterWithStaleSlaves
			a.Description = "Master cannot be reached by orchestrator and has running yet stale replicas"
			//
		} else if a.IsMaster && !a.LastCheckValid && a.CountValidReplicas > 0 && a.CountValidReplicatingReplicas > 0 {
			a.Analysis = UnreachableMaster
			a.Description = "Master cannot be reached by orchestrator but it has replicating replicas; possibly a network/host issue"
			//
		} else if a.IsMaster && a.LastCheckValid && a.CountReplicas == 1 && a.CountValidReplicas == a.CountReplicas && a.CountValidReplicatingReplicas == 0 {
			a.Analysis = MasterSingleSlaveNotReplicating
			a.Description = "Master is reachable but its single slave is not replicating"
			//
		} else if a.IsMaster && a.LastCheckValid && a.CountReplicas == 1 && a.CountValidReplicas == 0 {
			a.Analysis = MasterSingleSlaveDead
			a.Description = "Master is reachable but its single slave is dead"
			//
		} else if a.IsMaster && a.LastCheckValid && a.CountReplicas > 1 && a.CountValidReplicas == a.CountReplicas && a.CountValidReplicatingReplicas == 0 {
			a.Analysis = AllMasterSlavesNotReplicating
			a.Description = "Master is reachable but none of its replicas is replicating"
			//
		} else if a.IsMaster && a.LastCheckValid && a.CountReplicas > 1 && a.CountValidReplicas < a.CountReplicas && a.CountValidReplicas > 0 && a.CountValidReplicatingReplicas == 0 {
			a.Analysis = AllMasterSlavesNotReplicatingOrDead
			a.Description = "Master is reachable but none of its replicas is replicating"
			//
		} else if a.IsMaster && a.LastCheckValid && a.CountReplicas > 1 && a.CountStaleReplicas == a.CountReplicas && a.CountValidReplicas > 0 && a.CountValidReplicatingReplicas > 0 {
			a.Analysis = AllMasterSlavesStale
			a.Description = "Master is reachable but all of its replicas are stale, although attempting to replicate"
			//
		} else /* co-master */ if a.IsCoMaster && !a.LastCheckValid && a.CountReplicas > 0 && a.CountValidReplicas == a.CountReplicas && a.CountValidReplicatingReplicas == 0 {
			a.Analysis = DeadCoMaster
			a.Description = "Co-master cannot be reached by orchestrator and none of its replicas is replicating"
			//
		} else if a.IsCoMaster && !a.LastCheckValid && a.CountReplicas > 0 && a.CountValidReplicas < a.CountReplicas && a.CountValidReplicas > 0 && a.CountValidReplicatingReplicas == 0 {
			a.Analysis = DeadCoMasterAndSomeSlaves
			a.Description = "Co-master cannot be reached by orchestrator; some of its replicas are unreachable and none of its reachable replicas is replicating"
			//
		} else if a.IsCoMaster && !a.LastCheckValid && a.CountValidReplicas > 0 && a.CountValidReplicatingReplicas > 0 {
			a.Analysis = UnreachableCoMaster
			a.Description = "Co-master cannot be reached by orchestrator but it has replicating replicas; possibly a network/host issue"
			//
		} else if a.IsCoMaster && a.LastCheckValid && a.CountReplicas > 0 && a.CountValidReplicatingReplicas == 0 {
			a.Analysis = AllCoMasterSlavesNotReplicating
			a.Description = "Co-master is reachable but none of its replicas is replicating"
			//
		} else /* intermediate-master */ if !a.IsMaster && !a.LastCheckValid && a.CountReplicas == 1 && a.CountValidReplicas == a.CountReplicas && a.CountReplicasFailingToConnectToMaster == a.CountReplicas && a.CountValidReplicatingReplicas == 0 {
			a.Analysis = DeadIntermediateMasterWithSingleSlaveFailingToConnect
			a.Description = "Intermediate master cannot be reached by orchestrator and its (single) slave is failing to connect"
			//
		} else /* intermediate-master */ if !a.IsMaster && !a.LastCheckValid && a.CountReplicas == 1 && a.CountValidReplicas == a.CountReplicas && a.CountValidReplicatingReplicas == 0 {
			a.Analysis = DeadIntermediateMasterWithSingleSlave
			a.Description = "Intermediate master cannot be reached by orchestrator and its (single) slave is not replicating"
			//
		} else /* intermediate-master */ if !a.IsMaster && !a.LastCheckValid && a.CountReplicas > 1 && a.CountValidReplicas == a.CountReplicas && a.CountValidReplicatingReplicas == 0 {
			a.Analysis = DeadIntermediateMaster
			a.Description = "Intermediate master cannot be reached by orchestrator and none of its replicas is replicating"
			//
		} else if !a.IsMaster && !a.LastCheckValid && a.CountValidReplicas < a.CountReplicas && a.CountValidReplicas > 0 && a.CountValidReplicatingReplicas == 0 {
			a.Analysis = DeadIntermediateMasterAndSomeSlaves
			a.Description = "Intermediate master cannot be reached by orchestrator; some of its replicas are unreachable and none of its reachable replicas is replicating"
			//
		} else if !a.IsMaster && !a.LastCheckValid && a.CountReplicas > 0 && a.CountValidReplicas == 0 {
			a.Analysis = DeadIntermediateMasterAndSlaves
			a.Description = "Intermediate master cannot be reached by orchestrator and all of its replicas are unreachable"
			//
		} else if !a.IsMaster && !a.LastCheckValid && a.CountValidReplicas > 0 && a.CountValidReplicatingReplicas > 0 {
			a.Analysis = UnreachableIntermediateMaster
			a.Description = "Intermediate master cannot be reached by orchestrator but it has replicating replicas; possibly a network/host issue"
			//
		} else if !a.IsMaster && a.LastCheckValid && a.CountReplicas > 1 && a.CountValidReplicatingReplicas == 0 &&
			a.CountReplicasFailingToConnectToMaster > 0 && a.CountReplicasFailingToConnectToMaster == a.CountValidReplicas {
			// All replicas are either failing to connect to master (and at least one of these have to exist)
			// or completely dead.
			// Must have at least two replicas to reach such conclusion -- do note that the intermediate master is still
			// reachable to orchestrator, so we base our conclusion on replicas only at this point.
			a.Analysis = AllIntermediateMasterSlavesFailingToConnectOrDead
			a.Description = "Intermediate master is reachable but all of its replicas are failing to connect"
			//
		} else if !a.IsMaster && a.LastCheckValid && a.CountReplicas > 0 && a.CountValidReplicatingReplicas == 0 {
			a.Analysis = AllIntermediateMasterSlavesNotReplicating
			a.Description = "Intermediate master is reachable but none of its replicas is replicating"
			//
		} else if a.IsBinlogServer && a.IsFailingToConnectToMaster {
			a.Analysis = BinlogServerFailingToConnectToMaster
			a.Description = "Binlog server is unable to connect to its master"
			//
		} else if a.ReplicationDepth == 1 && a.IsFailingToConnectToMaster {
			a.Analysis = FirstTierSlaveFailingToConnectToMaster
			a.Description = "1st tier slave (directly replicating from topology master) is unable to connect to the master"
			//
		}
		//		 else if a.IsMaster && a.CountReplicas == 0 {
		//			a.Analysis = MasterWithoutSlaves
		//			a.Description = "Master has no replicas"
		//		}

		appendAnalysis := func(analysis *ReplicationAnalysis) {
			if a.Analysis == NoProblem && len(a.StructureAnalysis) == 0 {
				return
			}
			skipThisHost := false
			for _, filter := range config.Config.RecoveryIgnoreHostnameFilters {
				if matched, _ := regexp.MatchString(filter, a.AnalyzedInstanceKey.Hostname); matched {
					skipThisHost = true
				}
			}
			if a.IsDowntimed && !includeDowntimed {
				skipThisHost = true
			}
			if a.CountReplicas == a.CountDowntimedReplicas {
				switch a.Analysis {
				case AllMasterSlavesNotReplicating,
					AllMasterSlavesNotReplicatingOrDead,
					AllCoMasterSlavesNotReplicating,
					AllIntermediateMasterSlavesFailingToConnectOrDead,
					AllIntermediateMasterSlavesNotReplicating:
					a.IsReplicasDowntimed = true
				}
			}
			if a.IsReplicasDowntimed && !includeDowntimed {
				skipThisHost = true
			}
			if !skipThisHost {
				result = append(result, a)
			}
		}

		{
			// Moving on to structure analysis
			// We also do structural checks. See if there's potential danger in promotions
			if a.IsMaster && a.CountStatementBasedLoggingReplicas > 0 && a.CountMixedBasedLoggingReplicas > 0 {
				a.StructureAnalysis = append(a.StructureAnalysis, StatementAndMixedLoggingSlavesStructureWarning)
			}
			if a.IsMaster && a.CountStatementBasedLoggingReplicas > 0 && a.CountRowBasedLoggingReplicas > 0 {
				a.StructureAnalysis = append(a.StructureAnalysis, StatementAndRowLoggingSlavesStructureWarning)
			}
			if a.IsMaster && a.CountMixedBasedLoggingReplicas > 0 && a.CountRowBasedLoggingReplicas > 0 {
				a.StructureAnalysis = append(a.StructureAnalysis, MixedAndRowLoggingSlavesStructureWarning)
			}
			if a.IsMaster && a.CountDistinctMajorVersionsLoggingReplicas > 1 {
				a.StructureAnalysis = append(a.StructureAnalysis, MultipleMajorVersionsLoggingSlaves)
			}
		}
		appendAnalysis(&a)

		if a.CountReplicas > 0 && auditAnalysis {
			// Interesting enough for analysis
			go auditInstanceAnalysisInChangelog(&a.AnalyzedInstanceKey, a.Analysis)
		}
		return nil
	})

	if err != nil {
		return result, log.Errore(err)
	}
	// TODO: result, err = getConcensusReplicationAnalysis(result)
	return result, log.Errore(err)
}

func getConcensusReplicationAnalysis(analysisEntries []ReplicationAnalysis) ([]ReplicationAnalysis, error) {
	if !orcraft.IsRaftEnabled() {
		return analysisEntries, nil
	}
	if !config.Config.ExpectFailureAnalysisConcensus {
		return analysisEntries, nil
	}
	concensusAnalysisEntries := []ReplicationAnalysis{}
	peerAnalysisMap, err := ReadPeerAnalysisMap()
	if err != nil {
		return analysisEntries, err
	}
	quorumSize, err := orcraft.QuorumSize()
	if err != nil {
		return analysisEntries, err
	}

	for _, analysisEntry := range analysisEntries {
		instanceAnalysis := NewInstanceAnalysis(&analysisEntry.AnalyzedInstanceKey, analysisEntry.Analysis)
		analysisKey := instanceAnalysis.String()

		peerAnalysisCount := peerAnalysisMap[analysisKey]
		if 1+peerAnalysisCount >= quorumSize {
			// this node and enough other nodes in agreement
			concensusAnalysisEntries = append(concensusAnalysisEntries, analysisEntry)
		}
	}
	return concensusAnalysisEntries, nil
}

// auditInstanceAnalysisInChangelog will write down an instance's analysis in the database_instance_analysis_changelog table.
// To not repeat recurring analysis code, the database_instance_last_analysis table is used, so that only changes to
// analysis codes are written.
func auditInstanceAnalysisInChangelog(instanceKey *InstanceKey, analysisCode AnalysisCode) error {
	if lastWrittenAnalysis, found := recentInstantAnalysis.Get(instanceKey.DisplayString()); found {
		if lastWrittenAnalysis == analysisCode {
			// Surely nothing new.
			// And let's expand the timeout
			recentInstantAnalysis.Set(instanceKey.DisplayString(), analysisCode, cache.DefaultExpiration)
			return nil
		}
	}
	// Passed the cache; but does database agree that there's a change? Here's a persistent cache; this comes here
	// to verify no two orchestrator services are doing this without coordinating (namely, one dies, the other taking its place
	// and has no familiarity of the former's cache)
	analysisChangeWriteAttemptCounter.Inc(1)

	lastAnalysisChanged := false
	{
		sqlResult, err := db.ExecOrchestrator(`
			update database_instance_last_analysis set
				analysis = ?,
				analysis_timestamp = now()
			where
				hostname = ?
				and port = ?
				and analysis != ?
			`,
			string(analysisCode), instanceKey.Hostname, instanceKey.Port, string(analysisCode),
		)
		if err != nil {
			return log.Errore(err)
		}
		rows, err := sqlResult.RowsAffected()
		if err != nil {
			return log.Errore(err)
		}
		lastAnalysisChanged = (rows > 0)
	}
	if !lastAnalysisChanged {
		_, err := db.ExecOrchestrator(`
			insert ignore into database_instance_last_analysis (
					hostname, port, analysis_timestamp, analysis
				) values (
					?, ?, now(), ?
				)
			`,
			instanceKey.Hostname, instanceKey.Port, string(analysisCode),
		)
		if err != nil {
			return log.Errore(err)
		}
	}
	recentInstantAnalysis.Set(instanceKey.DisplayString(), analysisCode, cache.DefaultExpiration)
	if !lastAnalysisChanged {
		return nil
	}

	_, err := db.ExecOrchestrator(`
			insert into database_instance_analysis_changelog (
					hostname, port, analysis_timestamp, analysis
				) values (
					?, ?, now(), ?
				)
			`,
		instanceKey.Hostname, instanceKey.Port, string(analysisCode),
	)
	if err == nil {
		analysisChangeWriteCounter.Inc(1)
	}
	return log.Errore(err)
}

// ExpireInstanceAnalysisChangelog removes old-enough analysis entries from the changelog
func ExpireInstanceAnalysisChangelog() error {
	_, err := db.ExecOrchestrator(`
			delete
				from database_instance_analysis_changelog
			where
				analysis_timestamp < now() - interval ? hour
			`,
		config.Config.UnseenInstanceForgetHours,
	)
	return log.Errore(err)
}

// ReadReplicationAnalysisChangelog
func ReadReplicationAnalysisChangelog() (res [](*ReplicationAnalysisChangelog), err error) {
	query := `
		select
      hostname,
      port,
			analysis_timestamp,
			analysis
		from
			database_instance_analysis_changelog
		order by
			hostname, port, changelog_id
		`
	analysisChangelog := &ReplicationAnalysisChangelog{}
	err = db.QueryOrchestratorRowsMap(query, func(m sqlutils.RowMap) error {
		key := InstanceKey{Hostname: m.GetString("hostname"), Port: m.GetInt("port")}

		if !analysisChangelog.AnalyzedInstanceKey.Equals(&key) {
			analysisChangelog = &ReplicationAnalysisChangelog{AnalyzedInstanceKey: key, Changelog: []string{}}
			res = append(res, analysisChangelog)
		}
		analysisEntry := fmt.Sprintf("%s;%s,", m.GetString("analysis_timestamp"), m.GetString("analysis"))
		analysisChangelog.Changelog = append(analysisChangelog.Changelog, analysisEntry)

		return nil
	})

	if err != nil {
		log.Errore(err)
	}
	return res, err
}

// ReadPeerAnalysisMap reads raft-peer failure analysis, and returns a PeerAnalysisMap,
// indicating how many peers see which analysis
func ReadPeerAnalysisMap() (peerAnalysisMap PeerAnalysisMap, err error) {
	peerAnalysisMap = make(PeerAnalysisMap)
	query := `
		select
      hostname,
      port,
			analysis
		from
			database_instance_peer_analysis
		order by
			peer, hostname, port
		`
	err = db.QueryOrchestratorRowsMap(query, func(m sqlutils.RowMap) error {
		instanceKey := InstanceKey{Hostname: m.GetString("hostname"), Port: m.GetInt("port")}
		analysis := m.GetString("analysis")
		instanceAnalysis := NewInstanceAnalysis(&instanceKey, AnalysisCode(analysis))
		mapKey := instanceAnalysis.String()
		peerAnalysisMap[mapKey] = peerAnalysisMap[mapKey] + 1

		return nil
	})
	return peerAnalysisMap, log.Errore(err)
}
