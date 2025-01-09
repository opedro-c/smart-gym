package rfid

import (
	"gym-core-service/internal/postgres/sqlc"
)

type RfidData struct {
	ID     int32  `json:"id"`
	UserID int32  `json:"user_id"`
	CardID string `json:"card_id"`
}

func FromRfidModel(rfid sqlc.Rfid) RfidData {
	return RfidData{
		ID:     rfid.ID,
		UserID: rfid.UserID,
		CardID: rfid.CardID,
	}
}
