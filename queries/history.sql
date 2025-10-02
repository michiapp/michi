
-- name: InsertHistoryEntry :exec
INSERT INTO history (query, provider_id, provider_tag) VALUES (?, ?, ?);

-- name: GetRecentHistory :many
SELECT * FROM history ORDER BY timestamp DESC LIMIT ?;

-- name: ListHistory :many
SELECT * FROM history ORDER BY timestamp DESC;

-- name: DeleteHistoryEntry :exec
DELETE FROM history WHERE id = ?;
