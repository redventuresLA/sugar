package sugar

// Response The response from a validation of url values from sugar.
type Response struct {
	FieldErrors      []FieldError
	ExtraFieldErrors []ExtraFieldError
	ValidationErrors []ValidationError
}

// HasError Specifies whether the validation failed. In this case, you most likely want to return a 400 level status
// code.
func (r Response) HasError() bool {
	return len(r.FieldErrors) > 0 || len(r.ExtraFieldErrors) > 0
}

// ValidationError An error arising from validation using the Sugarable.Validate method.
type ValidationError struct {
	Field  string
	Reason string
}

// FieldError An error arising from
type FieldError struct {
	Field  string
	Reason string
}

// ExtraFieldError An error arising from the user including an extra field in the request that was not expected.
type ExtraFieldError struct {
	Field string
}

func (r Response) addFieldError(field string, reason string) {
	r.FieldErrors = append(r.FieldErrors, FieldError{Field: field, Reason: reason})
}

func (r Response) addExtraFieldError(field string) {
	r.ExtraFieldErrors = append(r.ExtraFieldErrors, ExtraFieldError{Field: field})
}
