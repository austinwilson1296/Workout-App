-- +goose Up
CREATE TABLE experience_level (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,  
    description TEXT             
);


-- +goose Down
DROP TABLE experience_level;