package user

import (
	"context"
	"gym-core-service/internal/postgres/sqlc"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user UserData) (UserEntity, error)
	GetById(ctx context.Context, id int32) (UserEntity, error)
	GetAllUsers(ctx context.Context) ([]UserEntity, error)
	ExistsUserById(ctx context.Context, id int32) (bool, error)
	ExistsUserByEmail(ctx context.Context, email string) (bool, error)
	UpdateUser(ctx context.Context, id int32, user UserData) error
	UpdateAdminUser(ctx context.Context, id int32, user UserData) error
	UpdateUserPassword(ctx context.Context, id int32, password string) error
	UpdateAdminUserPassword(ctx context.Context, id int32, password string) error
	UpdateUserRfid(ctx context.Context, id int32, rfid string) error
	GetUserIdByRfid(ctx context.Context, rfid string) (int32, error)
}

type SqlcUserRepository struct {
	db *sqlc.Queries
}

func NewSqlcUserRepository(db *sqlc.Queries) UserRepository {
	return &SqlcUserRepository{
		db: db,
	}
}

func (r *SqlcUserRepository) CreateUser(ctx context.Context, user UserData) (UserEntity, error) {
	userToCreate := sqlc.CreateUserParams{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	model_created, err := r.db.CreateUser(ctx, userToCreate)

	return FromUserModel(model_created), err
}

func (r *SqlcUserRepository) GetById(ctx context.Context, id int32) (UserEntity, error) {
	model, err := r.db.GetUserById(ctx, id)
	return FromUserModel(model), err
}

func (r *SqlcUserRepository) GetAllUsers(ctx context.Context) ([]UserEntity, error) {
	models, err := r.db.GetAllUsers(ctx)
	users := make([]UserEntity, len(models))
	for i, model := range models {
		users[i] = FromUserModel(model)
	}
	return users, err
}

func (r *SqlcUserRepository) ExistsUserById(ctx context.Context, id int32) (bool, error) {
	return r.db.ExistsUserById(ctx, id)
}

func (r *SqlcUserRepository) ExistsUserByEmail(ctx context.Context, email string) (bool, error) {
	return r.db.ExistsUserByEmail(ctx, email)
}

func (r *SqlcUserRepository) UpdateAdminUser(ctx context.Context, id int32, user UserData) error {
	return r.db.UpdateAdminUser(ctx, sqlc.UpdateAdminUserParams{
		ID:       id,
		Username: user.Username,
		Email:    user.Email,
	})
}

func (r *SqlcUserRepository) UpdateUser(ctx context.Context, id int32, user UserData) error {
	return r.db.UpdateUserData(ctx, sqlc.UpdateUserDataParams{
		ID:       id,
		Username: user.Username,
		Email:    user.Email,
	})
}

func (r *SqlcUserRepository) UpdateUserPassword(ctx context.Context, id int32, password string) error {
	return r.db.UpdateUserPassword(ctx, sqlc.UpdateUserPasswordParams{
		ID:       id,
		Password: password,
	})
}

func (r *SqlcUserRepository) UpdateAdminUserPassword(ctx context.Context, id int32, password string) error {
	return r.db.UpdateAdminUserPassword(ctx, sqlc.UpdateAdminUserPasswordParams{
		ID:       id,
		Password: password,
	})
}

func (r *SqlcUserRepository) UpdateUserRfid(ctx context.Context, id int32, rfid string) error {
	return r.db.UpdateUserRfid(ctx, sqlc.UpdateUserRfidParams{
		ID:   id,
		Rfid: rfid,
	})
}

func (r *SqlcUserRepository) GetUserIdByRfid(ctx context.Context, rfid string) (int32, error) {
	return r.db.GetUserIdByRfidId(ctx, rfid)
}
