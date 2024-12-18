-- +goose Up
ALTER TABLE users
    ADD COLUMN username VARCHAR(255) NOT NULL UNIQUE,
    ADD COLUMN password_hash TEXT NOT NULL,
    ADD COLUMN experience_level_id INT REFERENCES experience_level(id);


-- +goose Down
ALTER TABLE users
    DROP COLUMN username,
    DROP COLUMN password_hash,
    DROP COLUMN experience_level_id;

