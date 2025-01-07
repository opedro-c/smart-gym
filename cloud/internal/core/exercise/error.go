package exercise

import (
	s "cloud-gym/pkg/service"
	"errors"
)

var (
	CannotCreateExerciseError = s.NewServiceError(500, errors.New("cannot create exercise"))
)
