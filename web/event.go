package web

import (
	"log"
	"net/http"
	"strconv"

	"github.com/umemak/eventsite_go/model/event"
	"github.com/umemak/eventsite_go/model/eventUser"
	"github.com/umemak/eventsite_go/model/user"
)

func GetEventCreate(w http.ResponseWriter, r *http.Request) {
	u, err := user.BuildFromContext(r.Context())
	if err != nil {
		log.Printf("user.BuildFromContext: %v", err)
	}
	err = tpls["event_create.html"].Execute(w, struct {
		Header header
	}{
		Header: header{Title: "イベント作成", User: u},
	})
	if err != nil {
		log.Fatalf("Execute: %v", err)
	}
}

func GetEventDetail(w http.ResponseWriter, r *http.Request) {
	u, err := user.BuildFromContext(r.Context())
	if err != nil {
		log.Printf("user.BuildFromContext: %v", err)
	}
	values := r.URL.Query()
	id, err := strconv.ParseInt(values.Get("id"), 10, 64)
	e, err := event.Find(id)
	if err != nil {
		log.Fatalf("event.Find: %v", err)
	}
	eu, err := eventUser.FindByEvent(id)
	if err != nil {
		log.Fatalf("event.Find: %v", err)
	}
	err = tpls["event_detail.html"].Execute(w, struct {
		Header     header
		Event      event.Event
		EventUsers []eventUser.EventUser
	}{
		Header:     header{Title: "イベント詳細", User: u},
		Event:      *e,
		EventUsers: eu,
	})
	if err != nil {
		log.Fatalf("Execute: %v", err)
	}
}
