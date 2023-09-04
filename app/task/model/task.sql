SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `task`;
CREATE TABLE `task`
(
    `id`         bigint(20) UNSIGNED                                     NOT NULL AUTO_INCREMENT COMMENT '主键',
    `created_at` timestamp                                               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp                                               NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime(3)                                             NULL     DEFAULT NULL COMMENT '删除时间',
    `uid`        bigint(20) UNSIGNED                                     NOT NULL COMMENT '用户id',
    `title`      varchar(191) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '任务标题',
    `status`     tinyint(20)                                             NOT NULL DEFAULT 0 COMMENT '任务状态 0代办 1已完成',
    `content`    longtext CHARACTER SET utf8 COLLATE utf8_general_ci     NULL COMMENT '任务内容',
    `start_time` timestamp                                               NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '任务开始时间',
    `end_time`   timestamp                                               NULL     DEFAULT NULL COMMENT '任务结束时间',
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `idx_tasks_deleted_at` (`deleted_at`) USING BTREE,
    INDEX `idx_tasks_title` (`title`) USING BTREE,
    INDEX `idx_tasks_user` (`uid`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  CHARACTER SET = utf8
  COLLATE = utf8_general_ci
  ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
