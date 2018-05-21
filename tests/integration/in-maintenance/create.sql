INSERT INTO database_instance_maintenance (
    hostname, port, maintenance_active, begin_timestamp, end_timestamp, owner, reason
  ) values (
    'testhost', 22293, 1, current_timestamp, current_timestamp, 'test', 'integration test'
  );
UPDATE database_instance_maintenance SET end_timestamp=current_timestamp + interval 1 minute;
