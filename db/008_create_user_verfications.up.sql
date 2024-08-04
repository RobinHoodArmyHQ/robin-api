CREATE TABLE IF NOT EXISTS `user_verifications` (
    `id` BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    `user_id` VARCHAR(12) UNIQUE NOT NULL,
    `email_id` VARCHAR(255) NOT NULL,
    `otp` INT UNSIGNED,
    `otp_generated_at` DATETIME,
    `otp_expires_at` DATETIME,
    `otp_retry_count` INT UNSIGNED,
    `is_verified` TINYINT(1) DEFAULT 0,
    `extra_details` JSON
)