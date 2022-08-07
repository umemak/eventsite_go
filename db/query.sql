-- name: CreateEvent :execresult
INSERT INTO events (title, start, place, open, close, author) VALUES (?, ?, ?, ?, ?, ?);

-- name: ListEvents :many
SELECT id, title, start, place, open, close, author FROM events ORDER BY id;

-- name: GetEvent :one
SELECT id, title, start, place, open, close, author FROM events WHERE id = ? LIMIT 1;

-- name: CreateEventUser :execresult
INSERT INTO events_users (event_id, user_id, cancelled) VALUES (?, ?, ?);

-- name: CreateUser :execresult
INSERT INTO users (uid, name) VALUES (?, ?);

-- name: ListUsers :many
SELECT id, uid, name FROM users ORDER BY id;

-- name: GetUserByUID :one
SELECT id, uid, name FROM users WHERE uid = ? LIMIT 1;

-- name: ListCommentsTree :many
WITH RECURSIVE r AS (
    SELECT * FROM comments WHERE comments.id = ?
    UNION ALL
    SELECT comments.* FROM comments, r WHERE comments.parent_id = r.id
)
SELECT * FROM r;
