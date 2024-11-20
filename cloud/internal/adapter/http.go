package adapter

import (
	"cloud-gym/internal/core/exercise"
	appHttp "cloud-gym/pkg/http"

	"github.com/go-chi/chi/v5"
)

func MakeAppRouter() chi.Router {
	appRouter := chi.NewRouter()

	// appRouter.Post("/exercises", handlerFn http.HandlerFunc)
	appRouter.Post(
		"/exercises",
		appHttp.MakeRouteHandler(exercise.CreateExerciseHandler),
	)

	return appRouter
}
