package usecases

import (
	"cloud-gym/internal/core/exercise"
	"log"
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
	log.Println("Creating exercise")
	log.Println(exerciseDatas)
	result, err := c.exerciseRepository.CreateExercises(exerciseDatas)
	if err != nil {
		return nil, &exercise.CannotCreateExerciseError
	}
	return result, nil
}
