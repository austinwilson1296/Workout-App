-- +goose Up
CREATE TABLE exercise_muscle_group (
    exercise_id INT NOT NULL REFERENCES exercise(id) ON DELETE CASCADE,
    muscle_group_id INT NOT NULL REFERENCES muscle_group(id) ON DELETE CASCADE,
    PRIMARY KEY (exercise_id, muscle_group_id) -- Composite primary key
);

-- +goose Down
DROP TABLE exercise_muscle_group;