package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	configValidator "github.com/layerpro/upload-download-backend/clients/validator"
)

func Validator(data any) (string, error) {

	validate := configValidator.Validate
	err := validate.Struct(data)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		firstError := validationErrors[0]
		customMessage := getCustomErrorMessage(firstError)
		return customMessage, err
	}
	return "", nil
}

// Fungsi untuk menghasilkan pesan error kustom
func getCustomErrorMessage(err validator.FieldError) string {
	msg, exists := validationMessages[err.Tag()]
	if exists {
		return safeSprintf(msg, err.Field(), err.Param())
	}
	return err.Error()
}

var validationMessages = map[string]string{
	"required":          "%s is required",
	"min":               "%s length must be at least %s characters",
	"max":               "%s length must be less than or equal to %s characters",
	"email":             "%s must be a valid email",
	"alpha":             "%s must only contain alpha characters",
	"alphaSpace":        "%s must only contain alpha characters and space",
	"alphanumericSpace": "%s must only contain alphanumeric characters and space",
	"oneof":             "%s must be one of %s",
}

// handle singel of multiple params
func safeSprintf(format string, params ...interface{}) string {
	expectedPlaceholders := strings.Count(format, "%s")
	if len(params) > expectedPlaceholders {
		params = params[:expectedPlaceholders]
	}
	return fmt.Sprintf(format, params...)
}
