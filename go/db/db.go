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
	"github.com/outbrain/orchestrator/go/config"
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
          UNIQUE KEY maintenance_uidx (maintenance_active, hostname, port)
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
          KEY host_port_idx (hostname, port, audit_timestamp)
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
		  KEY is_successful_idx (is_successful, start_timestamp)
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
		  KEY agent_seed_idx (agent_seed_id, state_timestamp)
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
		DROP VIEW IF EXISTS _whats_wrong
	`,
	`
		DROP VIEW IF EXISTS whats_wrong
	`,
	`
		DROP VIEW IF EXISTS whats_wrong_summary
	`,
	`
		CREATE TABLE IF NOT EXISTS topology_recovery (
          recovery_id bigint unsigned not null auto_increment,
		  hostname varchar(128) NOT NULL,
		  port smallint unsigned NOT NULL,
          in_active_period tinyint unsigned NOT NULL DEFAULT 0,
          start_active_period timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
          end_active_period_unixtime int unsigned,
          end_recovery timestamp NULL,
		  processing_node_hostname varchar(128) CHARACTER SET ascii NOT NULL,
		  processcing_node_token varchar(128) NOT NULL,
		  successor_hostname varchar(128) DEFAULT NULL,
		  successor_port smallint unsigned DEFAULT NULL,
		  PRIMARY KEY (recovery_id),
          UNIQUE KEY hostname_port_active_period_uidx(hostname, port, in_active_period, end_active_period_unixtime),
		  KEY in_active_start_period_idx (in_active_period, start_active_period),
		  KEY start_active_period_idx (start_active_period)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		CREATE TABLE IF NOT EXISTS hostname_unresolve (
		  hostname varchar(128) NOT NULL,
		  unresolved_hostname varchar(128) NOT NULL,
		  PRIMARY KEY (hostname),
		  KEY unresolved_hostname_idx (unresolved_hostname)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		CREATE TABLE IF NOT EXISTS database_instance_pool (
          hostname varchar(128) CHARACTER SET ascii NOT NULL,
          port smallint(5) unsigned NOT NULL,
		  pool varchar(128) NOT NULL,
		  PRIMARY KEY (hostname, port, pool),
		  KEY pool_idx (pool)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
        CREATE TABLE IF NOT EXISTS database_instance_topology_history (
          snapshot_unix_timestamp INT UNSIGNED NOT NULL,
          hostname varchar(128) CHARACTER SET ascii NOT NULL,
          port smallint(5) unsigned NOT NULL,
          master_host varchar(128) CHARACTER SET ascii NOT NULL,
          master_port smallint(5) unsigned NOT NULL,
          cluster_name tinytext CHARACTER SET ascii NOT NULL,
          PRIMARY KEY (snapshot_unix_timestamp, hostname, port),
          KEY cluster_name_idx (snapshot_unix_timestamp, cluster_name(128))
        ) ENGINE=InnoDB DEFAULT CHARSET=ascii

	`,
	`
        CREATE TABLE IF NOT EXISTS candidate_database_instance (
          hostname varchar(128) CHARACTER SET ascii NOT NULL,
          port smallint(5) unsigned NOT NULL,
          last_suggested TIMESTAMP NOT NULL,
          PRIMARY KEY (hostname, port),
          KEY last_suggested_idx (last_suggested)
        ) ENGINE=InnoDB DEFAULT CHARSET=ascii

	`,
	`
        CREATE TABLE IF NOT EXISTS database_instance_downtime (
          hostname varchar(128) NOT NULL,
          port smallint(5) unsigned NOT NULL,
          downtime_active tinyint(4) DEFAULT NULL,
          begin_timestamp timestamp DEFAULT CURRENT_TIMESTAMP,
          end_timestamp timestamp,
          owner varchar(128) CHARACTER SET utf8 NOT NULL,
          reason text CHARACTER SET utf8 NOT NULL,
          PRIMARY KEY (hostname, port)
        ) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
        CREATE TABLE IF NOT EXISTS topology_failure_detection (
          detection_id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
          hostname varchar(128) NOT NULL,
          port smallint unsigned NOT NULL,
          in_active_period tinyint unsigned NOT NULL DEFAULT '0',
          start_active_period timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
          end_active_period_unixtime int unsigned NOT NULL,
          processing_node_hostname varchar(128) NOT NULL,
          processcing_node_token varchar(128) NOT NULL,
          analysis varchar(128) NOT NULL,
          cluster_name varchar(128) NOT NULL,
          cluster_alias varchar(128) NOT NULL,
          count_affected_slaves int unsigned NOT NULL,
          slave_hosts text NOT NULL,
          PRIMARY KEY (detection_id),
          UNIQUE KEY hostname_port_active_period_uidx (hostname, port, in_active_period, end_active_period_unixtime),
          KEY in_active_start_period_idx (in_active_period, start_active_period)
        ) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		CREATE TABLE IF NOT EXISTS hostname_resolve_history (
		  resolved_hostname varchar(128) NOT NULL,
		  hostname varchar(128) NOT NULL,
		  resolved_timestamp timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		  PRIMARY KEY (resolved_hostname),
		  KEY (hostname),
		  KEY resolved_timestamp_idx (resolved_timestamp)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		CREATE TABLE IF NOT EXISTS hostname_unresolve_history (
		  unresolved_hostname varchar(128) NOT NULL,
		  hostname varchar(128) NOT NULL,
		  last_registered TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		  PRIMARY KEY (unresolved_hostname),
		  KEY (hostname),
		  KEY last_registered_idx (last_registered)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
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
			ADD COLUMN last_sql_error TEXT NOT NULL AFTER exec_master_log_pos
	`,
	`
		ALTER TABLE 
			database_instance
			ADD COLUMN last_io_error TEXT NOT NULL AFTER last_sql_error
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
	`
		ALTER TABLE 
			database_instance
			ADD COLUMN data_center varchar(32) CHARACTER SET ascii NOT NULL AFTER cluster_name
	`,
	`
		ALTER TABLE 
			database_instance
			ADD COLUMN physical_environment varchar(32) CHARACTER SET ascii NOT NULL AFTER data_center
	`,
	`
		ALTER TABLE 
			database_instance_maintenance
			ADD KEY active_timestamp_idx (maintenance_active, begin_timestamp)
	`,
	`
		ALTER TABLE 
			database_instance
			ADD COLUMN uptime INT UNSIGNED NOT NULL AFTER last_seen
	`,
	`
		ALTER TABLE
			cluster_alias
			ADD UNIQUE KEY alias_uidx (alias)
	`,
	`
		ALTER TABLE 
			database_instance
			ADD COLUMN is_co_master TINYINT UNSIGNED NOT NULL AFTER replication_depth
	`,
	`
		ALTER TABLE 
			database_instance_maintenance
			ADD KEY active_end_timestamp_idx (maintenance_active, end_timestamp)
	`,
	`
		ALTER TABLE 
			database_instance
			ADD COLUMN sql_delay INT UNSIGNED NOT NULL AFTER slave_lag_seconds
	`,
	`
		ALTER TABLE 
			topology_recovery
			ADD COLUMN analysis              varchar(128) CHARACTER SET ascii NOT NULL,
			ADD COLUMN cluster_name          varchar(128) CHARACTER SET ascii NOT NULL,
			ADD COLUMN cluster_alias         varchar(128) CHARACTER SET ascii NOT NULL,
			ADD COLUMN count_affected_slaves int unsigned NOT NULL,
			ADD COLUMN slave_hosts text CHARACTER SET ascii NOT NULL
	`,
	`
		ALTER TABLE hostname_unresolve
			ADD COLUMN last_registered TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			ADD KEY last_registered_idx (last_registered)
	`,
	`
		ALTER TABLE topology_recovery
			ADD KEY cluster_name_in_active_idx (cluster_name, in_active_period)
	`,
	`
		ALTER TABLE topology_recovery
			ADD KEY end_recovery_idx (end_recovery)
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
		if !config.Config.SkipOrchestratorDatabaseUpdate {
			initOrchestratorDB(db)
		}
		db.SetMaxIdleConns(10)
	}
	return db, err
}

// initOrchestratorDB attempts to create/upgrade the orchestrator backend database. It is created once in the
// application's lifetime.
func initOrchestratorDB(db *sql.DB) error {
	log.Debug("Initializing orchestrator")
	for _, query := range generateSQL {
		_, err := execInternal(db, query)
		if err != nil {
			return log.Fatalf("Cannot initiate orchestrator: %+v", err)
		}
	}
	for _, query := range generateSQLPatches {
		// Patches are allowed to fail.
		_, _ = execInternalSilently(db, query)
	}
	return nil
}

// ExecOrchestrator will execute given query on the orchestrator backend database.
func execInternalSilently(db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	res, err := sqlutils.ExecSilently(db, query, args...)
	return res, err
}

// ExecOrchestrator will execute given query on the orchestrator backend database.
func execInternal(db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
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
