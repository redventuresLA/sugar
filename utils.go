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
	case "string":
		return handleParseString(input, t)
	case "float64":
		return handleParseFloat(input, t)
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

func handleParseString(input string, v reflect.Value) bool {
	if v.Kind() == reflect.Ptr {
		v.Set(reflect.ValueOf(&input))
	} else {
		v.Set(reflect.ValueOf(input))
	}
	return true
}

func handleParseFloat(input string, v reflect.Value) bool {
	f, e := strconv.ParseFloat(input, 64)
	if e != nil {
		return false
	}
	if v.Kind() == reflect.Ptr {
		v.Set(reflect.ValueOf(&f))
	} else {
		v.Set(reflect.ValueOf(f))
	}
	return true
}

func getFieldName(sf reflect.StructField) string {
	tag := sf.Tag.Get("sugar")
	if tag == "" {
		return sf.Name
	}
	return tag
}
