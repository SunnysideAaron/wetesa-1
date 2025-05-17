-- name: GetClient :one
SELECT * FROM client
WHERE client_id = $1 LIMIT 1;

-- name: ListClients :many
SELECT * FROM client
ORDER BY name;

-- name: CreateClient :one
INSERT INTO client (
  name, address
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateClient :exec
UPDATE client
  set name = $2,
  address = $3
WHERE client_id = $1;

-- name: DeleteClient :exec
DELETE FROM client
WHERE client_id = $1;