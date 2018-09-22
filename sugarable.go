package sugar

// Sugarable The struct containing the values to be parsed from the request
type Sugarable interface {

	// Validate Custom validation on the parsed values of your request struct. You can just return nil if you don't want
	// to perform any custom validation.
	Validate() []ValidationError
}
