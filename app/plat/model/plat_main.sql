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

 Date: 12/06/2023 23:50:14
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for plat_main
-- ----------------------------
DROP TABLE IF EXISTS `plat_main`;
CREATE TABLE `plat_main` (
  `id` bigint(20) unsigned NOT NULL,
  `appid` varchar(80) NOT NULL DEFAULT '' COMMENT '对外应用标识',
  `secret` varchar(80) NOT NULL DEFAULT '' COMMENT '对外应用密钥',
  `sys_plat_state_id` tinyint(4) NOT NULL DEFAULT '0' COMMENT '应用状态',
  `renter_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '租户id',
  `name` varchar(80) NOT NULL DEFAULT '' COMMENT '应用名称',
  `sys_plat_clas_id` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '应用类型',
  `expire_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '应用到期时间戳',
  `create_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `update_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `delete_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间戳',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
