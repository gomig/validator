package validations

import (
	"github.com/go-playground/validator/v10"
	v "github.com/gomig/validator"
)

func ibanValidation(fl validator.FieldLevel) bool {
	return v.IsIBAN(fl.Field().String())
}

// RegisterIBAN register iban number and it translations
func RegisterIBAN(val v.Validator) {
	val.AddValidation("iban", ibanValidation)
	val.AddTranslation("en", "iban", "Must be a valid iran iban number")
	val.AddTranslation("fa", "iban", "شماره شبا وارد شده نامعتبر است")
}
