-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT,
    url TEXT UNIQUE,
    description TEXT,
    published_at TIMESTAMP,
    unread BOOL NOT NULL DEFAULT TRUE,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);

CREATE INDEX idx_posts_feed_unread ON posts(feed_id, unread) WHERE unread = TRUE;

-- +goose Down
DROP TABLE posts;