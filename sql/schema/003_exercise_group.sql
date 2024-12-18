-- +goose Up
CREATE TABLE exercise_group (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,  
    description TEXT              
);


-- +goose Down
DROP TABLE exercise_group;