CREATE TABLE `users` (
    `user_id` varchar(255) PRIMARY KEY,
    `username` varchar(255) NOT NULL,
    `hashed_passward` varchar(255) NOT NULL,
    `campus` varchar(255) NOT NULL,
    `email` varchar(255) UNIQUE NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `accounts` (
  `account_id` bigint PRIMARY KEY,
  `owner` varchar(255) NOT NULL,
  `blockchain_id` varchar(255) UNIQUE NOT NULL,
  `token_type` varchar(255) NOT NULL,
  `balance` float NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `items` (
  `item_id` bigint PRIMARY KEY,
  `item_name` varchar(255) NOT NULL,
  `collection_id` bigint NOT NULL,
  `image` varchar(255) NOT NULL,
  `owner` varchar(255) NOT NULL,
  `creater` varchar(255) NOT NULL,
  `token_uri` varchar(255) UNIQUE NOT NULL,
  `favorites` bigint NOT NULL,
  `view_times` bigint NOT NULL,
  `about` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `collections` (
  `collection_id` bigint PRIMARY KEY,
  `collection_name` varchar(255) NOT NULL,
  `creater` varchar(255) NOT NULL,
  `image` varchar(255) NOT NULL,
  `items` bigint NOT NULL,
  `owners` bigint NOT NULL,
  `about` varchar(255) NOT NULL,
  `contract_address` varchar(255) UNIQUE NOT NULL,
  `token_standard` varchar(255) NOT NULL,
  `blockchain` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `favorites` (
  `favorites_id` bigint PRIMARY KEY,
  `user_id` varchar(255) NOT NULL,
  `item_id` bigint NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `events` (
  `event_id` bigint PRIMARY KEY,
  `event_name` varchar(255) NOT NULL,
  `event_description` varchar(255) NOT NULL,
  `creater` varchar(255) NOT NULL,
  `start_time` timestamp NOT NULL,
  `end_time` timestamp NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE `accounts` ADD FOREIGN KEY (`owner`) REFERENCES `users` (`user_id`);

ALTER TABLE `items` ADD FOREIGN KEY (`collection_id`) REFERENCES `collections` (`collection_id`);

ALTER TABLE `items` ADD FOREIGN KEY (`owner`) REFERENCES `users` (`user_id`);

ALTER TABLE `items` ADD FOREIGN KEY (`creater`) REFERENCES `users` (`user_id`);

ALTER TABLE `collections` ADD FOREIGN KEY (`creater`) REFERENCES `users` (`user_id`);

ALTER TABLE `favorites` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`);

ALTER TABLE `favorites` ADD FOREIGN KEY (`item_id`) REFERENCES `items` (`item_id`);

ALTER TABLE `events` ADD FOREIGN KEY (`creater`) REFERENCES `users` (`user_id`);