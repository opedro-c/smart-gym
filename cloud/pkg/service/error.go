package service

import "cloud-gym/pkg/http"

type ServiceError struct {
	Code    int
	Message string
}

func (e *ServiceError) Error() string {
	return e.Message
}

func NewServiceError(code int, message string) ServiceError {
	return ServiceError{
		Code:    code,
		Message: message,
	}
}

func (e *ServiceError) ToHTTPError() http.HTTPError {
	return http.NewHTTPError(e.Code, e.Message)
}
