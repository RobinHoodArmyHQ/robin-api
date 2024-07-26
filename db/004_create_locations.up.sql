CREATE TABLE IF NOT EXISTS `locations` (
    `location_id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255),
    `latitude` DECIMAL(9, 6),
    `longitude` DECIMAL(9, 6)
);
