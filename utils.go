package validator

import (
	"reflect"
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func parseFieldTag(s any, field string, tag string) (string, bool) {
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Struct {
		if f, ok := t.FieldByName(field); ok {
			if tg, ok := f.Tag.Lookup(tag); ok {
				return tg, true
			}
		}
	}
	return "", false
}

func suffixedTagOrFallback(s any, field string, tag string, suffix string, fallback string) string {
	if v, ok := parseFieldTag(s, field, tag+"_"+suffix); ok {
		return v
	} else if v, ok := parseFieldTag(s, field, tag); ok {
		return v
	} else {
		return fallback
	}
}

func formatNumericParam(param string) string {
	fmt := message.NewPrinter(language.English)
	if v, err := strconv.ParseFloat(param, 64); err != nil {
		return fmt.Sprintf("%d", v)
	} else if v, err := strconv.ParseUint(param, 10, 64); err != nil {
		return fmt.Sprintf("%d", v)
	} else if v, err := strconv.ParseInt(param, 10, 64); err != nil {
		return fmt.Sprintf("%d", v)
	}
	return param
}
