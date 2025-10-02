-- name: AddSessionUrl :one
INSERT INTO session_urls (session_id, url) VALUES (?, ?) RETURNING *;

-- name: ListSessionUrls :many
SELECT * FROM session_urls WHERE session_id = ? ORDER BY created_at ASC;

-- name: DeleteSessionUrls :exec
DELETE FROM session_urls WHERE session_id = ?;
