package validations

import (
	"github.com/go-playground/validator/v10"
	"github.com/gomig/utils"
	v "github.com/gomig/validator"
)

func alNumFaValidation(fl validator.FieldLevel) bool {
	v := fl.Field().String()
	return utils.ExtractAlphaNumPersian(v, fl.Param()) == v
}

// RegisterAlNumFaValidation register validations with translations
func RegisterAlNumFaValidation(val v.Validator) {
	val.AddValidation("alnumfa", alNumFaValidation)
	val.AddTranslation("en", "alnumfa", "Only alpha (en-fa) and numbers valid")
	val.AddTranslation("fa", "alnumfa", "فقط حروف فارسی و انگلیسی و عدد مجاز است")
}
