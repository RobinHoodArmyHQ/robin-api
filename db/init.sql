CREATE DATABASE IF NOT EXISTS rha;

CREATE TABLE events (
    event_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    description TEXT,
    start_time DATETIME,
    event_type INT,
    event_location_id BIGINT,
    min_robins TINYINT,
    max_robins TINYINT,
    created_by VARCHAR(255),
    updated_by VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME
);

CREATE TABLE locations (
    location_id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    latitude DOUBLE,
    longitude DOUBLE
);
