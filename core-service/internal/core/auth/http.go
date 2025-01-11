package auth

import (
	"context"
	"gym-core-service/internal/postgres/connection"
	utils "gym-core-service/pkg"
	"gym-core-service/pkg/controller"
	"net/http"
)

func authService(ctx context.Context) *AuthService {
	return NewAuthService(ctx, NewSqlcAuthRepository(connection.GetConnection()))
}

type Credentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// @Router			/auth/login/ [post]
// @Success			200	{object}	any
// @Param			credentials body auth.Credentials true "Credentials"
// @Accept			json
// @Produce		json
func LoginHandler(w http.ResponseWriter, r *http.Request) error {
	var input Credentials
	if err := controller.ParseAndValidateBody(r, &input); err != nil {
		return err
	}

	user, err := authService(r.Context()).LoginUser(input.Email, input.Password)
	if err != nil {
		return err
	}

	utils.WriteJSON(w, http.StatusOK, user)
	return nil
}

// @Router			/auth/admin/login/ [post]
// @Success			200	{object}	any
// @Param			credentials body auth.Credentials true "Credentials"
// @Accept			json
// @Produce		json
func LoginAdminHandler(w http.ResponseWriter, r *http.Request) error {
	var input Credentials
	if err := controller.ParseAndValidateBody(r, &input); err != nil {
		return err
	}

	user, err := authService(r.Context()).LoginAdminUser(input.Email, input.Password)
	if err != nil {
		return err
	}

	utils.WriteJSON(w, http.StatusOK, user)
	return nil
}
