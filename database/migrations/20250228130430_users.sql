-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    uid TEXT NOT NULL PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    password BLOB NOT NULL,
    salt BLOB NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
