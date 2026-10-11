-- +goose Up
CREATE TABLE feeds(
    id TEXT PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url TEXT UNIQUE,
    user_id TEXT NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(id)  ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;
