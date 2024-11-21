package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

var (
	client *mongo.Client
	once   sync.Once
)

func GetConnection() *mongo.Client {
	once.Do(func() {
		uri := "mongodb://gym:gym@localhost:27017"
		clientConnected, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
			panic(err)
		}
		client = clientConnected
	})
	return client
}
