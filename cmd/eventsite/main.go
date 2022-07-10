package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/umemak/eventsite_go/web"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", web.GetRoot)
	r.Post("/", web.PostRoot)
	r.Get("/login", web.GetLogin)
	r.Get("/event_create", web.GetEventCreate)
	r.Get("/event_detail", web.GetEventDetail)

	port := ":8081"
	log.Printf("http://localhost%s", port)
	http.ListenAndServe(port, r)
}
