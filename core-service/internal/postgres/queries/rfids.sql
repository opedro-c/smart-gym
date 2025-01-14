-- name: GetUserIdByRfidId :one
SELECT id
FROM users
WHERE rfid = $1;

-- name: UpdateUserRfid :exec
UPDATE users
  set rfid = $2
WHERE id = $1;
