package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Open() (*sql.DB, error) {
	dsn := os.Getenv("EVENTSITE_DSN")
	if dsn == "" {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
		dsn = os.Getenv("EVENTSITE_DSN")
	}
	db, err := sql.Open("mysql", dsn+"?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		return nil, err
	}
	return db, nil
}
