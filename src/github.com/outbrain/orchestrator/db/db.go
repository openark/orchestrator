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

package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/outbrain/golib/log"
	"github.com/outbrain/golib/sqlutils"
	"github.com/outbrain/orchestrator/config"
)

// generateSQL & generateSQLPatches are lists of SQL statements required to build the orchestrator backend
var generateSQL = []string{
	`
        CREATE TABLE IF NOT EXISTS database_instance (
          hostname varchar(128) CHARACTER SET ascii NOT NULL,
          port smallint(5) unsigned NOT NULL,
          last_checked timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
          last_seen timestamp NULL DEFAULT NULL,
          server_id int(10) unsigned NOT NULL,
          version varchar(128) CHARACTER SET ascii NOT NULL,
          binlog_format varchar(16) CHARACTER SET ascii NOT NULL,
          log_bin tinyint(3) unsigned NOT NULL,
          log_slave_updates tinyint(3) unsigned NOT NULL,
          binary_log_file varchar(128) CHARACTER SET ascii NOT NULL,
          binary_log_pos bigint(20) unsigned NOT NULL,
          master_host varchar(128) CHARACTER SET ascii NOT NULL,
          master_port smallint(5) unsigned NOT NULL,
          slave_sql_running tinyint(3) unsigned NOT NULL,
          slave_io_running tinyint(3) unsigned NOT NULL,
          master_log_file varchar(128) CHARACTER SET ascii NOT NULL,
          read_master_log_pos bigint(20) unsigned NOT NULL,
          relay_master_log_file varchar(128) CHARACTER SET ascii NOT NULL,
          exec_master_log_pos bigint(20) unsigned NOT NULL,
          seconds_behind_master bigint(20) unsigned DEFAULT NULL,
          slave_lag_seconds bigint(20) unsigned DEFAULT NULL,
          num_slave_hosts int(10) unsigned NOT NULL,
          slave_hosts text CHARACTER SET ascii NOT NULL,
          cluster_name tinytext CHARACTER SET ascii NOT NULL,
          PRIMARY KEY (hostname,port),
          KEY cluster_name_idx (cluster_name(128)),
          KEY last_checked_idx (last_checked),
          KEY last_seen_idx (last_seen)
        ) ENGINE=InnoDB DEFAULT CHARSET=ascii

	`,
	`
        CREATE TABLE IF NOT EXISTS database_instance_maintenance (
          database_instance_maintenance_id int(10) unsigned NOT NULL AUTO_INCREMENT,
          hostname varchar(128) NOT NULL,
          port smallint(5) unsigned NOT NULL,
          maintenance_active tinyint(4) DEFAULT NULL,
          begin_timestamp timestamp NULL DEFAULT NULL,
          end_timestamp timestamp NULL DEFAULT NULL,
          owner varchar(128) CHARACTER SET utf8 NOT NULL,
          reason text CHARACTER SET utf8 NOT NULL,
          PRIMARY KEY (database_instance_maintenance_id),
          UNIQUE KEY maintenance_uidx (maintenance_active,hostname,port)
        ) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
        CREATE TABLE IF NOT EXISTS database_instance_long_running_queries (
          hostname varchar(128) NOT NULL,
          port smallint(5) unsigned NOT NULL,
          process_id bigint(20) NOT NULL,
          process_started_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
          process_user varchar(16) CHARACTER SET utf8 NOT NULL,
          process_host varchar(128) CHARACTER SET utf8 NOT NULL,
          process_db varchar(128) CHARACTER SET utf8 NOT NULL,
          process_command varchar(16) CHARACTER SET utf8 NOT NULL,
          process_time_seconds int(11) NOT NULL,
          process_state varchar(128) CHARACTER SET utf8 NOT NULL,
          process_info varchar(1024) CHARACTER SET utf8 NOT NULL,
          PRIMARY KEY (hostname,port,process_id),
          KEY process_started_at_idx (process_started_at)
        ) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
        CREATE TABLE IF NOT EXISTS audit (
          audit_id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
          audit_timestamp timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
          audit_type varchar(128) CHARACTER SET ascii NOT NULL,
          hostname varchar(128) CHARACTER SET ascii NOT NULL DEFAULT '',
          port smallint(5) unsigned NOT NULL,
          message text CHARACTER SET utf8 NOT NULL,
          PRIMARY KEY (audit_id),
          KEY audit_timestamp_idx (audit_timestamp),
          KEY host_port_idx (hostname,port,audit_timestamp)
        ) ENGINE=InnoDB DEFAULT CHARSET=latin1 
	`,
	`
		CREATE TABLE IF NOT EXISTS host_agent (
		  hostname varchar(128) NOT NULL,
		  port smallint(5) unsigned NOT NULL,
		  token varchar(128) NOT NULL,
		  last_submitted timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		  last_checked timestamp NULL DEFAULT NULL,
		  last_seen timestamp NULL DEFAULT NULL,
		  mysql_port smallint(5) unsigned DEFAULT NULL,
		  count_mysql_snapshots smallint(5) unsigned NOT NULL,
		  PRIMARY KEY (hostname),
		  KEY token_idx (token(32)),
		  KEY last_submitted_idx (last_submitted),
		  KEY last_checked_idx (last_checked),
		  KEY last_seen_idx (last_seen)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		CREATE TABLE IF NOT EXISTS agent_seed (
		  agent_seed_id int(10) unsigned NOT NULL AUTO_INCREMENT,
		  target_hostname varchar(128) NOT NULL,
		  source_hostname varchar(128) NOT NULL,
		  start_timestamp timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		  end_timestamp timestamp NOT NULL,
		  is_complete tinyint(3) unsigned NOT NULL DEFAULT '0',
		  is_successful tinyint(3) unsigned NOT NULL DEFAULT '0',
		  PRIMARY KEY (agent_seed_id),
		  KEY target_hostname_idx (target_hostname,is_complete),
		  KEY source_hostname_idx (source_hostname,is_complete),
		  KEY start_timestamp_idx (start_timestamp),
		  KEY is_complete_idx (is_complete,start_timestamp),
		  KEY is_successful_idx (is_successful,start_timestamp)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		CREATE TABLE IF NOT EXISTS agent_seed_state (
		  agent_seed_state_id int(10) unsigned NOT NULL AUTO_INCREMENT,
		  agent_seed_id int(10) unsigned NOT NULL,
		  state_timestamp timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		  state_action varchar(127) NOT NULL,
		  error_message varchar(255) NOT NULL,
		  PRIMARY KEY (agent_seed_state_id),
		  KEY agent_seed_idx (agent_seed_id,state_timestamp)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		CREATE TABLE IF NOT EXISTS host_attributes (
		  hostname varchar(128) NOT NULL,
		  attribute_name varchar(128) NOT NULL,
		  attribute_value varchar(128) NOT NULL,
		  submit_timestamp timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		  expire_timestamp timestamp NULL DEFAULT NULL,
		  PRIMARY KEY (hostname,attribute_name),
		  KEY attribute_name_idx (attribute_name),
		  KEY attribute_value_idx (attribute_value),
		  KEY submit_timestamp_idx (submit_timestamp),
		  KEY expire_timestamp_idx (expire_timestamp)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		CREATE TABLE IF NOT EXISTS hostname_resolve (
		  hostname varchar(128) NOT NULL,
		  resolved_hostname varchar(128) NOT NULL,
		  resolved_timestamp timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		  PRIMARY KEY (hostname),
		  KEY resolved_timestamp_idx (resolved_timestamp)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		CREATE TABLE IF NOT EXISTS cluster_alias (
		  cluster_name varchar(128) CHARACTER SET ascii NOT NULL,
		  alias varchar(128) NOT NULL,
		  PRIMARY KEY (cluster_name)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		CREATE TABLE IF NOT EXISTS active_node (
		  anchor tinyint unsigned NOT NULL,
		  hostname varchar(128) CHARACTER SET ascii NOT NULL,
		  token varchar(128) NOT NULL,
		  last_seen_active timestamp NOT NULL,
		  PRIMARY KEY (anchor)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		INSERT IGNORE INTO active_node (anchor, hostname, token, last_seen_active)
			VALUES (1, '', '', NOW())
	`,
	`
		CREATE TABLE IF NOT EXISTS node_health (
		  hostname varchar(128) CHARACTER SET ascii NOT NULL,
		  token varchar(128) NOT NULL,
		  last_seen_active timestamp NOT NULL,
		  PRIMARY KEY (hostname)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		CREATE OR REPLACE VIEW _whats_wrong AS
		SELECT 
		    master_instance.hostname,
		    master_instance.port,
		    MIN(master_instance.cluster_name) AS cluster_name,
		    MIN(master_instance.last_checked <= master_instance.last_seen)
		        IS TRUE AS is_last_check_valid,
		    MIN(master_instance.master_host IN ('' , '_')
		        OR master_instance.master_port = 0) AS is_master,
		    MIN(CONCAT(master_instance.hostname,
		            ':',
		            master_instance.port) = master_instance.cluster_name) AS is_cluster_master,
		    COUNT(slave_instance.server_id) AS count_slaves,
		    IFNULL(SUM(slave_instance.last_checked <= slave_instance.last_seen),
		            0) AS count_valid_slaves,
		    IFNULL(SUM(slave_instance.last_checked <= slave_instance.last_seen
		                AND slave_instance.slave_io_running != 0
		                AND slave_instance.slave_sql_running != 0),
		            0) AS count_valid_replicating_slaves
		FROM
		    database_instance master_instance
		        LEFT JOIN
		    hostname_resolve ON (master_instance.hostname = hostname_resolve.hostname)
		        LEFT JOIN
		    database_instance slave_instance ON (COALESCE(hostname_resolve.resolved_hostname,
		            master_instance.hostname) = slave_instance.master_host
		        AND master_instance.port = slave_instance.master_port)
		GROUP BY 
			master_instance.hostname, 
			master_instance.port
		ORDER BY 
			is_master DESC , 
			is_cluster_master DESC, 
			count_slaves DESC
	`,
	`
		CREATE OR REPLACE VIEW whats_wrong AS
			(
				SELECT 
					hostname,
					port,
					cluster_name,
					'dead_master' as analysis,
					'This master cannot be reached by orchestrator and none of its slaves is replicating' as description,
					is_last_check_valid,
					count_slaves,
					count_valid_slaves,
					count_valid_replicating_slaves
				FROM _whats_wrong
				WHERE
					is_master
					AND is_last_check_valid = 0
					AND count_valid_slaves = count_slaves
					AND count_valid_replicating_slaves = 0
			)
			UNION ALL
			(
				SELECT 
					hostname,
					port,
					cluster_name,
					'dead_master_and_slaves' as analysis,
					'This master cannot be reached by orchestrator; all of its slaves are unreachable' as description,
					is_last_check_valid,
					count_slaves,
					count_valid_slaves,
					count_valid_replicating_slaves
				FROM _whats_wrong
				WHERE
					is_master
					AND count_slaves > 0
					AND is_last_check_valid = 0
					AND count_valid_slaves = 0
					AND count_valid_replicating_slaves = 0
			)
			UNION ALL
			(
				SELECT 
					hostname,
					port,
					cluster_name,
					'dead_master_and_some_slaves' as analysis,
					'This master cannot be reached by orchestrator; some of its slaves are unreachable and none of its reachable slaves is replicating' as description,
					is_last_check_valid,
					count_slaves,
					count_valid_slaves,
					count_valid_replicating_slaves
				FROM _whats_wrong
				WHERE
					is_master
					AND is_last_check_valid = 0
					AND count_valid_slaves < count_slaves
					AND count_valid_slaves > 0
					AND count_valid_replicating_slaves = 0
			)
			UNION ALL
			(
				SELECT 
					hostname,
					port,
					cluster_name,
					'unreachable_master' as analysis,
					'This master cannot be reached by orchestrator but it has replicating slaves; possibly a network/host issue' as description,
					is_last_check_valid,
					count_slaves,
					count_valid_slaves,
					count_valid_replicating_slaves
				FROM _whats_wrong
				WHERE
					is_master
					AND is_last_check_valid = 0
					AND count_valid_slaves > 0
					AND count_valid_replicating_slaves > 0
			)
			UNION ALL
			(
				SELECT 
					hostname,
					port,
					cluster_name,
					'all_slaves_not_replicating' as analysis,
					'The master is reachable but none of its slaves is replicating' as description,
					is_last_check_valid,
					count_slaves,
					count_valid_slaves,
					count_valid_replicating_slaves
				FROM _whats_wrong
				WHERE
					is_master
					AND is_last_check_valid = 1
					AND count_slaves > 0
					AND count_valid_replicating_slaves = 0
			)
			UNION ALL
			(
				SELECT 
					hostname,
					port,
					cluster_name,
					'master_without_slaves' as analysis,
					'The master does not have any slaves' as description,
					is_last_check_valid,
					count_slaves,
					count_valid_slaves,
					count_valid_replicating_slaves
				FROM _whats_wrong
				WHERE
					is_master
					AND count_slaves = 0
			)
	`,
}

var generateSQLPatches = []string{
	`
		ALTER TABLE 
			database_instance
			ADD COLUMN read_only TINYINT UNSIGNED NOT NULL AFTER version
	`,
	`
		ALTER TABLE 
			database_instance
			ADD COLUMN last_sql_error TEXT NOT NULL DEFAULT '' AFTER exec_master_log_pos
	`,
	`
		ALTER TABLE 
			database_instance
			ADD COLUMN last_io_error TEXT NOT NULL DEFAULT '' AFTER last_sql_error
	`,
	`
		ALTER TABLE 
			database_instance
			ADD COLUMN last_attempted_check TIMESTAMP AFTER last_checked
	`,
	`
		ALTER TABLE 
			database_instance
			ADD COLUMN oracle_gtid TINYINT UNSIGNED NOT NULL AFTER slave_io_running
	`,
	`
		ALTER TABLE 
			database_instance
			ADD COLUMN mariadb_gtid TINYINT UNSIGNED NOT NULL AFTER oracle_gtid
	`,
	`
		ALTER TABLE 
			database_instance
			ADD COLUMN relay_log_file varchar(128) CHARACTER SET ascii NOT NULL AFTER exec_master_log_pos
	`,
	`
		ALTER TABLE 
			database_instance
			ADD COLUMN relay_log_pos bigint unsigned NOT NULL AFTER relay_log_file
	`,
	`
		ALTER TABLE 
			database_instance
			ADD INDEX master_host_port_idx (master_host, master_port)
	`,
	`
		ALTER TABLE 
			database_instance
			ADD COLUMN pseudo_gtid TINYINT UNSIGNED NOT NULL AFTER mariadb_gtid
	`,
	`
		ALTER TABLE 
			database_instance
			ADD COLUMN replication_depth TINYINT UNSIGNED NOT NULL AFTER cluster_name
	`,
	`
		ALTER TABLE 
			database_instance
			ADD COLUMN has_replication_filters TINYINT UNSIGNED NOT NULL AFTER slave_io_running
	`,
}

// OpenTopology returns a DB instance to access a topology instance
func OpenTopology(host string, port int) (*sql.DB, error) {
	mysql_uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/?timeout=%ds", config.Config.MySQLTopologyUser, config.Config.MySQLTopologyPassword, host, port, config.Config.MySQLConnectTimeoutSeconds)
	db, _, err := sqlutils.GetDB(mysql_uri)
	db.SetMaxOpenConns(config.Config.MySQLTopologyMaxPoolConnections)
	db.SetMaxIdleConns(config.Config.MySQLTopologyMaxPoolConnections)
	return db, err
}

// OpenTopology returns the DB instance for the orchestrator backed database
func OpenOrchestrator() (*sql.DB, error) {
	mysql_uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=%ds", config.Config.MySQLOrchestratorUser, config.Config.MySQLOrchestratorPassword,
		config.Config.MySQLOrchestratorHost, config.Config.MySQLOrchestratorPort, config.Config.MySQLOrchestratorDatabase, config.Config.MySQLConnectTimeoutSeconds)
	db, fromCache, err := sqlutils.GetDB(mysql_uri)
	if err == nil && !fromCache {
		initOrchestratorDB(db)
		db.SetMaxIdleConns(10)
	}
	return db, err
}

// initOrchestratorDB attempts to create/upgrade the orchestrator backend database. It is created once in the
// application's lifetime.
func initOrchestratorDB(db *sql.DB) error {
	log.Debug("Initializing orchestrator")
	for _, query := range generateSQL {
		_, err := ExecOrchestrator(query)
		if err != nil {
			return log.Fatalf("Cannot initiate orchestrator: %+v", err)
		}
	}
	for _, query := range generateSQLPatches {
		// Patches are allowed to fail.
		_, _ = execOrchestratorSilently(query)
	}
	return nil
}

// ExecOrchestrator will execute given query on the orchestrator backend database.
func execOrchestratorSilently(query string, args ...interface{}) (sql.Result, error) {
	db, err := OpenOrchestrator()
	if err != nil {
		return nil, err
	}
	res, err := sqlutils.ExecSilently(db, query, args...)
	return res, err
}

// ExecOrchestrator will execute given query on the orchestrator backend database.
func ExecOrchestrator(query string, args ...interface{}) (sql.Result, error) {
	db, err := OpenOrchestrator()
	if err != nil {
		return nil, err
	}
	res, err := sqlutils.Exec(db, query, args...)
	return res, err
}
