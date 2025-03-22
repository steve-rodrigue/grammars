package suites

type suite struct {
	name   string
	input  []byte
	isFail bool
}

func createSuite(
	name string,
	input []byte,
	isFail bool,
) Suite {
	out := suite{
		name:   name,
		input:  input,
		isFail: isFail,
	}

	return &out
}

// Name returns the name
func (obj *suite) Name() string {
	return obj.name
}

// Input returns the input
func (obj *suite) Input() []byte {
	return obj.input
}

// IsFail returns true if expected to fail, false otherwise
func (obj *suite) IsFail() bool {
	return obj.isFail
}
