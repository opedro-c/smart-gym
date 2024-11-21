package exercise

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Interface
type ExerciseRepository interface {
	CreateExercise(exercise ExerciseRecord) (string, error)
	// CreateManyExercises(exercises []ExerciseRecord) ([]string, error)
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

func (r *MongoExerciseRepository) CreateExercise(exercise ExerciseRecord) (string, error) {
	coll := r.client.Database("db").Collection("books")
	result, err := coll.InsertOne(context.TODO(), exercise)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).String(), nil
}

// func (r *MongoExerciseRepository) CreateManyExercises(exercises []ExerciseRecord) ([]string, error) {
// 	coll := r.client.Database("db").Collection("books")
// 	result, err := coll.InsertMany(context.TODO(), exercises)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var ids []string
// 	for _, id := range result.InsertedIDs {
// 		ids = append(ids, id.(string))
// 	}

// 	return ids, nil
// }
