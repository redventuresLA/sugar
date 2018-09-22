package sugar

// Response The response from a validation of url values from sugar.
type Response struct {
	ParseErrors      []ParseError
	ExtraFieldErrors []ExtraFieldError
	ValidationErrors []ValidationError
}

// HumanReadableError Returns the error formatted as a map[string]interface{} so that it can be returned to the user in
// json form.
func (r Response) HumanReadableError() map[string]interface{} {
	if !r.HasError() {
		return nil
	}
	output := make(map[string]interface{})
	output["description"] = "Errors validating input parameters"
	if len(r.ParseErrors) > 0 {
		parseErrorsOutput := make([]map[string]interface{}, 0)
		for _, parseError := range r.ParseErrors {
			parseErrorsOutput = append(parseErrorsOutput, parseError.humanReadable())
		}
		output["parse_errors"] = map[string]interface{}{
			"errors": parseErrorsOutput,
			"action": "make sure that your input has the correct data type and that it is included.",
		}
	}
	if len(r.ValidationErrors) > 0 {
		validationErrorsOutput := make([]map[string]interface{}, 0)
		for _, validationError := range r.ValidationErrors {
			validationErrorsOutput = append(validationErrorsOutput, validationError.humanReadable())
		}
		output["validation_errors"] = map[string]interface{}{
			"errors": validationErrorsOutput,
			"action": "check your request to make sure you don't send invalid parameters",
		}
	}
	if len(r.ExtraFieldErrors) > 0 {
		extraFieldErrorsOutput := make([]map[string]interface{}, 0)
		for _, extraFieldError := range r.ExtraFieldErrors {
			extraFieldErrorsOutput = append(extraFieldErrorsOutput, extraFieldError.humanReadable())
		}
		output["extra_field_errors"] = map[string]interface{}{
			"errors": extraFieldErrorsOutput,
			"action": "don't send this field in future requests. It is unexpected.",
		}
	}
	return output
}

// HasError Specifies whether the validation failed. In this case, you most likely want to return a 400 level status
// code.
func (r Response) HasError() bool {
	return len(r.ParseErrors) > 0 || len(r.ExtraFieldErrors) > 0 || len(r.ValidationErrors) > 0
}

// ValidationError An error arising from validation using the Sugarable.Validate method.
type ValidationError struct {
	Field  string
	Reason string
}

func (ve ValidationError) humanReadable() map[string]interface{} {
	return map[string]interface{}{
		"field":  ve.Field,
		"reason": ve.Reason,
	}
}

// ParseError An error arising from invalid parsing of the input values
type ParseError struct {
	Field  string
	Reason string
}

func (pe ParseError) humanReadable() map[string]interface{} {
	return map[string]interface{}{
		"field":  pe.Field,
		"reason": pe.Reason,
	}
}

// ExtraFieldError An error arising from the user including an extra field in the request that was not expected.
type ExtraFieldError struct {
	Field string
}

func (efe ExtraFieldError) humanReadable() map[string]interface{} {
	return map[string]interface{}{
		"field": efe.Field,
	}
}

func (r *Response) addFieldError(field string, reason string) {
	r.ParseErrors = append(r.ParseErrors, ParseError{Field: field, Reason: reason})
}

func (r *Response) addExtraFieldError(field string) {
	r.ExtraFieldErrors = append(r.ExtraFieldErrors, ExtraFieldError{Field: field})
}
