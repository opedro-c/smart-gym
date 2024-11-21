package main

import (
	"cloud-gym/internal/adapter"
	"cloud-gym/internal/mongo"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	mongo.GetConnection()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Mount("/api/v1", adapter.MakeAppRouter())

	http.ListenAndServe(":3030", r)
}
