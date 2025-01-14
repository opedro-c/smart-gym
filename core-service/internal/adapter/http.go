package adapter

import (
	"gym-core-service/internal/core/auth"
	"gym-core-service/internal/core/machine"
	user "gym-core-service/internal/core/user"
	u "gym-core-service/pkg/controller"

	"github.com/go-chi/chi/v5"
)

func MakeAppRouter() chi.Router {
	appRouter := chi.NewRouter()

	// Public for Esp
	appRouter.Get("/rfids/{id}/user", u.MakeRouteHandler(user.GetUserIdOfRfidHandler))

	// For users
	appRouter.Get("/users/{id}", u.MakeRouteHandler(user.GetUserHandler))
	appRouter.Get("/machines", u.MakeRouteHandler(machine.GetMachines))
	appRouter.Post("/auth/login", u.MakeRouteHandler(auth.LoginHandler))

	// For admin
	appRouter.Route("/admin", func(r chi.Router) {
		r.Post("/auth/login", u.MakeRouteHandler(auth.LoginAdminHandler))
		r.Get("/users", u.MakeRouteHandler(user.GetAllUserHandler))
		r.Post("/users", u.MakeRouteHandler(user.CreateUserHandler))
		r.Put("/users/{id}", u.MakeRouteHandler(user.UpdateUserHandler))
		r.Put("/users/{id}/rfids/{rfid}", u.MakeRouteHandler(user.UpdateUserRfidsHandler))

		r.Post("/machines", u.MakeRouteHandler(machine.CreateMachine))
		r.Put("/machines/{id}", u.MakeRouteHandler(machine.UpdateMachine))
		r.Delete("/machines/{id}", u.MakeRouteHandler(machine.DeleteMachine))
	})

	return appRouter
}
