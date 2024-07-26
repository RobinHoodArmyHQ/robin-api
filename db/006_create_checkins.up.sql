CREATE TABLE IF NOT EXISTS `checkins` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `checkin_id` VARCHAR(12) UNIQUE,
    `user_id` VARCHAR(12),
    `event_id` VARCHAR(12),
    `photo_urls` JSON,
    `description` TEXT,
    `no_of_people_served` BIGINT,
    `no_of_student_taught` BIGINT,
    `created_at` DATETIME,
    `updated_at` DATETIME,
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_event_id` (`event_id`)
);