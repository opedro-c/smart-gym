package controller

import (
	"encoding/json"
	"gym-core-service/pkg/error/service_error"
	"gym-core-service/pkg/validator"
	"net/http"
)

func ParseJson(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func ParseAndValidateBody(r *http.Request, v interface{}) error {
	if err := ParseJson(r, v); err != nil {
		return service.NewServiceError(400, err)
	}

	if err := validator.ValidateJsonStruct(v); err != nil {
		return service.NewServiceError(400, err)
	}

	return nil
}
