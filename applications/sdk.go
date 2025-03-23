package applications

import (
	"github.com/steve-care-software/grammars/domain/asts"
	"github.com/steve-care-software/grammars/domain/grammars"
	"github.com/steve-care-software/grammars/domain/walkers/elements"
)

// NewBuilder creates a new application builder
func NewBuilder(
	grammarRepository grammars.Repository,
) Builder {
	elementsAdapter := asts.NewElementsAdapter()
	astAdapter := asts.NewAdapter(
		grammarRepository,
	)

	elementAdapter := elements.NewAdapter(
		grammarRepository,
	)

	tokensBuilder := asts.NewTokensBuilder()
	return createBuilder(
		elementsAdapter,
		astAdapter,
		elementAdapter,
		tokensBuilder,
	)
}

// Builder represents an application builder
type Builder interface {
	Create() Builder
	WithElement(ins elements.Element) Builder
	Now() (Application, error)
}

// Application represents the interpreter application
type Application interface {
	// Execute executes the parser
	Execute(input []byte, grammar grammars.Grammar) (any, []byte, error)

	// Suites executes all the test suites of the grammar
	Suites(grammar grammars.Grammar) error
}
