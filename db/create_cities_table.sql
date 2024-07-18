CREATE TABLE IF NOT EXISTS `cities` (
    `id` MEDIUMINT UNSIGNED NOT NULL AUTO_INCREMENT ,
    `city_id` VARCHAR(15) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL ,
    `name` VARCHAR(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL ,
    `state` VARCHAR(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL ,
    `country_id` TINYINT UNSIGNED NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE `idx_city_id` (`city_id`)
);

CREATE TABLE IF NOT EXISTS `countries` (
    `id` TINYINT UNSIGNED NOT NULL AUTO_INCREMENT ,
    `country_id` VARCHAR(15) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL ,
    `name` VARCHAR(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL ,
    PRIMARY KEY (`id`)
);