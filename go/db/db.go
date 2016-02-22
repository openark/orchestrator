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

	"github.com/go-sql-driver/mysql"
	"github.com/outbrain/golib/log"
	"github.com/outbrain/golib/sqlutils"
	"github.com/outbrain/orchestrator/go/config"
	"github.com/outbrain/orchestrator/go/ssl"
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

var internalDBDeploymentSQL = []string{
	`
		CREATE TABLE IF NOT EXISTS _orchestrator_db_deployment (
		  deployment_id int unsigned NOT NULL AUTO_INCREMENT,
		  deployment_type enum('base', 'patch'),
		  deploy_timestamp timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		  sql_statement TEXT CHARSET ascii NOT NULL,
		  statement_digest VARCHAR(128) CHARSET ascii NOT NULL,
		  statement_index INT UNSIGNED NOT NULL,
		  PRIMARY KEY (deployment_id),
		  UNIQUE KEY sql_index_uidx (statement_digest, statement_index)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
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
	`
		CREATE TABLE IF NOT EXISTS cluster_domain_name (
		  cluster_name varchar(128) CHARACTER SET ascii NOT NULL,
		  domain_name varchar(128) NOT NULL,
		  PRIMARY KEY (cluster_name),
		  KEY domain_name_idx(domain_name(32))
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
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
          last_suggested TIMESTAMP NOT NULL,
          PRIMARY KEY (equivalence_id),
          UNIQUE KEY equivalence_uidx (master1_hostname, master1_port, master1_binary_log_file, master1_binary_log_pos, master2_hostname, master2_port),
          KEY master2_idx (master2_hostname, master2_port, master2_binary_log_file, master2_binary_log_pos),
          KEY last_suggested_idx(last_suggested)
        ) ENGINE=InnoDB DEFAULT CHARSET=ascii
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
		  PRIMARY KEY (request_id),
		  KEY begin_timestamp_idx (begin_timestamp),
		  KEY end_timestamp_idx (end_timestamp)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		CREATE TABLE IF NOT EXISTS blocked_topology_recovery (
		  hostname varchar(128) NOT NULL,
		  port smallint(5) unsigned NOT NULL,
		  cluster_name varchar(128) NOT NULL,
		  analysis varchar(128) NOT NULL,
		  last_blocked_timestamp timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		  blocking_recovery_id bigint unsigned,
		  PRIMARY KEY (hostname, port),
		  KEY cluster_blocked_idx (cluster_name, last_blocked_timestamp)
		) ENGINE=InnoDB CHARSET=ascii
	`,
	`
		CREATE TABLE IF NOT EXISTS database_instance_last_analysis (
		  hostname varchar(128) NOT NULL,
		  port smallint(5) unsigned NOT NULL,
		  analysis_timestamp timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		  analysis varchar(128) NOT NULL,
		  PRIMARY KEY (hostname, port),
		  KEY analysis_timestamp_idx(analysis_timestamp)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		CREATE TABLE IF NOT EXISTS database_instance_analysis_changelog (
          changelog_id bigint unsigned not null auto_increment,
		  hostname varchar(128) NOT NULL,
		  port smallint(5) unsigned NOT NULL,
		  analysis_timestamp timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
		  analysis varchar(128) NOT NULL,
		  PRIMARY KEY (changelog_id),
		  KEY analysis_timestamp_idx(analysis_timestamp)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		CREATE TABLE IF NOT EXISTS node_health_history (
		  history_id bigint unsigned not null auto_increment,
		  hostname varchar(128) CHARACTER SET ascii NOT NULL,
		  token varchar(128) NOT NULL,
		  first_seen_active timestamp NOT NULL,
		  extra_info varchar(128) CHARACTER SET utf8 NOT NULL,
		  PRIMARY KEY (history_id),
		  UNIQUE KEY hostname_token_idx(hostname, token),
		  KEY first_seen_active_idx(first_seen_active)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
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
		  PRIMARY KEY (history_id),
		  KEY hostname_port_recorded_timestmp_idx (hostname, port, recorded_timestamp),
		  KEY recorded_timestmp_idx (recorded_timestamp)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
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
		  PRIMARY KEY (history_id),
		  UNIQUE KEY hostname_port_file_idx (hostname, port, binary_log_file),
		  KEY last_seen_idx (last_seen)
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
	  CREATE TABLE IF NOT EXISTS access_token (
	    access_token_id bigint unsigned not null auto_increment,
	    public_token varchar(128) NOT NULL,
	    secret_token varchar(128) NOT NULL,
	    generated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	    generated_by varchar(128) CHARACTER SET utf8 NOT NULL,
			is_acquired tinyint unsigned NOT NULL DEFAULT '0',
	    PRIMARY KEY (access_token_id),
	    UNIQUE KEY public_token_idx (public_token),
	    KEY generated_at_idx (generated_at)
	  ) ENGINE=InnoDB DEFAULT CHARSET=ascii
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
			PRIMARY KEY (hostname, port),
		  KEY current_seen_idx (current_seen)
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
	`
		ALTER TABLE
			database_instance
			ADD COLUMN binlog_server TINYINT UNSIGNED NOT NULL AFTER version
	`,
	`
		ALTER TABLE cluster_domain_name
			ADD COLUMN last_registered TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
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
			ADD COLUMN last_registered TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
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
			ADD COLUMN acknowledged TINYINT UNSIGNED NOT NULL DEFAULT 0,
			ADD COLUMN acknowledged_by varchar(128) CHARACTER SET utf8 NOT NULL,
			ADD COLUMN acknowledge_comment text CHARACTER SET utf8 NOT NULL
	`,
	`
		ALTER TABLE
			topology_recovery
			ADD COLUMN participating_instances text CHARACTER SET ascii NOT NULL after slave_hosts,
			ADD COLUMN lost_slaves text CHARACTER SET ascii NOT NULL after participating_instances,
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
			ADD COLUMN acknowledged_at TIMESTAMP NULL after acknowledged,
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
		ALTER TABLE node_health
			DROP PRIMARY KEY,
			ADD PRIMARY KEY (hostname, token)
	`,
	`
		ALTER TABLE node_health
			ADD COLUMN extra_info varchar(128) CHARACTER SET utf8 NOT NULL
	`,
	`
		ALTER TABLE agent_seed
			MODIFY end_timestamp timestamp NOT NULL DEFAULT '1971-01-01 00:00:00'
	`,
	`
		ALTER TABLE active_node
			MODIFY last_seen_active timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
	`,

	`
		ALTER TABLE node_health
			MODIFY last_seen_active timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
	`,
	`
		ALTER TABLE candidate_database_instance
			MODIFY last_suggested timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
	`,
	`
		ALTER TABLE master_position_equivalence
			MODIFY last_suggested timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
	`,
	`
		ALTER TABLE
			database_instance
			ADD COLUMN last_attempted_check TIMESTAMP NOT NULL DEFAULT '1971-01-01 00:00:00' AFTER last_checked
	`,
	`
		ALTER TABLE
			database_instance
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
			ADD COLUMN last_detection_id bigint unsigned NOT NULL,
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
}

// Track if a TLS has already been configured for topology
var topologyTLSConfigured bool = false

// Track if a TLS has already been configured for Orchestrator
var orchestratorTLSConfigured bool = false

// OpenTopology returns a DB instance to access a topology instance
func OpenTopology(host string, port int) (*sql.DB, error) {
	mysql_uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/?timeout=%ds", config.Config.MySQLTopologyUser, config.Config.MySQLTopologyPassword, host, port, config.Config.MySQLConnectTimeoutSeconds)
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
func OpenOrchestrator() (*sql.DB, error) {
	if config.Config.DatabaselessMode__experimental {
		return nil, nil
	}
	mysql_uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=%ds", config.Config.MySQLOrchestratorUser, config.Config.MySQLOrchestratorPassword,
		config.Config.MySQLOrchestratorHost, config.Config.MySQLOrchestratorPort, config.Config.MySQLOrchestratorDatabase, config.Config.MySQLConnectTimeoutSeconds)
	if config.Config.MySQLOrchestratorUseMutualTLS {
		mysql_uri, _ = SetupMySQLOrchestratorTLS(mysql_uri)
	}
	db, fromCache, err := sqlutils.GetDB(mysql_uri)
	if err == nil && !fromCache {
		if !config.Config.SkipOrchestratorDatabaseUpdate {
			initOrchestratorDB(db)
		}
		db.SetMaxIdleConns(10)
	}
	return db, err
}

// readInternalDeployments reads orchestrator db deployment statements that are known to have been executed
func readInternalDeployments() (baseDeployments []string, patchDeployments []string, err error) {
	if !config.Config.SmartOrchestratorDatabaseUpdate {
		return baseDeployments, patchDeployments, nil
	}
	query := fmt.Sprintf(`
		select
			deployment_type,
			sql_statement
		from
			_orchestrator_db_deployment
		order by
			deployment_id
		`)
	db, err := OpenOrchestrator()
	if err != nil {
		log.Fatalf("Cannot initiate orchestrator internal deployment: %+v", err)
	}

	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		deploymentType := m.GetString("deployment_type")
		sqlStatement := m.GetString("sql_statement")

		if deploymentType == "base" {
			baseDeployments = append(baseDeployments, sqlStatement)
		} else if deploymentType == "patch" {
			patchDeployments = append(patchDeployments, sqlStatement)
		} else {
			log.Fatalf("Unknown deployment type (%+v) encountered in _orchestrator_db_deployment", deploymentType)
		}

		return nil
	})
	if err != nil {
		log.Debugf("Deploying internal orchestrator tables to fix the above; this is a one time operation")
		// Table does not exist? Create it for first time
		for _, query := range internalDBDeploymentSQL {
			if _, err = execInternal(db, query); err != nil {
				log.Fatalf("Cannot initiate orchestrator internal deployment: %+v", err)
			}
		}
	}
	return baseDeployments, patchDeployments, nil
}

// writeInternalDeployment will persist a successful deployment
func writeInternalDeployment(db *sql.DB, deploymentType string, sqlStatement string, statementIndex int) error {
	if !config.Config.SmartOrchestratorDatabaseUpdate {
		return nil
	}
	query := `
        	insert ignore into _orchestrator_db_deployment (
				deployment_type, sql_statement, statement_digest, statement_index) VALUES (
				?, ?, CONCAT(SHA1(?), ':', LEFT(REPLACE(REPLACE(REPLACE(?, ' ', ''), '\n', ' '), '\t', ''), 60)), ?)
				`
	if _, err := execInternal(db, query, deploymentType, sqlStatement, sqlStatement, sqlStatement, statementIndex); err != nil {
		log.Fatalf("Unable to write to _orchestrator_db_deployment: %+v", err)
	}
	return nil
}

// ResetInternalDeployment will clear existing deployment history (and the effect will be a complete re-deployment
// the next run). This is made available for possible operational errors, a red button
func ResetInternalDeployment() error {
	db, err := OpenOrchestrator()
	if err != nil {
		log.Fatalf("Cannot reset orchestrator internal deployment: %+v", err)
	}
	if _, err := execInternal(db, `delete from _orchestrator_db_deployment`); err != nil {
		log.Fatalf("Unable to clear _orchestrator_db_deployment: %+v", err)
	}
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

// deployIfNotAlreadyDeployed will issue given sql queries that are not already known to be deployed.
// This iterates both lists (to-run and already-deployed) and also verifies no contraditions.
func deployIfNotAlreadyDeployed(db *sql.DB, queries []string, deployedQueries []string, deploymentType string, fatalOnError bool) error {
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
	err = tx.QueryRow(`select @@session.sql_mode`).Scan(&originalSqlMode)
	if _, err := tx.Exec(`set @@session.sql_mode=REPLACE(@@session.sql_mode, 'NO_ZERO_DATE', '')`); err != nil {
		log.Fatale(err)
	}
	if _, err := tx.Exec(`set @@session.sql_mode=REPLACE(@@session.sql_mode, 'NO_ZERO_IN_DATE', '')`); err != nil {
		log.Fatale(err)
	}

	for i, query := range queries {
		queryAlreadyExecuted := false
		// While iterating 'queries', also iterate 'deployedQueries'. Expect identity
		if len(deployedQueries) > i {
			if deployedQueries[i] != query {
				log.Fatalf("initOrchestratorDB() PANIC: non matching %s queries between deployment requests and _orchestrator_db_deployment. Please execute 'orchestrator -c reset-internal-db-deployment'", deploymentType)
			}
			queryAlreadyExecuted = true
		}
		if queryAlreadyExecuted {
			continue
		}
		if i == 0 {
			//log.Debugf("sql_mode is: %+v", originalSqlMode)
		}
		if config.Config.SmartOrchestratorDatabaseUpdate {
			log.Debugf("initOrchestratorDB executing: %.80s", strings.TrimSpace(strings.Replace(query, "\n", "", -1)))
		}

		if fatalOnError {
			if _, err := tx.Exec(query); err != nil {
				return log.Fatalf("Cannot initiate orchestrator: %+v", err)
			}
		} else {
			tx.Exec(query)
			// And ignore any error
		}
		writeInternalDeployment(db, deploymentType, query, i)
	}
	if _, err := tx.Exec(`set session sql_mode=?`, originalSqlMode); err != nil {
		log.Fatale(err)
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

	baseDeployments, patchDeployments, _ := readInternalDeployments()
	deployIfNotAlreadyDeployed(db, generateSQLBase, baseDeployments, "base", true)
	deployIfNotAlreadyDeployed(db, generateSQLPatches, patchDeployments, "patch", false)

	return nil
}

// execInternalSilently
func execInternalSilently(db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	res, err := sqlutils.ExecSilently(db, query, args...)
	return res, err
}

// execInternal
func execInternal(db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	res, err := sqlutils.ExecSilently(db, query, args...)
	return res, err
}

// ExecOrchestrator will execute given query on the orchestrator backend database.
func ExecOrchestrator(query string, args ...interface{}) (sql.Result, error) {
	if config.Config.DatabaselessMode__experimental {
		return DummySqlResult{}, nil
	}
	db, err := OpenOrchestrator()
	if err != nil {
		return nil, err
	}
	res, err := sqlutils.Exec(db, query, args...)
	return res, err
}

// QueryRowsMapOrchestrator
func QueryOrchestratorRowsMap(query string, on_row func(sqlutils.RowMap) error) error {
	if config.Config.DatabaselessMode__experimental {
		return nil
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
	db, err := OpenOrchestrator()
	if err != nil {
		return err
	}

	if argsArray == nil {
		argsArray = EmptyArgs
	}
	return log.Criticale(sqlutils.QueryRowsMapBuffered(db, query, on_row, argsArray...))
}
