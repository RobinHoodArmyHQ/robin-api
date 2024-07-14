CREATE DATABASE IF NOT EXISTS rha;

CREATE TABLE events (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    event_id VARCHAR(12) DEFAULT NULL,
    name VARCHAR(255),
    description TEXT,
    start_time DATETIME,
    event_type TINYINT UNSIGNED,
    event_location_id BIGINT UNSIGNED,
    min_robins TINYINT UNSIGNED,
    max_robins TINYINT UNSIGNED,
    created_by BIGINT UNSIGNED,
    created_at DATETIME,
    INDEX (event_id)
);

CREATE TABLE locations (
    location_id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    latitude DECIMAL(9, 6),
    longitude DECIMAL(9, 6)
);
