package web

import (
	"log"
	"net/http"

	"github.com/umemak/eventsite_go/model/user"
)

func GetSearch(w http.ResponseWriter, r *http.Request) {
	u, err := user.BuildFromContext(r.Context())
	if err != nil {
		log.Printf("user.BuildFromContext: %v", err)
	}
	err = tpls["search.html"].Execute(w, struct {
		Header header
	}{
		Header: header{Title: "検索", User: u},
	})
	if err != nil {
		log.Fatalf("Execute: %v", err)
	}
}
