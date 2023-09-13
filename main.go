package validator

import (
	"github.com/gomig/translator"
)

// NewValidator create new validator
func NewValidator(t translator.Translator, locale string) Validator {
	v := new(validatorDriver)
	v.init(t, locale)
	return v
}

// NewErrorResponse create new error response instance
func NewErrorResponse() ErrorResponse {
	r := new(errorResponseDriver)
	r.init()
	return r
}

// Invalidate generate invalid state for field
func Invalidate(field, err string) ErrorResponse {
	e := NewErrorResponse()
	e.AddError(field, err, "")
	return e
}
