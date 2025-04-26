-- name: GetAllItems :many
SELECT *
from list;
-- name: CreateNewItem :one
INSERT INTO list(
    id,
    title,
    description,
    priority,
    created_at,
    updated_at
  )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;
-- name: UpdateItem :exec
UPDATE list(
    title,
    description,
    priority,
    updated_at
  )
SET ($2, $3, $4, $5)
WHERE id = $1