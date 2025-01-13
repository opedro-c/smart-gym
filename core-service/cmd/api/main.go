package main

import (
	"gym-core-service/internal/adapter"
	"gym-core-service/internal/postgres/connection"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/jackc/pgx/v5/stdlib"
  "github.com/go-chi/cors"

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

  r.Use(cors.Handler(cors.Options{
    // AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
    AllowedOrigins:   []string{"https://*", "http://*"},
    // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300, // Maximum value not ignored by any of major browsers
  }))

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:3030/swagger/doc.json"), //The url pointing to API definition
	))

	// r.Mount("/api/v1", adapter.MakeAppRouter())
	r.Mount("/", adapter.MakeAppRouter())

	log.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
