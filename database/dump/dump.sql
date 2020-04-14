create database emem;
use emem;
CREATE TABLE `mems` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `title` varchar(200) DEFAULT NULL,
    `description` text  DEFAULT NULL,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp,
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp on update current_timestamp,
    PRIMARY KEY (`id`),
    UNIQUE KEY `ID_UNIQUE` (`id`)
);

