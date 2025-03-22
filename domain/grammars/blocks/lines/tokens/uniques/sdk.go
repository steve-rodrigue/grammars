package uniques

import "github.com/steve-care-software/grammars/domain/grammars/blocks/lines/tokens/elements"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a unique builder
type Builder interface {
	Create() Builder
	WithElement(element elements.Element) Builder
	MustBe() Builder
	MustNot() Builder
	WithIndex(index uint) Builder
	Now() (Unique, error)
}

// Unique represents if a token must be unique or not
type Unique interface {
	Element() elements.Element
	MustBe() bool
	MustNot() bool
	Index() uint
}
