package eventUser

import (
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
		return 0, err
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO eventUser (eventid, userid) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(e.EventID, e.USerID)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func FindByEvent(id int64) ([]EventUser, error) {
	db, err := db.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT id, eventid, userid FROM eventUser WHERE eventid = ? ORDER BY id", id)
	if err != nil {
		return nil, err
	}
	eventUsers := []EventUser{}
	for rows.Next() {
		var e EventUser
		err := rows.Scan(&e.ID, &e.EventID, &e.USerID)
		if err != nil {
			return nil, err
		}
		eventUsers = append(eventUsers, e)
	}
	return eventUsers, nil
}
