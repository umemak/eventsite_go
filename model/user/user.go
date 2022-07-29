package user

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-chi/jwtauth/v5"
	"github.com/koron/go-dproxy"
	"github.com/umemak/eventsite_go/db"
	"github.com/umemak/eventsite_go/pb"
	"github.com/umemak/eventsite_go/sqlc"
)

func Create(email string, password string, passwordConfirm string, name string) (sqlc.User, error) {
	body, err := pb.CreateUser(email, password, passwordConfirm)
	if err != nil {
		return sqlc.User{}, fmt.Errorf("pb.CreateUser: %w", err)
	}
	var v any
	json.Unmarshal(body, &v)
	uid, err := dproxy.New(v).M("id").String()
	if err != nil {
		return sqlc.User{}, fmt.Errorf("dproxy.New.M.String: %w", err)
	}
	ret := sqlc.CreateUserParams{Uid: uid, Name: name}
	id, err := createDB(ret)
	if err != nil {
		return sqlc.User{}, fmt.Errorf("createDB: %w", err)
	}
	return sqlc.User{ID: id, Uid: uid, Name: name}, nil
}

func createDB(u sqlc.CreateUserParams) (int64, error) {
	db, err := db.Open()
	if err != nil {
		return 0, fmt.Errorf("db.Open: %w", err)
	}
	defer db.Close()
	queries := sqlc.New(db)
	ctx := context.Background()
	res, err := queries.CreateUser(ctx, u)
	if err != nil {
		return 0, fmt.Errorf("queries.CreateUser: %w", err)
	}
	return res.LastInsertId()
}

func List() ([]sqlc.User, error) {
	db, err := db.Open()
	if err != nil {
		return nil, fmt.Errorf("db.Open: %w", err)
	}
	defer db.Close()
	queries := sqlc.New(db)
	ctx := context.Background()
	users, err := queries.ListUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("queries.ListUsers: %w", err)
	}
	return users, nil
}

func GetByUID(uid string) (sqlc.User, error) {
	db, err := db.Open()
	if err != nil {
		return sqlc.User{}, fmt.Errorf("db.Open: %w", err)
	}
	defer db.Close()
	queries := sqlc.New(db)
	ctx := context.Background()
	u, err := queries.GetUserByUID(ctx, uid)
	if err != nil {
		return sqlc.User{}, fmt.Errorf("queries.ListUsers: %w", err)
	}
	return u, nil
}

func AuthViaEmail(email string, password string) (sqlc.User, error) {
	body, err := pb.AuthViaEmail(email, password)
	if err != nil {
		return sqlc.User{}, fmt.Errorf("pb.AuthViaEmail: %w", err)
	}
	var v any
	json.Unmarshal(body, &v)
	uid, err := dproxy.New(v).M("user").M("id").String()
	if err != nil {
		return sqlc.User{}, fmt.Errorf("dproxy.New.M.M.String: %w", err)
	}
	return GetByUID(uid)
}

func BuildFromContext(ctx context.Context) (sqlc.User, error) {
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return sqlc.User{}, fmt.Errorf("jwtauth.FromContext: %w", err)
	}
	return sqlc.User{
		ID:   int64(claims["id"].(float64)),
		Uid:  claims["uid"].(string),
		Name: claims["name"].(string),
	}, err
}
