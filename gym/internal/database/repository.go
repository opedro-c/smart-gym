package database

import (
	"gym/internal/entities"
)

type ExerciseRepository interface {
	SaveExercise(exercise entities.Exercise) error
}
