UPDATE database_instance SET data_center='seattle', physical_environment='prod';
UPDATE database_instance SET data_center='ny', physical_environment='prod' where port in (22293, 22296);


INSERT INTO candidate_database_instance (hostname, port, promotion_rule) VALUES ('testhost', 22295, 'prefer');
