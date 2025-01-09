package main

import (
	"context"
	"database/sql"
	"fmt"
	"gym-core-service/internal/postgres/sqlc"
	"log"
	"net/http"
	"os"

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
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	ctx := context.Background()

	var db *sql.DB
	db, err := sql.Open("pgx", "postgres://gym:gym@localhost:5432/gym")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	queries := sqlc.New(db)

	user := sqlc.CreateUserParams{
		Username: "Talison",
		Email:    "email",
		Password: "123",
	}
	newUser, err := queries.CreateUser(ctx, user)

	if err != nil {
		fmt.Printf("Unable to create user: %v\n", err)
	} else {
		fmt.Printf("Created user: %v\n", newUser)
	}

	// r.Mount("/api/v1", adapter.MakeAppRouter())

	log.Println("Server is running on port 3030")
	http.ListenAndServe(":3030", r)
}
