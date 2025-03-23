package elements

import (
	"errors"

	"github.com/steve-care-software/grammars/domain/grammars/blocks/lines/tokens/elements/references"
)

type elementBuilder struct {
	rule      string
	block     string
	constant  string
	reference references.Reference
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		rule:      "",
		block:     "",
		constant:  "",
		reference: nil,
	}

	return &out
}

// Create initializes the builder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithRule adds a rule to the builder
func (app *elementBuilder) WithRule(rule string) ElementBuilder {
	app.rule = rule
	return app
}

// WithBlock adds a block to the builder
func (app *elementBuilder) WithBlock(block string) ElementBuilder {
	app.block = block
	return app
}

// WithConstant adds a constant to the builder
func (app *elementBuilder) WithConstant(constant string) ElementBuilder {
	app.constant = constant
	return app
}

// WithReference adds a reference to the builder
func (app *elementBuilder) WithReference(reference references.Reference) ElementBuilder {
	app.reference = reference
	return app
}

// Now builds a new Element
func (app *elementBuilder) Now() (Element, error) {
	if app.rule != "" {
		return createElementWithRule(app.rule), nil
	}

	if app.block != "" {
		return createElementWithBlock(app.block), nil
	}

	if app.constant != "" {
		return createElementWithConstant(app.constant), nil
	}

	if app.reference != nil {
		return createElementWithReference(app.reference), nil
	}

	return nil, errors.New("the Element is invalid")
}
