package web

import (
	"log"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/umemak/eventsite_go/model/event"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	isLogin := isLogin(r.Context())
	name := ""
	if isLogin {
		_, claims, _ := jwtauth.FromContext(r.Context())
		name = claims["name"].(string)
	}

	events, err := event.List()
	if err != nil {
		log.Fatalf("event.List: %v", err)
	}
	err = tpls["index.html"].Execute(w, struct {
		Header  header
		Events  []event.Event
		IsLogin bool
		Name    string
	}{
		Header:  header{Title: "トップ"},
		Events:  events,
		IsLogin: isLogin,
		Name:    name,
	})
	if err != nil {
		log.Fatalf("Execute: %v", err)
	}
}
