package exercise

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Interface
type ExerciseRepository interface {
	CreateExercise(exercise ExerciseRecord) (string, error)
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
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}
