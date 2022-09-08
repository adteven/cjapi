
-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `avatar` varchar(200) DEFAULT NULL COMMENT '头像',
    `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '邮箱',
    `enabled` tinyint(1) DEFAULT NULL COMMENT '状态：1启用、0禁用',
    `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '密码',
    `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '用户名',
    `dept_id` bigint DEFAULT NULL COMMENT '部门名称',
    `phone` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '手机号码',
    `job_id` bigint DEFAULT NULL COMMENT '岗位名称',
    `create_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建日期',
    `nick_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
    `sex` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
    `update_time` datetime DEFAULT CURRENT_TIMESTAMP,
    `is_del` tinyint(1) DEFAULT '0',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `username` (`username`) USING BTREE,
    KEY `FK5rwmryny6jthaaxkogownknqp` (`dept_id`) USING BTREE,
    KEY `FKfftoc2abhot8f2wu6cl9a5iky` (`job_id`) USING BTREE,
    KEY `FKpq2dhypk2qgt68nauh2by22jb` (`avatar`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10033 DEFAULT CHARSET=utf8mb3 ROW_FORMAT=COMPACT COMMENT='系统用户';