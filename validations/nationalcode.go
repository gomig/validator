package validations

import (
	"github.com/go-playground/validator/v10"
	v "github.com/gomig/validator"
)

func nationalCodeValidation(fl validator.FieldLevel) bool {
	return v.IsNationalCode(fl.Field().String())
}

// RegisterNationalCodeValidation register validations with translations
func RegisterNationalCodeValidation(val v.Validator) {
	val.AddValidation("nationalcode", nationalCodeValidation)
	val.AddTranslation("en", "nationalcode", "Must be a valid national code")
	val.AddTranslation("fa", "nationalcode", "کدملی وارد شده معتبر نیست")
}
