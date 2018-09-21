package sugar

type Sugarable interface {
	Validate() []ValidationError
}
