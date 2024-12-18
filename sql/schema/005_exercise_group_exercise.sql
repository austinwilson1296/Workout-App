-- +goose Up
CREATE TABLE exercise_group_exercise (
    exercise_id INT REFERENCES exercise(id),
    group_id INT REFERENCES exercise_group(id),
    PRIMARY KEY (exercise_id, group_id)
);


-- +goose Down
DROP TABLE exercise_group_exercise;