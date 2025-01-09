package user

import (
	"errors"
	s "gym-core-service/pkg/service"
)

var (
	ErrUserNotFound      = s.NewServiceError(404, errors.New("user not found"))
	ErrUserAlreadyExists = s.NewServiceError(409, errors.New("user already exists"))
	ErrPasswordMismatch  = s.NewServiceError(400, errors.New("password mismatch"))
	ErrPasswordTooShort  = s.NewServiceError(400, errors.New("password too short"))
)
