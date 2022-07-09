package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/umemak/eventsite_go/web"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", web.GetRoot)
	r.Post("/", web.PostRoot)

	port := ":8081"
	log.Printf("http://localhost%s", port)
	http.ListenAndServe(port, r)
}
