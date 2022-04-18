CREATE TABLE IF NOT EXISTS todos (
     id VARCHAR(36) NOT NULL PRIMARY KEY,
     title VARCHAR(255),
     body TEXT,
     completed boolean DEFAULT FALSE,
     created_at TIMESTAMP,
     updated_at TIMESTAMP,
     deleted_at TIMESTAMP
);