package auth

import (
	"context"
	"gym-core-service/internal/core/user"
	"gym-core-service/internal/postgres/sqlc"
)

type AuthRepository interface {
	GetUserByEmailAndPassword(ctx context.Context, email string, password string) (user.UserEntity, error)
	GetAdminUserByEmailAndPassword(ctx context.Context, email string, password string) (user.UserEntity, error)
}

type SqlcAuthRepository struct {
	db *sqlc.Queries
}

func NewSqlcAuthRepository(db *sqlc.Queries) AuthRepository {
	return &SqlcAuthRepository{
		db: db,
	}
}

func (r *SqlcAuthRepository) GetUserByEmailAndPassword(ctx context.Context, email string, password string) (user.UserEntity, error) {
	userModel, err := r.db.GetUserByEmailAndPassword(ctx, sqlc.GetUserByEmailAndPasswordParams{
		Email:    email,
		Password: password,
	})

	return user.FromUserModel(userModel), err
}

func (r *SqlcAuthRepository) GetAdminUserByEmailAndPassword(ctx context.Context, email string, password string) (user.UserEntity, error) {
	userModel, err := r.db.GetAdminByEmailAndPassword(ctx, sqlc.GetAdminByEmailAndPasswordParams{
		Email:    email,
		Password: password,
	})

	return user.FromAdminUserModel(userModel), err
}
