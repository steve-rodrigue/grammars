package lines

import "github.com/steve-care-software/grammars/domain/roles/permissions/lines/tokens"

// Lines represents a lines of tokens
type Lines interface {
	List() []Line
}

// Line represents a line
type Line interface {
	Tokens() []tokens.Tokens
	HasComment() bool
	Comment() string
}
