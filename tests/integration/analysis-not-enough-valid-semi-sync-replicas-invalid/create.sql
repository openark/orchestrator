UPDATE database_instance SET semi_sync_master_enabled=1, semi_sync_master_wait_for_slave_count=2 where port=22293;
UPDATE database_instance SET semi_sync_replica_enabled=1 where port in (22294, 22296);
UPDATE database_instance SET last_seen=last_checked - interval 1 minute where port=22296;
