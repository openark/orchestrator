--
-- This is a suggestion for Pseudo-GTID injection and configuration.
-- The Pseudo-GTID injected here is using both DDL & DML. This includes:
-- - Generation of `meta` schema
-- - Generation of `pseudo_gtid_status` table
-- - Generation of `create_pseudo_gtid_event` event (executes every 5 second)
-- - Enabling of event_scheduler (you will need to make sure you have "event_scheduler=1" in my.cnf)
--
-- replace "meta" with any other schema you prefer


--
--  The following to be added to orchestrator.conf.json
--
--  "PseudoGTIDPattern": "drop view if exists .*?`_pseudo_gtid_hint__",
--  "PseudoGTIDMonotonicHint": "asc:",
--  "DetectPseudoGTIDQuery": "select count(*) as pseudo_gtid_exists from meta.pseudo_gtid_status where anchor = 1 and time_generated > now() - interval 2 day",
--


create database if not exists meta;
use meta;

create table if not exists pseudo_gtid_status (
  anchor                      int unsigned not null,
  originating_mysql_host      varchar(128) charset ascii not null,
  originating_mysql_port      int unsigned not null,
  originating_server_id       int unsigned not null,
  time_generated              timestamp not null default current_timestamp,
  pseudo_gtid_uri             varchar(255) charset ascii not null,
  pseudo_gtid_hint            varchar(255) charset ascii not null,
  PRIMARY KEY (anchor)
);

drop event if exists create_pseudo_gtid_event;
delimiter $$
create event if not exists
  create_pseudo_gtid_event
  on schedule every 5 second starts current_timestamp
  on completion preserve
  enable
  do
    begin
      DECLARE lock_result INT;
      DECLARE CONTINUE HANDLER FOR SQLEXCEPTION BEGIN END;

      SET lock_result = GET_LOCK('pseudo_gtid_status', 0);
      IF lock_result = 1 THEN
        set @connection_id := connection_id();
        set @now := now();
        set @rand := floor(rand()*(1 << 32));
        set @pseudo_gtid_hint := concat_ws(':', lpad(hex(unix_timestamp(@now)), 8, '0'), lpad(hex(@connection_id), 16, '0'), lpad(hex(@rand), 8, '0'));
        set @_create_statement := concat('drop ', 'view if exists `meta`.`_pseudo_gtid_', 'hint__asc:', @pseudo_gtid_hint, '`');
        PREPARE st FROM @_create_statement;
        EXECUTE st;
        DEALLOCATE PREPARE st;

        /*!50600
        SET innodb_lock_wait_timeout = 1;
        */
        set @serverid := @@server_id;
        set @hostname := @@hostname;
        set @port := @@port;
        set @pseudo_gtid := concat('pseudo-gtid://', @hostname, ':', @port, '/', @serverid, '/', date(@now), '/', time(@now), '/', @rand);
        insert into pseudo_gtid_status (
             anchor,
             originating_mysql_host,
             originating_mysql_port,
             originating_server_id,
             time_generated,
             pseudo_gtid_uri,
             pseudo_gtid_hint
          )
      	  values (1, @hostname, @port, @serverid, @now, @pseudo_gtid, @pseudo_gtid_hint)
      	  on duplicate key update
      		  originating_mysql_host = values(originating_mysql_host),
      		  originating_mysql_port = values(originating_mysql_port),
      		  originating_server_id = values(originating_server_id),
      		  time_generated = values(time_generated),
       		  pseudo_gtid_uri = values(pseudo_gtid_uri),
       		  pseudo_gtid_hint = values(pseudo_gtid_hint)
        ;
        SET lock_result = RELEASE_LOCK('pseudo_gtid_status');
      END IF;
    end
$$

delimiter ;

set global event_scheduler := 1;
