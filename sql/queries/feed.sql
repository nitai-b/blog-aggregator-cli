-- name: CreateFeed :one
INSERT INTO feeds (id, user_id, name, url, created_at, updated_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT *
  FROM feeds
  ORDER BY feeds.updated_at DESC;

-- name: GetFeedsWithUser :many
SELECT feeds.id, user_id, feeds.name, url, feeds.created_at, feeds.updated_at, u.id AS user_id, u.created_at AS user_created_at, u.updated_at AS user_updated_at, u.name AS user_name
  FROM feeds
  INNER JOIN users u ON feeds.user_id = u.id
  ORDER BY feeds.updated_at DESC;
