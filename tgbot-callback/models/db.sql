create database callback_info;
use callback_info;
CREATE TABLE `tg_callback_acl` (
	`id` bigint(20) NOT NULL AUTO_INCREMENT,
	`bot_name`  varchar(254) COLLATE utf8mb4_general_ci DEFAULT '' NOT NULL,
	`tg_name` varchar(254) COLLATE utf8mb4_general_ci DEFAULT '' NOT NULL,
	`tg_username` varchar(254) COLLATE utf8mb4_general_ci DEFAULT '' NOT NULL,
	`status` bigint(20) NOT NULL DEFAULT '0' NOT NULL,
	`create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	`update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`delete_time` timestamp NULL ,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


CREATE TABLE `project_groups` (
	`id` bigint(20) NOT NULL AUTO_INCREMENT,
	`job_name` varchar(254) COLLATE utf8mb4_general_ci DEFAULT '' NOT NULL,
	`project_staff` JSON,
	`status` bigint(20) NOT NULL DEFAULT '0' NOT NULL,
	`remark`  varchar(254) COLLATE utf8mb4_general_ci DEFAULT '' NOT NULL,
	`cascades` varchar(254) COLLATE utf8mb4_general_ci DEFAULT '' NOT NULL,
	`create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	`update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`delete_time` timestamp NULL ,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
-- INSERT INTO callback_info.project_groups ( job_name, project_staff, STATUS, remark, cascades ) VALUES( "testA-background", '["Mark","nike"]', 1, "这是一个测试", "Online" );
-- INSERT INTO callback_info.project_groups ( job_name, project_staff, STATUS, remark, cascades ) VALUES( "测试", '["Mark","nike"]', 1, "这是一个测试", "Online" );