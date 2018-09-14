UPDATE database_instance SET last_seen=last_checked - interval 1 minute where port=22293;
UPDATE database_instance SET slave_lag_seconds=600, sql_delay=600 where port in (22294, 22296, 22297);
