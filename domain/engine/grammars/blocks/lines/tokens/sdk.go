package tokens

import (
	"github.com/steve-care-software/grammars/domain/engine/grammars/blocks/lines/tokens/cardinalities"
	"github.com/steve-care-software/grammars/domain/engine/grammars/blocks/lines/tokens/elements"
	"github.com/steve-care-software/grammars/domain/engine/grammars/blocks/lines/tokens/reverses"
	"github.com/steve-care-software/grammars/domain/engine/grammars/blocks/lines/tokens/uniques"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewTokenBuilder creates a token builder
func NewTokenBuilder() TokenBuilder {
	return createTokenBuilder()
}

// Builder represents a tokens list
type Builder interface {
	Create() Builder
	WithList(list []Token) Builder
	Now() (Tokens, error)
}

// Tokens represents tokens
type Tokens interface {
	List() []Token
	Fetch(name string, idx uint) (Token, error)
}

// TokenBuilder represents the token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithElement(element elements.Element) TokenBuilder
	WithCardinality(cardinality cardinalities.Cardinality) TokenBuilder
	WithReverse(reverse reverses.Reverse) TokenBuilder
	WithUnique(unique uniques.Unique) TokenBuilder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Name() string
	Element() elements.Element
	Cardinality() cardinalities.Cardinality
	HasReverse() bool
	Reverse() reverses.Reverse
	HasUnique() bool
	Unique() uniques.Unique
}
