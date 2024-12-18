-- name: GetExercisesPerCategoryByExperience :many
WITH RankedExercises AS (
    SELECT 
        e.id AS exercise_id,
        e.name AS exercise_name,
        eg.name AS category_name,
        RANDOM() AS random_sort
    FROM 
        exercise e
    JOIN 
        exercise_group_exercise ege ON e.id = ege.exercise_id
    JOIN 
        exercise_group eg ON ege.group_id = eg.id
    JOIN 
        exercise_level_mapping elm ON e.id = elm.exercise_id
    WHERE 
        eg.name IN ('core_hips_legs', 'core_spinal', 'thoracic_spine_mobility', 'scapulo_thoracic', 'shoulders_scapula')
        AND elm.level_id = $1
    ORDER BY 
        random_sort
    LIMIT $2 * (SELECT COUNT(DISTINCT name) FROM exercise_group WHERE name IN ('core_hips_legs', 'core_spinal', 'thoracic_spine_mobility', 'scapulo_thoracic', 'shoulders_scapula'))
)
SELECT 
    exercise_id,
    exercise_name,
    category_name
FROM 
    RankedExercises
GROUP BY 
    category_name, 
    exercise_id, 
    exercise_name
LIMIT $2;
