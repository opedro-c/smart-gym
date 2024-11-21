package migrations

import (
	m "cloud-gym/internal/mongo"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Migrate(client *mongo.Client) {
	db := client.Database(m.DATABASE_NAME)
	tso := options.TimeSeries()
	tso.SetTimeField("started_at")
	tso.SetMetaField("user_id")
	tso.SetGranularity("minutes")

	opts := options.CreateCollection().SetTimeSeriesOptions(tso)
	db.CreateCollection(context.TODO(), m.EXERCISES_COLLECTION_NAME, opts)
}
