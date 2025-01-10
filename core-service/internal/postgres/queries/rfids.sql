-- name: CreateRfid :one
INSERT INTO rfids (
    user_id, card_id
) VALUES (
    $1, $2
)
RETURNING *;

-- name: GetRfidsByUserId :many
SELECT *
FROM rfids
WHERE user_id = $1;

-- name: GetUserIdByRfidId :one
SELECT user_id
FROM rfids
WHERE id = $1;

-- name: DeleteRfids :exec
DELETE FROM rfids
WHERE id = ANY($1::int[]) AND user_id = $2;
