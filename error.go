package sugar

type Response struct {
	FieldErrors []FieldError
	ExtraFieldErrors []ExtraFieldError
}

func (r Response) HasError() bool {
	return len(r.FieldErrors) > 0 || len(r.ExtraFieldErrors) > 0
}

type FieldError struct {
	Field string
	Reason string
}

type ExtraFieldError struct {
	Field string
}
