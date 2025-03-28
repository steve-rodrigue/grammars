package tokens

import "github.com/steve-care-software/grammars/domain/roles/roles/lines/tokens/elements"

// Tokens represents a tokens list
type Tokens interface {
	List() []Token
}

// Token represents a role token
type Token interface {
	Element() elements.Element
	IsNot() bool
	MustNot() bool
	MustBe() bool
}
