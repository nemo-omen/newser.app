-- +goose Up
CREATE TABLE IF NOT EXISTS
    sessions (
        token  TEXT PRIMARY KEY,
        data   BYTEA NOT NULL,
        expiry TIMESTAMPTZ NOT NULL
    );
-- +goose Down
DROP TABLE IF EXISTS sessions;
