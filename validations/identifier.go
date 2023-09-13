package validations

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	v "github.com/gomig/validator"
)

func identifierValidation(fl validator.FieldLevel) bool {
	return v.IsIdentifier(fmt.Sprint(fl.Field().Interface()))
}

// RegisterIdentifierValidation register identifier validator and it translations
func RegisterIdentifierValidation(val v.Validator) {
	val.AddValidation("identifier", identifierValidation)
	val.AddTranslation("en", "identifier", "Must be a valid (numeric) identifier")
	val.AddTranslation("fa", "identifier", "شناسه وارد شده معتبر نیست")
}
