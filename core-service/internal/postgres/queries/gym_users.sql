-- name: GetUserById :many
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  username, email, password
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
  set username = $2,
  email = $3,
  enabled = $4
WHERE id = $1;