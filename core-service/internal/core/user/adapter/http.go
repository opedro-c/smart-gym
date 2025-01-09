package adapter

import (
	"context"
	"errors"
	"gym-core-service/internal/core/user"
	"gym-core-service/internal/postgres/connection"
	utils "gym-core-service/pkg"
	s "gym-core-service/pkg/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func userService(ctx context.Context) *user.UserService {
	repository := user.NewSqlcUserRepository(connection.GetConnection())
	return user.NewUserService(ctx, repository)
}

// @Router			/users/{id} [get]
// @Success			200	{object}	any
// @Param			id   path      int  true  "Account ID"
// @Param			userData body user.UserData true "User Data"
// @Accept			json
// @Produce		json
func GetUserHandler(w http.ResponseWriter, r *http.Request) error {
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return s.NewServiceError(400, errors.New("user id invalid"))
	}

	user, err := userService(r.Context()).GetUserById(int32(userId))
	if err != nil {
		return err
	}

	utils.WriteJSON(w, http.StatusOK, user)

	return nil
}

// @Router			/users/{id} [put]
// @Success		200	{object}	any
// @Param			id   path      int  true  "Account ID"
// @Param			userData body user.UserData true "User Data"
// @Accept			json
// @Produce		json
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) error {
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return s.NewServiceError(400, errors.New("user id invalid"))
	}

	var input user.UserData
	if err := utils.ParseJson(r, &input); err != nil {
		return s.NewServiceError(400, err)
	}

	if err := utils.ValidateJsonStruct(&input); err != nil {
		return s.NewServiceError(400, err)
	}

	// Execute service
	if err := userService(r.Context()).UpdateUserData(int32(userId), input); err != nil {
		return err
	}

	utils.WriteJSON(w, http.StatusOK, input)

	return nil
}
