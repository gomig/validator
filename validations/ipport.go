package validations

import (
	"github.com/go-playground/validator/v10"
	v "github.com/gomig/validator"
)

func ipPortValidation(fl validator.FieldLevel) bool {
	return v.IsIPPort(fl.Field().String())
}

// RegisterIPPortValidation register validations with translations
func RegisterIPPortValidation(val v.Validator) {
	val.AddValidation("ipport", ipPortValidation)
	val.AddTranslation("en", "ipport", "Must be a valid IP:Port combination")
	val.AddTranslation("fa", "ipport", "IP:Port وارد شده معتبر نیست")
}
