UPDATE database_instance SET
  last_checked=current_timestamp,
  last_attempted_check=current_timestamp,
  last_seen=current_timestamp,
  seconds_behind_master=7,
  slave_lag_seconds=7
where port in (22294, 22297);
