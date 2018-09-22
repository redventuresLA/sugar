package sugar

import (
	"reflect"
	"strconv"
)

func parseInputToType(input string, t reflect.Value) bool {
	switchType := t.Type()
	if switchType.Kind() == reflect.Ptr {
		switchType = switchType.Elem()
	}

	switch name := switchType.Name(); name {
	case "int":
		return handleParseInt(input, t)
	default:
		return false
	}
}

func handleParseInt(input string, v reflect.Value) bool {
	i, e := strconv.Atoi(input)
	if e != nil {
		return false
	}
	if v.Kind() == reflect.Ptr {
		v.Set(reflect.ValueOf(&i))
	} else {
		v.Set(reflect.ValueOf(i))
	}
	return true
}

func getFieldName(sf reflect.StructField) string {
	tag := sf.Tag.Get("sugar")
	if tag == "" {
		return sf.Name
	} else {
		return tag
	}
}
