package sugar_test

import (
	"github.com/redventuresLA/sugar"
	"net/url"
	"testing"
)

func GetUrlValues(input map[string]string) url.Values {
	output := url.Values{}
	for k, v := range input {
		output.Set(k, v)
	}
	return output
}

type testType1 struct {
	Field1 int `sugar:"field_1"`
	Field2 int `sugar:"field_2"`
}

func (tt1 testType1) Validate() []sugar.ValidationError {
	return nil
}

func TestParseValues1_1(t *testing.T) {
	output := testType1{}
	input := GetUrlValues(map[string]string{
		"field_1": "10",
		"field_2": "101",
	})
	result := sugar.ParseValues(input, &output)
	if result.HasError() {
		t.Error("Not expecting error from result")
	}
	if output.Field1 != 10 || output.Field2 != 101 {
		t.Error("The parsing was incorrect", output.Field1, output.Field2)
	}
}

type testType2 struct {
	Field2 *int `sugar:"field_2"`
}

func (tt1 testType2) Validate() []sugar.ValidationError {
	return nil
}

func TestParseValues2_1(t *testing.T) {
	output := testType2{}
	input := GetUrlValues(map[string]string{
		"field_2": "101",
	})
	result := sugar.ParseValues(input, &output)
	if result.HasError() {
		t.Error("Not expecting error from result")
	}
	if output.Field2 == nil || *output.Field2 != 101 {
		t.Error("The parsing was incorrect", output.Field2)
	}
}

func TestParseValues2_2(t *testing.T) {
	output := testType2{}
	input := GetUrlValues(map[string]string{})
	result := sugar.ParseValues(input, &output)
	if result.HasError() {
		t.Error("Not expecting error from result")
	}
	if output.Field2 != nil {
		t.Error("The parsing was incorrect", output.Field2)
	}
}
