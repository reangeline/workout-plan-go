-- name: FindUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: CreateUser :exec
INSERT INTO users (id_user, full_name, email, profile_picture) 
VALUES ($1,$2,$3,$4)
RETURNING *;