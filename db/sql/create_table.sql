-- Adminer 4.8.1 MySQL 8.0.29 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `eventid` bigint NOT NULL,
  `userid` bigint NOT NULL,
  `comment` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `postedAt` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `eventid` (`eventid`),
  KEY `userid` (`userid`),
  CONSTRAINT `comment_ibfk_1` FOREIGN KEY (`eventid`) REFERENCES `event` (`id`),
  CONSTRAINT `comment_ibfk_2` FOREIGN KEY (`userid`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


DROP TABLE IF EXISTS `event`;
CREATE TABLE `event` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `start` datetime NOT NULL,
  `place` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `open` datetime NOT NULL,
  `close` datetime NOT NULL,
  `author` bigint NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


DROP TABLE IF EXISTS `eventUser`;
CREATE TABLE `eventUser` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `eventid` bigint NOT NULL,
  `userid` bigint NOT NULL,
  PRIMARY KEY (`id`),
  KEY `eventid` (`eventid`),
  KEY `userid` (`userid`),
  CONSTRAINT `eventUser_ibfk_1` FOREIGN KEY (`eventid`) REFERENCES `event` (`id`),
  CONSTRAINT `eventUser_ibfk_2` FOREIGN KEY (`userid`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


DROP TABLE IF EXISTS `upload`;
CREATE TABLE `upload` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `eventid` bigint NOT NULL,
  `url` text COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `eventid` (`eventid`),
  CONSTRAINT `upload_ibfk_1` FOREIGN KEY (`eventid`) REFERENCES `event` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` text COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- 2022-07-09 05:16:08
