// source: query_ex.sql

package sqlc

import (
	"context"
)

const listEventUsers = `-- name: ListEventUsers :many
SELECT eu.id, eu.event_id, eu.user_id, eu.cancelled, u.name
FROM (
    SELECT id, event_id, user_id, cancelled,
    row_number() OVER (PARTITION BY event_id, user_id ORDER BY id DESC) AS num
    FROM events_users
) eu, users u
WHERE eu.event_id = ?
  AND eu.num = 1
  AND eu.user_id = u.id
ORDER BY eu.id
`

type ListEventUsersRow struct {
	ID        int64
	EventID   int64
	UserID    int64
	Cancelled bool
	Name      string
}

func (q *Queries) ListEventUsers(ctx context.Context, eventID int64) ([]ListEventUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, listEventUsers, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListEventUsersRow
	for rows.Next() {
		var i ListEventUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.EventID,
			&i.UserID,
			&i.Cancelled,
			&i.Name,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
