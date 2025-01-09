-- name: GetAdminUserById :many
SELECT * FROM admin_users
WHERE id = $1 LIMIT 1;

-- name: CreateAdminUser :one
INSERT INTO admin_users (
  username, email, password
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateAdminUser :exec
UPDATE admin_users
  set username = $2,
  email = $3
WHERE id = $1;