-- name: InsertShortcut :one
INSERT INTO shortcuts (
  alias, url
) VALUES (?, ?)
RETURNING *;

-- name: GetShortcutByAlias :one
SELECT * FROM shortcuts
WHERE alias = ?;

-- name: ListShortcuts :many
SELECT * FROM shortcuts;

-- name: DeleteShortcut :exec
DELETE FROM shortcuts
WHERE id = ?;

-- name: DeleteShortcutFromAlias :exec
DELETE FROM shortcuts
WHERE alias = ?;
