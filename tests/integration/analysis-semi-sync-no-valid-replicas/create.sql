UPDATE database_instance SET semi_sync_master_enabled=1, semi_sync_master_status=0, semi_sync_master_wait_for_slave_count=1, semi_sync_master_clients=0 where port=22293;
UPDATE database_instance SET semi_sync_replica_enabled=1;
