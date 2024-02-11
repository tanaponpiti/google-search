package boothstrap

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()
	Validate.RegisterValidation("username", validateUsername)
	Validate.RegisterValidation("password", validatePassword)
}

func validateUsername(fl validator.FieldLevel) bool {
	usernamePattern := regexp.MustCompile(`^[a-zA-Z0-9_]{3,16}$`)
	return usernamePattern.MatchString(fl.Field().String())
}

func validatePassword(fl validator.FieldLevel) bool {
	passwordPattern := regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()_+={}\[\]:;"'<>,.?\/~-]{8,}$`)
	return passwordPattern.MatchString(fl.Field().String())
}
