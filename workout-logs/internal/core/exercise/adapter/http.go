package adapter

import (
	"cloud-gym/internal/core/exercise"
	"cloud-gym/internal/core/exercise/usecases"
	"cloud-gym/internal/mongo"
	utils "cloud-gym/pkg"
	s "cloud-gym/pkg/service"
	"net/http"
)

// Create Exercises Handler
//
//	@Summary		Create a couple of exercises
//	@Tags			exercises
//	@Accept			json
//	@Produce		json
//	@Param			exercises body []exercise.ExerciseRecord true "Exercises"
//	@Success		200	{object}	any
//	@Router			/exercises [post]
func CreateExerciseHandler(w http.ResponseWriter, r *http.Request) error {
	var input []exercise.ExerciseRecord
	if err := utils.ParseJson(r, &input); err != nil {
		return s.NewServiceError(400, err)
	}

	for _, data := range input {
		if err := utils.ValidateJsonStruct(&data); err != nil {
			return s.NewServiceError(400, err)
		}
	}

	useCase := usecases.NewCreateExercises(getExerciseRepository())

	result, err := useCase.Execute(input)
	if err != nil {
		return err
	}

	utils.WriteJSON(w, http.StatusCreated, result)

	return nil
}

func getExerciseRepository() exercise.ExerciseRepository {
	return exercise.NewMongoExerciseRepository(mongo.GetConnection())
}
