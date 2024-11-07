package business

import (
	"gym/internal/database"
	"gym/internal/entities"
	"gym/pkg/logger"
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
		logger.Logger().Printf("RfId %s not found", input.Exercise.UserRfId)
		return output, nil
	}

	if err = u.exerciseRepo.SaveExercise(input.Exercise); err != nil {
		logger.Logger().Println()
		return output, err
	}

	if err = u.exerciseRepo.SaveSeriesInBatch(input.Series); err != nil {
		return output, err
	}

	return output, nil
}
