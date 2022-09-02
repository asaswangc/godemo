-- 创建资源类别表
DROP TABLE IF EXISTS `admin_cmdb_class`;
CREATE TABLE IF NOT EXISTS `admin_cmdb_class`
(
    `id`         INT(11)      NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `tenant_id`  INT(11)      NOT NULL COMMENT '租户id',
    `is_enabled` VARCHAR(10)  NOT NULL DEFAULT 'true' COMMENT '是否启用',
    `class_name` VARCHAR(10)  NOT NULL COMMENT '资源类名称',
    `comments`   VARCHAR(200) NOT NULL DEFAULT '' COMMENT '备注',
    `created_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_unique` (`class_name`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1001
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

-- 创建资源模型表
DROP TABLE IF EXISTS `admin_cmdb_model`;
CREATE TABLE IF NOT EXISTS `admin_cmdb_model`
(
    `id`            INT(11)      NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `tenant_id`     INT(11)      NOT NULL COMMENT '租户id',
    `is_enabled`    VARCHAR(10)  NOT NULL DEFAULT 'true' COMMENT '是否启用',
    `class_id`      INT(11)      NOT NULL COMMENT '资源类别id',
    `comments`      VARCHAR(200) NOT NULL DEFAULT '' COMMENT '备注',
    `model_name`    VARCHAR(100) NOT NULL COMMENT '资源模型表的名称',
    `model_name_zh` VARCHAR(20)  NOT NULL COMMENT '资源模型表的名称(中文)',
    `created_at`    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT ' 最后更新时间 ',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_unique` (`model_name`) -- 所有表名均不可重复
) ENGINE = InnoDB
  AUTO_INCREMENT = 1001
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

-- 创建资源模型字段表
DROP TABLE IF EXISTS `admin_cmdb_field`;
CREATE TABLE IF NOT EXISTS `admin_cmdb_field`
(
    `id`             INT(11)      NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `tenant_id`      INT(11)      NOT NULL COMMENT '租户id',
    `field_name`     VARCHAR(50)  NOT NULL COMMENT '字段名称',
    `field_name_zh`  VARCHAR(50)  NOT NULL COMMENT '字段名称',
    `field_type`     VARCHAR(10)  NOT NULL COMMENT '字段类型',
    `field_length`   int(11)      NOT NULL COMMENT '字段长度',
    `allow_not_null` VARCHAR(10)  NOT NULL DEFAULT 'true' COMMENT '允许为空',
    `verify_id`      INT(11)      NOT NULL DEFAULT 0 COMMENT '验证器id',
    `model_id`       INT(11)      NOT NULL COMMENT '资源模型表的id',
    `comments`       VARCHAR(200) NOT NULL DEFAULT '' COMMENT '备注',
    `created_at`     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_unique` (`field_name`, `model_id`) -- 同一张资源模型表中不可以出现名字相同的字段
) ENGINE = InnoDB
  AUTO_INCREMENT = 1001
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

-- 创建资源模型字段验证器表
DROP TABLE IF EXISTS `admin_cmdb_verify`;
CREATE TABLE IF NOT EXISTS `admin_cmdb_verify`
(
    `id`          INT(11)      NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `tenant_id`   INT(11)      NOT NULL COMMENT '租户id',
    `is_enabled`  VARCHAR(10)  NOT NULL DEFAULT 'true' COMMENT '是否启用',
    `verify_name` VARCHAR(50)  NOT NULL COMMENT '验证器名称',
    `verify_type` VARCHAR(50)  NOT NULL COMMENT '验证器类型',
    `verify_body` VARCHAR(100) NOT NULL COMMENT '验证器内容',
    `comments`    VARCHAR(200) NOT NULL DEFAULT '' COMMENT '备注',
    `created_at`  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_unique` (`verify_name`) -- 所有验证器名称均不可重复
) ENGINE = InnoDB
  AUTO_INCREMENT = 1001
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

-- 创建操作审计表
DROP TABLE IF EXISTS `admin_cmdb_audit`;
CREATE TABLE IF NOT EXISTS `admin_cmdb_audit`
(
    `id`                    INT(11)      NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `tenant_id`             INT(11)      NOT NULL COMMENT '租户id',
    `operate_type`          varchar(20)  NOT NULL COMMENT '操作动作',
    `operate_after_string`  VARCHAR(500) NOT NULL DEFAULT '' COMMENT '操作后数据',
    `operate_before_string` VARCHAR(500) NOT NULL DEFAULT '' COMMENT '操作前数据',
    `operate_object_name`   VARCHAR(20)  NOT NULL DEFAULT '' COMMENT '操作对象名称',
    `operate_instance_name` VARCHAR(20)  NOT NULL DEFAULT '' COMMENT '操作实例名称',
    `comments`              VARCHAR(200) NOT NULL DEFAULT '' COMMENT '备注(这个字段可以拼装一下字段,拼装成一句话)',
    `created_at`            TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`            TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1001
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

-- 表视图展示【定制列】
DROP TABLE IF EXISTS `admin_user_field`;
CREATE TABLE IF NOT EXISTS `admin_user_field`
(
    `id`         INT(11)     NOT NULL AUTO_INCREMENT COMMENT '自增id',
    `users_id`   INT(11)     NOT NULL COMMENT '用户id',
    `tenant_id`  INT(11)     NOT NULL COMMENT '租户id',
    `model_id`   INT(11)     NOT NULL COMMENT '模型id',
    `field_name` VARCHAR(20) NOT NULL COMMENT '字段id',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_unique` (`users_id`, `tenant_id`, `field_name`, `model_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1001
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;
