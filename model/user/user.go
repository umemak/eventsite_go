package user

import (
	"github.com/umemak/eventsite_go/db"
)

type User struct {
	ID   int64
	UID  string
	Name string
}

func Create(u User) (int64, error) {
	db, err := db.Open()
	if err != nil {
		return 0, err
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO user (uid, name) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(u.UID, u.Name)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func List() ([]User, error) {
	db, err := db.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT id, uid, name FROM user ORDER BY id")
	if err != nil {
		return nil, err
	}
	users := []User{}
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.UID, &u.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
