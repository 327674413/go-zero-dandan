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

 Date: 12/06/2023 23:50:33
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_main
-- ----------------------------
DROP TABLE IF EXISTS `user_main`;
CREATE TABLE `user_main` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_union_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '平台层用户唯一表示',
  `state_em` tinyint(4) NOT NULL DEFAULT '0' COMMENT '用户状态枚举',
  `account` varchar(80) NOT NULL DEFAULT '' COMMENT '登录账号',
  `password` char(64) NOT NULL DEFAULT '' COMMENT '登录密码',
  `uid` varchar(40) NOT NULL DEFAULT '' COMMENT '用户编号',
  `nickname` varchar(80) NOT NULL DEFAULT '' COMMENT '昵称',
  `phone` varchar(40) NOT NULL DEFAULT '' COMMENT '手机号',
  `phone_area` varchar(10) NOT NULL DEFAULT '' COMMENT '手机区号',
  `email` varchar(80) NOT NULL DEFAULT '' COMMENT '邮箱地址',
  `avatar` varchar(200) NOT NULL DEFAULT '' COMMENT '头像',
  `sex_em` tinyint(4) NOT NULL DEFAULT '0' COMMENT '性别枚举',
  `plat_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '应用id',
  `create_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间戳',
  `update_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间戳',
  `delete_at` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间戳',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
