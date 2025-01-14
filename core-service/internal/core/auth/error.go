package auth

import (
	"errors"
	s "gym-core-service/pkg/error/service_error"
)

var (
	ErrNotAllowed      = s.NewServiceError(403, errors.New("not allowed"))
	ErrInvalidPassword = s.NewServiceError(400, errors.New("invalid password"))
)
