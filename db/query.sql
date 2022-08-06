-- name: CreateEvent :execresult
INSERT INTO events (title, start, place, open, close, author) VALUES (?, ?, ?, ?, ?, ?);

-- name: ListEvents :many
SELECT id, title, start, place, open, close, author FROM events ORDER BY id;

-- name: GetEvent :one
SELECT id, title, start, place, open, close, author FROM events WHERE id = ? LIMIT 1;

-- name: CreateEventUser :execresult
INSERT INTO events_users (event_id, user_id, cancelled) VALUES (?, ?, ?);

-- name: ListEventUsers :many
SELECT eu.id, eu.event_id, eu.user_id, eu.cancelled, u.name
FROM (
    SELECT id, event_id, user_id, cancelled,
    row_number() OVER (PARTITION BY event_id, user_id ORDER BY id DESC) AS num
    FROM events_users
) eu, users u
WHERE eu.event_id = ?
  AND eu.num = 1
  AND eu.user_id = u.id
ORDER BY eu.id;

-- name: CreateUser :execresult
INSERT INTO users (uid, name) VALUES (?, ?);

-- name: ListUsers :many
SELECT id, uid, name FROM users ORDER BY id;

-- name: GetUserByUID :one
SELECT id, uid, name FROM users WHERE uid = ? LIMIT 1;

-- -- name: ListCommentsTree :many
-- WITH RECURSIVE r AS (
--     SELECT * FROM comments WHERE id = ?
--     UNION ALL
--     SELECT comments.* FROM comments, r WHERE comments.parent_id = r.id
-- )
-- SELECT * FROM r;
