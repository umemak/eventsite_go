package event

import (
	"context"
	"fmt"

	"github.com/umemak/eventsite_go/db"
	"github.com/umemak/eventsite_go/sqlc"
)

func Create(e sqlc.CreateEventParams) (int64, error) {
	db, err := db.Open()
	if err != nil {
		return 0, fmt.Errorf("db.Open: %w", err)
	}
	defer db.Close()
	queries := sqlc.New(db)
	ctx := context.Background()
	res, err := queries.CreateEvent(ctx, e)
	if err != nil {
		return 0, fmt.Errorf("queries.CreateEvent: %w", err)
	}
	return res.LastInsertId()
}

func List() ([]sqlc.Event, error) {
	db, err := db.Open()
	if err != nil {
		return nil, fmt.Errorf("db.Open: %w", err)
	}
	defer db.Close()
	queries := sqlc.New(db)
	ctx := context.Background()
	events, err := queries.ListEvents(ctx)
	if err != nil {
		return nil, fmt.Errorf("queries.ListEvents: %w", err)
	}
	return events, nil
}

func Find(id int64) (*sqlc.Event, error) {
	db, err := db.Open()
	if err != nil {
		return nil, fmt.Errorf("db.Open: %w", err)
	}
	defer db.Close()
	queries := sqlc.New(db)
	ctx := context.Background()
	event, err := queries.GetEvent(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("queries.GetEvent: %w", err)
	}
	return &event, nil
}
