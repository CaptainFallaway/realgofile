-- name: insertUser :exec
INSERT INTO users (uid, username, password, salt)
VALUES (?, ?, ?, ?);

-- name: updateUser :exec
UPDATE users SET username = ?, password = ?, salt = ? WHERE uid = ?;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = ?;

-- name: GetUserByUid :one
SELECT * FROM users WHERE uid = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE uid = ?;
