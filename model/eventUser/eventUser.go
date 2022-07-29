package eventUser

import (
	"context"
	"fmt"

	"github.com/umemak/eventsite_go/db"
	"github.com/umemak/eventsite_go/sqlc"
)

func Create(e sqlc.CreateEventUserParams) (int64, error) {
	db, err := db.Open()
	if err != nil {
		return 0, fmt.Errorf("db.Open: %w", err)
	}
	defer db.Close()
	queries := sqlc.New(db)
	ctx := context.Background()
	res, err := queries.CreateEventUser(ctx, e)
	if err != nil {
		return 0, fmt.Errorf("queries.CreateEventUser: %w", err)
	}
	return res.LastInsertId()
}

func FindByEvent(id int64) ([]sqlc.Eventuser, error) {
	db, err := db.Open()
	if err != nil {
		return nil, fmt.Errorf("db.Open: %w", err)
	}
	defer db.Close()
	queries := sqlc.New(db)
	ctx := context.Background()
	eventUsers, err := queries.ListEventUsers(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("queries.ListEventUsers: %w", err)
	}
	return eventUsers, nil
}
