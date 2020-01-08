insert into database_instance_tags (
      hostname, port, tag_name, tag_value
) values
('testhost', 22293, 'role', 'backup'),
('testhost', 22293, 'in_service', ''),
('testhost', 22294, 'role', 'delayed'),
('testhost', 22295, 'in_service', '');
