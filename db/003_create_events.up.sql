CREATE TABLE IF NOT EXISTS `events` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `event_id` varchar(15) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `start_time` datetime DEFAULT NULL,
  `event_type` enum('MEAL_DRIVE','ACADEMY') CHARACTER SET ascii COLLATE ascii_general_ci DEFAULT NULL,
  `location_id` bigint unsigned DEFAULT NULL,
  `min_robins` tinyint unsigned DEFAULT NULL,
  `max_robins` tinyint unsigned DEFAULT NULL,
  `created_by` bigint unsigned DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `event_id` (`event_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
