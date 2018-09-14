UPDATE database_instance SET master_uuid='00020192-1111-1111-1111-111111111111', gtid_errant='00020194-3333-3333-3333-333333333333:1-2' where port=22294;
INSERT INTO database_instance_downtime (
    hostname, port, downtime_active, begin_timestamp, end_timestamp, owner, reason
  ) values (
    'testhost', 22294, 1, current_timestamp, current_timestamp, 'test', 'integration test'
  );
UPDATE database_instance_downtime SET end_timestamp=current_timestamp + interval 1 minute;
