/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50725
 Source Host           : localhost:3306
 Source Schema         : outofmemory_dev

 Target Server Type    : MySQL
 Target Server Version : 50725
 File Encoding         : 65001

 Date: 20/06/2019 17:38:42
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(11) unsigned NOT NULL COMMENT '文章ID',
  `title` varchar(100) NOT NULL COMMENT '标题',
  `content` text NOT NULL COMMENT '内容',
  `type` varchar(50) NOT NULL DEFAULT 'md' COMMENT '文章类型',
  `state` tinyint(1) unsigned NOT NULL DEFAULT '2' COMMENT '文章状态（0: 删除，1: 已发布，2: 草稿）',
  `author_id` int(11) unsigned NOT NULL COMMENT '作者ID',
  `author_name` varchar(100) NOT NULL COMMENT '作者名称',
  `author_avatar` varchar(255) DEFAULT NULL COMMENT '作者头像链接',
  `created_at` bigint(20) DEFAULT NULL,
  `updated_at` bigint(20) DEFAULT NULL,
  `deleted_at` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `category_id` int(11) unsigned NOT NULL COMMENT '目录ID',
  `name` varchar(100) NOT NULL COMMENT '目录名',
  `state` tinyint(1) NOT NULL COMMENT '目录状态（0: 停用，1:正常，2:删除）',
  `created_at` bigint(20) DEFAULT NULL,
  `updated_at` bigint(20) DEFAULT NULL,
  `deleted_at` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of category
-- ----------------------------
BEGIN;
INSERT INTO `category` VALUES (1, 0, 'default', 1, 1561020588, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for category_article
-- ----------------------------
DROP TABLE IF EXISTS `category_article`;
CREATE TABLE `category_article` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `author_id` int(11) unsigned NOT NULL COMMENT '作者ID',
  `category_id` int(11) unsigned NOT NULL COMMENT '目录ID',
  `article_id` int(11) unsigned NOT NULL COMMENT '文章ID',
  `created_at` bigint(20) DEFAULT NULL,
  `updated_at` bigint(20) DEFAULT NULL,
  `deleted_at` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int(11) unsigned NOT NULL COMMENT '标签ID',
  `name` varchar(100) NOT NULL COMMENT '标签名',
  `state` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '标签状态（0：停用，1:正常，2:删除）',
  `created_at` bigint(20) DEFAULT NULL,
  `updated_at` bigint(20) DEFAULT NULL,
  `deleted_at` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for tag_article
-- ----------------------------
DROP TABLE IF EXISTS `tag_article`;
CREATE TABLE `tag_article` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `author_id` int(11) unsigned NOT NULL COMMENT '作者ID',
  `tag_id` int(11) unsigned NOT NULL COMMENT '标签ID',
  `created_at` bigint(20) DEFAULT NULL,
  `updated_at` bigint(20) DEFAULT NULL,
  `deleted_at` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(11) unsigned NOT NULL COMMENT '用户ID',
  `username` varchar(100) NOT NULL COMMENT '用户名',
  `gender` varchar(5) NOT NULL DEFAULT 'm' COMMENT '性别',
  `email` varchar(255) DEFAULT NULL COMMENT '邮箱',
  `birthday` bigint(20) DEFAULT NULL COMMENT '生日',
  `website` varchar(255) DEFAULT NULL COMMENT '个人站点',
  `phone` varchar(50) DEFAULT NULL COMMENT '手机号码',
  `avatar_url` varchar(255) DEFAULT NULL COMMENT '头像链接',
  `created_at` bigint(20) DEFAULT NULL,
  `updated_at` bigint(20) DEFAULT NULL,
  `deleted_at` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for user_auth
-- ----------------------------
DROP TABLE IF EXISTS `user_auth`;
CREATE TABLE `user_auth` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(11) unsigned NOT NULL COMMENT '用户ID',
  `identifier` varchar(100) NOT NULL COMMENT '用户唯一标识',
  `credential` varchar(100) NOT NULL COMMENT '认证凭据',
  `identifier_from` tinyint(1) unsigned NOT NULL COMMENT '认证来源（0: 站内， 1: 站外）',
  `identity_type` varchar(50) NOT NULL COMMENT '认证类型',
  `verified` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:未认证，1:已认证',
  `verify_date` bigint(20) unsigned DEFAULT NULL COMMENT '完成认证时间',
  `last_login_time` bigint(20) DEFAULT NULL COMMENT '上次登录时间',
  `created_at` bigint(20) DEFAULT NULL,
  `updated_at` bigint(20) DEFAULT NULL,
  `deleted_at` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
