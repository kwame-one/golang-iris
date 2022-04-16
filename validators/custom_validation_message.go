package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type validationError struct {
	ActualTag string `json:"tag"`
	Field     string `json:"field"`
	Message   string `json:"message"`
}

func CustomMessages(errs validator.ValidationErrors) []validationError {
	validationErrors := make([]validationError, 0, len(errs))
	for _, validationErr := range errs {

		validationErrors = append(validationErrors, validationError{

			ActualTag: validationErr.ActualTag(),
			Field:     validationErr.Field(),
			Message:   fmt.Sprintf("validation failed for %s field on the %s tag", validationErr.Field(), validationErr.ActualTag()),
		})
	}

	return validationErrors
}
