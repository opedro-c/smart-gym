package adapter

import (
	"cloud-gym/internal/core/exercise"
	"cloud-gym/internal/core/exercise/usecases"
	"cloud-gym/internal/mongo"
	utils "cloud-gym/pkg"
	"fmt"
	"net/http"
)

// func GetExercisesHandler(w http.ResponseWriter, r *http.Request) error {
// 	useCase := usecases.NewCreateExercise(getExerciseRepository())

// 	result, err := useCase.Execute(data)
// 	if err != nil {
// 		return err
// 	}

// 	return utils.WriteJSON(w, http.StatusCreated, result)
// }

func CreateExerciseHandler(w http.ResponseWriter, r *http.Request) error {

	if r.Body == nil {
		fmt.Println("FOIIIII")
	}

	var data exercise.ExerciseRecord
	if err := utils.ParseJson(r, &data); err != nil {
		return err
	}

	useCase := usecases.NewCreateExercise(getExerciseRepository())

	result, err := useCase.Execute(data)
	if err != nil {
		return err
	}

	fmt.Println("result", result)

	return utils.WriteJSON(w, http.StatusCreated, result)
}

func getExerciseRepository() exercise.ExerciseRepository {
	return exercise.NewMongoExerciseRepository(mongo.GetConnection())
}
