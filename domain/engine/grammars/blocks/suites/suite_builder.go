package suites

import (
	"errors"
)

type suiteBuilder struct {
	name   string
	input  []byte
	isFail bool
}

func createSuiteBuilder() SuiteBuilder {
	out := suiteBuilder{
		name:   "",
		input:  nil,
		isFail: false,
	}

	return &out
}

// Create initializes the builder
func (app *suiteBuilder) Create() SuiteBuilder {
	return createSuiteBuilder()
}

// WithName adds a name to the builder
func (app *suiteBuilder) WithName(name string) SuiteBuilder {
	app.name = name
	return app
}

// WithInput adds a input to the builder
func (app *suiteBuilder) WithInput(input []byte) SuiteBuilder {
	app.input = input
	return app
}

// IsFail flags the suite as fail
func (app *suiteBuilder) IsFail() SuiteBuilder {
	app.isFail = true
	return app
}

// Now builds a new Suite instance
func (app *suiteBuilder) Now() (Suite, error) {
	if app.input != nil && len(app.input) <= 0 {
		app.input = nil
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Suite instance")
	}

	return createSuite(app.name, app.input, app.isFail), nil
}
