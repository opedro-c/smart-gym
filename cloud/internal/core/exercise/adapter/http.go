package adapter

import (
	"cloud-gym/internal/core/exercise"
	"cloud-gym/internal/core/exercise/usecases"
	"cloud-gym/internal/mongo"
	"cloud-gym/pkg/service"
	"cloud-gym/pkg/service/user"
	"net/http"
)

func CreateExerciseHandler(w http.ResponseWriter, r *http.Request) error {
	exerciseRepository := exercise.NewMongoExerciseRepository(mongo.GetConnection())

	useCase := usecases.NewCreateExercise(exerciseRepository)

	result, err := useCase.Execute()
	if err != nil {
		return err
	}

	return http.WriteJSON(w, result)
}
