package event

import (
	"time"

	"github.com/umemak/eventsite_go/db"
)

type Event struct {
	ID     int64
	Title  string
	Start  *time.Time
	Place  string
	Open   *time.Time
	Close  *time.Time
	Author int64
}

func Create(e Event) (int64, error) {
	db, err := db.Open()
	if err != nil {
		return 0, err
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO event (title, start, place, open, close, author) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(e.Title, e.Start, e.Place, e.Open, e.Close, e.Author)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func List() ([]Event, error) {
	db, err := db.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT id, title, start, place, open, close, author FROM event ORDER BY id")
	if err != nil {
		return nil, err
	}
	events := []Event{}
	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Title, &e.Start, &e.Place, &e.Open, &e.Close, &e.Author)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func Find(id int64) (*Event, error) {
	db, err := db.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	row := db.QueryRow("SELECT id, title, start, place, open, close, author FROM event WHERE id = ?", id)
	var e Event
	err = row.Scan(&e.ID, &e.Title, &e.Start, &e.Place, &e.Open, &e.Close, &e.Author)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
