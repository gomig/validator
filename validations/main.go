package validations

import "github.com/gomig/validator"

// RegisterExtraValidations register extra validations to validator
func RegisterExtraValidations(v validator.Validator) {
	RegisterAlNumValidation(v)
	RegisterAlNumFaValidation(v)
	RegisterCreditCardValidation(v)
	RegisterIBAN(v)
	RegisterIdentifierValidation(v)
	RegisterIDNumberValidation(v)
	RegisterIPPortValidation(v)
	RegisterJalaaliValidation(v)
	RegisterMobileValidation(v)
	RegisterNationalCodeValidation(v)
	RegisterPostalCodeValidation(v)
	RegisterTelValidation(v)
	RegisterUnsignedValidation(v)
	RegisterUsernameValidation(v)
	RegisterUUIDValidation(v)
}
