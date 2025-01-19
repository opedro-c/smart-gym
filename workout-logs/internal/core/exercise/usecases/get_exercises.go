package usecases

import (
	"cloud-gym/internal/core/exercise"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InputGetExercises struct {
	StartedAt  primitive.DateTime
	FinishedAt primitive.DateTime
	OriginId   string
	UserId     uint64
}

type GetExercises struct {
	exerciseRepository exercise.ExerciseRepository
}

func NewGetExercises(exerciseRepository exercise.ExerciseRepository) GetExercises {
	return GetExercises{
		exerciseRepository: exerciseRepository,
	}
}

func (g *GetExercises) Execute(input InputGetExercises) ([]exercise.OutputGetExercises, error) {
	return g.exerciseRepository.GetExercises(input.StartedAt, input.FinishedAt, input.OriginId, input.UserId)
}
