-- name: CreateFeedFollow :one
INSERT INTO feed_follows (
 id, created_at, updated_at, user_id, feed_id 
) VALUES (
  $1, $2, $3, $4, $5
  )
RETURNING 
  *,
  (select name as user_name from users where users.id = $4),
  (select name as feed_name from feeds where feeds.id = $5)
;

-- name: GetFeedFollowsForUser :many
SELECT * FROM feed_follows
where user_id = $1;
