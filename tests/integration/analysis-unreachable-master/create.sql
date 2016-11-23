UPDATE database_instance SET last_seen=last_checked - interval 1 minute where port=22293;
