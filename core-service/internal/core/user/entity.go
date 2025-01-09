package user

import (
	"gym-core-service/internal/postgres/sqlc"
	"time"
)

type UserEntity struct {
	ID        int32     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Enabled   bool      `json:"enabled"`
	Admin     bool      `json:"admin"`
	Data      UserData  `json:"data" validate:"required,dive"`
}

type UserData struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Email    string `json:"email" validate:"required,email"`
}

func FromAdminUserModel(user sqlc.AdminUser) UserEntity {
	return UserEntity{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		Enabled:   true,
		Admin:     true,
		Data: UserData{
			Username: user.Username,
			Email:    user.Email,
		},
	}
}

func FromUserModel(user sqlc.User) UserEntity {
	return UserEntity{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		Enabled:   user.Enabled,
		Admin:     false,
		Data: UserData{
			Username: user.Username,
			Email:    user.Email,
		},
	}
}
