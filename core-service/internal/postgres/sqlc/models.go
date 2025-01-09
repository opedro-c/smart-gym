// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"time"
)

type AdminUser struct {
	ID        int32
	CreatedAt time.Time
	Username  string
	Email     string
	Password  string
}

type Rfid struct {
	ID        int32
	CreatedAt time.Time
	CardID    string
	UserID    int32
}

type User struct {
	ID        int32
	CreatedAt time.Time
	Enabled   bool
	Username  string
	Email     string
	Password  string
}
