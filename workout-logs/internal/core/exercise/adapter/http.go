package adapter

import (
	"cloud-gym/internal/core/exercise"
	"cloud-gym/internal/core/exercise/usecases"
	"cloud-gym/internal/mongo"
	utils "cloud-gym/pkg"
	s "cloud-gym/pkg/service"
	"errors"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log/slog"
	"net/http"
	"strconv"
	"time"
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
//func CreateExerciseHandler(w http.ResponseWriter, r *http.Request) error {
//	var input []exercise.ExerciseRecord
//	if err := utils.ParseJson(r, &input); err != nil {
//		return s.NewServiceError(400, err)
//	}
//
//	for _, data := range input {
//		if err := utils.ValidateJsonStruct(&data); err != nil {
//			return s.NewServiceError(400, err)
//		}
//	}
//
//	useCase := usecases.NewCreateExercises(getExerciseRepository())
//
//	result, err := useCase.Execute(input)
//	if err != nil {
//		return err
//	}
//
//	utils.WriteJSON(w, http.StatusCreated, result)
//
//	return nil
//}

// GetExercisesHandler handles fetching exercises based on query parameters.
//
// @Summary      Get exercises based on query parameters
// @Tags         exercises
// @Accept       json
// @Produce      json
// @Param        started_at query uint64 true "Start time"
// @Param        finished_at query uint64 true "Finish time"
// @Param        origin_id query string true "Origin ID"
// @Param        user_id query uint64 true "User ID"
// @Success      200 {object} any
func GetExercisesHandler(w http.ResponseWriter, r *http.Request) error {
	startedAt, err := strconv.ParseInt(r.URL.Query().Get("started_at"), 10, 64)
	finishedAt, err := strconv.ParseInt(r.URL.Query().Get("finished_at"), 10, 64)

	if err != nil {
		return s.NewServiceError(400, err)
	}

	_userId := chi.URLParam(r, "user_id")
	userId, err := strconv.ParseUint(_userId, 10, 64)
	if err != nil {
		return s.NewServiceError(400, err)
	}

	originId := chi.URLParam(r, "origin_id")
	if originId == "" {
		return s.NewServiceError(400, errors.New("origin_id or user_id is empty on path parameter"))
	}

	useCase := usecases.NewGetExercises(getExerciseRepository())

	input := usecases.InputGetExercises{
		StartedAt:  primitive.NewDateTimeFromTime(time.Unix(startedAt, 0)),
		FinishedAt: primitive.NewDateTimeFromTime(time.Unix(finishedAt, 0)),
		OriginId:   originId,
		UserId:     userId,
	}
	result, err := useCase.Execute(input)
	if err != nil {
		return err
	}

	err = utils.WriteJSON(w, http.StatusOK, result)
	if err != nil {
		slog.Debug(err.Error())
	}

	return nil
}

func getExerciseRepository() exercise.ExerciseRepository {
	return exercise.NewMongoExerciseRepository(mongo.GetConnection())
}
