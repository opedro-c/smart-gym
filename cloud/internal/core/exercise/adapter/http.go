package adapter

import (
	"cloud-gym/internal/core/exercise"
	"cloud-gym/internal/core/exercise/usecases"
	"cloud-gym/internal/mongo"
	utils "cloud-gym/pkg"
	"net/http"
)

func CreateExerciseHandler(w http.ResponseWriter, r *http.Request) error {
	var data []exercise.ExerciseRecord
	if err := utils.ParseJson(r, &data); err != nil {
		return err
	}

	if err := utils.ValidateJsonStruct(&data); err != nil {
		return err
	}

	useCase := usecases.NewCreateExercises(getExerciseRepository())

	result, err := useCase.Execute(data)
	if err != nil {
		return err
	}

	utils.WriteJSON(w, http.StatusCreated, result)

	return nil
}

func getExerciseRepository() exercise.ExerciseRepository {
	return exercise.NewMongoExerciseRepository(mongo.GetConnection())
}
