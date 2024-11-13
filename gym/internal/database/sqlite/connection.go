package sqlite

import (
	"database/sql"
	"log/slog"
	"path/filepath"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetConnection() *sql.DB {
	once.Do(func() {
		db = openConnection()
	})
	return db
}

func openConnection() *sql.DB {
	slog.Info("Opening database connection")
	dbPath := filepath.Join("gym.db")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		slog.Error("Error while opening database connection")
		panic(err)
	}
	return db
}

