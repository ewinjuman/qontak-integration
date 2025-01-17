package utils

import (
	"fmt"
	Error "github.com/ewinjuman/go-lib/error"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"net/http"
	"strings"
)

// NewValidator func for create a new validator for model fields.
func NewValidator() *validator.Validate {
	// Create a new validator for a Book model.
	validate := validator.New()

	// Custom validation for uuid.UUID fields.
	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return true
		}
		return false
	})

	return validate
}

// ValidatorErrors func for show validation errors for each invalid fields.
func ValidatorErrors(err error) map[string]string {
	// Define fields map.
	fields := map[string]string{}

	// Make error message for each invalid field.
	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Field()] = err.Error()
	}

	return fields
}

// ValidatorMapError func for show validation errors from map error.
func ValidatorMapError(mapErr map[string]interface{}) error {
	var notValid []string

	for k, v := range mapErr {
		myMap := v.(validator.ValidationErrors)
		notValid = append(notValid, k+" "+fmt.Sprintf("%v", myMap.Error()))
	}
	er := strings.Join(notValid, ", ")
	return Error.New(http.StatusBadRequest, "FAILED", er)
}
