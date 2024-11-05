package sqlite

import (
	"database/sql"
	"gym/internal/entities"
)

type ExerciseRepository struct {
	db *sql.DB
}

func NewExerciseRepository(db *sql.DB) *ExerciseRepository {
	return &ExerciseRepository{db}
}

func (r *ExerciseRepository) SaveExercise(exercise entities.Exercise, series entities.Series[]) error {


	stmt, err := r.db.Prepare("INSERT INTO exercises(name, description) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(exercise.Name, exercise.Description)
	if err != nil {
		return err
	}

	return nil
}
