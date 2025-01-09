package auth

import (
	"errors"
	s "gym-core-service/pkg/service"
)

var (
	ErrNotAllowed = s.NewServiceError(403, errors.New("not allowed"))
)
