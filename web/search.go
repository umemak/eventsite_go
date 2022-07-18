package web

import (
	"log"
	"net/http"
)

func GetSearch(w http.ResponseWriter, r *http.Request) {
	err := tpls["search.html"].Execute(w, struct {
		Header header
	}{
		Header: header{Title: "検索"},
	})
	if err != nil {
		log.Fatalf("Execute: %v", err)
	}
}
