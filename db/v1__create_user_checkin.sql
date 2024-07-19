
CREATE TABLE users (
   id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
   user_id VARCHAR(12) UNIQUE,
   first_name VARCHAR(255),
   last_name VARCHAR(255),
   avatar_url TEXT,
   mobile_number JSON,
   email_id VARCHAR(255),
   facebook_id VARCHAR(255),
   twitter_id VARCHAR(255),
   instagram_id VARCHAR(255),
   level JSON,
   default_city JSON
);

CREATE TABLE check_ins (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    check_in_id VARCHAR(12) UNIQUE,
    user_id VARCHAR(12),
    event_id VARCHAR(12),
    photo_urls JSON,
    description TEXT,
    no_of_people_served BIGINT,
    no_of_student_taught BIGINT,
    created_at DATETIME,
    updated_at DATETIME,
    INDEX idx_user_id (user_id),
    INDEX idx_event_id (event_id)
);
