package validations

import (
	"github.com/go-playground/validator/v10"
	"github.com/gomig/jalaali"
	v "github.com/gomig/validator"
)

func jalaaliValidation(fl validator.FieldLevel) bool {
	if d := jalaali.Parse(fl.Field().String()); d != nil {
		return true
	}
	return false
}

// RegisterJalaaliValidation register validations with translations
func RegisterJalaaliValidation(val v.Validator) {
	val.AddValidation("jalaali", jalaaliValidation)
	val.AddTranslation("en", "jalaali", "Must be a valid jalaali date")
	val.AddTranslation("fa", "jalaali", "تاریخ وارد شده معتبر نیست")
}
