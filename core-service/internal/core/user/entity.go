package user

import (
	"gym-core-service/internal/postgres/sqlc"
	"time"
)

type UserEntity struct {
	ID        int32     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Admin     bool      `json:"admin"`
	Data      UserData  `json:"data" validate:"required,dive"`
	Rfid      string    `json:"rfid"`
}

type UserData struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3,max=32"`
}

func FromAdminUserModel(user sqlc.AdminUser) UserEntity {
	return UserEntity{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		Admin:     true,
		Rfid:      "",
		Data: UserData{
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
		},
	}
}

func FromUserModel(user sqlc.User) UserEntity {
	return UserEntity{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		Admin:     false,
		Rfid:      user.Rfid,
		Data: UserData{
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
		},
	}
}
