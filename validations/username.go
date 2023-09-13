package validations

import (
	"github.com/go-playground/validator/v10"
	v "github.com/gomig/validator"
)

func usernameValidation(fl validator.FieldLevel) bool {
	return v.IsUsername(fl.Field().String())
}

// RegisterUsernameValidation register validations with translations
func RegisterUsernameValidation(val v.Validator) {
	val.AddValidation("username", usernameValidation)
	val.AddTranslation("en", "username", "Must be a valid username")
	val.AddTranslation("fa", "username", "مقدار ورودی نمی تواند شامل کاراکترهای غیرمجاز باشد")
}
