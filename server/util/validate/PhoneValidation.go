package validate

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func PhoneValidation(fl validator.FieldLevel) bool {
	phoneRegex := regexp.MustCompile(`^\d{3}-\d{4}-\d{4}$`)
	return phoneRegex.MatchString(fl.Field().String())
}
