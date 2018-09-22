package sugar

import (
	"net/url"
	"reflect"
)

func ParseValues(input url.Values, output Sugarable) Response {

	r := Response{}

	elem := reflect.ValueOf(output).Elem()
	typeOfOutput := elem.Type()

	handledUrlValues := make(map[string]bool)

	for i := 0; i < elem.NumField(); i++ {

		structField := typeOfOutput.Field(i)
		valueField := elem.Field(i)

		fieldName := getFieldName(structField)

		handledUrlValues[fieldName] = true

		rawInput := input.Get(fieldName)

		if rawInput == "" {
			if valueField.Kind() != reflect.Ptr {
				r.addFieldError(fieldName, FieldMissingID)
			}
			continue
		}

		ok := parseInputToType(rawInput, valueField)
		if !ok {
			r.addFieldError(fieldName, ValidateFailedID)
			continue
		}
	}

	for key := range input {
		if _, exists := handledUrlValues[key]; !exists {
			r.addExtraFieldError(key)
		}
	}

	if r.HasError() {
		return r
	}

	r.ValidationErrors = output.Validate()

	return r
}
