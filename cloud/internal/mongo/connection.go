package mongo

import (
	"context"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	once   sync.Once
)

func GetConnection() *mongo.Client {
	once.Do(func() {
		uri := os.Getenv("MONGODB_URL") // "mongodb://gym:gym@localhost:27017"
		if uri == "" {
			panic("Missing MONGODB_URI environment variable")
		}

		clientConnected, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
			panic(err)
		}
		client = clientConnected
	})
	return client
}
