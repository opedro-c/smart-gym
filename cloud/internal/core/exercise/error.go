package exercise

import (
	s "cloud-gym/pkg/service"
)

var (
	CannotCreateExerciseError = s.NewServiceError(500, "cannot create exercise")
)
