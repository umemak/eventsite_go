package web

import (
	"log"
	"net/http"
	"text/template"
)

func GetSearch(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/template/search.html")
	if err != nil {
		log.Fatalf("template.ParseFiles: %v", err)
	}
	t.Execute(w, nil)
}
