package validator

import (
	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func ValidateJsonStruct(v interface{}) error {
	err := validate.Struct(v)
	if err != nil {
		return err
	}
	return nil
}
