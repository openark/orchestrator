UPDATE database_instance SET semi_sync_master_enabled=1, semi_sync_master_wait_for_slave_count=2 where port=22293;
UPDATE database_instance SET semi_sync_replica_enabled=1 where master_port=22293 limit 1;
