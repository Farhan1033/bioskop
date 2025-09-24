package validation

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) error {
	var errMessage []string

	for _, err := range err.(validator.ValidationErrors) {
		switch err.Tag() {
		case "required":
			errMessage = append(errMessage, fmt.Sprintf("%s is required", strings.ToLower(err.Field())))
		case "min":
			errMessage = append(errMessage, fmt.Sprintf("%s must be at least %s characters/value", strings.ToLower(err.Field()), err.Param()))
		case "max":
			errMessage = append(errMessage, fmt.Sprintf("%s must be at most %s characters/value", strings.ToLower(err.Field()), err.Param()))
		case "email":
			errMessage = append(errMessage, fmt.Sprintf("invalid format %s", strings.ToLower(err.Field())))
		case "url":
			errMessage = append(errMessage, fmt.Sprintf("invalid format %s", strings.ToLower(err.Field())))
		default:
			errMessage = append(errMessage, fmt.Sprintf("%s is invalid", strings.ToLower(err.Field())))
		}
	}

	return errors.New(strings.Join(errMessage, ", "))
}
