package main

import (
	"database/sql"
	"embed"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"os"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	var db *sql.DB
	db, err := sql.Open("pgx", "postgres://gym:gym@localhost:5432/gym")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("pgx"); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to set dialect in Goose: %v\n", err)
		os.Exit(1)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to migrate: %v\n", err)
		os.Exit(1)
	}

	fmt.Print("[OK] All migrations applied")
}
