package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gomig/translator"
)

type validatorDriver struct {
	V      *validator.Validate
	T      translator.Translator
	locale string
}

func (v *validatorDriver) init(t translator.Translator, locale string) {
	v.V = validator.New()
	v.V.RegisterTagNameFunc(func(fld reflect.StructField) string {
		var name string
		if name = strings.SplitN(fld.Tag.Get("field"), ",", 2)[0]; name != "" {
			return name
		} else if name = strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]; name != "" {
			return name
		} else if name = strings.SplitN(fld.Tag.Get("form"), ",", 2)[0]; name != "" {
			return name
		} else if name = strings.SplitN(fld.Tag.Get("xml"), ",", 2)[0]; name != "" {
			return name
		}
		return fld.Name
	})
	v.T = t
	v.locale = locale
}

func (v *validatorDriver) proccessStructValidation(locale string, s any, err any) ErrorResponse {
	res := NewErrorResponse()
	if errors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errors {
			title := suffixedTagOrFallback(s, e.StructField(), "vTitle", locale, e.Field())
			param := suffixedTagOrFallback(s, e.StructField(), "vParam", locale, "")
			if param == "" {
				param = e.Param()
				if _, ok := parseFieldTag(s, e.StructField(), "vFormat"); ok {
					param = formatNumericParam(param)
				}
			}
			res.AddError(e.Field(), e.Tag(), v.TranslateStruct(s, locale, e.Tag(), e.ActualTag(), map[string]string{
				"field": title,
				"param": param,
			}))
		}
	}
	return res
}

func (v *validatorDriver) proccessVarValidation(locale string, params ValidatorParam, messages map[string]string, err any) ErrorResponse {
	res := NewErrorResponse()
	if errors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errors {
			if msg, ok := messages[e.Tag()]; ok {
				res.AddError(params.Name, e.Tag(), msg)
			} else {
				title := params.Title
				if title == "" {
					title = params.Name
				}
				param := params.ParamTitle
				if param == "" {
					param = e.Param()
					if params.Format {
						param = formatNumericParam(param)
					}
				}
				res.AddError(params.Name, e.Tag(), v.Translate(locale, e.Tag(), map[string]string{
					"field": title,
					"param": param,
				}))
			}
		}
	}
	return res
}

func (v *validatorDriver) Validator() *validator.Validate {
	return v.V
}

func (v *validatorDriver) AddValidation(tag string, val validator.Func) {
	v.V.RegisterValidation(tag, val)
}

func (v *validatorDriver) AddTranslation(locale string, key string, message string) {
	v.T.Register(locale, key, message)
}

func (v *validatorDriver) Translate(locale string, key string, placeholders map[string]string) string {
	return v.T.Translate(locale, key, placeholders)
}

func (v *validatorDriver) TranslateStruct(s any, locale string, key string, field string, placeholders map[string]string) string {
	return v.T.TranslateStruct(s, locale, key, field, placeholders)
}

func (v *validatorDriver) StructLocale(locale string, s any) ErrorResponse {
	err := v.V.Struct(s)
	if err != nil {
		return v.proccessStructValidation(locale, s, err)
	}
	return NewErrorResponse()
}

func (v *validatorDriver) Struct(s any) ErrorResponse {
	return v.StructLocale(v.locale, s)
}

func (v *validatorDriver) StructExceptLocale(locale string, s any, fields ...string) ErrorResponse {
	err := v.V.StructExcept(s, fields...)
	if err != nil {
		return v.proccessStructValidation(locale, s, err)
	}
	return NewErrorResponse()
}

func (v *validatorDriver) StructExcept(s any, fields ...string) ErrorResponse {
	return v.StructExceptLocale(v.locale, s, fields...)
}

func (v *validatorDriver) StructPartialLocale(locale string, s any, fields ...string) ErrorResponse {
	err := v.V.StructPartial(s, fields...)
	if err != nil {
		return v.proccessStructValidation(locale, s, err)
	}
	return NewErrorResponse()
}

func (v *validatorDriver) StructPartial(s any, fields ...string) ErrorResponse {
	return v.StructPartialLocale(v.locale, s, fields...)
}

func (v *validatorDriver) VarLocale(locale string, params ValidatorParam, field any, tag string, messages map[string]string) ErrorResponse {
	if params.Name == "" {
		panic("Name field must passed to ValidatorParam")
	}
	err := v.V.Var(field, tag)
	if err != nil {
		return v.proccessVarValidation(locale, params, messages, err)
	}
	return NewErrorResponse()
}

func (v *validatorDriver) Var(params ValidatorParam, field any, tag string, messages map[string]string) ErrorResponse {
	return v.VarLocale(v.locale, params, field, tag, messages)
}

func (v *validatorDriver) VarWithValueLocale(locale string, params ValidatorParam, field any, other any, tag string, messages map[string]string) ErrorResponse {
	if params.Name == "" {
		panic("Name field must passed to ValidatorParam")
	}
	err := v.V.VarWithValue(field, other, tag)
	if err != nil {
		return v.proccessVarValidation(locale, params, messages, err)
	}
	return NewErrorResponse()
}

func (v *validatorDriver) VarWithValue(params ValidatorParam, field any, other any, tag string, messages map[string]string) ErrorResponse {
	return v.VarWithValueLocale(v.locale, params, field, other, tag, messages)
}
