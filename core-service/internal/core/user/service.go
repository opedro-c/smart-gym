package user

import (
	"context"
)

type UserService struct {
	ctx        context.Context
	repository UserRepository
}

func NewUserService(ctx context.Context, repository UserRepository) *UserService {
	return &UserService{ctx, repository}
}

func (u *UserService) GetUserById(id int32) (*UserEntity, error) {
	exists, err := u.ExistsUserById(id)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, ErrUserNotFound
	}

	user, err := u.repository.GetById(u.ctx, id)
	return &user, err
}

func (u *UserService) ExistsUserById(id int32) (bool, error) {
	is_exist, err := u.repository.ExistsUserById(u.ctx, id)
	return is_exist, err
}

func (u *UserService) ExistsUserByEmail(email string) (bool, error) {
	is_exist, err := u.repository.ExistsUserByEmail(u.ctx, email)
	return is_exist, err
}

func (u *UserService) CreateUser(userData UserData, password string) (*UserEntity, error) {
	if len(password) < 3 {
		return nil, ErrPasswordTooShort
	}

	exists, err := u.ExistsUserByEmail(userData.Email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, ErrUserAlreadyExists
	}

	userCreated, err := u.repository.CreateUser(u.ctx, userData, password)

	return &userCreated, err
}

func (u *UserService) UpdateUserData(id int32, user UserData) error {
	return u.repository.UpdateUser(u.ctx, id, user)
}

func (u *UserService) UpdateUserPassword(id int32, password string) error {
	return u.repository.UpdateUserPassword(u.ctx, id, password)
}
