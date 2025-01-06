package exercise

import (
	"context"

	db "cloud-gym/internal/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Interface
type ExerciseRepository interface {
	CreateExercises(exercises []ExerciseRecord) ([]string, error)
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
