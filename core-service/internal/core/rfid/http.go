package rfid

import (
	"context"
	"errors"
	"gym-core-service/internal/postgres/connection"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func rfidService(ctx context.Context) *RfidService {
	repository := NewSqlcRfidRepository(connection.GetConnection())
	return NewRfidService(ctx, repository)
}

func GetUserIdOfRfidHandler(w http.ResponseWriter, r *http.Request) error {
	rfidId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return s.NewServiceError(400, errors.New("rfid id invalid"))
	}

	_, err = rfidService(r.Context()).GetRfidById(int32(rfidId))
	if err != nil {
		return err
	}

	userId, err := rfidService(r.Context()).GetUserIdOfRfid(int32(rfidId))
	if err != nil {
		return err
	}

	utils.WriteJSON(w, http.StatusOK, userId)

	return nil
}
