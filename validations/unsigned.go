package validations

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	v "github.com/gomig/validator"
)

func unsignedValidation(fl validator.FieldLevel) bool {
	return v.IsUnsigned(fmt.Sprint(fl.Field().Interface()))
}

// RegisterUnsignedValidation register validations with translations
func RegisterUnsignedValidation(val v.Validator) {
	val.AddValidation("unsigned", unsignedValidation)
	val.AddTranslation("en", "unsigned", "Must be a unsigned number")
	val.AddTranslation("fa", "unsigned", "باید یک عدد صحیح مثبت باشد")
}
