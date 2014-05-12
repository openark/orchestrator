package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/outbrain/sqlutils"
	"github.com/outbrain/orchestrator/config"
)


var generateSQL = []string{
	`
		CREATE TABLE database_instance (
		  hostname varchar(128) CHARACTER SET ascii NOT NULL,
		  port smallint(5) unsigned NOT NULL,
		  last_checked timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		  last_seen timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
		  is_last_seen_valid tinyint(3) unsigned NOT NULL,
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
		  num_slave_hosts int(10) unsigned NOT NULL,
		  slave_hosts text CHARACTER SET ascii NOT NULL,
		  cluster_name tinytext CHARACTER SET ascii NOT NULL,
		  PRIMARY KEY (hostname,port),
		  KEY cluster_name_idx (cluster_name(128))
		) ENGINE=InnoDB DEFAULT CHARSET=ascii
	`,
	`
		CREATE TABLE database_instance_maintenance (
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
}


func OpenTopology(host string, port int) (*sql.DB, error) {
	mysql_uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/", config.Config.MySQLTopologyUser, config.Config.MySQLTopologyPassword, host, port)
	return sqlutils.GetDB(mysql_uri)
}

func OpenOrchestrator() (*sql.DB, error) {
	mysql_uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Config.MySQLOrchestratorUser, config.Config.MySQLOrchestratorPassword, 
		config.Config.MySQLOrchestratorHost, config.Config.MySQLOrchestratorPort, config.Config.MySQLOrchestratorDatabase)
	return sqlutils.GetDB(mysql_uri)
}

func ExecOrchestrator(query string, args ...interface{}) (sql.Result, error) {
	db,	err	:=	OpenOrchestrator()
	if err != nil {
		return nil, err
	}
	res, err := sqlutils.Exec(db, query, args...)
	return res, err
}

