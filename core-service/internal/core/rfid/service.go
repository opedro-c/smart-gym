package rfid

import (
	"context"
)

type RfidService struct {
	ctx        context.Context
	repository RfidRepository
}

func NewRfidService(ctx context.Context, repository RfidRepository) *RfidService {
	return &RfidService{ctx, repository}
}

func (u *RfidService) CreateRfid(userID int32, cardID string) (*RfidData, error) {
	data, err := u.repository.CreateRfid(u.ctx, userID, cardID)
	return &data, err
}

func (u *RfidService) FindRfidByUserId(id int32) ([]RfidData, error) {
	return u.repository.FindRfidByUserId(u.ctx, id)
}

func (u *RfidService) DeleteRfidsByIds(ids []int32, userId int32) error {
	return u.repository.DeleteRfidsByIds(u.ctx, ids, userId)
}
