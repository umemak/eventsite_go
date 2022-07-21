package web

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/umemak/eventsite_go/model/event"
	"github.com/umemak/eventsite_go/model/eventUser"
	"github.com/umemak/eventsite_go/model/user"
)

func GetEventCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	u, err := user.BuildFromContext(r.Context())
	if err != nil {
		log.Printf("user.BuildFromContext: %v", err)
	}
	var buf bytes.Buffer
	err = tpls["event_create.html"].Execute(&buf, struct {
		Header  header
		Message string
	}{
		Header: header{Title: "イベント作成", User: u},
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("tpls.Execute: %v", err), http.StatusInternalServerError)
		return
	}
	buf.WriteTo(w)
}

func PostEventCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	html_ng := "event_create.html"
	u, err := user.BuildFromContext(r.Context())
	if err != nil {
		log.Printf("user.BuildFromContext: %v", err)
	}
	if u.UID == "" {
		http.Redirect(w, r, "/login", 302)
	}
	err = r.ParseForm()
	if err != nil {
		log.Fatalf("ParseForm: %v", err)
	}
	now := time.Now()
	e := event.Event{
		Title:  r.PostFormValue("title"),
		Start:  &now, //r.PostFormValue("start"),
		Place:  r.PostFormValue("place"),
		Open:   &now, //r.PostFormValue("open"),
		Close:  &now, //r.PostFormValue("close"),
		Author: u.ID,
	}
	_, err = event.Create(e)
	if err != nil {
		var buf bytes.Buffer
		err = tpls[html_ng].Execute(&buf, struct {
			Header  header
			Message string
		}{
			Header:  header{Title: "イベント作成", User: u},
			Message: err.Error(),
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("tpls.Execute: %v", err), http.StatusInternalServerError)
			return
		}
		buf.WriteTo(w)
		return
	}
	http.Redirect(w, r, "/", 302)
}

func GetEventDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	u, err := user.BuildFromContext(r.Context())
	if err != nil {
		log.Printf("user.BuildFromContext: %v", err)
	}
	values := r.URL.Query()
	id, err := strconv.ParseInt(values.Get("id"), 10, 64)
	e, err := event.Find(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("event.Find: %v", err), http.StatusInternalServerError)
		return
	}
	eu, err := eventUser.FindByEvent(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("eventUser.FindByEvent: %v", err), http.StatusInternalServerError)
		return
	}
	var buf bytes.Buffer
	err = tpls["event_detail.html"].Execute(&buf, struct {
		Header     header
		Event      event.Event
		EventUsers []eventUser.EventUser
	}{
		Header:     header{Title: "イベント詳細", User: u},
		Event:      *e,
		EventUsers: eu,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("tpls.Execute: %v", err), http.StatusInternalServerError)
		return
	}
	buf.WriteTo(w)
}
