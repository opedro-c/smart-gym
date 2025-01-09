package adapter

import (
	"github.com/go-chi/chi/v5"
	user "gym-core-service/internal/core/user"
	u "gym-core-service/pkg/http"
)

func MakeAppRouter() chi.Router {
	appRouter := chi.NewRouter()

	appRouter.Put("/users/{id}", u.MakeRouteHandler(user.UpdateUserHandler))
	appRouter.Get("/users/{id}", u.MakeRouteHandler(user.GetUserHandler))

	appRouter.Get("/users/{id}/rfids", u.MakeRouteHandler(user.GetUserRfidsHandler))
	appRouter.Post("/users/{id}/rfids", u.MakeRouteHandler(user.CreateUserRfidsHandler))
	appRouter.Delete("/users/{id}/rfids", u.MakeRouteHandler(user.DeleteRfidsUserHandler))

	return appRouter
}
