CREATE TABLE IF NOT EXISTS `locations` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `latitude` decimal(9,6) DEFAULT NULL,
  `longitude` decimal(9,6) DEFAULT NULL,
  `google_place_id` varchar(50) CHARACTER SET ascii COLLATE ascii_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;