/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50617
Source Host           : localhost:3306
Source Database       : jeedev

Target Server Type    : MYSQL
Target Server Version : 50617
File Encoding         : 65001

Date: 2015-08-06 09:30:49
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `app`
-- ----------------------------
DROP TABLE IF EXISTS `app`;
CREATE TABLE `app` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `create_date` datetime NOT NULL,
  `is_delete` bit(1) DEFAULT b'0',
  `re_order` int(11) DEFAULT 0,
  `update_date` datetime NOT NULL,
  `app_code` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `app_name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `publish_date` date DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `app_code` (`app_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Records of app
-- ----------------------------
INSERT INTO `app` VALUES ('1', '2014-08-20 09:17:43', '', '0', '2014-08-20 09:17:48', '100000', '神庙逃亡', '2014-08-01');
INSERT INTO `app` VALUES ('2', '2014-08-20 09:18:54', '', '0', '2014-08-20 09:18:49', '100001', '愤怒的小鸟', '2014-07-08');
