package exercise

import (
	db "cloud-gym/internal/mongo"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Interface
type ExerciseRepository interface {
	CreateExercises(exercises []ExerciseRecord) ([]string, error)
	GetExercises(startedAt primitive.DateTime, finishedAt primitive.DateTime, originId string, userId uint64) ([]OutputGetExercises, error)
}

// Implementation
type MongoExerciseRepository struct {
	client *mongo.Client
}

func NewMongoExerciseRepository(client *mongo.Client) ExerciseRepository {
	return &MongoExerciseRepository{
		client: client,
	}
}

func (r *MongoExerciseRepository) CreateExercises(exercises []ExerciseRecord) ([]string, error) {
	exercisesModels := NewExerciseCollectionRecord(exercises)

	// Convert []ExerciseCollectionRecord to []interface{}
	documents := make([]interface{}, len(exercisesModels))
	for i := range exercisesModels {
		documents[i] = exercisesModels[i]
	}

	coll := r.client.Database(db.DATABASE_NAME).Collection(db.EXERCISES_COLLECTION_NAME)
	result, err := coll.InsertMany(context.TODO(), documents)
	if err != nil {
		return nil, err
	}

	objectIds := make([]string, 0, len(result.InsertedIDs))
	for _, id := range result.InsertedIDs {
		objectIds = append(objectIds, id.(primitive.ObjectID).Hex())
	}

	return objectIds, nil
}

func (r *MongoExerciseRepository) GetExercises(startedAt primitive.DateTime, finishedAt primitive.DateTime, originId string, userId uint64) ([]OutputGetExercises, error) {
	coll := r.client.Database(db.DATABASE_NAME).Collection(db.EXERCISES_COLLECTION_NAME)
	filter := bson.M{
		"started_at":  bson.M{"$gte": startedAt},
		"finished_at": bson.M{"$lte": finishedAt},
		"origin_id":   originId,
		"user_id":     userId,
	}

	opts := options.Find().SetSort(bson.D{{"started_at", 1}})
	cursor, err := coll.Find(context.TODO(), filter, opts)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			slog.Debug("Error closing cursor", err)
		}
	}(cursor, context.TODO())

	var output []OutputGetExercises
	if err = cursor.All(context.TODO(), &output); err != nil {
		return nil, err
	}

	return output, nil
}

type OutputGetExercises struct {
	StartedAt  primitive.DateTime `bson:"started_at" json:"started_at"`
	FinishedAt primitive.DateTime `bson:"finished_at" json:"finished_at"`
	OriginID   string             `bson:"origin_id" json:"origin_id"`
	UserID     uint64             `bson:"user_id" json:"user_id"`
	Data       struct {
		Weight uint32 `bson:"weight" json:"weight"`
	} `bson:"data" json:"data"`
	ID primitive.ObjectID `bson:"_id" json:"id"`
}
