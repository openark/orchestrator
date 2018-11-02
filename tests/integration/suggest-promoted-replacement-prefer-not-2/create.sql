UPDATE database_instance SET master_port=22294 where port not in (22293, 22294);
-- topology:
--
-- 22293
-- + 22294
--   + 22295
--   + 22296
--   + 22297
--
UPDATE database_instance SET data_center='seattle', physical_environment='prod';
UPDATE database_instance SET data_center='ny', physical_environment='prod' where port in (22294, 22295);
UPDATE database_instance SET data_center='sf', physical_environment='prod' where port in (22297);

INSERT INTO candidate_database_instance (hostname, port, promotion_rule) VALUES ('testhost', 22294, 'prefer_not');
INSERT INTO candidate_database_instance (hostname, port, promotion_rule) VALUES ('testhost', 22295, 'neutral');
INSERT INTO candidate_database_instance (hostname, port, promotion_rule) VALUES ('testhost', 22297, 'neutral');
