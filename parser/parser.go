package parser

import (
	"reflect"
)

// ParseInputToType Turns a string into the type that is need by "t" and assigns the value to "t". Will return a boolean
// specifying if the operation was successful or not.
func ParseInputToType(input string, t reflect.Value) bool {
	switchType := t.Type()

	if switchType.Kind() == reflect.Slice {
		return parseSlice(input, t, switchType.Elem().Name())
	}

	if switchType.Kind() == reflect.Ptr {
		switchType = switchType.Elem()
	}

	return parsePrimitive(input, t, switchType.Name())

}
