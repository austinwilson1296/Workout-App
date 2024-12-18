-- +goose Up
CREATE TABLE muscle_group (
    id SERIAL PRIMARY KEY,         
    name VARCHAR(255) NOT NULL UNIQUE,  
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE muscle_group;