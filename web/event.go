package web

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/umemak/eventsite_go/model/event"
	"github.com/umemak/eventsite_go/model/eventUser"
)

func GetEventCreate(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/template/event_create.html")
	if err != nil {
		log.Fatalf("template.ParseFiles: %v", err)
	}
	t.Execute(w, nil)
}

func GetEventDetail(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/template/event_detail.html")
	if err != nil {
		log.Fatalf("template.ParseFiles: %v", err)
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
	t.Execute(w, struct {
		Event      event.Event
		EventUsers []eventUser.EventUser
	}{
		Event:      *e,
		EventUsers: eu,
	})
	t.Execute(w, nil)
}
