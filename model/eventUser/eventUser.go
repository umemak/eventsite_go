package eventUser

import (
	"fmt"

	"github.com/umemak/eventsite_go/db"
)

type EventUser struct {
	ID      int64
	EventID int64
	USerID  int64
}

func Create(e EventUser) (int64, error) {
	db, err := db.Open()
	if err != nil {
		return 0, fmt.Errorf("db.Open: %w", err)
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO eventUser (eventid, userid) VALUES (?, ?)")
	if err != nil {
		return 0, fmt.Errorf("db.Prepare: %w", err)
	}
	defer stmt.Close()
	res, err := stmt.Exec(e.EventID, e.USerID)
	if err != nil {
		return 0, fmt.Errorf("stmt.Exec: %w", err)
	}
	return res.LastInsertId()
}

func FindByEvent(id int64) ([]EventUser, error) {
	db, err := db.Open()
	if err != nil {
		return nil, fmt.Errorf("db.Open: %w", err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT id, eventid, userid FROM eventUser WHERE eventid = ? ORDER BY id", id)
	if err != nil {
		return nil, fmt.Errorf("db.Query: %w", err)
	}
	eventUsers := []EventUser{}
	for rows.Next() {
		var e EventUser
		err := rows.Scan(&e.ID, &e.EventID, &e.USerID)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}
		eventUsers = append(eventUsers, e)
	}
	return eventUsers, nil
}
