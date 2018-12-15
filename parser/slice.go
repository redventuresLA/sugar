package parser

import (
	"reflect"
	"strconv"
	"strings"
)

func parseSlice(input string, v reflect.Value, name string) bool {
	split := strings.Split(input, ",")
	switch name {
	case intType:
		return handleParseIntSlice(split, v)
	case stringType:
		return handleParseStringSlice(split, v)
	case float64Type:
		return handleParseFloatSlice(split, v)
	case boolType:
		return handleBoolSlice(split, v)
	default:
		return false
	}
}

func handleParseStringSlice(input []string, v reflect.Value) bool {
	v.Set(reflect.ValueOf(input))
	return true
}

func handleParseFloatSlice(input []string, v reflect.Value) bool {
	output := make([]float64, len(input))

	for idx, val := range input {
		if parsed, err := strconv.ParseFloat(val, 64); err == nil {
			output[idx] = parsed
		} else {
			return false
		}
	}
	v.Set(reflect.ValueOf(output))
	return true
}

func handleParseIntSlice(input []string, v reflect.Value) bool {
	output := make([]int, len(input))

	for idx, val := range input {
		if parsed, err := strconv.Atoi(val); err == nil {
			output[idx] = parsed
		} else {
			return false
		}
	}

	v.Set(reflect.ValueOf(output))
	return true
}

func handleBoolSlice(input []string, v reflect.Value) bool {
	output := make([]bool, len(input))

	for idx, val := range input {
		if parsed, err := strconv.ParseBool(val); err == nil {
			output[idx] = parsed
		} else {
			return false
		}
	}
	v.Set(reflect.ValueOf(output))
	return true
}
