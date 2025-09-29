-- name: InsertProvider :exec
INSERT INTO search_providers (
    tag, url, category, domain, rank, site_name, subcategory
) VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetProviderByTag :one
SELECT * FROM search_providers
WHERE tag = ?;

-- name: ListProviders :many
SELECT * FROM search_providers
ORDER BY rank DESC;

-- name: DeleteProvider :exec
DELETE FROM search_providers
WHERE id = ?;
