package exercise

import (
	"context"
	"encoding/json"
	"log"

	m "cloud-gym/internal/mongo"

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
	for i, model := range exercisesModels {
		documents[i] = model
	}

	log.Println("Creating exercise")
	a, _ := json.Marshal(documents)
	log.Println(string(a[:]))

	coll := r.client.Database(m.DATABASE_NAME).Collection(m.EXERCISES_COLLECTION_NAME)
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
