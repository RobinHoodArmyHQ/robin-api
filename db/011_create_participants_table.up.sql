CREATE TABLE IF NOT EXISTS `participants` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `event_id` bigint unsigned NOT NULL,
    `user_id` bigint unsigned NOT NULL,
    `status` varchar(20) CHARACTER SET ascii COLLATE ascii_general_ci DEFAULT 'INTERESTED',
    `role` varchar(20) CHARACTER SET ascii COLLATE ascii_general_ci DEFAULT 'VOLUNTEER',
    `created_at` datetime DEFAULT NULL,
    `updated_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `event_id` (`event_id`),
    KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
