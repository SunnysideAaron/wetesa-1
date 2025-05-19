-- name: GetClient :one
SELECT * FROM client
WHERE client_id = $1 LIMIT 1;


-- name: ListClients :many
SELECT * FROM client
WHERE (name ILIKE sqlc.narg(name) OR sqlc.narg(name) IS NULL)
ORDER BY name ASC
LIMIT sqlc.arg(lim)
OFFSET sqlc.arg(off);



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