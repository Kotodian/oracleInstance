/*
 Navicat MySQL Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 50721
 Source Host           : localhost:3306
 Source Schema         : oracle

 Target Server Type    : MySQL
 Target Server Version : 50721
 File Encoding         : 65001

 Date: 21/07/2020 13:55:45
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tb_oracle_gather
-- ----------------------------
DROP TABLE IF EXISTS `tb_oracle_gather`;
CREATE TABLE `tb_oracle_gather`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `bytes_sent_via_sql_net_to_client` double NULL DEFAULT NULL,
  `bytes_received_via_sql_net_from_client` double NULL DEFAULT NULL,
  `dual_time` bigint(20) NULL DEFAULT NULL,
  `iops` double(20, 0) NULL DEFAULT NULL,
  `mbps` double(20, 0) NULL DEFAULT NULL,
  `use_total_PGA` double NULL DEFAULT NULL,
  `share_pool_size` double NULL DEFAULT NULL,
  `sql_pin_hit_ratio` double NULL DEFAULT NULL,
  `buffer_hit` double NULL DEFAULT NULL,
  `sort_memory` double NULL DEFAULT NULL,
  `parse_count_hard` bigint(20) NULL DEFAULT NULL,
  `redo_buffer_allocation_retries` double NULL DEFAULT NULL,
  `user_inactive_sessions` int(11) NULL DEFAULT NULL,
  `user_active_sessions` int(11) NULL DEFAULT NULL,
  `execute_count` int(11) NULL DEFAULT NULL,
  `background_sessions` int(11) NULL DEFAULT NULL,
  `sql_net_roundtrips_tfrom_client` double NULL DEFAULT NULL,
  `time` bigint(20) NULL DEFAULT NULL,
  `instance_id` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 43 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_oracle_gather
-- ----------------------------
INSERT INTO `tb_oracle_gather` VALUES (42, 0, 29368, 0, 9, 0, 147828736, 469762048, 97.74, 99.73, 0.9999792044177379, -768749044, 0, 4, 26, 157, 0, 141, 1595296740, 1);

-- ----------------------------
-- Table structure for tb_oracle_instance
-- ----------------------------
DROP TABLE IF EXISTS `tb_oracle_instance`;
CREATE TABLE `tb_oracle_instance`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `host` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `dbname` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_oracle_instance
-- ----------------------------
INSERT INTO `tb_oracle_instance` VALUES (1, '192.168.1.67', 'system', 'oracle', 'test');

-- ----------------------------
-- Table structure for tb_oracle_last_data
-- ----------------------------
DROP TABLE IF EXISTS `tb_oracle_last_data`;
CREATE TABLE `tb_oracle_last_data`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `bytes_sent_via_sql_net_to_client` double NULL DEFAULT NULL,
  `bytes_received_via_sql_net_from_client` double NULL DEFAULT NULL,
  `parse_count_hard` int(11) NULL DEFAULT NULL,
  `execute_count` int(11) NULL DEFAULT NULL,
  `iops` double(20, 0) NULL DEFAULT NULL,
  `mbps` double(20, 0) NULL DEFAULT NULL,
  `sql_net_round_trips_from_client` double NULL DEFAULT NULL,
  `instance_id` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_oracle_last_data
-- ----------------------------
INSERT INTO `tb_oracle_last_data` VALUES (1, 0, 240386864, 768871957, 2017690, 381996, 2984, 722177, 1);

SET FOREIGN_KEY_CHECKS = 1;
