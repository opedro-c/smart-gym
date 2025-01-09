package rfid

import (
	"context"
	"gym-core-service/internal/postgres/sqlc"
)

type RfidRepository interface {
	CreateRfid(ctx context.Context, userID int32, cardID string) (RfidData, error)
	FindRfidByUserId(ctx context.Context, id int32) ([]int32, error)
	DeleteRfidsByIds(ctx context.Context, ids []int32, userId int32) error
}

type SqlcRfidRepository struct {
	db *sqlc.Queries
}

func NewSqlcRfidRepository(db *sqlc.Queries) RfidRepository {
	return &SqlcRfidRepository{
		db: db,
	}
}

func (r *SqlcRfidRepository) CreateRfid(ctx context.Context, userID int32, cardID string) (RfidData, error) {
	model, err := r.db.CreateRfid(ctx, sqlc.CreateRfidParams{
		UserID: userID,
		CardID: cardID,
	})

	return FromRfidModel(model), err
}

func (r *SqlcRfidRepository) FindRfidByUserId(ctx context.Context, id int32) ([]int32, error) {
	rfids, err := r.db.GetRfidsByUserId(ctx, id)
	if err != nil {
		return nil, err
	}

	ids := make([]int32, len(rfids))
	for i, rfid := range rfids {
		ids[i] = rfid.ID
	}

	return ids, nil
}

func (r *SqlcRfidRepository) DeleteRfidsByIds(ctx context.Context, ids []int32, userId int32) error {
	return r.db.DeleteRfids(ctx, sqlc.DeleteRfidsParams{
		UserID:  userId,
		Column1: ids,
	})
}
