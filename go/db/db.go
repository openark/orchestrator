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
	"crypto/tls"
	"database/sql"
	"fmt"
	"strings"

	"github.com/github/orchestrator/go/config"
	"github.com/github/orchestrator/go/ssl"
	"github.com/go-sql-driver/mysql"
	"github.com/openark/golib/log"
	"github.com/openark/golib/sqlutils"
)

var (
	EmptyArgs []interface{}
)

type DummySqlResult struct {
}

func (this DummySqlResult) LastInsertId() (int64, error) {
	return 0, nil
}

func (this DummySqlResult) RowsAffected() (int64, error) {
	return 1, nil
}

// generateSQLBase & generateSQLPatches are lists of SQL statements required to build the orchestrator backend
var generateSQLBase = []string{
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
          cluster_name varchar(128) CHARACTER SET ascii NOT NULL,
          PRIMARY KEY (hostname,port)
        ) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
				DROP INDEX cluster_name_idx ON database_instance
	`,
	`
				CREATE INDEX cluster_name_idx_database_instance ON database_instance(cluster_name)
	`,
	`
				DROP INDEX last_checked_idx ON database_instance
	`,
	`
				CREATE INDEX last_checked_idx_database_instance ON database_instance(last_checked)
	`,
	`
				DROP INDEX last_seen_idx ON database_instance
	`,
	`
				CREATE INDEX last_seen_idx_database_instance ON database_instance(last_seen)
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
          PRIMARY KEY (database_instance_maintenance_id)
        ) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
				DROP INDEX maintenance_uidx ON database_instance_maintenance
	`,
	`
				CREATE UNIQUE INDEX maintenance_uidx_database_instance_maintenance ON database_instance_maintenance (maintenance_active, hostname, port)
	`,
	`
        CREATE TABLE IF NOT EXISTS database_instance_long_running_queries (
          hostname varchar(128) NOT NULL,
          port smallint(5) unsigned NOT NULL,
          process_id bigint(20) NOT NULL,
          process_started_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
          process_user varchar(16) CHARACTER SET utf8 NOT NULL,
          process_host varchar(128) CHARACTER SET utf8 NOT NULL,
          process_db varchar(128) CHARACTER SET utf8 NOT NULL,
          process_command varchar(16) CHARACTER SET utf8 NOT NULL,
          process_time_seconds int(11) NOT NULL,
          process_state varchar(128) CHARACTER SET utf8 NOT NULL,
          process_info varchar(1024) CHARACTER SET utf8 NOT NULL,
          PRIMARY KEY (hostname,port,process_id)
        ) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
				DROP INDEX process_started_at_idx ON database_instance_long_running_queries
	`,
	`
				CREATE INDEX process_started_at_idx_database_instance_long_running_queries ON database_instance_long_running_queries (process_started_at)
	`,
	`
        CREATE TABLE IF NOT EXISTS audit (
          audit_id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
          audit_timestamp timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
          audit_type varchar(128) CHARACTER SET ascii NOT NULL,
          hostname varchar(128) CHARACTER SET ascii NOT NULL DEFAULT '',
          port smallint(5) unsigned NOT NULL,
          message text CHARACTER SET utf8 NOT NULL,
          PRIMARY KEY (audit_id)
        ) ENGINE=InnoDB DEFAULT CHARSET=latin1
	`,
	`
				DROP INDEX audit_timestamp_idx ON audit
	`,
	`
				CREATE INDEX audit_timestamp_idx_audit ON audit (audit_timestamp)
	`,
	`
				DROP INDEX host_port_idx ON audit
	`,
	`
				CREATE INDEX host_port_idx_audit ON audit (hostname, port, audit_timestamp)
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
		  PRIMARY KEY (hostname)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
				DROP INDEX token_idx ON host_agent
	`,
	`
				CREATE INDEX token_idx_host_agent ON host_agent (token)
	`,
	`
				DROP INDEX last_submitted_idx ON host_agent
	`,
	`
				CREATE INDEX last_submitted_idx_host_agent ON host_agent (last_submitted)
	`,
	`
				DROP INDEX last_checked_idx ON host_agent
	`,
	`
				CREATE INDEX last_checked_idx_host_agent ON host_agent (last_checked)
	`,
	`
				DROP INDEX last_seen_idx ON host_agent
	`,
	`
				CREATE INDEX last_seen_idx_host_agent ON host_agent (last_seen)
	`,
	`
		CREATE TABLE IF NOT EXISTS agent_seed (
		  agent_seed_id int(10) unsigned NOT NULL AUTO_INCREMENT,
		  target_hostname varchar(128) NOT NULL,
		  source_hostname varchar(128) NOT NULL,
		  start_timestamp timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		  end_timestamp timestamp NOT NULL DEFAULT '1971-01-01 00:00:00',
		  is_complete tinyint(3) unsigned NOT NULL DEFAULT '0',
		  is_successful tinyint(3) unsigned NOT NULL DEFAULT '0',
		  PRIMARY KEY (agent_seed_id)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
				DROP INDEX target_hostname_idx ON agent_seed
	`,
	`
				CREATE INDEX target_hostname_idx_agent_seed ON agent_seed (target_hostname,is_complete)
	`,
	`
				DROP INDEX source_hostname_idx ON agent_seed
	`,
	`
				CREATE INDEX source_hostname_idx_agent_seed ON agent_seed (source_hostname,is_complete)
	`,
	`
				DROP INDEX start_timestamp_idx ON agent_seed
	`,
	`
				CREATE INDEX start_timestamp_idx_agent_seed ON agent_seed (start_timestamp)
	`,
	`
				DROP INDEX is_complete_idx ON agent_seed
	`,
	`
				CREATE INDEX is_complete_idx_agent_seed ON agent_seed (is_complete,start_timestamp)
	`,
	`
				DROP INDEX is_successful_idx ON agent_seed
	`,
	`
				CREATE INDEX is_successful_idx_agent_seed ON agent_seed (is_successful, start_timestamp)
	`,
	`
		CREATE TABLE IF NOT EXISTS agent_seed_state (
		  agent_seed_state_id int(10) unsigned NOT NULL AUTO_INCREMENT,
		  agent_seed_id int(10) unsigned NOT NULL,
		  state_timestamp timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		  state_action varchar(127) NOT NULL,
		  error_message varchar(255) NOT NULL,
		  PRIMARY KEY (agent_seed_state_id)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
				DROP INDEX agent_seed_idx ON agent_seed_state
	`,
	`
				CREATE INDEX agent_seed_idx_agent_seed_state ON agent_seed_state (agent_seed_id, state_timestamp)
	`,
	`
		CREATE TABLE IF NOT EXISTS host_attributes (
		  hostname varchar(128) NOT NULL,
		  attribute_name varchar(128) NOT NULL,
		  attribute_value varchar(128) NOT NULL,
		  submit_timestamp timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		  expire_timestamp timestamp NULL DEFAULT NULL,
		  PRIMARY KEY (hostname,attribute_name)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		DROP INDEX attribute_name_idx ON host_attributes
	`,
	`
		CREATE INDEX attribute_name_idx_host_attributes ON host_attributes (attribute_name)
	`,
	`
		DROP INDEX attribute_value_idx ON host_attributes
	`,
	`
		CREATE INDEX attribute_value_idx_host_attributes ON host_attributes (attribute_value)
	`,
	`
		DROP INDEX submit_timestamp_idx ON host_attributes
	`,
	`
		CREATE INDEX submit_timestamp_idx_host_attributes ON host_attributes (submit_timestamp)
	`,
	`
		DROP INDEX expire_timestamp_idx ON host_attributes
	`,
	`
		CREATE INDEX expire_timestamp_idx_host_attributes ON host_attributes (expire_timestamp)
	`,
	`
		CREATE TABLE IF NOT EXISTS hostname_resolve (
		  hostname varchar(128) NOT NULL,
		  resolved_hostname varchar(128) NOT NULL,
		  resolved_timestamp timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		  PRIMARY KEY (hostname)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		DROP INDEX resolved_timestamp_idx ON hostname_resolve
	`,
	`
		CREATE INDEX resolved_timestamp_idx_hostname_resolve ON hostname_resolve (resolved_timestamp)
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
		  last_seen_active timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
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
		  last_seen_active timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		  PRIMARY KEY (hostname, token)
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
			PRIMARY KEY (recovery_id)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		DROP INDEX in_active_start_period_idx ON topology_recovery
	`,
	`
		CREATE INDEX in_active_start_period_idx_topology_recovery ON topology_recovery (in_active_period, start_active_period)
	`,
	`
		DROP INDEX start_active_period_idx ON topology_recovery
	`,
	`
		CREATE INDEX start_active_period_idx_topology_recovery ON topology_recovery (start_active_period)
	`,
	`
		DROP INDEX hostname_port_active_period_uidx ON topology_recovery
	`,
	`
		CREATE UNIQUE INDEX hostname_port_active_period_uidx_topology_recovery ON topology_recovery (hostname, port, in_active_period, end_active_period_unixtime)
	`,
	`
		CREATE TABLE IF NOT EXISTS hostname_unresolve (
		  hostname varchar(128) NOT NULL,
		  unresolved_hostname varchar(128) NOT NULL,
		  PRIMARY KEY (hostname)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		DROP INDEX unresolved_hostname_idx ON hostname_unresolve
	`,
	`
		CREATE INDEX unresolved_hostname_idx_hostname_unresolve ON hostname_unresolve (unresolved_hostname)
	`,
	`
		CREATE TABLE IF NOT EXISTS database_instance_pool (
			hostname varchar(128) CHARACTER SET ascii NOT NULL,
			port smallint(5) unsigned NOT NULL,
			pool varchar(128) NOT NULL,
			PRIMARY KEY (hostname, port, pool)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		DROP INDEX pool_idx ON database_instance_pool
	`,
	`
		CREATE INDEX pool_idx_database_instance_pool ON database_instance_pool (pool)
	`,
	`
		CREATE TABLE IF NOT EXISTS database_instance_topology_history (
			snapshot_unix_timestamp INT UNSIGNED NOT NULL,
			hostname varchar(128) CHARACTER SET ascii NOT NULL,
			port smallint(5) unsigned NOT NULL,
			master_host varchar(128) CHARACTER SET ascii NOT NULL,
			master_port smallint(5) unsigned NOT NULL,
			cluster_name tinytext CHARACTER SET ascii NOT NULL,
			PRIMARY KEY (snapshot_unix_timestamp, hostname, port)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		DROP INDEX cluster_name_idx ON database_instance_topology_history
	`,
	`
		CREATE INDEX cluster_name_idx_database_instance_topology_history ON database_instance_topology_history (snapshot_unix_timestamp, cluster_name(128))
	`,
	`
		CREATE TABLE IF NOT EXISTS candidate_database_instance (
			hostname varchar(128) CHARACTER SET ascii NOT NULL,
			port smallint(5) unsigned NOT NULL,
			last_suggested TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (hostname, port)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		DROP INDEX last_suggested_idx ON candidate_database_instance
	`,
	`
		CREATE INDEX last_suggested_idx_candidate_database_instance ON candidate_database_instance (last_suggested)
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
			PRIMARY KEY (detection_id)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		DROP INDEX hostname_port_active_period_uidx ON topology_failure_detection
	`,
	`
		CREATE UNIQUE INDEX hostname_port_active_period_uidx_topology_failure_detection ON topology_failure_detection (hostname, port, in_active_period, end_active_period_unixtime)
	`,
	`
		DROP INDEX in_active_start_period_idx ON topology_failure_detection
	`,
	`
		CREATE INDEX in_active_start_period_idx_topology_failure_detection ON topology_failure_detection (in_active_period, start_active_period)
	`,
	`
		CREATE TABLE IF NOT EXISTS hostname_resolve_history (
			resolved_hostname varchar(128) NOT NULL,
			hostname varchar(128) NOT NULL,
			resolved_timestamp timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (resolved_hostname)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		DROP INDEX hostname ON hostname_resolve_history
	`,
	`
		CREATE INDEX hostname_idx_hostname_resolve_history ON hostname_resolve_history (hostname)
	`,
	`
		DROP INDEX resolved_timestamp_idx ON hostname_resolve_history
	`,
	`
		CREATE INDEX resolved_timestamp_idx_hostname_resolve_history ON hostname_resolve_history (resolved_timestamp)
	`,
	`
		CREATE TABLE IF NOT EXISTS hostname_unresolve_history (
			unresolved_hostname varchar(128) NOT NULL,
			hostname varchar(128) NOT NULL,
			last_registered TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (unresolved_hostname)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		DROP INDEX hostname ON hostname_unresolve_history
	`,
	`
		CREATE INDEX hostname_idx_hostname_unresolve_history ON hostname_unresolve_history (hostname)
	`,
	`
		DROP INDEX last_registered_idx ON hostname_unresolve_history
	`,
	`
		CREATE INDEX last_registered_idx_hostname_unresolve_history ON hostname_unresolve_history (last_registered)
	`,
	`
		CREATE TABLE IF NOT EXISTS cluster_domain_name (
			cluster_name varchar(128) CHARACTER SET ascii NOT NULL,
			domain_name varchar(128) NOT NULL,
			PRIMARY KEY (cluster_name)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		DROP INDEX domain_name_idx ON cluster_domain_name
	`,
	`
		CREATE INDEX domain_name_idx_cluster_domain_name ON cluster_domain_name (domain_name(32))
	`,
	`
		CREATE TABLE IF NOT EXISTS master_position_equivalence (
			equivalence_id bigint unsigned not null auto_increment,
			master1_hostname varchar(128) CHARACTER SET ascii NOT NULL,
			master1_port smallint(5) unsigned NOT NULL,
			master1_binary_log_file varchar(128) CHARACTER SET ascii NOT NULL,
			master1_binary_log_pos bigint(20) unsigned NOT NULL,
			master2_hostname varchar(128) CHARACTER SET ascii NOT NULL,
			master2_port smallint(5) unsigned NOT NULL,
			master2_binary_log_file varchar(128) CHARACTER SET ascii NOT NULL,
			master2_binary_log_pos bigint(20) unsigned NOT NULL,
			last_suggested TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (equivalence_id)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		DROP INDEX equivalence_uidx ON master_position_equivalence
	`,
	`
		CREATE UNIQUE INDEX equivalence_uidx_master_position_equivalence ON master_position_equivalence (master1_hostname, master1_port, master1_binary_log_file, master1_binary_log_pos, master2_hostname, master2_port)
	`,
	`
		DROP INDEX master2_idx ON master_position_equivalence
	`,
	`
		CREATE INDEX master2_idx_master_position_equivalence ON master_position_equivalence (master2_hostname, master2_port, master2_binary_log_file, master2_binary_log_pos)
	`,
	`
		DROP INDEX last_suggested_idx ON master_position_equivalence
	`,
	`
		CREATE INDEX last_suggested_idx_master_position_equivalence ON master_position_equivalence (last_suggested)
	`,
	`
		CREATE TABLE IF NOT EXISTS async_request (
			request_id bigint unsigned NOT NULL AUTO_INCREMENT,
			command varchar(128) charset ascii not null,
			hostname varchar(128) NOT NULL,
			port smallint(5) unsigned NOT NULL,
			destination_hostname varchar(128) NOT NULL,
			destination_port smallint(5) unsigned NOT NULL,
			pattern text CHARACTER SET utf8 NOT NULL,
			gtid_hint varchar(32) charset ascii not null,
			begin_timestamp timestamp NULL DEFAULT NULL,
			end_timestamp timestamp NULL DEFAULT NULL,
			story text CHARACTER SET utf8 NOT NULL,
			PRIMARY KEY (request_id)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		DROP INDEX begin_timestamp_idx ON async_request
	`,
	`
		CREATE INDEX begin_timestamp_idx_async_request ON async_request (begin_timestamp)
	`,
	`
		DROP INDEX end_timestamp_idx ON async_request
	`,
	`
		CREATE INDEX end_timestamp_idx_async_request ON async_request (end_timestamp)
	`,
	`
		CREATE TABLE IF NOT EXISTS blocked_topology_recovery (
			hostname varchar(128) NOT NULL,
			port smallint(5) unsigned NOT NULL,
			cluster_name varchar(128) NOT NULL,
			analysis varchar(128) NOT NULL,
			last_blocked_timestamp timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			blocking_recovery_id bigint unsigned,
			PRIMARY KEY (hostname, port)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		DROP INDEX cluster_blocked_idx ON blocked_topology_recovery
	`,
	`
		CREATE INDEX cluster_blocked_idx_blocked_topology_recovery ON blocked_topology_recovery (cluster_name, last_blocked_timestamp)
	`,
	`
		CREATE TABLE IF NOT EXISTS database_instance_last_analysis (
		  hostname varchar(128) NOT NULL,
		  port smallint(5) unsigned NOT NULL,
		  analysis_timestamp timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		  analysis varchar(128) NOT NULL,
		  PRIMARY KEY (hostname, port)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		DROP INDEX analysis_timestamp_idx ON database_instance_last_analysis
	`,
	`
		CREATE INDEX analysis_timestamp_idx_database_instance_last_analysis ON database_instance_last_analysis (analysis_timestamp)
	`,
	`
		CREATE TABLE IF NOT EXISTS database_instance_analysis_changelog (
			changelog_id bigint unsigned not null auto_increment,
			hostname varchar(128) NOT NULL,
			port smallint(5) unsigned NOT NULL,
			analysis_timestamp timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			analysis varchar(128) NOT NULL,
			PRIMARY KEY (changelog_id)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		DROP INDEX analysis_timestamp_idx ON database_instance_analysis_changelog
	`,
	`
		CREATE INDEX analysis_timestamp_idx_database_instance_analysis_changelog ON database_instance_analysis_changelog (analysis_timestamp)
	`,
	`
		CREATE TABLE IF NOT EXISTS node_health_history (
			history_id bigint unsigned not null auto_increment,
			hostname varchar(128) CHARACTER SET ascii NOT NULL,
			token varchar(128) NOT NULL,
			first_seen_active timestamp NOT NULL,
			extra_info varchar(128) CHARACTER SET utf8 NOT NULL,
			PRIMARY KEY (history_id)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		DROP INDEX first_seen_active_idx ON node_health_history
	`,
	`
		CREATE INDEX first_seen_active_idx_node_health_history ON node_health_history (first_seen_active)
	`,
	`
		DROP INDEX hostname_token_idx ON node_health_history
	`,
	`
		CREATE UNIQUE INDEX hostname_token_idx_node_health_history ON node_health_history (hostname, token)
	`,
	`
		CREATE TABLE IF NOT EXISTS database_instance_coordinates_history (
			history_id bigint unsigned not null auto_increment,
			hostname varchar(128) NOT NULL,
			port smallint(5) unsigned NOT NULL,
			recorded_timestamp timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			binary_log_file varchar(128) NOT NULL,
			binary_log_pos bigint(20) unsigned NOT NULL,
			relay_log_file varchar(128) NOT NULL,
			relay_log_pos bigint(20) unsigned NOT NULL,
			PRIMARY KEY (history_id)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		DROP INDEX hostname_port_recorded_timestmp_idx ON database_instance_coordinates_history
	`,
	`
		CREATE INDEX hostname_port_recorded_idx_database_instance_coordinates_history ON database_instance_coordinates_history (hostname, port, recorded_timestamp)
	`,
	`
		DROP INDEX recorded_timestmp_idx ON database_instance_coordinates_history
	`,
	`
		CREATE INDEX recorded_timestmp_idx_database_instance_coordinates_history ON database_instance_coordinates_history (recorded_timestamp)
	`,
	`
		CREATE TABLE IF NOT EXISTS database_instance_binlog_files_history (
			history_id bigint unsigned not null auto_increment,
			hostname varchar(128) NOT NULL,
			port smallint(5) unsigned NOT NULL,
			binary_log_file varchar(128) NOT NULL,
			binary_log_pos bigint(20) unsigned NOT NULL,
			first_seen timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			last_seen timestamp NOT NULL DEFAULT '1971-01-01 00:00:00',
			PRIMARY KEY (history_id)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		DROP INDEX hostname_port_file_idx ON database_instance_binlog_files_history
	`,
	`
		CREATE UNIQUE INDEX hostname_port_file_idx_database_instance_binlog_files_history ON database_instance_binlog_files_history (hostname, port, binary_log_file)
	`,
	`
		DROP INDEX last_seen_idx ON database_instance_binlog_files_history
	`,
	`
		CREATE INDEX last_seen_idx_database_instance_binlog_files_history ON database_instance_binlog_files_history (last_seen)
	`,
	`
		CREATE TABLE IF NOT EXISTS access_token (
			access_token_id bigint unsigned not null auto_increment,
			public_token varchar(128) NOT NULL,
			secret_token varchar(128) NOT NULL,
			generated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			generated_by varchar(128) CHARACTER SET utf8 NOT NULL,
			is_acquired tinyint unsigned NOT NULL DEFAULT '0',
			PRIMARY KEY (access_token_id)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		DROP INDEX public_token_idx ON access_token
	`,
	`
		CREATE UNIQUE INDEX public_token_uidx_access_token ON access_token (public_token)
	`,
	`
		DROP INDEX generated_at_idx ON access_token
	`,
	`
		CREATE INDEX generated_at_idx_access_token ON access_token (generated_at)
	`,
	`
		CREATE TABLE IF NOT EXISTS database_instance_recent_relaylog_history (
			hostname varchar(128) NOT NULL,
			port smallint(5) unsigned NOT NULL,
			current_relay_log_file varchar(128) NOT NULL,
			current_relay_log_pos bigint(20) unsigned NOT NULL,
			current_seen timestamp NOT NULL DEFAULT '1971-01-01 00:00:00',
			prev_relay_log_file varchar(128) NOT NULL,
			prev_relay_log_pos bigint(20) unsigned NOT NULL,
			prev_seen timestamp NOT NULL DEFAULT '1971-01-01 00:00:00',
			PRIMARY KEY (hostname, port)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		DROP INDEX current_seen_idx ON database_instance_recent_relaylog_history
	`,
	`
		CREATE INDEX current_seen_idx_database_instance_recent_relaylog_history ON database_instance_recent_relaylog_history (current_seen)
	`,
	`
		CREATE TABLE IF NOT EXISTS orchestrator_metadata (
			anchor tinyint unsigned NOT NULL,
			last_deployed_version varchar(128) CHARACTER SET ascii NOT NULL,
			last_deployed_timestamp timestamp NOT NULL,
			PRIMARY KEY (anchor)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		CREATE TABLE IF NOT EXISTS orchestrator_db_deployments (
			deployed_version varchar(128) CHARACTER SET ascii NOT NULL,
			deployed_timestamp timestamp NOT NULL,
			PRIMARY KEY (deployed_version)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		CREATE TABLE IF NOT EXISTS global_recovery_disable (
			disable_recovery tinyint unsigned NOT NULL COMMENT 'Insert 1 to disable recovery globally',
			PRIMARY KEY (disable_recovery)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		CREATE TABLE IF NOT EXISTS cluster_alias_override (
			cluster_name varchar(128) CHARACTER SET ascii NOT NULL,
			alias varchar(128) NOT NULL,
			PRIMARY KEY (cluster_name)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		CREATE TABLE IF NOT EXISTS topology_recovery_steps (
			recovery_step_id bigint unsigned not null auto_increment,
			recovery_uid varchar(128) CHARACTER SET ascii NOT NULL,
			audit_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			message text CHARACTER SET utf8 NOT NULL,
			PRIMARY KEY (recovery_step_id)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
}

// generateSQLPatches contains DDLs for patching schema to the latest version.
// Add new statements at the end of the list so they form a changelog.
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
		DROP INDEX master_host_port_idx ON database_instance
	`,
	`
		ALTER TABLE
			database_instance
			ADD INDEX master_host_port_idx_database_instance (master_host, master_port)
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
			ADD COLUMN analysis              varchar(128) CHARACTER SET ascii NOT NULL
	`,
	`
		ALTER TABLE
			topology_recovery
			ADD COLUMN cluster_name          varchar(128) CHARACTER SET ascii NOT NULL
	`,
	`
		ALTER TABLE
			topology_recovery
			ADD COLUMN cluster_alias         varchar(128) CHARACTER SET ascii NOT NULL
	`,
	`
		ALTER TABLE
			topology_recovery
			ADD COLUMN count_affected_slaves int unsigned NOT NULL
	`,
	`
		ALTER TABLE
			topology_recovery
			ADD COLUMN slave_hosts text CHARACTER SET ascii NOT NULL
	`,
	`
		ALTER TABLE hostname_unresolve
			ADD COLUMN last_registered TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	`,
	`
		ALTER TABLE hostname_unresolve
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
	`
		ALTER TABLE
			database_instance
			ADD COLUMN binlog_server TINYINT UNSIGNED NOT NULL AFTER version
	`,
	`
		ALTER TABLE cluster_domain_name
			ADD COLUMN last_registered TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	`,
	`
		ALTER TABLE cluster_domain_name
			ADD KEY last_registered_idx (last_registered)
	`,
	`
		ALTER TABLE
			database_instance
			ADD COLUMN supports_oracle_gtid TINYINT UNSIGNED NOT NULL AFTER oracle_gtid
	`,
	`
		ALTER TABLE
			database_instance
			ADD COLUMN executed_gtid_set text CHARACTER SET ascii NOT NULL AFTER oracle_gtid
	`,
	`
		ALTER TABLE
			database_instance
			ADD COLUMN server_uuid varchar(64) CHARACTER SET ascii NOT NULL AFTER server_id
	`,
	`
		ALTER TABLE
			database_instance
			ADD COLUMN suggested_cluster_alias varchar(128) CHARACTER SET ascii NOT NULL AFTER cluster_name
	`,
	`
		ALTER TABLE cluster_alias
			ADD COLUMN last_registered TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	`,
	`
		ALTER TABLE cluster_alias
			ADD KEY last_registered_idx (last_registered)
	`,
	`
		ALTER TABLE
			topology_recovery
			ADD COLUMN is_successful TINYINT UNSIGNED NOT NULL DEFAULT 0 AFTER processcing_node_token
	`,
	`
		ALTER TABLE
			topology_recovery
			ADD COLUMN acknowledged TINYINT UNSIGNED NOT NULL DEFAULT 0
	`,
	`
		ALTER TABLE
			topology_recovery
			ADD COLUMN acknowledged_by varchar(128) CHARACTER SET utf8 NOT NULL
	`,
	`
		ALTER TABLE
			topology_recovery
			ADD COLUMN acknowledge_comment text CHARACTER SET utf8 NOT NULL
	`,
	`
		ALTER TABLE
			topology_recovery
			ADD COLUMN participating_instances text CHARACTER SET ascii NOT NULL after slave_hosts
	`,
	`
		ALTER TABLE
			topology_recovery
			ADD COLUMN lost_slaves text CHARACTER SET ascii NOT NULL after participating_instances
	`,
	`
		ALTER TABLE
			topology_recovery
			ADD COLUMN all_errors text CHARACTER SET ascii NOT NULL after lost_slaves
	`,
	`
		ALTER TABLE audit
			ADD COLUMN cluster_name varchar(128) CHARACTER SET ascii NOT NULL DEFAULT '' AFTER port
	`,
	`
		ALTER TABLE candidate_database_instance
			ADD COLUMN priority TINYINT SIGNED NOT NULL DEFAULT 1 comment 'positive promote, nagative unpromotes'
	`,
	`
		ALTER TABLE
			topology_recovery
			ADD COLUMN acknowledged_at TIMESTAMP NULL after acknowledged
	`,
	`
		ALTER TABLE
			topology_recovery
			ADD KEY acknowledged_idx (acknowledged, acknowledged_at)
	`,
	`
		ALTER TABLE
			blocked_topology_recovery
			ADD KEY last_blocked_idx (last_blocked_timestamp)
	`,
	`
		ALTER TABLE candidate_database_instance
			ADD COLUMN promotion_rule enum('must', 'prefer', 'neutral', 'prefer_not', 'must_not') NOT NULL DEFAULT 'neutral'
	`,
	`
		ALTER TABLE node_health /* sqlite3-skip */
			DROP PRIMARY KEY,
			ADD PRIMARY KEY (hostname, token)
	`,
	`
		ALTER TABLE node_health
			ADD COLUMN extra_info varchar(128) CHARACTER SET utf8 NOT NULL
	`,
	`
		ALTER TABLE agent_seed /* sqlite3-skip */
			MODIFY end_timestamp timestamp NOT NULL DEFAULT '1971-01-01 00:00:00'
	`,
	`
		ALTER TABLE active_node /* sqlite3-skip */
			MODIFY last_seen_active timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
	`,

	`
		ALTER TABLE node_health /* sqlite3-skip */
			MODIFY last_seen_active timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
	`,
	`
		ALTER TABLE candidate_database_instance /* sqlite3-skip */
			MODIFY last_suggested timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
	`,
	`
		ALTER TABLE master_position_equivalence /* sqlite3-skip */
			MODIFY last_suggested timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
	`,
	`
		ALTER TABLE
			database_instance
			ADD COLUMN last_attempted_check TIMESTAMP NOT NULL DEFAULT '1971-01-01 00:00:00' AFTER last_checked
	`,
	`
		ALTER TABLE
			database_instance /* sqlite3-skip */
			MODIFY last_attempted_check TIMESTAMP NOT NULL DEFAULT '1971-01-01 00:00:00'
	`,
	`
		ALTER TABLE
			database_instance_analysis_changelog
			ADD KEY instance_timestamp_idx (hostname, port, analysis_timestamp)
	`,
	`
		ALTER TABLE
			topology_recovery
			ADD COLUMN last_detection_id bigint unsigned NOT NULL
	`,
	`
		ALTER TABLE
			topology_recovery
			ADD KEY last_detection_idx (last_detection_id)
	`,
	`
		ALTER TABLE node_health_history
			ADD COLUMN command varchar(128) CHARACTER SET utf8 NOT NULL
	`,
	`
		ALTER TABLE node_health
			ADD COLUMN command varchar(128) CHARACTER SET utf8 NOT NULL
	`,
	`
		ALTER TABLE database_instance_topology_history
			ADD COLUMN version varchar(128) CHARACTER SET ascii NOT NULL
	`,
	`
		ALTER TABLE
			database_instance
			ADD COLUMN gtid_purged text CHARACTER SET ascii NOT NULL AFTER executed_gtid_set
	`,
	`
		ALTER TABLE
			database_instance_coordinates_history
			ADD COLUMN last_seen timestamp NOT NULL DEFAULT '1971-01-01 00:00:00' AFTER recorded_timestamp
	`,
	`
		ALTER TABLE
			access_token
			ADD COLUMN is_reentrant TINYINT UNSIGNED NOT NULL default 0
	`,
	`
		ALTER TABLE
			access_token
			ADD COLUMN acquired_at timestamp NOT NULL DEFAULT '1971-01-01 00:00:00'
	`,
	`
		ALTER TABLE
			database_instance_pool
			ADD COLUMN registered_at timestamp NOT NULL DEFAULT '1971-01-01 00:00:00'
	`,
	`
		ALTER TABLE
			database_instance
			ADD COLUMN replication_credentials_available TINYINT UNSIGNED NOT NULL
	`,
	`
		ALTER TABLE
			database_instance
			ADD COLUMN has_replication_credentials TINYINT UNSIGNED NOT NULL
	`,
	`
		ALTER TABLE
			database_instance
			ADD COLUMN allow_tls TINYINT UNSIGNED NOT NULL AFTER sql_delay
	`,
	`
		ALTER TABLE
			database_instance
			ADD COLUMN semi_sync_enforced TINYINT UNSIGNED NOT NULL AFTER physical_environment
	`,
	`
		ALTER TABLE
			database_instance
			ADD COLUMN instance_alias varchar(128) CHARACTER SET ascii NOT NULL AFTER physical_environment
	`,
	`
		ALTER TABLE
			topology_recovery
			ADD COLUMN successor_alias varchar(128) DEFAULT NULL
	`,
	`
		ALTER TABLE
			database_instance /* sqlite3-skip */
			MODIFY cluster_name varchar(128) NOT NULL
	`,
	`
		ALTER TABLE
			node_health
			ADD INDEX last_seen_active_idx (last_seen_active)
	`,
	`
		ALTER TABLE
			database_instance_maintenance
			ADD COLUMN processing_node_hostname varchar(128) CHARACTER SET ascii NOT NULL
	`,
	`
		ALTER TABLE
			database_instance_maintenance
			ADD COLUMN processing_node_token varchar(128) NOT NULL
	`,
	`
		ALTER TABLE
			database_instance_maintenance
			ADD COLUMN explicitly_bounded TINYINT UNSIGNED NOT NULL
	`,
	`
		ALTER TABLE node_health_history
			ADD COLUMN app_version varchar(64) CHARACTER SET ascii NOT NULL DEFAULT ""
	`,
	`
		ALTER TABLE node_health
			ADD COLUMN app_version varchar(64) CHARACTER SET ascii NOT NULL DEFAULT ""
	`,
	`
		ALTER TABLE node_health_history /* sqlite3-skip */
			MODIFY app_version varchar(64) CHARACTER SET ascii NOT NULL DEFAULT ""
	`,
	`
		ALTER TABLE node_health /* sqlite3-skip */
			MODIFY app_version varchar(64) CHARACTER SET ascii NOT NULL DEFAULT ""
	`,
	`
		ALTER TABLE
			database_instance
			ADD COLUMN version_comment varchar(128) NOT NULL DEFAULT ''
	`,
	`
		ALTER TABLE active_node
			ADD COLUMN first_seen_active timestamp NOT NULL DEFAULT '1971-01-01 00:00:00'
	`,
	`
		ALTER TABLE node_health
			ADD COLUMN first_seen_active timestamp NOT NULL DEFAULT '1971-01-01 00:00:00'
	`,
	`
		ALTER TABLE database_instance
			ADD COLUMN major_version varchar(16) CHARACTER SET ascii NOT NULL
	`,
	`
		ALTER TABLE
			database_instance
			ADD COLUMN binlog_row_image varchar(16) CHARACTER SET ascii NOT NULL
	`,
	`
		ALTER TABLE topology_recovery
			ADD COLUMN uid varchar(128) CHARACTER SET ascii NOT NULL
	`,
	`
		CREATE INDEX uid_idx_topology_recovery ON topology_recovery(uid)
	`,
	`
		CREATE INDEX recovery_uid_idx_topology_recovery_steps ON topology_recovery_steps(recovery_uid)
	`,
	`
		ALTER TABLE
			database_instance
			ADD COLUMN last_discovery_latency bigint not null
	`,
	`
		CREATE INDEX end_timestamp_idx_database_instance_downtime ON database_instance_downtime(end_timestamp)
	`,
	`
		CREATE INDEX suggested_cluster_alias_idx_database_instance ON database_instance(suggested_cluster_alias)
	`,
}

// Track if a TLS has already been configured for topology
var topologyTLSConfigured bool = false

// Track if a TLS has already been configured for Orchestrator
var orchestratorTLSConfigured bool = false

// OpenDiscovery returns a DB instance to access a topology instance.
// It has lower read timeout than OpenTopology and is intended to
// be used with low-latency discovery queries.
func OpenDiscovery(host string, port int) (*sql.DB, error) {
	return openTopology(host, port, config.Config.MySQLDiscoveryReadTimeoutSeconds)
}

// OpenTopology returns a DB instance to access a topology instance.
func OpenTopology(host string, port int) (*sql.DB, error) {
	return openTopology(host, port, config.Config.MySQLTopologyReadTimeoutSeconds)
}

func openTopology(host string, port int, readTimeout int) (*sql.DB, error) {
	mysql_uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/?timeout=%ds&readTimeout=%ds",
		config.Config.MySQLTopologyUser,
		config.Config.MySQLTopologyPassword,
		host, port,
		config.Config.MySQLConnectTimeoutSeconds,
		readTimeout,
	)
	if config.Config.MySQLTopologyUseMutualTLS {
		mysql_uri, _ = SetupMySQLTopologyTLS(mysql_uri)
	}
	db, _, err := sqlutils.GetDB(mysql_uri)
	db.SetMaxOpenConns(config.Config.MySQLTopologyMaxPoolConnections)
	db.SetMaxIdleConns(config.Config.MySQLTopologyMaxPoolConnections)
	return db, err
}

// Create a TLS configuration from the config supplied CA, Certificate, and Private key.
// Register the TLS config with the mysql drivers as the "topology" config
// Modify the supplied URI to call the TLS config
// TODO: Way to have password mixed with TLS for various nodes in the topology.  Currently everything is TLS or everything is password
func SetupMySQLTopologyTLS(uri string) (string, error) {
	if !topologyTLSConfigured {
		tlsConfig, err := ssl.NewTLSConfig(config.Config.MySQLTopologySSLCAFile, !config.Config.MySQLTopologySSLSkipVerify)
		// Drop to TLS 1.0 for talking to MySQL
		tlsConfig.MinVersion = tls.VersionTLS10
		if err != nil {
			return "", log.Fatalf("Can't create TLS configuration for Topology connection %s: %s", uri, err)
		}
		tlsConfig.InsecureSkipVerify = config.Config.MySQLTopologySSLSkipVerify
		if err = ssl.AppendKeyPair(tlsConfig, config.Config.MySQLTopologySSLCertFile, config.Config.MySQLTopologySSLPrivateKeyFile); err != nil {
			return "", log.Fatalf("Can't setup TLS key pairs for %s: %s", uri, err)
		}
		if err = mysql.RegisterTLSConfig("topology", tlsConfig); err != nil {
			return "", log.Fatalf("Can't register mysql TLS config for topology: %s", err)
		}
		topologyTLSConfigured = true
	}
	return fmt.Sprintf("%s&tls=topology", uri), nil
}

// OpenTopology returns the DB instance for the orchestrator backed database
func OpenOrchestrator() (db *sql.DB, err error) {
	if config.Config.DatabaselessMode__experimental {
		return nil, nil
	}
	var fromCache bool
	if config.Config.IsSQLite() {
		db, fromCache, err = sqlutils.GetSQLiteDB(config.Config.SQLite3DataFile)
		if err == nil && !fromCache {
			log.Debugf("Connected to orchestrator backend: sqlite on %v", config.Config.SQLite3DataFile)
		}
	} else {
		mysql_uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=%ds&readTimeout=%ds&interpolateParams=%t",
			config.Config.MySQLOrchestratorUser,
			config.Config.MySQLOrchestratorPassword,
			config.Config.MySQLOrchestratorHost,
			config.Config.MySQLOrchestratorPort,
			config.Config.MySQLOrchestratorDatabase,
			config.Config.MySQLConnectTimeoutSeconds,
			config.Config.MySQLOrchestratorReadTimeoutSeconds,
			config.Config.MySQLInterpolateParams,
		)
		if config.Config.MySQLOrchestratorUseMutualTLS {
			mysql_uri, _ = SetupMySQLOrchestratorTLS(mysql_uri)
		}
		db, fromCache, err = sqlutils.GetDB(mysql_uri)
		if err == nil && !fromCache {
			// do not show the password but do show what we connect to.
			safe_mysql_uri := fmt.Sprintf("%s:?@tcp(%s:%d)/%s?timeout=%ds", config.Config.MySQLOrchestratorUser,
				config.Config.MySQLOrchestratorHost, config.Config.MySQLOrchestratorPort, config.Config.MySQLOrchestratorDatabase, config.Config.MySQLConnectTimeoutSeconds)
			log.Debugf("Connected to orchestrator backend: %v", safe_mysql_uri)
			if config.Config.MySQLOrchestratorMaxPoolConnections > 0 {
				log.Debugf("Orchestrator pool SetMaxOpenConns: %d", config.Config.MySQLOrchestratorMaxPoolConnections)
				db.SetMaxOpenConns(config.Config.MySQLOrchestratorMaxPoolConnections)
			}
		}
	}
	if err == nil && !fromCache {
		if !config.Config.SkipOrchestratorDatabaseUpdate {
			initOrchestratorDB(db)
		}
		db.SetMaxIdleConns(10)
	}
	return db, err
}

func translateStatement(statement string) (string, error) {
	if config.Config.IsSQLite() {
		statement = sqlutils.ToSqlite3Dialect(statement)
	}
	return statement, nil
}

// versionIsDeployed checks if given version has already been deployed
func versionIsDeployed(db *sql.DB) (result bool, err error) {
	query := `
		select
			count(*) as is_deployed
		from
			orchestrator_db_deployments
		where
			deployed_version = ?
		`
	err = db.QueryRow(query, config.RuntimeCLIFlags.ConfiguredVersion).Scan(&result)
	// err means the table 'orchestrator_db_deployments' does not even exist, in which case we proceed
	// to deploy.
	// If there's another error to this, like DB gone bad, then we're about to find out anyway.
	return result, err
}

// registerOrchestratorDeployment updates the orchestrator_metadata table upon successful deployment
func registerOrchestratorDeployment(db *sql.DB) error {
	query := `
    	replace into orchestrator_db_deployments (
				deployed_version, deployed_timestamp
			) values (
				?, NOW()
			)
				`
	if _, err := execInternal(db, query, config.RuntimeCLIFlags.ConfiguredVersion); err != nil {
		log.Fatalf("Unable to write to orchestrator_metadata: %+v", err)
	}
	log.Debugf("Migrated database schema to version [%+v]", config.RuntimeCLIFlags.ConfiguredVersion)
	return nil
}

// Create a TLS configuration from the config supplied CA, Certificate, and Private key.
// Register the TLS config with the mysql drivers as the "orchestrator" config
// Modify the supplied URI to call the TLS config
func SetupMySQLOrchestratorTLS(uri string) (string, error) {
	if !orchestratorTLSConfigured {
		tlsConfig, err := ssl.NewTLSConfig(config.Config.MySQLOrchestratorSSLCAFile, true)
		// Drop to TLS 1.0 for talking to MySQL
		tlsConfig.MinVersion = tls.VersionTLS10
		if err != nil {
			return "", log.Fatalf("Can't create TLS configuration for Orchestrator connection %s: %s", uri, err)
		}
		tlsConfig.InsecureSkipVerify = config.Config.MySQLOrchestratorSSLSkipVerify
		if err = ssl.AppendKeyPair(tlsConfig, config.Config.MySQLOrchestratorSSLCertFile, config.Config.MySQLOrchestratorSSLPrivateKeyFile); err != nil {
			return "", log.Fatalf("Can't setup TLS key pairs for %s: %s", uri, err)
		}
		if err = mysql.RegisterTLSConfig("orchestrator", tlsConfig); err != nil {
			return "", log.Fatalf("Can't register mysql TLS config for orchestrator: %s", err)
		}
		orchestratorTLSConfigured = true
	}
	return fmt.Sprintf("%s&tls=orchestrator", uri), nil
}

// deployStatements will issue given sql queries that are not already known to be deployed.
// This iterates both lists (to-run and already-deployed) and also verifies no contraditions.
func deployStatements(db *sql.DB, queries []string) error {
	tx, err := db.Begin()
	if err != nil {
		log.Fatale(err)
	}
	// Ugly workaround ahead.
	// Origin of this workaround is the existence of some "timestamp NOT NULL," column definitions,
	// where in NO_ZERO_IN_DATE,NO_ZERO_DATE sql_mode are invalid (since default is implicitly "0")
	// This means installation of orchestrator fails on such configured servers, and in particular on 5.7
	// where this setting is the dfault.
	// For purpose of backwards compatability, what we do is force sql_mode to be more relaxed, create the schemas
	// along with the "invalid" definition, and then go ahead and fix those definitions via following ALTER statements.
	// My bad.
	originalSqlMode := ""
	if config.Config.IsMySQL() {
		err = tx.QueryRow(`select @@session.sql_mode`).Scan(&originalSqlMode)
		if _, err := tx.Exec(`set @@session.sql_mode=REPLACE(@@session.sql_mode, 'NO_ZERO_DATE', '')`); err != nil {
			log.Fatale(err)
		}
		if _, err := tx.Exec(`set @@session.sql_mode=REPLACE(@@session.sql_mode, 'NO_ZERO_IN_DATE', '')`); err != nil {
			log.Fatale(err)
		}
	}
	for i, query := range queries {
		if i == 0 {
			//log.Debugf("sql_mode is: %+v", originalSqlMode)
		}

		query, err := translateStatement(query)
		if err != nil {
			return log.Fatalf("Cannot initiate orchestrator: %+v; query=%+v", err, query)
		}
		if _, err := tx.Exec(query); err != nil {
			if strings.Contains(err.Error(), "syntax error") {
				return log.Fatalf("Cannot initiate orchestrator: %+v; query=%+v", err, query)
			}
			if !sqlutils.IsAlterTable(query) && !sqlutils.IsCreateIndex(query) && !sqlutils.IsDropIndex(query) {
				return log.Fatalf("Cannot initiate orchestrator: %+v; query=%+v", err, query)
			}
			if !strings.Contains(err.Error(), "duplicate column name") &&
				!strings.Contains(err.Error(), "Duplicate column name") &&
				!strings.Contains(err.Error(), "check that column/key exists") &&
				!strings.Contains(err.Error(), "already exists") &&
				!strings.Contains(err.Error(), "Duplicate key name") {
				log.Errorf("Error initiating orchestrator: %+v; query=%+v", err, query)
			}
		}
	}
	if config.Config.IsMySQL() {
		if _, err := tx.Exec(`set session sql_mode=?`, originalSqlMode); err != nil {
			log.Fatale(err)
		}
	}
	if err := tx.Commit(); err != nil {
		log.Fatale(err)
	}
	return nil
}

// initOrchestratorDB attempts to create/upgrade the orchestrator backend database. It is created once in the
// application's lifetime.
func initOrchestratorDB(db *sql.DB) error {
	log.Debug("Initializing orchestrator")

	versionAlreadyDeployed, err := versionIsDeployed(db)
	if versionAlreadyDeployed && config.RuntimeCLIFlags.ConfiguredVersion != "" && err == nil {
		// Already deployed with this version
		return nil
	}
	if config.Config.PanicIfDifferentDatabaseDeploy && config.RuntimeCLIFlags.ConfiguredVersion != "" && !versionAlreadyDeployed {
		log.Fatalf("PanicIfDifferentDatabaseDeploy is set. Configured version %s is not the version found in the database", config.RuntimeCLIFlags.ConfiguredVersion)
	}
	log.Debugf("Migrating database schema")
	deployStatements(db, generateSQLBase)
	deployStatements(db, generateSQLPatches)
	registerOrchestratorDeployment(db)

	if config.Config.IsSQLite() {
		ExecOrchestrator(`PRAGMA journal_mode = WAL`)
		ExecOrchestrator(`PRAGMA synchronous = NORMAL`)
	}

	return nil
}

// execInternalSilently
func execInternalSilently(db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	var err error
	query, err = translateStatement(query)
	if err != nil {
		return nil, err
	}
	res, err := sqlutils.ExecSilently(db, query, args...)
	return res, err
}

// execInternal
func execInternal(db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	var err error
	query, err = translateStatement(query)
	if err != nil {
		return nil, err
	}
	res, err := sqlutils.ExecSilently(db, query, args...)
	return res, err
}

// ExecOrchestrator will execute given query on the orchestrator backend database.
func ExecOrchestrator(query string, args ...interface{}) (sql.Result, error) {
	if config.Config.DatabaselessMode__experimental {
		return DummySqlResult{}, nil
	}
	var err error
	query, err = translateStatement(query)
	if err != nil {
		return nil, err
	}
	db, err := OpenOrchestrator()
	if err != nil {
		return nil, err
	}
	dbexec := sqlutils.Exec
	if config.Config.MySQLInterpolateParams {
		dbexec = sqlutils.ExecNoPrepare
	}
	res, err := dbexec(db, query, args...)
	return res, err
}

// QueryRowsMapOrchestrator
func QueryOrchestratorRowsMap(query string, on_row func(sqlutils.RowMap) error) error {
	if config.Config.DatabaselessMode__experimental {
		return nil
	}
	query, err := translateStatement(query)
	if err != nil {
		return log.Fatalf("Cannot query orchestrator: %+v; query=%+v", err, query)
	}
	db, err := OpenOrchestrator()
	if err != nil {
		return err
	}

	return sqlutils.QueryRowsMap(db, query, on_row)
}

// QueryOrchestrator
func QueryOrchestrator(query string, argsArray []interface{}, on_row func(sqlutils.RowMap) error) error {
	if config.Config.DatabaselessMode__experimental {
		return nil
	}
	query, err := translateStatement(query)
	if err != nil {
		return log.Fatalf("Cannot query orchestrator: %+v; query=%+v", err, query)
	}
	db, err := OpenOrchestrator()
	if err != nil {
		return err
	}

	return log.Criticale(sqlutils.QueryRowsMap(db, query, on_row, argsArray...))
}

// QueryOrchestratorRowsMapBuffered
func QueryOrchestratorRowsMapBuffered(query string, on_row func(sqlutils.RowMap) error) error {
	if config.Config.DatabaselessMode__experimental {
		return nil
	}
	query, err := translateStatement(query)
	if err != nil {
		return log.Fatalf("Cannot query orchestrator: %+v; query=%+v", err, query)
	}
	db, err := OpenOrchestrator()
	if err != nil {
		return err
	}

	return sqlutils.QueryRowsMapBuffered(db, query, on_row)
}

// QueryOrchestratorBuffered
func QueryOrchestratorBuffered(query string, argsArray []interface{}, on_row func(sqlutils.RowMap) error) error {
	if config.Config.DatabaselessMode__experimental {
		return nil
	}
	query, err := translateStatement(query)
	if err != nil {
		return log.Fatalf("Cannot query orchestrator: %+v; query=%+v", err, query)
	}
	db, err := OpenOrchestrator()
	if err != nil {
		return err
	}

	if argsArray == nil {
		argsArray = EmptyArgs
	}
	return log.Criticale(sqlutils.QueryRowsMapBuffered(db, query, on_row, argsArray...))
}
