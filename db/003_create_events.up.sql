CREATE TABLE IF NOT EXISTS `events` (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `event_id` VARCHAR(12),
    `name` VARCHAR(255),
    `description` TEXT,
    `start_time` DATETIME,
    `event_type` TINYINT UNSIGNED,
    `event_location_id` BIGINT UNSIGNED,
    `min_robins` TINYINT UNSIGNED,
    `max_robins` TINYINT UNSIGNED,
    `created_by` BIGINT UNSIGNED,
    `created_at` DATETIME,
    UNIQUE (`event_id`)
);