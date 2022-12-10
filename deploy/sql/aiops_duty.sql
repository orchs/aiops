
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `duty_conf`;
CREATE TABLE `ltw_duty_conf` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `group_id` bigint(20) unsigned DEFAULT NULL,
    `name` varchar(64) DEFAULT NULL,
    `start_at` varchar(64) DEFAULT '00:00',
    `end_at` varchar(64) DEFAULT '24:00',
    `priority` tinyint(3) unsigned DEFAULT NULL,
    `create_at` bigint(20) NOT NULL DEFAULT '0',
    `create_by` varchar(64) NOT NULL DEFAULT '',
    `update_at` bigint(20) NOT NULL DEFAULT '0',
    `update_by` varchar(64) NOT NULL DEFAULT '',
    `using` tinyint(4) NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`),
    KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `duty_roster`;
CREATE TABLE `duty_roster` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `duty_conf_id` bigint(20) unsigned DEFAULT NULL,
    `user_id` bigint(20) unsigned DEFAULT NULL,
    `start_at` bigint(20) NOT NULL DEFAULT '0',
    `end_at` bigint(20) NOT NULL DEFAULT '0',
    `create_at` bigint(20) NOT NULL DEFAULT '0',
    `create_by` varchar(64) NOT NULL DEFAULT '',
    `update_at` bigint(20) NOT NULL DEFAULT '0',
    `update_by` varchar(64) NOT NULL DEFAULT '',
    `duty_date` int(11) DEFAULT '20060102',
    `group_id` int(11) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `key_start_end` (`start_at`,`end_at`),
    KEY `duty_conf_id` (`duty_conf_id`),
    KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=307 DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
