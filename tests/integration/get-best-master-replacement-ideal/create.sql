UPDATE database_instance SET data_center='ny', physical_environment='prod';

INSERT INTO candidate_database_instance (hostname, port, promotion_rule) VALUES ('testhost', 22297, 'prefer');
