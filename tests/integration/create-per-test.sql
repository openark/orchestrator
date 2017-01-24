DROP TABLE IF EXISTS database_instance;
CREATE TABLE database_instance (
  hostname varchar(128) NOT NULL,
  port smallint(5) unsigned NOT NULL,
  last_checked timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  last_attempted_check timestamp NOT NULL DEFAULT '1970-12-31 23:00:00',
  last_seen timestamp NULL DEFAULT NULL,
  uptime int(10) unsigned NOT NULL,
  server_id int(10) unsigned NOT NULL,
  server_uuid varchar(64) NOT NULL,
  version varchar(128) NOT NULL,
  binlog_server tinyint(3) unsigned NOT NULL,
  read_only tinyint(3) unsigned NOT NULL,
  binlog_format varchar(16) NOT NULL,
  log_bin tinyint(3) unsigned NOT NULL,
  log_slave_updates tinyint(3) unsigned NOT NULL,
  binary_log_file varchar(128) NOT NULL,
  binary_log_pos bigint(20) unsigned NOT NULL,
  master_host varchar(128) NOT NULL,
  master_port smallint(5) unsigned NOT NULL,
  slave_sql_running tinyint(3) unsigned NOT NULL,
  slave_io_running tinyint(3) unsigned NOT NULL,
  has_replication_filters tinyint(3) unsigned NOT NULL,
  oracle_gtid tinyint(3) unsigned NOT NULL,
  executed_gtid_set text NOT NULL,
  gtid_purged text NOT NULL,
  supports_oracle_gtid tinyint(3) unsigned NOT NULL,
  mariadb_gtid tinyint(3) unsigned NOT NULL,
  pseudo_gtid tinyint(3) unsigned NOT NULL,
  master_log_file varchar(128) NOT NULL,
  read_master_log_pos bigint(20) unsigned NOT NULL,
  relay_master_log_file varchar(128) NOT NULL,
  exec_master_log_pos bigint(20) unsigned NOT NULL,
  relay_log_file varchar(128) NOT NULL,
  relay_log_pos bigint(20) unsigned NOT NULL,
  last_sql_error text NOT NULL,
  last_io_error text NOT NULL,
  seconds_behind_master bigint(20) unsigned DEFAULT NULL,
  slave_lag_seconds bigint(20) unsigned DEFAULT NULL,
  sql_delay int(10) unsigned NOT NULL,
  allow_tls tinyint(3) unsigned NOT NULL,
  num_slave_hosts int(10) unsigned NOT NULL,
  slave_hosts text NOT NULL,
  cluster_name varchar(128) NOT NULL,
  suggested_cluster_alias varchar(128) NOT NULL,
  data_center varchar(32) NOT NULL,
  physical_environment varchar(32) NOT NULL,
  instance_alias varchar(128) NOT NULL,
  semi_sync_enforced tinyint(3) unsigned NOT NULL,
  replication_depth tinyint(3) unsigned NOT NULL,
  is_co_master tinyint(3) unsigned NOT NULL,
  replication_credentials_available tinyint(3) unsigned NOT NULL,
  has_replication_credentials tinyint(3) unsigned NOT NULL,
  PRIMARY KEY (hostname,port),
  KEY cluster_name_idx (cluster_name),
  KEY last_checked_idx (last_checked),
  KEY last_seen_idx (last_seen),
  KEY master_host_port_idx (master_host,master_port)
) ENGINE=InnoDB DEFAULT CHARSET=ascii;

-- topology:
--
-- 22293
-- + 22294
--   + 22295
-- + 22296
-- + 22297
--
INSERT INTO database_instance VALUES ('testhost',22293,'2016-11-22 09:57:42','2016-11-22 09:57:42','2016-11-22 09:57:42',70189,1,'00022293-1111-1111-1111-111111111111','5.6.28-log',0,0,'STATEMENT',1,1,'mysql-bin.000155',4440872,'',0,0,0,0,0,'','',0,0,1,'',0,'',0,'',0,'','',NULL,NULL,0,0,3,'[{\"Hostname\":\"testhost\",\"Port\":22294},{\"Hostname\":\"testhost\",\"Port\":22297},{\"Hostname\":\"testhost\",\"Port\":22296}]','testhost:22293','testhost','ny','','',0,0,0,0,0);
INSERT INTO database_instance VALUES ('testhost',22294,'2016-11-22 09:57:42','2016-11-22 09:57:42','2016-11-22 09:57:42',79929,101,'9a996060-6b8f-11e6-903f-6191b3cde928','5.6.28-log',0,0,'STATEMENT',1,1,'mysql-bin.000117',9762775,'testhost',22293,1,1,0,0,'','',0,0,1,'mysql-bin.000155',4440872,'mysql-bin.000155',4440872,'mysql-relay.000013',4441035,'\"\"','\"\"',0,0,0,0,1,'[{\"Hostname\":\"testhost\",\"Port\":22295}]','testhost:22293','','seattle','','',0,1,0,0,1);
INSERT INTO database_instance VALUES ('testhost',22295,'2016-11-22 09:57:42','2016-11-22 09:57:42','2016-11-22 09:57:42',483348,102,'9dc85926-6b8f-11e6-903f-85211507e568','5.6.28-log',0,0,'STATEMENT',1,1,'mysql-bin.000117',59871358,'testhost',22294,1,1,0,0,'','',0,0,1,'mysql-bin.000117',9762775,'mysql-bin.000117',9762775,'mysql-relay.000005',4667810,'\"\"','\"\"',0,0,0,0,0,'[]','testhost:22293','','seattle','','',0,2,0,0,1);
INSERT INTO database_instance VALUES ('testhost',22296,'2016-11-22 09:57:42','2016-11-22 09:57:42','2016-11-22 09:57:42',483346,103,'00022296-4444-4444-4444-444444444444','5.6.28',0,0,'STATEMENT',0,1,'',0,'testhost',22293,1,1,0,0,'','',0,0,1,'mysql-bin.000155',4440872,'mysql-bin.000155',4440872,'mysql-relay.000011',4441035,'\"\"','\"\"',0,0,0,0,0,'[]','testhost:22293','','ny','','',0,1,0,0,1);
INSERT INTO database_instance VALUES ('testhost',22297,'2016-11-22 09:57:42','2016-11-22 09:57:42','2016-11-22 09:57:42',483344,104,'00022297-5555-5555-5555-555555555555','5.6.28-log',0,0,'STATEMENT',1,1,'mysql-bin.001001',30564325,'testhost',22293,1,1,0,0,'','',0,0,1,'mysql-bin.000155',4440872,'mysql-bin.000155',4440872,'mysql-relay.000101',4441035,'\"\"','\"\"',0,0,0,0,0,'[]','testhost:22293','','ny','','',0,1,0,0,1);

DELETE FROM candidate_database_instance;
