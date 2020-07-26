UPDATE database_instance SET semi_sync_master_enabled=1, semi_sync_master_status=1, semi_sync_master_wait_for_slave_count=2, semi_sync_master_clients=1 where port=22293;
UPDATE database_instance SET last_seen=last_checked - interval 1 minute where port=22296;
INSERT INTO database_instance_stale_binlog_coordinates (
    hostname, port,	binary_log_file, binary_log_pos, first_seen
  )
  values (
    'testhost', 22293, 'mysql-bin.000167', 137086726, CURRENT_TIMESTAMP
  );
