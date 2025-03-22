package tokens

import (
	"github.com/steve-care-software/grammars/domain/grammars/blocks/lines/tokens/cardinalities"
	"github.com/steve-care-software/grammars/domain/grammars/blocks/lines/tokens/elements"
	"github.com/steve-care-software/grammars/domain/grammars/blocks/lines/tokens/reverses"
	"github.com/steve-care-software/grammars/domain/grammars/blocks/lines/tokens/uniques"
)

type token struct {
	element     elements.Element
	cardinality cardinalities.Cardinality
	reverse     reverses.Reverse
	unique      uniques.Unique
}

func createToken(
	element elements.Element,
	cardinality cardinalities.Cardinality,
) Token {
	return createTokenInternally(element, cardinality, nil, nil)
}

func createTokenWithReverse(
	element elements.Element,
	cardinality cardinalities.Cardinality,
	reverse reverses.Reverse,
) Token {
	return createTokenInternally(element, cardinality, reverse, nil)
}

func createTokenWithReverseWithUnique(
	element elements.Element,
	cardinality cardinalities.Cardinality,
	unique uniques.Unique,
) Token {
	return createTokenInternally(element, cardinality, nil, unique)
}

func createTokenWithReverseAndUnique(
	element elements.Element,
	cardinality cardinalities.Cardinality,
	reverse reverses.Reverse,
	unique uniques.Unique,
) Token {
	return createTokenInternally(element, cardinality, reverse, unique)
}

func createTokenInternally(
	element elements.Element,
	cardinality cardinalities.Cardinality,
	reverse reverses.Reverse,
	unique uniques.Unique,
) Token {
	out := token{
		element:     element,
		cardinality: cardinality,
		reverse:     reverse,
		unique:      unique,
	}

	return &out
}

// Name returns the name
func (obj *token) Name() string {
	return obj.element.Name()
}

// Element returns the element
func (obj *token) Element() elements.Element {
	return obj.element
}

// Cardinality returns the cardinality
func (obj *token) Cardinality() cardinalities.Cardinality {
	return obj.cardinality
}

// HasReverse returns true if there is a reverse, false otherwise
func (obj *token) HasReverse() bool {
	return obj.reverse != nil
}

// Reverse returns the reverse, if any
func (obj *token) Reverse() reverses.Reverse {
	return obj.reverse
}

// HasUnique returns true if there is a unique, false otherwise
func (obj *token) HasUnique() bool {
	return obj.unique != nil
}

// Unique returns the unique, if any
func (obj *token) Unique() uniques.Unique {
	return obj.unique
}
