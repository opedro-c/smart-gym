-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: ExistsUserById :one
SELECT EXISTS(SELECT 1 FROM users WHERE id = $1 LIMIT 1);

-- name: ExistsUserByEmail :one
SELECT EXISTS(SELECT 1 FROM users WHERE email = $1 LIMIT 1);


-- name: CreateUser :one
INSERT INTO users (
  username, email, password
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateUserData :exec
UPDATE users
  set username = $2,
  email = $3
WHERE id = $1;

-- name: UpdateUserPassword :exec
UPDATE users
  set password = $2
WHERE id = $1;
