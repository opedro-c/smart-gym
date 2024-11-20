package usecases

import (
	"cloud-gym/internal/core/exercise"
)

type CreateExercise struct {
	exerciseRepository exercise.ExerciseRepository
}

func NewCreateExercise(exerciseRepository exercise.ExerciseRepository) CreateExercise {
	return CreateExercise{
		exerciseRepository: exerciseRepository,
	}
}

func (c *CreateExercise) Execute(exerciseData exercise.ExerciseRecord) (string, error) {
	result, err := c.exerciseRepository.CreateExercise(exerciseData)
	if err != nil {
		return "", &exercise.CannotCreateExerciseError
	}
	return result, nil
}
