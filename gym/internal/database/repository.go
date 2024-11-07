package database

import (
	"gym/internal/entities"
)

type ExerciseRepository interface {
	SaveExercise(exercise entities.Exercise) error
	SaveSeriesInBatch(series []entities.Series) error
	GetRfId(rfId string) (string, error)
}
