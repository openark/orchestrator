UPDATE database_instance SET master_port=22294 where port not in (22293, 22294);
-- topology:
--
-- 22293
-- + 22294
--   + 22295
--   + 22296
--   + 22297
--
