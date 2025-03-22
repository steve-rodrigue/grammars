package uniques

import (
	"errors"

	"github.com/steve-care-software/grammars/domain/grammars/blocks/lines/tokens/elements"
)

type builder struct {
	element elements.Element
	mustBe  bool
	mustNot bool
	pIndex  *uint
}

func createBuilder() Builder {
	out := builder{
		element: nil,
		mustBe:  false,
		mustNot: false,
		pIndex:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithElement adds an element to the builder
func (app *builder) WithElement(element elements.Element) Builder {
	app.element = element
	return app
}

// MustBe flags the builder as must be unique
func (app *builder) MustBe() Builder {
	app.mustBe = true
	return app
}

// MustNot flags the builder as must NOT be unique
func (app *builder) MustNot() Builder {
	app.mustNot = true
	return app
}

// WithIndex adds an index to the builder
func (app *builder) WithIndex(index uint) Builder {
	app.pIndex = &index
	return app
}

// Now builds a new Unique instance
func (app *builder) Now() (Unique, error) {
	if app.element == nil {
		return nil, errors.New("the element is mandatory in order to build a Unique instance")
	}

	if app.pIndex == nil {
		return nil, errors.New("the index is mandatory in order to build a Unique instance")
	}

	if app.mustBe && app.mustNot {
		return nil, errors.New("the mustBe and mustNot cannot be both true while building a Unique instance")
	}

	if app.mustBe {
		return createUniqueWithMustBe(app.element, *app.pIndex), nil
	}

	if app.mustNot {
		return createUniqueWithMustNot(app.element, *app.pIndex), nil
	}

	return nil, errors.New("the Unique is invalid")
}
