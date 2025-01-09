package user

import (
	"context"
	"errors"
	"gym-core-service/internal/core/rfid"
	"gym-core-service/internal/postgres/connection"
	utils "gym-core-service/pkg"
	s "gym-core-service/pkg/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func userService(ctx context.Context) *UserService {
	repository := NewSqlcUserRepository(connection.GetConnection())
	return NewUserService(ctx, repository)
}

func rfidService(ctx context.Context) *rfid.RfidService {
	repository := rfid.NewSqlcRfidRepository(connection.GetConnection())
	return rfid.NewRfidService(ctx, repository)
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

	var input UserData
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

// @Router			/users/{id}/rfids [get]
// @Success			200	{object}	any
// @Param			id   path      int  true  "Account ID"
// @Accept			json
// @Produce		json
func GetUserRfidsHandler(w http.ResponseWriter, r *http.Request) error {
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return s.NewServiceError(400, errors.New("user id invalid"))
	}

	_, err = userService(r.Context()).GetUserById(int32(userId))
	if err != nil {
		return err
	}

	rfids, err := rfidService(r.Context()).FindRfidByUserId(int32(userId))
	if err != nil {
		return err
	}

	utils.WriteJSON(w, http.StatusOK, rfids)

	return nil
}

// @Router			/users/{id}/rfids [post]
// @Success			200	{object}	any
// @Param			id   path      int  true  "Account ID"
// @Accept			json
// @Produce		json
func CreateUserRfidsHandler(w http.ResponseWriter, r *http.Request) error {
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return s.NewServiceError(400, errors.New("user id invalid"))
	}

	_, err = userService(r.Context()).GetUserById(int32(userId))
	if err != nil {
		return err
	}

	cardId := r.URL.Query().Get("card_id")
	if cardId == "" {
		return s.NewServiceError(400, errors.New("card_id is required"))
	}

	createdData, err := rfidService(r.Context()).CreateRfid(int32(userId), cardId)
	if err != nil {
		return err
	}

	utils.WriteJSON(w, http.StatusAccepted, createdData)

	return nil
}

// @Router			/users/{id}/rfids [delete]
// @Success			200	{object}	any
// @Param			id   path      int  true  "Account ID"
// @Param			rfidsIds body []int32 true "Rfid IDS"
// @Accept			json
// @Produce		json
func DeleteRfidsUserHandler(w http.ResponseWriter, r *http.Request) error {
	userId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return s.NewServiceError(400, errors.New("user id invalid"))
	}

	_, err = userService(r.Context()).GetUserById(int32(userId))
	if err != nil {
		return err
	}

	var input []int32
	if err := utils.ParseJson(r, &input); err != nil {
		return s.NewServiceError(400, err)
	}

	err = rfidService(r.Context()).DeleteRfidsByIds(input, int32(userId))
	if err != nil {
		return err
	}

	utils.WriteJSON(w, http.StatusAccepted, nil)

	return nil
}
