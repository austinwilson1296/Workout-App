-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, username, password_hash, experience_level_id)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: GetUserByUsername :one
SELECT * 
FROM users
WHERE username = $1;