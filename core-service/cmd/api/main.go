package main

import (
	"gym-core-service/internal/adapter"
	"gym-core-service/internal/postgres/connection"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/jackc/pgx/v5/stdlib"

	// swagger
	"github.com/swaggo/http-swagger"
	_ "gym-core-service/docs"
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
	// r.Use(middleware.RedirectSlashes)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Recoverer)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3030/swagger/doc.json"), //The url pointing to API definition
	))

	// r.Mount("/api/v1", adapter.MakeAppRouter())
	r.Mount("/", adapter.MakeAppRouter())

	log.Println("Server is running on port 3030")
	err := http.ListenAndServe(":3030", r)
	if err != nil {
		log.Fatal(err)
	}
}
