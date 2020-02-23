CREATE TABLE  IF NOT EXISTS `activities` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `source` varchar(15) DEFAULT NULL COMMENT 'consumer',
  `type` varchar(20) DEFAULT NULL COMMENT 'action type like ORDER_UPDATE',
  `entity_id` varchar(11) DEFAULT NULL COMMENT 'on which consumer taking action such as ride_id',
  `entity_type` varchar(11) DEFAULT NULL COMMENT 'which servic''s data consumer is editing like foo/ride/parcel',
  `field` varchar(255) DEFAULT NULL COMMENT 'the keys/fields which are being updated like pickup_addess/dropoff_address',
  `new_value` text,
  `old_value` text,
  `data` varchar(30) DEFAULT NULL,
  `description` text COMMENT 'full request body',
  `action_by` varchar(30) DEFAULT NULL COMMENT 'the user by whom the action is taken',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `source` (`source`),
  KEY `type` (`type`),
  KEY `entity_id` (`entity_id`),
  KEY `action_by` (`action_by`),
  KEY `created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `activities` WRITE;
/*!40000 ALTER TABLE `activities` DISABLE KEYS */;

INSERT INTO `activities` (`source`, `type`, `entity_id`, `entity_type`, `field`, `new_value`, `old_value`, `data`, `description`, `action_by`, `created_at`, `updated_at`, `deleted_at`)
VALUES
  ('engineroom','RIDE_DURATION','20BGIJH','ride','ride_duration','20','20','',NULL,'imran@gmail.com','2020-02-21 06:32:41','2020-02-21 06:32:41',NULL),
  ('engineroom','RIDE_DISTANCE','20BGKG6','ride','ride_distance','5','5.6','','','tuhin@gmail.com','2020-02-21 06:33:14','2020-02-21 06:33:14',NULL),
  ('payout','DUE_ADJUSTMENT','1','driver','amount_in_cents','1290','3000','payment has been collected for','{\"transaction_log_id\":330,\"driver_id\":1566,\"amount\":-315.73,\"transaction_id\":12950}','rakib@gmail.com','2020-02-21 07:08:15','2020-02-21 07:08:15',NULL);

/*!40000 ALTER TABLE `activities` ENABLE KEYS */;
UNLOCK TABLES;


