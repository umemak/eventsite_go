package web

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/umemak/eventsite_go/model/event"
	"github.com/umemak/eventsite_go/model/user"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	u, err := user.BuildFromContext(r.Context())
	if err != nil {
		log.Printf("user.BuildFromContext: %v", err)
	}
	events, err := event.List()
	if err != nil {
		http.Error(w, fmt.Sprintf("event.List: %v", err), http.StatusInternalServerError)
		return
	}
	var buf bytes.Buffer
	err = tpls["index.html"].Execute(&buf, struct {
		Header header
		Events []event.Event
	}{
		Header: header{Title: "トップ", User: u},
		Events: events,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("tpls.Execute: %v", err), http.StatusInternalServerError)
		return
	}
	buf.WriteTo(w)
}
