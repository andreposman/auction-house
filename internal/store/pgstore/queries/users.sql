-- name: CreateUser :one
INSERT INTO users ("user_name", "email", "password_hash", "bio")
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;
