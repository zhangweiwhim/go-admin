CREATE TABLE `tb_teach_all` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `tech_name` varchar(128) DEFAULT NULL COMMENT '名称',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `grade` varchar(100) DEFAULT NULL,
  `class` varchar(100) DEFAULT NULL,
  `subject` varchar(100) DEFAULT NULL,
  `tech_no` varchar(100) DEFAULT NULL,
  `tech_term` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_tb_teach_all_deleted_at` (`deleted_at`),
  KEY `idx_tb_teach_all_create_by` (`create_by`),
  KEY `idx_tb_teach_all_update_by` (`update_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



-- goadmin.tb_student definition

CREATE TABLE `tb_student` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `name` varchar(128) DEFAULT NULL COMMENT '名称',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',

  `tech_no` varchar(100) DEFAULT NULL,
  `tec_now_sub` varchar(100) DEFAULT NULL,
  `tec_now_grade` varchar(100) DEFAULT NULL,
  `tec_now_class` varchar(100) DEFAULT NULL,

  `sex` int DEFAULT NULL,
  `age` int DEFAULT NULL,
  `id_no` varchar(100) DEFAULT NULL,
  `address` varchar(100) DEFAULT NULL,
  `tel` varchar(100) DEFAULT NULL,
  `parents_name` varchar(100) DEFAULT NULL,
  `parents_tel` varchar(100) DEFAULT NULL,
  `in_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_tb_student_deleted_at` (`deleted_at`),
  KEY `idx_tb_student_create_by` (`create_by`),
  KEY `idx_tb_student_update_by` (`update_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



CREATE TABLE `tb_teacher` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
  `name` varchar(128) DEFAULT NULL COMMENT '名称',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  `create_by` bigint DEFAULT NULL COMMENT '创建者',
  `update_by` bigint DEFAULT NULL COMMENT '更新者',
  `grade` varchar(100) DEFAULT NULL,
  `class` varchar(100) DEFAULT NULL,
  `head_teacher` varchar(100) DEFAULT NULL,
  `stu_no` varchar(100) DEFAULT NULL,
  `sex` int DEFAULT NULL,
  `in_time` timestamp NULL DEFAULT NULL,
  `age` int DEFAULT NULL,
  `id_no` varchar(100) DEFAULT NULL,
  `address` varchar(100) DEFAULT NULL,
  `tel` varchar(100) DEFAULT NULL,
  `honorary` varchar(100) DEFAULT NULL,
  `resume` varchar(100) DEFAULT NULL,
  `other_tel` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_tb_teacher_deleted_at` (`deleted_at`),
  KEY `idx_tb_teacher_create_by` (`create_by`),
  KEY `idx_tb_teacher_update_by` (`update_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;




CREATE TABLE `tb_exam` (
                           `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
                           `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
                           `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
                           `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
                           `create_by` bigint DEFAULT NULL COMMENT '创建者',
                           `update_by` bigint DEFAULT NULL COMMENT '更新者',

                           `exam_grade` varchar(100) DEFAULT NULL,
                           `exam_class` varchar(100) DEFAULT NULL,
                           `exam_start` varchar(100) DEFAULT NULL,
                           `exam_end` varchar(100) DEFAULT NULL,

                           `exam_type` int DEFAULT NULL,
                           `exam_sub` int DEFAULT NULL,
                           `term_name` varchar(100) DEFAULT NULL,
                           PRIMARY KEY (`id`),
                           KEY `idx_tb_exam_deleted_at` (`deleted_at`),
                           KEY `idx_tb_exam_create_by` (`create_by`),
                           KEY `idx_tb_exam_update_by` (`update_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `tb_sub` (
                          `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
                          `name` varchar(128) DEFAULT NULL COMMENT '名称',
                          `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
                          `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
                          `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
                          `create_by` bigint DEFAULT NULL COMMENT '创建者',
                          `update_by` bigint DEFAULT NULL COMMENT '更新者',

                          PRIMARY KEY (`id`),
                          KEY `idx_tb_sub_deleted_at` (`deleted_at`),
                          KEY `idx_tb_sub_create_by` (`create_by`),
                          KEY `idx_tb_sub_update_by` (`update_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `tb_term` (
                           `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键编码',
                           `name` varchar(128) DEFAULT NULL COMMENT '名称',
                           `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
                           `updated_at` datetime(3) DEFAULT NULL COMMENT '最后更新时间',
                           `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
                           `create_by` bigint DEFAULT NULL COMMENT '创建者',
                           `update_by` bigint DEFAULT NULL COMMENT '更新者',
                           `start_time` varchar(128) DEFAULT NULL COMMENT '开始时间',
                           `end_time` varchar(128) DEFAULT NULL COMMENT '结束时间',
                           PRIMARY KEY (`id`),
                           KEY `idx_tb_term_deleted_at` (`deleted_at`),
                           KEY `idx_tb_term_create_by` (`create_by`),
                           KEY `idx_tb_term_update_by` (`update_by`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

