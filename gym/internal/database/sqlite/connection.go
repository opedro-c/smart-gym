package sqlite

import (
	"database/sql"
	"gym/pkg/logger"
	"path/filepath"
	 _ "github.com/mattn/go-sqlite3"
	"sync"
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
	logger.Logger().Println("Opening database connection")
	dbPath := filepath.Join("gym.db")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		logger.Logger().Fatalln("Error while opening database connection", err)
	}
	return db
}

