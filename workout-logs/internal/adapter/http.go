package adapter

import (
	exercise "cloud-gym/internal/core/exercise/adapter"
	u "cloud-gym/pkg/http"

	"github.com/go-chi/chi/v5"
)

func MakeAppRouter() chi.Router {
	appRouter := chi.NewRouter()

	//appRouter.Post("/exercises", u.MakeRouteHandler(exercise.CreateExerciseHandler))
	appRouter.Get("/users/{user_id}/origins/{origin_id}/exercises", u.MakeRouteHandler(exercise.GetExercisesHandler))
	return appRouter
}
