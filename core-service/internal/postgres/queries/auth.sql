-- name: GetUserByEmailAndPassword :one
SELECT * FROM users
WHERE email = $1
    AND password = $2;

-- name: GetAdminByEmailAndPassword :one
SELECT * FROM admin_users
WHERE email = $1
    AND password = $2;