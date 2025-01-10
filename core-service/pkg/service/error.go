package service

import (
	"gym-core-service/pkg/controller"
)

type ServiceError struct {
	Code int
	Err  error
}

func (e ServiceError) Error() string {
	return e.Err.Error()
}

func NewServiceError(code int, err error) ServiceError {
	return ServiceError{
		Code: code,
		Err:  err,
	}
}

func (e ServiceError) ToHTTPError() controller.HTTPError {
	return controller.NewHTTPError(e.Code, e.Error())
}
