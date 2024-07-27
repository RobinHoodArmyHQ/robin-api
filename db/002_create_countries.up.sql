CREATE TABLE IF NOT EXISTS `countries` (
    `id` TINYINT UNSIGNED NOT NULL AUTO_INCREMENT ,
    `country_id` VARCHAR(15) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL ,
    `name` VARCHAR(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL ,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;