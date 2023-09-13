package validations

import (
	"github.com/go-playground/validator/v10"
	"github.com/gomig/utils"
	v "github.com/gomig/validator"
)

func alNumValidation(fl validator.FieldLevel) bool {
	v := fl.Field().String()
	return utils.ExtractAlphaNum(v, fl.Param()) == v
}

// RegisterAlNumValidation register validations with translations
func RegisterAlNumValidation(val v.Validator) {
	val.AddValidation("alnum", alNumValidation)
	val.AddTranslation("en", "alnum", "Only alpha and numbers valid")
	val.AddTranslation("fa", "alnum", "فقط حروف و اعداد انگلیسی مجاز است")
}
