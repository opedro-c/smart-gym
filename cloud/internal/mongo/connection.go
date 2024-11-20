package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"sync"
)

var (
	client *mongo.Client
	once   sync.Once
)

func GetConnection() *mongo.Client {
	once.Do(func() {
		uri := os.Getenv("MONGODB_URI")
		clientConnected, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
			panic(err)
		}
		client = clientConnected
	})
	return client
}
