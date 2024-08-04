CREATE TABLE IF NOT EXISTS `photo_uploads` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `s3_prefix` varchar(500) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL,
  `s3_key` varchar(500) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL,
  `user_id` bigint unsigned NOT NULL,
  `event_id` bigint unsigned DEFAULT NULL,
  `created_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
