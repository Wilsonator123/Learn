-- name: GetAllTasks :many
SELECT *
from task;
-- name: CreateNewTask :one
INSERT INTO task(
    id,
    title,
    position,
    description,
    created_at,
    updated_at
  )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: UpdateTask :exec
UPDATE task
SET title = COALESCE($2, title),
  description = COALESCE($3, description),
  position = COALESCE($4, position),
  updated_at = $5
WHERE id = $1;
-- name: GetTask :one
SELECT *
from task
WHERE id = $1;
-- name: DeleteTask :exec
DELETE FROM task
WHERE id = $1;