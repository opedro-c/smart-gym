package adapter

import (
	"gym-core-service/internal/core/machine"
	"gym-core-service/internal/core/rfid"
	user "gym-core-service/internal/core/user"
	u "gym-core-service/pkg/http"

	"github.com/go-chi/chi/v5"
)

func MakeAppRouter() chi.Router {
	appRouter := chi.NewRouter()

	appRouter.Put("/users/{id}", u.MakeRouteHandler(user.UpdateUserHandler))
	appRouter.Get("/users/{id}", u.MakeRouteHandler(user.GetUserHandler))

	appRouter.Get("/users/{id}/rfids", u.MakeRouteHandler(user.GetUserRfidsHandler))
	appRouter.Post("/users/{id}/rfids", u.MakeRouteHandler(user.CreateUserRfidsHandler))
	appRouter.Delete("/users/{id}/rfids", u.MakeRouteHandler(user.DeleteRfidsUserHandler))

	appRouter.Get("/rfids/{id}/user", u.MakeRouteHandler(rfid.GetUserIdOfRfidHandler))

	appRouter.Get("/machines", u.MakeRouteHandler(machine.GetMachines))
	appRouter.Post("/machines", u.MakeRouteHandler(machine.CreateMachine))
	appRouter.Put("/machines/{id}", u.MakeRouteHandler(machine.UpdateMachine))
	appRouter.Delete("/machines/{id}", u.MakeRouteHandler(machine.UpdateMachine))

	return appRouter
}
