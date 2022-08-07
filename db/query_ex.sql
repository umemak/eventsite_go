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
