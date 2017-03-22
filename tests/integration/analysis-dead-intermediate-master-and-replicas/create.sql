-- 22295 replicates from 22294
UPDATE database_instance SET last_seen=last_checked - interval 1 minute where master_port=22294;
UPDATE database_instance SET last_seen=last_checked - interval 1 minute where port=22294;
