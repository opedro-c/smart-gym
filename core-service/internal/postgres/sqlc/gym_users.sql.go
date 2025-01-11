// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: gym_users.sql

package sqlc

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  username, email, password
) VALUES (
  $1, $2, $3
)
RETURNING id, created_at, username, email, password, rfid
`

type CreateUserParams struct {
	Username string
	Email    string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.Email, arg.Password)
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

const existsUserByEmail = `-- name: ExistsUserByEmail :one
SELECT EXISTS(SELECT 1 FROM users WHERE email = $1 LIMIT 1)
`

func (q *Queries) ExistsUserByEmail(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRowContext(ctx, existsUserByEmail, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const existsUserById = `-- name: ExistsUserById :one
SELECT EXISTS(SELECT 1 FROM users WHERE id = $1 LIMIT 1)
`

func (q *Queries) ExistsUserById(ctx context.Context, id int32) (bool, error) {
	row := q.db.QueryRowContext(ctx, existsUserById, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, created_at, username, email, password, rfid FROM users
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.Username,
			&i.Email,
			&i.Password,
			&i.Rfid,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserById = `-- name: GetUserById :one
SELECT id, created_at, username, email, password, rfid FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
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

const updateUserData = `-- name: UpdateUserData :exec
UPDATE users
  set username = $2,
  email = $3
WHERE id = $1
`

type UpdateUserDataParams struct {
	ID       int32
	Username string
	Email    string
}

func (q *Queries) UpdateUserData(ctx context.Context, arg UpdateUserDataParams) error {
	_, err := q.db.ExecContext(ctx, updateUserData, arg.ID, arg.Username, arg.Email)
	return err
}

const updateUserPassword = `-- name: UpdateUserPassword :exec
UPDATE users
  set password = $2
WHERE id = $1
`

type UpdateUserPasswordParams struct {
	ID       int32
	Password string
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error {
	_, err := q.db.ExecContext(ctx, updateUserPassword, arg.ID, arg.Password)
	return err
}
