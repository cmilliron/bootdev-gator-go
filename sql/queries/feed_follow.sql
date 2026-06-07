-- name: CreateFollowFeed :one
WITH inserted_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, feed_id, user_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT 
    inf.id,
    inf.created_at,
    inf.updated_at,
    inf.feed_id,
    inf.user_id,
    u.name AS username,
    f.name AS feed_name,
    f.url AS feed_url
FROM inserted_follow inf
INNER JOIN users u ON inf.user_id = u.id
INNER JOIN feeds f ON inf.feed_id = f.id;

-- name: GetFeedFollowsForUser :many
SELECT 
    ff.*,
    u.name AS username,
    f.name AS feed_name,
    f.url AS feed_url
FROM feed_follows ff
INNER JOIN users u ON ff.user_id = u.id
INNER JOIN feeds f ON ff.feed_id = f.id
WHERE ff.user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows 
WHERE feed_id = $1 AND user_id = $2;

