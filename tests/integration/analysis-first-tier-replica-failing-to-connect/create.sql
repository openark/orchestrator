UPDATE database_instance SET slave_io_running=0, last_io_error='error connecting to master' where port=22294;
