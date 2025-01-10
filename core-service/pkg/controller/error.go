package controller

import "fmt"

func NewHTTPError(code int, message string) HTTPError {
	return HTTPError{
		Code:    code,
		Message: message,
	}
}

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.Code, e.Message)
}

type ToHttpError interface {
	error
	ToHTTPError() HTTPError
}
