package asts

import (
	"errors"
	"fmt"
)

type builder struct {
	root Element
}

func createBuilder() Builder {
	out := builder{
		root: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithRoot adds a root to the builder
func (app *builder) WithRoot(root Element) Builder {
	app.root = root
	return app
}

// Now builds a new AST
func (app *builder) Now() (AST, error) {
	if app.root == nil {
		return nil, errors.New("the root is mandatory in order to build a AST instance")
	}

	_, err := app.root.Validate(map[string]BlockCount{})
	if err != nil {
		str := fmt.Sprintf("there was an error while validating the AST: %s", err.Error())
		return nil, errors.New(str)
	}

	return createAST(
		app.root,
	), nil
}
