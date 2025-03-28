package suites

// Suites represents suites
type Suites interface {
	List() []Suite
}

// Suite represents a user suite
type Suite interface {
	Name() string
	Roles() []string
	IsNegative() bool
	HasComment() bool
	Comment() string
}
