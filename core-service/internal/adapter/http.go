package adapter

import (
	"github.com/go-chi/chi/v5"
	user "gym-core-service/internal/core/user/adapter"
	u "gym-core-service/pkg/http"
)

func MakeAppRouter() chi.Router {
	appRouter := chi.NewRouter()

	appRouter.Put("/users/{id}", u.MakeRouteHandler(user.UpdateUserHandler))
	appRouter.Get("/users/{id}", u.MakeRouteHandler(user.GetUserHandler))

	return appRouter
}
