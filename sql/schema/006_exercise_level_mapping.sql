-- +goose Up
CREATE TABLE exercise_level_mapping (
    exercise_id INT REFERENCES exercise(id),
    level_id INT REFERENCES experience_level(id),
    PRIMARY KEY (exercise_id, level_id)
);


-- +goose Down
DROP TABLE exercise_level_mapping;