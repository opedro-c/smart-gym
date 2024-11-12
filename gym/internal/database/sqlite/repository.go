package sqlite

import (
	"database/sql"
	"gym/internal/entities"
)

type ExerciseRepository struct {
	tx *sql.Tx
}

func NewExerciseRepository(tx *sql.Tx) *ExerciseRepository {
	return &ExerciseRepository{tx}
}

func (r *ExerciseRepository) SaveExercise(exercise entities.Exercise) error {

	stmt, err := r.tx.Prepare(`
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

func (r *ExerciseRepository) SaveSeriesInBatch(series []entities.Series) error {
	sql := "INSERT INTO series(started_at, finished_at, repetitions, weight, exercise_id) VALUES"
	values := []interface{}{}
	for _, s := range series {
		sql += "(?, ?, ?, ?, ?),"
		values = append(values, s.StartedAt, s.FinishedAt, s.Repetitions, s.Weight, s.ExerciseId)
	}
	sql = sql[:len(sql)-1]
	stmt, err := r.tx.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(values...)
	if err != nil {
		return err
	}
	return nil
}

func (r *ExerciseRepository) GetRfId(rfId string) (rfIdFound string, err error) {

	err = r.tx.QueryRow("SELECT rf_id FROM valid_rf_ids WHERE rf_id = ?", rfId).Scan(&rfIdFound)
	if err != nil {
		return "", err
	}
	return rfIdFound, nil
}
