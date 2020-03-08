# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.28-log)
# Database: go_orm
# Generation Time: 2020-03-08 07:24:25 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table activities
# ------------------------------------------------------------

DROP TABLE IF EXISTS `activities`;

CREATE TABLE `activities` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `consumer_id` int(11) DEFAULT NULL,
  `source` varchar(15) DEFAULT NULL COMMENT 'consumer',
  `type` varchar(20) DEFAULT NULL COMMENT 'action type like ORDER_UPDATE',
  `entity_id` varchar(50) DEFAULT NULL COMMENT 'on which consumer taking action such as ride_id',
  `entity_name` varchar(60) DEFAULT NULL COMMENT 'which servic''s data consumer is editing like foo/ride/parcel',
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

INSERT INTO `activities` (`id`, `consumer_id`, `source`, `type`, `entity_id`, `entity_name`, `data`, `description`, `action_by`, `created_at`, `updated_at`, `deleted_at`)
VALUES
  (1,4,'engineroom','RIDE_DURATION','20BGIJH','ride',NULL,NULL,'imran@gmail.com','2020-03-07 05:04:29','2020-03-07 05:04:29',NULL),
  (2,4,'engineroom','RIDE_DURATION','20BGIJHk','ride',NULL,NULL,'imran@gmail.com','2020-03-07 05:04:31','2020-03-07 05:04:31',NULL),
  (3,4,'engineroom','RIDE_DURATION','20BGIJH','ride',NULL,NULL,'imran@gmail.com','2020-03-07 05:16:48','2020-03-07 05:16:48',NULL),
  (4,4,'engineroom','RIDE_DURATION','20BGIJH','ride',NULL,NULL,'imran@gmail.com','2020-03-07 05:30:34','2020-03-07 05:30:34',NULL),
  (5,4,'engineroom','RIDE_DURATION','20BGIJH','ride',NULL,NULL,'imran@gmail.com','2020-03-07 05:30:36','2020-03-07 05:30:36',NULL),
  (6,4,'engineroom','RIDE_DURATION','20BGIJH','ride',NULL,NULL,'imran@gmail.com','2020-03-07 05:30:36','2020-03-07 05:30:36',NULL),
  (7,4,'engineroom','RIDE_DURATION','20BGIJH','ride',NULL,NULL,'imran@gmail.com','2020-03-07 08:47:55','2020-03-07 08:47:55',NULL),
  (8,4,'engineroom','RIDE_DURATION','20BGIJH','ride',NULL,NULL,'imran@gmail.com','2020-03-07 08:51:43','2020-03-07 08:51:43',NULL);

/*!40000 ALTER TABLE `activities` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table changes
# ------------------------------------------------------------

DROP TABLE IF EXISTS `changes`;

CREATE TABLE `changes` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `activity_id` int(11) DEFAULT NULL,
  `field` text,
  `new_value` text,
  `old_value` text,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `changes` WRITE;
/*!40000 ALTER TABLE `changes` DISABLE KEYS */;

INSERT INTO `changes` (`id`, `activity_id`, `field`, `new_value`, `old_value`, `created_at`, `updated_at`, `deleted_at`)
VALUES
  (1,1,'ride_duration','20','20','2020-03-07 05:04:29','2020-03-07 05:04:29',NULL),
  (2,1,'ride_duration','20','20','2020-03-07 05:04:29','2020-03-07 05:04:29',NULL),
  (3,2,'ride_duration','20','20','2020-03-07 05:04:31','2020-03-07 05:04:31',NULL),
  (4,2,'ride_duration','20','20','2020-03-07 05:04:31','2020-03-07 05:04:31',NULL),
  (5,3,'ride_duration','20','20','2020-03-07 05:16:48','2020-03-07 05:16:48',NULL),
  (6,3,'ride_duration','20','20','2020-03-07 05:16:48','2020-03-07 05:16:48',NULL),
  (7,4,'ride_duration','20','20','2020-03-07 05:30:34','2020-03-07 05:30:34',NULL),
  (8,4,'ride_duration','20','20','2020-03-07 05:30:34','2020-03-07 05:30:34',NULL),
  (9,5,'ride_duration','20','20','2020-03-07 05:30:36','2020-03-07 05:30:36',NULL),
  (10,5,'ride_duration','20','20','2020-03-07 05:30:36','2020-03-07 05:30:36',NULL),
  (11,6,'ride_duration','20','20','2020-03-07 05:30:36','2020-03-07 05:30:36',NULL),
  (12,6,'ride_duration','20','20','2020-03-07 05:30:36','2020-03-07 05:30:36',NULL),
  (13,7,'ride_duration','20','20','2020-03-07 08:47:55','2020-03-07 08:47:55',NULL),
  (14,7,'ride_duration','20','20','2020-03-07 08:47:55','2020-03-07 08:47:55',NULL),
  (15,8,'ride_duration','20','20','2020-03-07 08:51:43','2020-03-07 08:51:43',NULL),
  (16,8,'ride_duration','20','20','2020-03-07 08:51:43','2020-03-07 08:51:43',NULL);

/*!40000 ALTER TABLE `changes` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table consumers
# ------------------------------------------------------------

DROP TABLE IF EXISTS `consumers`;

CREATE TABLE `consumers` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `app_name` varchar(25) DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `app_name` (`app_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `consumers` WRITE;
/*!40000 ALTER TABLE `consumers` DISABLE KEYS */;

INSERT INTO `consumers` (`id`, `app_name`, `token`, `created_at`, `updated_at`, `deleted_at`)
VALUES
  (4,'er','a1aac9f0a7718d1578ea0a4028e61fce','2020-02-24 09:22:22','2020-02-24 09:22:22',NULL),
  (11,'payout','c83526409ff4922fd56df7f306fbabee','2020-02-24 09:45:04','2020-02-24 09:45:04',NULL);

/*!40000 ALTER TABLE `consumers` ENABLE KEYS */;
UNLOCK TABLES;

DELIMITER ;;
/*!50003 SET SESSION SQL_MODE="ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION" */;;
/*!50003 CREATE */ /*!50017 DEFINER=`root`@`%` */ /*!50003 TRIGGER `BeforeConsumerInsert` BEFORE INSERT ON `consumers` FOR EACH ROW BEGIN
    IF NEW.token IS NULL THEN
       SET NEW.token = SUBSTR(SHA1(uuid()), 1, 32);
    END IF;
 
END */;;
DELIMITER ;
/*!50003 SET SESSION SQL_MODE=@OLD_SQL_MODE */;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
