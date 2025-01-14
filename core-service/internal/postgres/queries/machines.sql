-- name: CreateMachine :one
INSERT INTO machines (
    name, origin_id
) VALUES (
    $1, $2
)
RETURNING *;

-- name: GetMachines :many
SELECT *
FROM machines;

-- name: DeleteMachine :exec
DELETE FROM machines
WHERE id = $1;

-- name: UpdateMachine :one
UPDATE machines
    set name = $2,
    origin_id = $3
WHERE id = $1
RETURNING *;