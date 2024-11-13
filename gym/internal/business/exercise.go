package business

import (
	"gym/internal/database"
	"gym/internal/entities"
	"log/slog"
)

type SaveExerciseUseCase[T InputSaveExerciseUseCase] struct {
	exerciseRepo database.ExerciseRepository
}

type InputSaveExerciseUseCase struct {
	Exercise entities.Exercise `json:"exercise"`
	Series   []entities.Series `json:"series"`
}

type OutputSaveExerciseUseCase struct{}

func NewSaveExerciseUseCase[T InputSaveExerciseUseCase](exerciseRepo database.ExerciseRepository) *SaveExerciseUseCase[T] {
	return &SaveExerciseUseCase[T]{
		exerciseRepo: exerciseRepo,
	}
}

func (u *SaveExerciseUseCase[T]) Execute(input InputSaveExerciseUseCase) (output OutputSaveExerciseUseCase, err error) {
	rfId, err := u.exerciseRepo.GetRfId(input.Exercise.UserRfId)
	if err != nil {
		return output, err
	}
	if rfId == "" {
		slog.Debug("RfId not found", "RfId", input.Exercise.UserRfId)
		return output, nil
	}

	if err = u.exerciseRepo.SaveExercise(input.Exercise); err != nil {
		return output, err
	}

	if err = u.exerciseRepo.SaveSeriesInBatch(input.Series); err != nil {
		return output, err
	}

	return output, nil
}
