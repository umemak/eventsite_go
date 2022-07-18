package web

import (
	"log"
	"net/http"

	"github.com/umemak/eventsite_go/model/event"
	"github.com/umemak/eventsite_go/model/user"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	u, err := user.BuildFromContext(r.Context())
	if err != nil {
		log.Printf("user.BuildFromContext: %v", err)
	}
	events, err := event.List()
	if err != nil {
		log.Fatalf("event.List: %v", err)
	}
	err = tpls["index.html"].Execute(w, struct {
		Header header
		Events []event.Event
	}{
		Header: header{Title: "トップ", User: u},
		Events: events,
	})
	if err != nil {
		log.Fatalf("Execute: %v", err)
	}
}
