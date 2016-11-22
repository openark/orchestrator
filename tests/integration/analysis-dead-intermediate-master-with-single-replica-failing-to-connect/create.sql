-- 22295 replicates from 22294
UPDATE database_instance SET slave_io_running=0, last_io_error='error connecting to master' where master_port=22294;
UPDATE database_instance SET last_seen=last_checked - interval 1 minute where port=22294;
