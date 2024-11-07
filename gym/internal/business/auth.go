package business

import (
	"database/sql"
	"gym/internal/database"
	"gym/pkg/abstract"
)

type SaveExerciseUseCase[T InputSaveExerciseUseCase] struct {
	*abstract.AbstractUseCase[T]
	exerciseRepo database.ExerciseRepository
	tx           *sql.Tx
}

type InputSaveExerciseUseCase struct {
	RfId string
}

func NewSaveExerciseUseCase[T InputSaveExerciseUseCase](tx *sql.Tx, exerciseRepo database.ExerciseRepository) *SaveExerciseUseCase[T] {
	return &SaveExerciseUseCase[T]{
		AbstractUseCase: &abstract.AbstractUseCase[T]{Tx: tx},
		tx:           tx,
		exerciseRepo: exerciseRepo,
	}
}

func (u *SaveExerciseUseCase[T]) Execute(input InputSaveExerciseUseCase) error {
	_, err := u.exerciseRepo.GetRfId(u.tx, input.RfId)

	if err != nil {
		return err
	}
	return nil
}
