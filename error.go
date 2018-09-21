package sugar

type Response struct {
	FieldErrors []FieldError
	ExtraFieldErrors []ExtraFieldError
	ValidationErrors []ValidationError
}

func (r Response) HasError() bool {
	return len(r.FieldErrors) > 0 || len(r.ExtraFieldErrors) > 0
}

type ValidationError struct {
	Field string
	Reason string
}

type FieldError struct {
	Field string
	Reason string
}

type ExtraFieldError struct {
	Field string
}

func (r Response) addFieldError(field string, reason string) {
	r.FieldErrors = append(r.FieldErrors, FieldError{Field: field, Reason: reason})
}

func (r Response) addExtraFieldError(field string) {
	r.ExtraFieldErrors = append(r.ExtraFieldErrors, ExtraFieldError{Field: field})
}