package user

import (
	"encoding/json"
	"fmt"

	"github.com/koron/go-dproxy"
	"github.com/umemak/eventsite_go/db"
	"github.com/umemak/eventsite_go/pb"
)

type User struct {
	ID   int64
	UID  string
	Name string
}

func Create(email string, password string, passwordConfirm string, name string) (*User, error) {
	body, err := pb.CreateUser(email, password, passwordConfirm)
	if err != nil {
		return nil, fmt.Errorf("CreateUser: %+w", err)
	}
	var v any
	json.Unmarshal(body, &v)
	uid, err := dproxy.New(v).M("id").String()
	if err != nil {
		return nil, fmt.Errorf("dproxy: %+w", err)
	}
	ret := User{UID: uid, Name: name}
	id, err := createDB(ret)
	if err != nil {
		return &ret, fmt.Errorf("createDB: %+w", err)
	}
	ret.ID = id
	return &ret, nil
}

func createDB(u User) (int64, error) {
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

func GetByUID(uid string) (*User, error) {
	db, err := db.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	row := db.QueryRow("SELECT id, uid, name FROM user WHERE uid = ?", uid)
	var u User
	err = row.Scan(&u.ID, &u.UID, &u.Name)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func AuthViaEmail(email string, password string) (*User, error) {
	body, err := pb.AuthViaEmail(email, password)
	if err != nil {
		return nil, fmt.Errorf("AuthViaEmail: %+w", err)
	}
	var v any
	json.Unmarshal(body, &v)
	uid, err := dproxy.New(v).M("user").M("id").String()
	if err != nil {
		return nil, fmt.Errorf("dproxy: %+w", err)
	}
	u := &User{
		UID: uid,
	}
	return u, nil
}
