SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`              bigint(20) UNSIGNED                                     NOT NULL AUTO_INCREMENT COMMENT '主键',
    `created_at`      timestamp                                               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`      timestamp                                               NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at`      datetime(3)                                             NULL     DEFAULT NULL COMMENT '删除时间',
    `username`        varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NULL     DEFAULT NULL COMMENT '用户名',
    `password_digest` longtext CHARACTER SET utf8 COLLATE utf8_general_ci     NULL COMMENT '密码',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `username` (`username`) USING BTREE,
    INDEX `idx_users_deleted_at` (`deleted_at`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci
  ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;