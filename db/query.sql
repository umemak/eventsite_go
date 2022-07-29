-- name: CreateEvent :execresult
INSERT INTO event (title, start, place, open, close, author) VALUES (?, ?, ?, ?, ?, ?);

-- name: ListEvents :many
SELECT id, title, start, place, open, close, author FROM event ORDER BY id;

-- name: GetEvent :one
SELECT id, title, start, place, open, close, author FROM event WHERE id = ? LIMIT 1;

-- name: CreateEventUser :execresult
INSERT INTO eventUser (eventid, userid) VALUES (?, ?);

-- name: ListEventUsers :many
SELECT id, eventid, userid, status FROM eventUser WHERE eventid = ? ORDER BY id;

-- name: CreateUser :execresult
INSERT INTO user (uid, name) VALUES (?, ?);

-- name: ListUsers :many
SELECT id, uid, name FROM user ORDER BY id;

-- name: GetUserByUID :one
SELECT id, uid, name FROM user WHERE uid = ? LIMIT 1;
