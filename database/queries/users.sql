-- this is a sqlite3 queries file

-- name: insertUser :exec
INSERT INTO users (username, password, salt, created_at)
VALUES (?, ?, ?, ?);

-- name: updateUser :exec
UPDATE users SET username = ?, password = ?, salt = ?, created_at = ? WHERE username = ?;

-- name: selectUserByUsername :one
SELECT * FROM users WHERE username = ?;

-- name: deleteUserByUsername :exec
DELETE FROM users WHERE username = ?;
