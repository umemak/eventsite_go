package web

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
		Header header
	}{
		Header: header{Title: "イベント作成", User: u},
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("tpls.Execute: %v", err), http.StatusInternalServerError)
		return
	}
	buf.WriteTo(w)
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	eu, err := eventUser.FindByEvent(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
