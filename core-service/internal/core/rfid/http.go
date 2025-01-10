package rfid

import (
	"context"
	"errors"
	"gym-core-service/internal/postgres/connection"
	"net/http"
	"strconv"

	utils "gym-core-service/pkg"
	s "gym-core-service/pkg/service"

	"github.com/go-chi/chi/v5"
)

func rfidService(ctx context.Context) *RfidService {
	repository := NewSqlcRfidRepository(connection.GetConnection())
	return NewRfidService(ctx, repository)
}

// @Router			/rfids/{id}/user [get]
// @Success			200	{object}	any
// @Accept			json
// @Produce			json
func GetUserIdOfRfidHandler(w http.ResponseWriter, r *http.Request) error {
	rfidId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return s.NewServiceError(400, errors.New("rfid id invalid"))
	}

	userId, err := rfidService(r.Context()).GetUserIdOfRfid(int32(rfidId))
	if err != nil {
		return err
	}

	utils.WriteJSON(w, http.StatusOK, userId)

	return nil
}
