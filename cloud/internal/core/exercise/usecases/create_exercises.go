package usecases

import (
	"cloud-gym/internal/core/exercise"
)

type CreateExercises struct {
	exerciseRepository exercise.ExerciseRepository
}

func NewCreateExercises(exerciseRepository exercise.ExerciseRepository) CreateExercises {
	return CreateExercises{
		exerciseRepository: exerciseRepository,
	}
}

func (c *CreateExercises) Execute(exerciseDatas []exercise.ExerciseRecord) ([]string, error) {
	return c.exerciseRepository.CreateExercises(exerciseDatas)
}
