-- name: GetAllTabs :many
SELECT *
FROM tab;
-- name: GetTab :one
SELECT *
FROM tab
where id = $1;
-- name: CreateNewTab :one
INSERT INTO tab(
    id,
    title,
    layout,
    created_at,
    updated_at
  )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: UpdateTab :exec
UPDATE tab
SET title = COALESCE($2, title),
  layout = COALESCE($3, layout),
  updated_at = $4
WHERE id = $1;