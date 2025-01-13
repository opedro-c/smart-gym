package connection

import (
	"database/sql"
	"fmt"
	"gym-core-service/internal/postgres/sqlc"
	"os"
	"sync"
)

var (
	sqlDB *sql.DB
	once  sync.Once
)

func GetConnection() *sqlc.Queries {
	once.Do(func() {
		db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
		if err != nil {
			fmt.Printf("Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		sqlDB = db
	})

	return sqlc.New(sqlDB)
}
