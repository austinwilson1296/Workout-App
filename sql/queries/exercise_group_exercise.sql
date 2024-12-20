-- name: GetCoreHipsLegsExercises :many
SELECT 
    e.id AS exercise_id,
    e.name AS exercise_name,
    eg.name AS category_name
FROM 
    exercise e
JOIN 
    exercise_group_exercise ege ON e.id = ege.exercise_id
JOIN 
    exercise_group eg ON ege.group_id = eg.id
JOIN 
    exercise_level_mapping elm ON e.id = elm.exercise_id
WHERE 
    eg.name = 'core_hips_legs'
    AND elm.level_id = $1
ORDER BY RANDOM()
LIMIT $2;

-- name: GetCoreSpinalExercises :many
SELECT 
    e.id AS exercise_id,
    e.name AS exercise_name,
    eg.name AS category_name
FROM 
    exercise e
JOIN 
    exercise_group_exercise ege ON e.id = ege.exercise_id
JOIN 
    exercise_group eg ON ege.group_id = eg.id
JOIN 
    exercise_level_mapping elm ON e.id = elm.exercise_id
WHERE 
    eg.name = 'core_spinal'
    AND elm.level_id = $1
ORDER BY RANDOM()
LIMIT $2;

-- name: GetThoracicSpineMobilityExercises :many
SELECT 
    e.id AS exercise_id,
    e.name AS exercise_name,
    eg.name AS category_name
FROM 
    exercise e
JOIN 
    exercise_group_exercise ege ON e.id = ege.exercise_id
JOIN 
    exercise_group eg ON ege.group_id = eg.id
JOIN 
    exercise_level_mapping elm ON e.id = elm.exercise_id
WHERE 
    eg.name = 'thoracic_spine_mobility'
    AND elm.level_id = $1
ORDER BY RANDOM()
LIMIT $2;

-- name: GetScapuloThoracicExercises :many
SELECT 
    e.id AS exercise_id,
    e.name AS exercise_name,
    eg.name AS category_name
FROM 
    exercise e
JOIN 
    exercise_group_exercise ege ON e.id = ege.exercise_id
JOIN 
    exercise_group eg ON ege.group_id = eg.id
JOIN 
    exercise_level_mapping elm ON e.id = elm.exercise_id
WHERE 
    eg.name = 'scapulo_thoracic'
    AND elm.level_id = $1
ORDER BY RANDOM()
LIMIT $2;

-- name: GetShouldersScapulaExercises :many
SELECT 
    e.id AS exercise_id,
    e.name AS exercise_name,
    eg.name AS category_name
FROM 
    exercise e
JOIN 
    exercise_group_exercise ege ON e.id = ege.exercise_id
JOIN 
    exercise_group eg ON ege.group_id = eg.id
JOIN 
    exercise_level_mapping elm ON e.id = elm.exercise_id
WHERE 
    eg.name = 'shoulders_scapula'
    AND elm.level_id = $1
ORDER BY RANDOM()
LIMIT $2;
