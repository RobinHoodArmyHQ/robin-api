CREATE TABLE IF NOT EXISTS `users` (
   `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
   `user_id` VARCHAR(12) UNIQUE,
   `first_name` VARCHAR(255),
   `last_name` VARCHAR(255),
   `avatar_url` TEXT,
   `mobile_number` JSON,
   `email_id` VARCHAR(255),
   `facebook_id` VARCHAR(255),
   `twitter_id` VARCHAR(255),
   `instagram_id` VARCHAR(255),
   `level` JSON,
   `default_city` JSON
);