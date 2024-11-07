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

func (r *ExerciseRepository) SaveExercise(exercise entities.Exercise) error {

	stmt, err := r.db.Prepare(`
		INSERT INTO exercises(user_rf_id, started_at, finished_at, name)
		VALUES(?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(exercise.UserRfId, exercise.StartedAt, exercise.FinishedAt, exercise.Name)
	if err != nil {
		return err
	}

	return nil
}

func (r *ExerciseRepository) SaveSeries(series entities.Series) error {

	stmt, err := r.db.Prepare(`
		INSERT INTO series(started_at, finished_at, repetitions, weight, exercise_id)
		VALUES(?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(series.StartedAt, series.FinishedAt, series.Repetitions, series.Weight, series.ExerciseId)
	if err != nil {
		return err
	}

	return nil
}

func (r *ExerciseRepository) GetRfId(tx *sql.Tx, rfId string) (rfIdFound string, err error) {

	err = tx.QueryRow("SELECT rf_id FROM valid_rf_ids WHERE rf_id = ?", rfId).Scan(&rfIdFound)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	return rfIdFound, nil
}
