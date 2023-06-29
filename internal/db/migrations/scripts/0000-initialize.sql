-- +migrate Up

CREATE TABLE IF NOT EXISTS books (
    id BIGINT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    publication_year INT,
    price BIGINT,
    description TEXT,
    type ENUM('Fiction', 'Non-fiction', 'Sci-fi', 'Mystery', 'Thriller'),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL
);

-- +migrate Down

DROP TABLE IF EXISTS books;