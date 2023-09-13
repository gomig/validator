package validations

import (
	"github.com/go-playground/validator/v10"
	v "github.com/gomig/validator"
)

func postalCodeValidation(fl validator.FieldLevel) bool {
	return v.IsPostalcode(fl.Field().String())
}

// RegisterPostalCodeValidation register validations with translations
func RegisterPostalCodeValidation(val v.Validator) {
	val.AddValidation("postalcode", postalCodeValidation)
	val.AddTranslation("en", "postalcode", "Must be a valid postal-code")
	val.AddTranslation("fa", "postalcode", "کد پستی وارد شده معتبر نیست")
}
