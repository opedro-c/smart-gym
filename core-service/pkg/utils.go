package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ParseJson(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func WriteJSON(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

var validate = validator.New(validator.WithRequiredStructEnabled())

func ValidateJsonStruct(v interface{}) error {
	err := validate.Struct(v)
	if err != nil {
		return err
	}
	return nil
}
