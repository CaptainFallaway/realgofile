-- this is a sqlite3 queries file

-- name: insertUser :exec
INSERT INTO users (username, password, salt)
VALUES (?, ?, ?);

-- name: updateUser :exec
UPDATE users SET username = ?, password = ?, salt = ? WHERE username = ?;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE username = ?;
