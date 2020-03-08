TYPE=TRIGGERS
triggers='CREATE DEFINER=`root`@`%` TRIGGER `BeforeConsumerInsert` BEFORE INSERT ON `consumers` FOR EACH ROW BEGIN\n    IF NEW.token IS NULL THEN\n       SET NEW.token = SUBSTR(SHA1(uuid()), 1, 32);\n    END IF;\n \nEND'
sql_modes=1436549152
definers='root@%'
client_cs_names='utf8'
connection_cl_names='utf8_general_ci'
db_cl_names='utf8mb4_unicode_ci'
created=158315233653
