
CREATE TABLE `login`(
                        `email` VARCHAR(255) NOT NULL unique,
                        `passwd` VARCHAR(255) NOT NULL,

                        `emailToken` VARCHAR(255),
                        `genTime` DATETIME,
                        `activated` VARCHAR(20),
                        `verify_code` VARCHAR(255),
                        `codeTime` DATETIME,

                        `user_id` VARCHAR(255) unique,
                        PRIMARY KEY (`email`)
);

CREATE TABLE `accounts` (
                            `user_id` varchar(255) PRIMARY KEY,
                            `email` varchar(255),
                            `user_name` VARCHAR(255),
                            `banner_image` VARCHAR(255),
                            `avatar_image` VARCHAR(255),
                            `poison` VARCHAR(255),
                            `organization` VARCHAR(255),
                            `token` INT,
                            `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `collection` (
                              `collection_id` INT AUTO_INCREMENT PRIMARY KEY,
                              `collection_name` varchar(255),
                              `logo_image` varchar(255),
                              `feature_image` varchar(255),
                              `banner_image` varchar(255),
                              `items_count` int,
                              `description` varchar(255),
                              `owner` varchar(255),
                              `creater` varchar(255),
                              `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `collection_label` (
                                    `label_id` bigint NOT NULL AUTO_INCREMENT,
                                    `collection_id` int NOT NULL,
                                    `label` varchar(255) NOT NULL,
                                    PRIMARY KEY (`label_id`),
                                    UNIQUE KEY `item_label` (`item_id`, `label`)
);

CREATE TABLE `items` (
                         `item_id` varchar(255) PRIMARY KEY,
                         `item_name` varchar(255) NOT NULL,
                         `collection_id` bigint NOT NULL,
                         `item_data` varchar(255) NOT NULL,
                         `description` varchar(255) NOT NULL,
                         `owner_id` varchar(255) NOT NULL,
                         `creater_id` varchar(255) NOT NULL,
                         `category`varchar(255) NOT NULL,
                         `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `item_label` (
                              `label_id` bigint NOT NULL AUTO_INCREMENT,
                              `item_id` varchar(255) NOT NULL,
                              `label` varchar(255) NOT NULL,
                              PRIMARY KEY (`label_id`),
                              UNIQUE KEY `item_label` (`item_id`, `label`)
);

CREATE TABLE `item_history` (
                                `item_id` varchar(255) NOT NULL,
                                `from`    varchar(255) NOT NULL,
                                `to`      varchar(255) NOT NULL,
                                `operation` varchar(255) NOT NULL,
                                `time` timestamp NOT NULL
);