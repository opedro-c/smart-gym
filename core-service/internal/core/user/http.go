package user

import (
	"context"
	"errors"
	"gym-core-service/internal/postgres/connection"
	utils "gym-core-service/pkg"
	"gym-core-service/pkg/controller"
	s "gym-core-service/pkg/error/service_error"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func userService(ctx context.Context) *UserService {
	repository := NewSqlcUserRepository(connection.GetConnection())
	return NewUserService(ctx, repository)
}

// @Router			/users/{id} [get]
// @Success			200	{object}	any
// @Param			id   path      int  true  "Account ID"
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

// @Router			/admin/users/ [get]
// @Success			200	{object}	any
// @Accept			json
// @Produce		json
func GetAllUserHandler(w http.ResponseWriter, r *http.Request) error {
	user, err := userService(r.Context()).GetAllUsers()
	if err != nil {
		return err
	}

	utils.WriteJSON(w, http.StatusOK, user)

	return nil
}

// @Router			/admin/users/ [post]
// @Success			200	{object}	any
// @Param			userData body user.UserData true "User Data"
// @Accept			json
// @Produce		json
func CreateUserHandler(w http.ResponseWriter, r *http.Request) error {
	var input UserData
	if err := controller.ParseAndValidateBody(r, &input); err != nil {
		return err
	}

	user, err := userService(r.Context()).CreateUser(input)
	if err != nil {
		return err
	}

	utils.WriteJSON(w, http.StatusCreated, user)
	return nil
}

// @Router			/admin/users/{id} [put]
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

	var input UserData
	if err := controller.ParseAndValidateBody(r, &input); err != nil {
		return err
	}

	if err := userService(r.Context()).UpdateUserData(int32(userId), input); err != nil {
		return err
	}

	utils.WriteJSON(w, http.StatusOK, input)

	return nil
}

// @Router			/admin/users/{id}/rfids/{rfid} [put]
// @Success			200	{object}	any
// @Param			id   path      int  true  "Account ID"
// @Param			rfid   path      string  true  "RFID"
// @Accept			json
// @Produce			json
func UpdateUserRfidsHandler(w http.ResponseWriter, r *http.Request) error {
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return s.NewServiceError(400, errors.New("user id invalid"))
	}

	rfid := chi.URLParam(r, "rfid")
	if rfid == "" {
		return s.NewServiceError(400, errors.New("rfid id invalid"))
	}

	err = userService(r.Context()).UpdateUserRfid(int32(userId), rfid)
	if err != nil {
		return err
	}

	utils.WriteJSON(w, http.StatusNoContent, nil)

	return nil
}

// @Router			/rfids/{id}/user [get]
// @Success			200	{object}	any
// @Param			id   path      string  true  "Account ID"
// @Accept			json
// @Produce			json
func GetUserIdOfRfidHandler(w http.ResponseWriter, r *http.Request) error {
	rfid := chi.URLParam(r, "id")
	if rfid == "" {
		return s.NewServiceError(400, errors.New("rfid id invalid"))
	}

	userId, err := userService(r.Context()).GetUserIdByRfid(rfid)
	if err != nil {
		return err
	}

	utils.WriteJSON(w, http.StatusOK, userId)

	return nil
}
