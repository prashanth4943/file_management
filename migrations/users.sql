-- CREATE TABLE IF NOT EXISTS users (
--     id INT AUTO_INCREMENT PRIMARY KEY,
--     first_name VARCHAR(255) ,
--     last_name VARCHAR(255) ,
--     email VARCHAR(255) UNIQUE NOT NULL,
--     phone VARCHAR(15) ,
--     image_name VARCHAR(255) ,
--     uuid VARCHAR(55) ,
--     dl_url TEXT ,  -- Store the CDN URL
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
-- );


CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,             -- Unique identifier
    first_name VARCHAR(255),                       -- First name of the user
    last_name VARCHAR(255),                        -- Last name of the user
    email VARCHAR(255) UNIQUE NOT NULL,            -- Unique email address
    phone VARCHAR(15),                             -- Phone number
    image_name VARCHAR(255),                       -- Name of the uploaded image
    uuid VARCHAR(55),                              -- UUID for user identification
    dl_url TEXT,                                   -- Store the CDN URL of the image
    expiry_date  DATETIME,                          -- Date when the record expires
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Record creation time (only set once)
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- Update time (automatically updated on modification)
);
