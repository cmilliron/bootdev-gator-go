-- name: CreatePost :one
INSERT INTO posts (
    id,
    created_at,
    updated_at,
    title,
    url,
    description,
    published_at,
    unread,
    feed_id
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
ON CONFLICT (url) DO NOTHING
RETURNING *;

-- name: GetPostForUser :many
SELECT posts.* 
FROM posts
INNER JOIN feed_follows ff 
ON posts.feed_id = ff.feed_id
WHERE ff.user_id = $1
ORDER BY posts.published_at DESC NULLS LAST
LIMIT $2 OFFSET $3;
