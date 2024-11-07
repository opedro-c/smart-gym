package database

import (
	"database/sql"
	"gym/internal/entities"
)

type ExerciseRepository interface {
	SaveExercise(tx *sql.Tx, exercise entities.Exercise) error
	SaveSeries(tx *sql.Tx, series entities.Series) error
	GetRfId(tx *sql.Tx, rfId string) (string, error)
}
