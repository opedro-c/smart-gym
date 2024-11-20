package migrations

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Migrate(client *mongo.Client) {
	db := client.Database("db")
	tso := options.TimeSeries()
	tso.SetTimeField("startedAt")
	tso.SetMetaField("a")

	opts := options.CreateCollection().SetTimeSeriesOptions(tso)
	db.CreateCollection(context.TODO(), "march2022", opts)
}
