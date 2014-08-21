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
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/outbrain/log"
	"github.com/outbrain/sqlutils"
	"github.com/outbrain/orchestrator/config"
)


// generateSQL is a list of SQL statements required to build the orchestrator backend
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
		  end_timestamp timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
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
}


// OpenTopology returns a DB instance to access a topology instance
func OpenTopology(host string, port int) (*sql.DB, error) {
	mysql_uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/", config.Config.MySQLTopologyUser, config.Config.MySQLTopologyPassword, host, port)
	db, _, err := sqlutils.GetDB(mysql_uri)
	return db, err
}

// OpenTopology returns the DB instance for the orchestrator backed database
func OpenOrchestrator() (*sql.DB, error) {
	mysql_uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Config.MySQLOrchestratorUser, config.Config.MySQLOrchestratorPassword, 
		config.Config.MySQLOrchestratorHost, config.Config.MySQLOrchestratorPort, config.Config.MySQLOrchestratorDatabase)
	db, fromCache, err := sqlutils.GetDB(mysql_uri)
	if err == nil && !fromCache {
		initOrchestratorDB(db)
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
	return nil
}


// ExecOrchestrator will execute given query on the orchestrator backend database.
func ExecOrchestrator(query string, args ...interface{}) (sql.Result, error) {
	db,	err	:=	OpenOrchestrator()
	if err != nil {
		return nil, err
	}
	res, err := sqlutils.Exec(db, query, args...)
	return res, err
}

