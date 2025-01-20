-- +goose Up
ALTER TABLE exercise
    ADD COLUMN exclude_from_cooldown BOOLEAN;


-- +goose Down
ALTER TABLE users
    DROP COLUMN exclude_from_cooldown