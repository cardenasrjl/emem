create database emem;
use emem;

CREATE TABLE `mems` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL,
    `volume_id` bigint(20)  NULL,
    `title` varchar(200) DEFAULT NULL,
    `description` text  DEFAULT NULL,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp,
    `updated_at` timestamp NOT NULL DEFAULT current_timestamp on update current_timestamp,
    PRIMARY KEY (`id`),
    UNIQUE KEY `ID_UNIQUE` (`id`)
);

create table volumes
(
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    user_id bigint(20) not null,
    title VARCHAR(50) null,
    created_at TIMESTAMP default current_timestamp not null,
    updated_at TIMESTAMP default current_timestamp not null,
    PRIMARY KEY (`id`),
    UNIQUE KEY `ID_UNIQUE` (`id`)
);

create unique index _mem_id_volume_uindex
    on volumes (user_id, title);

create table users
(
    id bigint(20) auto_increment,
    username VARCHAR(255) not null,
    password VARCHAR(255) null,
    created_at TIMESTAMP default current_timestamp not null,
    updated_at TIMESTAMP not null on UPDATE current_timestamp,
    PRIMARY KEY (`id`),
    UNIQUE KEY `ID_UNIQUE` (`id`)
);

