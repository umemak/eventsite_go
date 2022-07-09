package user

import (
	"github.com/umemak/eventsite_go/db"
)

type User struct {
	ID   int64
	Name string
}

func Create(u User) (int64, error) {
	db, err := db.Open()
	if err != nil {
		return 0, err
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO user (name) VALUES (?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(u.Name)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func Find() ([]User, error) {
	db, err := db.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT id, name FROM user ORDER BY id")
	if err != nil {
		return nil, err
	}
	users := []User{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
