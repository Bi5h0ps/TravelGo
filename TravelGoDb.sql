/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80033 (8.0.33)
 Source Host           : localhost:3306
 Source Schema         : TravelGoDb

 Target Server Type    : MySQL
 Target Server Version : 80033 (8.0.33)
 File Encoding         : 65001

 Date: 30/07/2023 19:53:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
  `ID` bigint NOT NULL AUTO_INCREMENT,
  `post_id` bigint DEFAULT NULL,
  `username` longtext,
  `content` longtext,
  `created_at` datetime(3) DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of comments
-- ----------------------------
BEGIN;
INSERT INTO `comments` (`ID`, `post_id`, `username`, `content`, `created_at`, `is_deleted`) VALUES (1, 1, 'haogecsol', 'lol Toronto is no fun', '2023-07-23 10:45:01.646', 0);
INSERT INTO `comments` (`ID`, `post_id`, `username`, `content`, `created_at`, `is_deleted`) VALUES (2, 1, 'haogecsol', 'sounds like a plan', '2023-07-23 10:45:10.803', 0);
INSERT INTO `comments` (`ID`, `post_id`, `username`, `content`, `created_at`, `is_deleted`) VALUES (3, 2, 'haogecsol', 'Tokyo sounds like a plan', '2023-07-23 10:45:39.139', 0);
INSERT INTO `comments` (`ID`, `post_id`, `username`, `content`, `created_at`, `is_deleted`) VALUES (4, 2, 'haogecsol', 'Tokyo woule be so hot by then', '2023-07-23 10:45:48.989', 0);
INSERT INTO `comments` (`ID`, `post_id`, `username`, `content`, `created_at`, `is_deleted`) VALUES (5, 3, 'haogecsol', 'Shanghai is my favorite city!', '2023-07-23 10:46:03.431', 0);
COMMIT;

-- ----------------------------
-- Table structure for travel_posts
-- ----------------------------
DROP TABLE IF EXISTS `travel_posts`;
CREATE TABLE `travel_posts` (
  `ID` bigint NOT NULL AUTO_INCREMENT,
  `username` longtext,
  `post_title` longtext,
  `destination` longtext,
  `start_date` date DEFAULT NULL,
  `end_date` date DEFAULT NULL,
  `tags` longtext,
  `is_deleted` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of travel_posts
-- ----------------------------
BEGIN;
INSERT INTO `travel_posts` (`ID`, `username`, `post_title`, `destination`, `start_date`, `end_date`, `tags`, `is_deleted`) VALUES (1, 'haogecsol', 'Toronto Summer Travel Idea', 'Toronto', '2023-07-23', '2023-07-26', 'Canada|vacation|expensive', 0);
INSERT INTO `travel_posts` (`ID`, `username`, `post_title`, `destination`, `start_date`, `end_date`, `tags`, `is_deleted`) VALUES (2, 'haogecsol', 'Tokyo Summer Break', 'Tokyo', '2023-07-28', '2023-08-28', 'Japan|vacation|expensive', 0);
INSERT INTO `travel_posts` (`ID`, `username`, `post_title`, `destination`, `start_date`, `end_date`, `tags`, `is_deleted`) VALUES (3, 'haogecsol', 'Shanghai Staycation', 'Shanghai', '2023-08-05', '2023-08-12', 'China|vacation|expensive', 0);
INSERT INTO `travel_posts` (`ID`, `username`, `post_title`, `destination`, `start_date`, `end_date`, `tags`, `is_deleted`) VALUES (4, 'haogecsol', 'Shenzhen Staycation', 'Shenzhen', '2023-08-05', '2023-08-15', 'China|vacation|expensive', 0);
INSERT INTO `travel_posts` (`ID`, `username`, `post_title`, `destination`, `start_date`, `end_date`, `tags`, `is_deleted`) VALUES (5, 'haogecsol', 'Hong Kong Summer Travel Idea S', 'Hong Kong', '2023-07-21', '2023-07-30', 'summer|fun|sweat', 0);
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `ID` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(191) NOT NULL,
  `password` longtext,
  `email` longtext,
  PRIMARY KEY (`ID`,`username`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` (`ID`, `username`, `password`, `email`) VALUES (1, 'haogecsol', '$2a$10$ma4jatFd2XINoiF0g8SoseWkMMMTgzrjS5UJGm/Xug1dGMUiZA6MS', 'guohy716@cottt.hku.hk');
INSERT INTO `users` (`ID`, `username`, `password`, `email`) VALUES (2, 'testaccount1', '$2a$10$WmPhQLMLBqPlGU5ad1DnbuoWY/XMny1JPLbCUBl183noqOad2iTCW', 'test1@cottt.hku.hk');
INSERT INTO `users` (`ID`, `username`, `password`, `email`) VALUES (3, 'testaccount2', '$2a$10$ayYdoHTa6LZTYLHSocCAU.Cp0bZIO27TnVqmkNVKA1Q2IHP7HswyS', 'test2@cottt.hku.hk');
INSERT INTO `users` (`ID`, `username`, `password`, `email`) VALUES (4, 'testaccount3', '$2a$10$HP77Lde.T9H8X37X2bZi7OlQmqWkPQrajT00TKPBsO8boRKDzLuSu', 'test3@cottt.hku.hk');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
