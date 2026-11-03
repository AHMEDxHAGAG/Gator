-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url)
VALUES(
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: GetFeedFromURL :one
SELECT * FROM feeds
where url = $1;

-- name: GetFeedFromID :one
SELECT * FROM feeds
where id = $1;

-- name: ResetFeeds :exec
DELETE FROM feeds;