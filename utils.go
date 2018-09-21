package sugar

import (
	"reflect"
	"strconv"
)

func parseInputToType(input string, t reflect.Type) (interface{}, bool) {
	switch name := t.Name(); name {
	case "int":
		return parseStringToInt(input)
	default:
		return 0, false
	}
}

func parseStringToInt(input string) (int, bool) {
	i, e := strconv.Atoi(input)
	return i, e == nil
}

func getFieldName(sf reflect.StructField) string {
	tag := sf.Tag.Get("sugar")
	if tag == "" {
		return sf.Name
	} else {
		return tag
	}
}
