package web

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/umemak/eventsite_go/model/event"
	"github.com/umemak/eventsite_go/model/eventUser"
	"github.com/umemak/eventsite_go/model/user"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/template/index.html")
	if err != nil {
		log.Fatalf("template.ParseFiles: %v", err)
	}
	events, err := event.List()
	if err != nil {
		log.Fatalf("event.List: %v", err)
	}
	t.Execute(w, struct {
		Events []event.Event
	}{
		Events: events,
	})
}

func PostRoot(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/template/index.html")
	if err != nil {
		log.Fatalf("template.ParseFiles: %v", err)
	}
	err = r.ParseForm()
	if err != nil {
		log.Fatalf("ParseForm: %v", err)
	}
	_, err = user.Create(user.User{Name: r.PostFormValue("name")})
	if err != nil {
		log.Fatalf("user.Create: %v", err)
	}
	users, err := user.List()
	if err != nil {
		log.Fatalf("user.List: %v", err)
	}
	t.Execute(w, struct {
		Users []user.User
	}{
		Users: users,
	})
}

func GetLogin(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/template/login.html")
	if err != nil {
		log.Fatalf("template.ParseFiles: %v", err)
	}
	t.Execute(w, nil)
}

func GetSearch(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/template/search.html")
	if err != nil {
		log.Fatalf("template.ParseFiles: %v", err)
	}
	t.Execute(w, nil)
}

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
