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
UPDATE list
SET title = $2,
  description = $3,
  priority = $4,
  updated_at = $5
WHERE id = $1;
-- name: GetItem :one
SELECT *
from list
WHERE id = $1;