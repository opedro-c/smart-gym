package main

import (
	"gym-core-service/internal/adapter"
	"gym-core-service/internal/postgres/connection"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// @title Smart-Gym API
// @version 1.0
// @description This is the strongest server ever.
// @BasePath /api/v1
// @accept json
func main() {
	connection.GetConnection()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Mount("/api/v1", adapter.MakeAppRouter())

	log.Println("Server is running on port 3030")
	err := http.ListenAndServe(":3030", r)
	if err != nil {
		log.Fatal(err)
	}
}
