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

 Date: 12/06/2023 23:50:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `birth_date` date NOT NULL DEFAULT '0000-00-00' COMMENT '出生日期',
  `plat_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '应用id',
  `create_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `update_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `delete_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间戳',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
