package parser

import (
	"reflect"
	"strconv"
	"strings"
)

func parsePrimitive(input string, v reflect.Value, name string) bool {
	switch name{
	case intType:
		return handleParseInt(input, v)
	case stringType:
		return handleParseString(input, v)
	case float64Type:
		return handleParseFloat(input, v)
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
	} else if v.Kind() == reflect.Slice {
		output := strings.Split(input, ",")
		v.Set(reflect.ValueOf(output))
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
