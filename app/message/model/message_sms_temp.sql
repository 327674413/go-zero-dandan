/*
 Navicat Premium Data Transfer

 Source Server         : gozero
 Source Server Type    : MySQL
 Source Server Version : 50650
 Source Host           : 81.69.7.120:3306
 Source Schema         : gozero

 Target Server Type    : MySQL
 Target Server Version : 50650
 File Encoding         : 65001

 Date: 12/06/2023 23:48:55
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for message_sms_temp
-- ----------------------------
DROP TABLE IF EXISTS `message_sms_temp`;
CREATE TABLE `message_sms_temp` (
  `id` bigint(20) unsigned NOT NULL,
  `name` varchar(80) NOT NULL DEFAULT '',
  `secret_id` varchar(80) NOT NULL DEFAULT '' COMMENT 'SecretId',
  `secret_key` varchar(80) NOT NULL DEFAULT '' COMMENT 'SecretKey',
  `region` varchar(80) NOT NULL DEFAULT '' COMMENT 'region',
  `sms_sdk_appid` varchar(80) NOT NULL DEFAULT '' COMMENT 'SmsSdkAppId',
  `SignName` varchar(80) NOT NULL DEFAULT '' COMMENT 'SignName',
  `template_id` varchar(80) NOT NULL DEFAULT '' COMMENT 'TemplateId',
  `plat_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '应用id',
  `create_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `update_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `delete_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间戳',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
