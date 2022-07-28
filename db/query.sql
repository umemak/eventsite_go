-- name: GetUser :one
SELECT * FROM user WHERE id = ? LIMIT 1;

-- name: CreateEvent :execresult
INSERT INTO event (title, start, place, open, close, author) VALUES (?, ?, ?, ?, ?, ?);

-- name: ListEvents :many
SELECT id, title, start, place, open, close, author FROM event ORDER BY id;

-- name: GetEvent :one
SELECT id, title, start, place, open, close, author FROM event WHERE id = ? LIMIT 1;
