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
	"github.com/umemak/eventsite_go/sqlc"
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
	u, err := user.BuildFromContext(r.Context())
	if err != nil {
		log.Printf("user.BuildFromContext: %v", err)
	}
	if u.Uid == "" {
		http.Redirect(w, r, "/login", 302)
	}
	err = r.ParseForm()
	if err != nil {
		log.Fatalf("ParseForm: %v", err)
	}
	e, err := buildEvent(r, u.ID)
	if err != nil {
		eventCreateFailed(w, u, err.Error())
		return
	}
	_, err = event.Create(e)
	if err != nil {
		eventCreateFailed(w, u, err.Error())
		return
	}
	http.Redirect(w, r, "/", 302)
}

func buildEvent(r *http.Request, uid int64) (sqlc.CreateEventParams, error) {
	start, err := time.Parse("2006-01-02", r.PostFormValue("start"))
	if err != nil {
		return sqlc.CreateEventParams{}, fmt.Errorf("time.Parse: %w", err)
	}
	open, err := time.Parse("2006-01-02 15:04", r.PostFormValue("open"))
	if err != nil {
		return sqlc.CreateEventParams{}, fmt.Errorf("time.Parse: %w", err)
	}
	close, err := time.Parse("2006-01-02 15:04", r.PostFormValue("close"))
	if err != nil {
		return sqlc.CreateEventParams{}, fmt.Errorf("time.Parse: %w", err)
	}
	e := sqlc.CreateEventParams{
		Title:  r.PostFormValue("title"),
		Start:  start,
		Place:  r.PostFormValue("place"),
		Open:   open,
		Close:  close,
		Author: uid,
	}
	return e, nil
}

func eventCreateFailed(w http.ResponseWriter, u sqlc.User, errMsg string) {
	var buf bytes.Buffer
	err := tpls["event_create.html"].Execute(&buf, struct {
		Header  header
		Message string
	}{
		Header:  header{Title: "イベント作成", User: u},
		Message: errMsg,
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
		Event      sqlc.Event
		EventUsers []sqlc.Eventuser
	}{
		Header:     header{Title: "イベント詳細", User: u},
		Event:      e,
		EventUsers: eu,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("tpls.Execute: %v", err), http.StatusInternalServerError)
		return
	}
	buf.WriteTo(w)
}
