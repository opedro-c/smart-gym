package main

import (
	"cloud-gym/internal/mongo"
	"cloud-gym/internal/mongo/migrations"
)

func main() {
	client := mongo.GetConnection()
	migrations.Migrate(client)
}
