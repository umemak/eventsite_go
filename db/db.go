package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Open() (*sql.DB, error) {
	dsn := os.Getenv("EVENTSITE_DSN")
	if dsn == "" {
		err := godotenv.Load()
		if err != nil {
			return nil, fmt.Errorf("godotenv.Load: %w", err)
		}
		dsn = os.Getenv("EVENTSITE_DSN")
	}
	db, err := sql.Open("mysql", dsn+"?parseTime=true")
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}
	return db, nil
}
