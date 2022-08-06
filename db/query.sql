-- name: CreateEvent :execresult
INSERT INTO event (title, start, place, open, close, author) VALUES (?, ?, ?, ?, ?, ?);

-- name: ListEvents :many
SELECT id, title, start, place, open, close, author FROM event ORDER BY id;

-- name: GetEvent :one
SELECT id, title, start, place, open, close, author FROM event WHERE id = ? LIMIT 1;

-- name: CreateEventUser :execresult
INSERT INTO eventUser (eventid, userid, `status`) VALUES (?, ?, ?);

-- name: ListEventUsers :many
SELECT eu.id, eu.eventid, eu.userid, eu.`status`, u.name
FROM (
    SELECT id, eventid, userid, `status`,
    row_number() OVER (PARTITION BY eventid, userid ORDER BY id DESC) AS num
    FROM eventUser
) eu, user u
WHERE eu.eventid = ?
  AND eu.num = 1
  AND eu.userid = u.id
ORDER BY eu.id;

-- name: CreateUser :execresult
INSERT INTO user (uid, name) VALUES (?, ?);

-- name: ListUsers :many
SELECT id, uid, name FROM user ORDER BY id;

-- name: GetUserByUID :one
SELECT id, uid, name FROM user WHERE uid = ? LIMIT 1;
