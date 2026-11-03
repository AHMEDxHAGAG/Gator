-- +goose Up
CREATE TABLE feeds(
    id TEXT PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE feeds;
