package references

import "errors"

type builder struct {
	path     []string
	name     string
	pVersion *uint
}

func createBuilder() Builder {
	out := builder{
		path:     nil,
		name:     "",
		pVersion: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithPath adds a path to the builder
func (app *builder) WithPath(path []string) Builder {
	app.path = path
	return app
}

// WithName adds a name to the builder
func (app *builder) WithName(name string) Builder {
	app.name = name
	return app
}

// WithVersion adds a version to the builder
func (app *builder) WithVersion(version uint) Builder {
	app.pVersion = &version
	return app
}

// Now builds a new Reference instance
func (app *builder) Now() (Reference, error) {
	if app.path == nil && len(app.path) <= 0 {
		return nil, errors.New("the path is mandatory in order to build a Reference instance")
	}

	if app.pVersion == nil {
		return nil, errors.New("the version is mandatory in order to build a Reference instance")
	}

	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Reference instance")
	}

	return createReference(
		app.path,
		app.name,
		*app.pVersion,
	), nil
}
