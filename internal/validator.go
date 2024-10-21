package internal

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var requiredTagMessages = map[string]string{
	"required":      "Field is required",
	"requiredParam": "Parameter is required", // still use "required" for param tag on your struct
	"email":         "Email is not valid",
	"max":           "Exceeds maximum length",
	"min":           "Does not meet minimum length",
	"len":           "Does not meet length",
	"lte":           "Exceeds maximum value",
	"gte":           "Does not meet minimum value",
}

func ValidateRequest(req *interface{}, validatorType string) (*ValidateReturn, error) {
	var modelTag string
	var errorMessage string

	if validatorType == "body" {
		modelTag = "body"
		errorMessage = "Invalid body type"
	} else if validatorType == "param" {
		modelTag = "form"
		errorMessage = "Invalid param type"
	} else if validatorType == "uri" {
		modelTag = "uri"
		errorMessage = "Invalid uri type"
	} else {
		return nil, fmt.Errorf("invalid validator type")
	}

	val := validator.New()
	val.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get(modelTag), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := val.Struct(*req)
	if err == nil {
		return nil, nil
	}

	var valErrors validator.ValidationErrors
	if !errors.As(err, &valErrors) {
		return nil, err
	}

	errors := make([]RequiredField, len(valErrors))
	for i, valError := range valErrors {
		errorTag := valError.Tag()
		if errorTag == "required" && validatorType == "param" {
			errorTag = "required-param"
		}

		message, ok := requiredTagMessages[errorTag]
		if !ok {
			message = errorTag + " is not valid"
		}

		errors[i] = RequiredField{
			Field:   strings.ToLower(valError.Field()),
			Message: message,
		}
	}

	return &ValidateReturn{
		Message: errorMessage,
		Errors:  errors,
	}, nil
}
