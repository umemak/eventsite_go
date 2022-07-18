package web

import (
	"log"
	"net/http"
	"strconv"

	"github.com/umemak/eventsite_go/model/event"
	"github.com/umemak/eventsite_go/model/eventUser"
)

func GetEventCreate(w http.ResponseWriter, r *http.Request) {
	err := tpls["event_create.html"].Execute(w, struct {
		Header header
	}{
		Header: header{Title: "イベント作成"},
	})
	if err != nil {
		log.Fatalf("Execute: %v", err)
	}
}

func GetEventDetail(w http.ResponseWriter, r *http.Request) {
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
		Header:     header{Title: "イベント詳細"},
		Event:      *e,
		EventUsers: eu,
	})
	if err != nil {
		log.Fatalf("Execute: %v", err)
	}
}
