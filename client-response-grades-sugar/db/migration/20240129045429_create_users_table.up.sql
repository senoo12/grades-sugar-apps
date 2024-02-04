CREATE TABLE users
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    phone_number VARCHAR(13) NOT NULL UNIQUE,
    role ENUM('super user', 'admin', 'user') DEFAULT 'user' NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at timestamp NOT NULL DEFAULT (now()),
    updated_at timestamp NOT NULL DEFAULT (now())
);