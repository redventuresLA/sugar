package sugar_test

import (
	"net/url"
	"testing"

	"github.com/redventuresLA/sugar"
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

type testType3 struct {
	Field1 string `sugar:"field_1"`
}

func (tt testType3) Validate() []sugar.ValidationError {
	return nil
}

func TestParseValues3_1(t *testing.T) {
	output := testType3{}
	input := GetUrlValues(map[string]string{
		"field_1": "xyz",
	})
	result := sugar.ParseValues(input, &output)
	if result.HasError() {
		t.Error("Not expecting error")
	}
	if output.Field1 != "xyz" {
		t.Error("The parsing was incorrect", output.Field1)
	}
}

type testType4 struct {
	Field1 *string `sugar:"field_1"`
}

func (tt testType4) Validate() []sugar.ValidationError {
	return nil
}

func TestParseValues4_1(t *testing.T) {
	output := testType4{}
	input := GetUrlValues(map[string]string{
		"field_1": "abc",
	})
	result := sugar.ParseValues(input, &output)
	if result.HasError() {
		t.Error("Not expecting error")
	}
	if output.Field1 == nil || *output.Field1 != "abc" {
		t.Error("The parsing was incorrect", output.Field1)
	}
}

func TestParseValues4_2(t *testing.T) {
	output := testType4{}
	input := GetUrlValues(map[string]string{})
	result := sugar.ParseValues(input, &output)
	if result.HasError() {
		t.Error("Not expecting error")
	}
	if output.Field1 != nil {
		t.Error("The parsing was incorrect", output.Field1)
	}
}

type testType5 struct {
	Field1 float64 `sugar:"field_1"`
}

func (tt testType5) Validate() []sugar.ValidationError {
	return nil
}

func TestParseValues5_1(t *testing.T) {
	output := testType5{}
	input := GetUrlValues(map[string]string{
		"field_1": "3.14",
	})
	result := sugar.ParseValues(input, &output)
	if result.HasError() {
		t.Error("Not expecting error")
	}
	if output.Field1 != 3.14 {
		t.Error("The parsing was incorrect", output.Field1)
	}
}

type testType6 struct {
	Field1 *float64 `sugar:"field_1"`
}

func (tt testType6) Validate() []sugar.ValidationError {
	return nil
}

func TestParseValues6_1(t *testing.T) {
	output := testType6{}
	input := GetUrlValues(map[string]string{
		"field_1": "3.14",
	})
	result := sugar.ParseValues(input, &output)
	if result.HasError() {
		t.Error("Not expecting error")
	}
	if result.HumanReadableError() != nil {
		t.Error("Did not have human readable error")
		return
	}
	if output.Field1 == nil || *output.Field1 != 3.14 {
		t.Error("The parsing was incorrect", output.Field1)
	}
}

func TestParseValues6_2(t *testing.T) {
	output := testType6{}
	input := GetUrlValues(map[string]string{})
	result := sugar.ParseValues(input, &output)
	if result.HasError() {
		t.Error("Not expecting error")
	}
	if output.Field1 != nil {
		t.Error("The parsing was incorrect", output.Field1)
	}
}

type testType7 struct {
	Field1 int `sugar:"field_1"`
}

func (tt testType7) Validate() []sugar.ValidationError {
	return nil
}

func TestParseValues7_1(t *testing.T) {
	output := testType7{}
	input := GetUrlValues(map[string]string{})
	result := sugar.ParseValues(input, &output)
	if !result.HasError() {
		t.Error("Did not have error")
		return
	}
	if result.HumanReadableError() == nil {
		t.Error("Did not have human readable error")
		return
	}
	if len(result.ExtraFieldErrors) != 0 || len(result.ValidationErrors) != 0 || len(result.ParseErrors) != 1 {
		t.Error("Error counts were not right", result.ExtraFieldErrors, result.ValidationErrors, result.ParseErrors)
		return
	}
	e := result.ParseErrors[0]
	if e.Field != "field_1" || e.Reason != sugar.FieldMissingID {
		t.Error("Error was wrong type")
	}
}

func TestParseValues7_2(t *testing.T) {
	output := testType7{}
	input := GetUrlValues(map[string]string{
		"field_1": "something",
	})
	result := sugar.ParseValues(input, &output)
	if !result.HasError() {
		t.Error("Did not have error")
		return
	}
	if result.HumanReadableError() == nil {
		t.Error("Did not have human readable error")
		return
	}
	if len(result.ExtraFieldErrors) != 0 || len(result.ValidationErrors) != 0 || len(result.ParseErrors) != 1 {
		t.Error("Error counts were not right", result.ExtraFieldErrors, result.ValidationErrors, result.ParseErrors)
		return
	}
	e := result.ParseErrors[0]
	if e.Field != "field_1" || e.Reason != sugar.ValidateFailedID {
		t.Error("Error was wrong type")
	}
}

func TestParseValues7_3(t *testing.T) {
	output := testType7{}
	input := GetUrlValues(map[string]string{
		"field_1": "4.223",
	})
	result := sugar.ParseValues(input, &output)
	if !result.HasError() {
		t.Error("Did not have error")
		return
	}
	if result.HumanReadableError() == nil {
		t.Error("Did not have human readable error")
		return
	}
	if len(result.ExtraFieldErrors) != 0 || len(result.ValidationErrors) != 0 || len(result.ParseErrors) != 1 {
		t.Error("Error counts were not right", result.ExtraFieldErrors, result.ValidationErrors, result.ParseErrors)
		return
	}
	e := result.ParseErrors[0]
	if e.Field != "field_1" || e.Reason != sugar.ValidateFailedID {
		t.Error("Error was wrong type")
	}
}

type testType8 struct {
	field1 *float64
}

func (tt testType8) Validate() []sugar.ValidationError {
	return nil
}

func TestParseValues8_1(t *testing.T) {
	output := testType8{}
	input := GetUrlValues(map[string]string{
		"field1": "something",
	})
	result := sugar.ParseValues(input, &output)
	if !result.HasError() {
		t.Error("Did not have error")
		return
	}
	if result.HumanReadableError() == nil {
		t.Error("Did not have human readable error")
		return
	}
	if len(result.ExtraFieldErrors) != 0 || len(result.ValidationErrors) != 0 || len(result.ParseErrors) != 1 {
		t.Error("Error counts were not right", result.ExtraFieldErrors, result.ValidationErrors, result.ParseErrors)
		return
	}
	e := result.ParseErrors[0]
	if e.Field != "field1" || e.Reason != sugar.ValidateFailedID {
		t.Error("Error was wrong type", e.Field, e.Reason)
	}
}

type testType9 struct {
	Field1 byte `sugar:"field_1"`
}

func (tt testType9) Validate() []sugar.ValidationError {
	return nil
}

func TestParseValues9_1(t *testing.T) {
	output := testType9{}
	input := GetUrlValues(map[string]string{
		"field_1": "something",
	})
	result := sugar.ParseValues(input, &output)
	if !result.HasError() {
		t.Error("Did not have error")
		return
	}
	if result.HumanReadableError() == nil {
		t.Error("Did not have human readable error")
		return
	}
	if len(result.ExtraFieldErrors) != 0 || len(result.ValidationErrors) != 0 || len(result.ParseErrors) != 1 {
		t.Error("Error counts were not right", result.ExtraFieldErrors, result.ValidationErrors, result.ParseErrors)
		return
	}
	e := result.ParseErrors[0]
	if e.Field != "field_1" || e.Reason != sugar.ValidateFailedID {
		t.Error("Error was wrong type", e.Field, e.Reason)
	}
}

type testType10 struct {
	Field1 string `sugar:"field_1"`
}

func (tt testType10) Validate() []sugar.ValidationError {
	return nil
}

func TestParseValues10_1(t *testing.T) {
	output := testType10{}
	input := GetUrlValues(map[string]string{
		"field_1": "something",
		"field_2": "extra!",
	})
	result := sugar.ParseValues(input, &output)
	if !result.HasError() {
		t.Error("Did not have error")
		return
	}
	if result.HumanReadableError() == nil {
		t.Error("Did not have human readable error")
		return
	}
	if len(result.ExtraFieldErrors) != 1 || len(result.ValidationErrors) != 0 || len(result.ParseErrors) != 0 {
		t.Error("Error counts were not right", result.ExtraFieldErrors, result.ValidationErrors, result.ParseErrors)
		return
	}
	e := result.ExtraFieldErrors[0]
	if e.Field != "field_2" {
		t.Error("Error was wrong field", e.Field)
	}
}

type testType11 struct {
	Field1 string `sugar:"field_1"`
}

func (tt testType11) Validate() []sugar.ValidationError {
	e := sugar.ValidationError{
		Field:  "field_1",
		Reason: "My Reason",
	}
	return []sugar.ValidationError{e}
}

func TestParseValues11_1(t *testing.T) {
	output := testType11{}
	input := GetUrlValues(map[string]string{
		"field_1": "something",
	})
	result := sugar.ParseValues(input, &output)
	if !result.HasError() {
		t.Error("Did not have error")
		return
	}
	if result.HumanReadableError() == nil {
		t.Error("Did not have human readable error")
		return
	}
	if len(result.ExtraFieldErrors) != 0 || len(result.ValidationErrors) != 1 || len(result.ParseErrors) != 0 {
		t.Error("Error counts were not right", result.ExtraFieldErrors, result.ValidationErrors, result.ParseErrors)
		return
	}
	e := result.ValidationErrors[0]
	if e.Field != "field_1" || e.Reason != "My Reason" {
		t.Error("Error was wrong field", e.Field)
	}
}

type testType12 struct {
	Field1 []string `sugar:"field_1"`
}

func (tt testType12) Validate() []sugar.ValidationError {
	return nil
}

func TestParseValues12_1(t *testing.T) {
	output := testType12{}
	input := GetUrlValues(map[string]string{
		"field_1": "red,blue,green",
	})
	result := sugar.ParseValues(input, &output)
	if result.HasError() {
		t.Error("Should not have error", result)
		return
	}
	if len(output.Field1) != 3 {
		t.Error("Wrong length")
		return
	}
	if output.Field1[0] != "red" {
		t.Error("should be red")
	}
	if output.Field1[1] != "blue" {
		t.Error("should be blue")
	}
	if output.Field1[2] != "green" {
		t.Error("should be green")
	}
}

func TestParseValues12_2(t *testing.T) {
	output := testType12{}
	input := GetUrlValues(map[string]string{})
	result := sugar.ParseValues(input, &output)
	if result.HasError() {
		t.Error("should not have error")
		return
	}
	if len(output.Field1) != 0 {
		t.Error("Should hve 0 length")
	}
	if output.Field1 != nil {
		t.Error("Should be nil")
	}
}

type testType13 struct {
	Field1 []int `sugar:"field_1"`
}

func (tt testType13) Validate() []sugar.ValidationError {
	return nil
}

func TestParseValues13_1(t *testing.T) {
	output := testType13{}
	input := GetUrlValues(map[string]string{
		"field_1": "1,2,3",
	})
	result := sugar.ParseValues(input, &output)
	if result.HasError() {
		t.Error("should not have error")
	}
	if len(output.Field1) != 3 {
		t.Error("Invalid response")
	}
	if output.Field1[0] != 1 || output.Field1[1] != 2 || output.Field1[2] != 3 {
		t.Error("The fields were parsed incorrectly")
	}
}

func TestParseValues13_2(t *testing.T) {
	output := testType13{}
	input := GetUrlValues(map[string]string{
		"field_1": "1,2.2,3",
	})
	result := sugar.ParseValues(input, &output)
	if !result.HasError() {
		t.Error("should have error")
	}
}

type testType14 struct {
	Field1 []float64 `sugar:"field_1"`
}

func (tt testType14) Validate() []sugar.ValidationError {
	return nil
}

func TestParseValues14_1(t *testing.T) {
	output := testType14{}
	input := GetUrlValues(map[string]string{
		"field_1": "1.2,2.4,3",
	})
	result := sugar.ParseValues(input, &output)
	if result.HasError() {
		t.Error("should not have error")
	}
	if len(output.Field1) != 3 {
		t.Error("Invalid response")
	}
	if output.Field1[0] != 1.2 || output.Field1[1] != 2.4 || output.Field1[2] != 3 {
		t.Error("The fields were parsed incorrectly")
	}
}
