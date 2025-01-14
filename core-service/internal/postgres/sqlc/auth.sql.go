// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: auth.sql

package sqlc

import (
	"context"
)

const getAdminByEmailAndPassword = `-- name: GetAdminByEmailAndPassword :one
SELECT id, created_at, username, email, password FROM admin_users
WHERE email = $1
    AND password = $2
`

type GetAdminByEmailAndPasswordParams struct {
	Email    string
	Password string
}

func (q *Queries) GetAdminByEmailAndPassword(ctx context.Context, arg GetAdminByEmailAndPasswordParams) (AdminUser, error) {
	row := q.db.QueryRowContext(ctx, getAdminByEmailAndPassword, arg.Email, arg.Password)
	var i AdminUser
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Username,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const getUserByEmailAndPassword = `-- name: GetUserByEmailAndPassword :one
SELECT id, created_at, username, email, password, rfid FROM users
WHERE email = $1
    AND password = $2
`

type GetUserByEmailAndPasswordParams struct {
	Email    string
	Password string
}

func (q *Queries) GetUserByEmailAndPassword(ctx context.Context, arg GetUserByEmailAndPasswordParams) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmailAndPassword, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Rfid,
	)
	return i, err
}
