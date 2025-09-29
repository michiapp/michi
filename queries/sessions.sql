-- name: InsertSession :one
INSERT INTO sessions (
  alias
) VALUES (?)
RETURNING *;

-- name: GetSessionByAlias :one
SELECT * FROM sessions
WHERE alias = ?;

-- name: ListSessions :many 
SELECT * FROM sessions;


-- name: UpdateSession :one
UPDATE sessions
SET alias = ?
WHERE id = ?
RETURNING *;

-- name: DeleteSession :exec
DELETE FROM sessions
WHERE id = ?;

-- name: DeleteSessionByAlias :exec
DELETE FROM sessions
WHERE alias = ?;
