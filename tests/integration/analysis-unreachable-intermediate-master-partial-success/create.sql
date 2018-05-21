-- 22295 replicates from 22294
UPDATE database_instance SET last_seen=last_checked - interval 1 minute where port=22294;
UPDATE database_instance SET last_check_partial_success = 1 where port=22294;
