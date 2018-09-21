package sugar

import (
	"net/url"
	"reflect"
)

func ParseValues(input url.Values, output Sugarable) Response {

	r := Response{}

	elem := reflect.ValueOf(output)
	typeOfOutput := elem.Type()

	handledUrlValues := make(map[string]bool)

	for i := 0; i < elem.NumField(); i++ {

		structField := typeOfOutput.Field(i)
		valueField := elem.Field(i)

		fieldName := getFieldName(structField)

		handledUrlValues[fieldName] = true

		rawInput := input.Get(fieldName)

		if rawInput == "" {
			r.addFieldError(fieldName, FIELD_MISSING_ID)
			continue
		}
		if !valueField.CanSet() {
			// In this case, the output interface has an unexpected type
			r.addFieldError(fieldName, SERVER_ERROR_ID)
			continue
		}
		parsedInput, ok := parseInputToType(rawInput, valueField.Type())
		if !ok {
			r.addFieldError(fieldName, VALIDATE_FAILED_ID)
			continue
		}

		valueField.Set(reflect.ValueOf(parsedInput))
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
