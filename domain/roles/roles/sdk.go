package roles

import (
	"github.com/steve-care-software/grammars/domain/roles/roles/lines"
	"github.com/steve-care-software/grammars/domain/roles/roles/suites"
)

// Roles represents roles
type Roles interface {
	List() []Role
}

// Role represents a role
type Role interface {
	Name() string
	HasLines() bool
	Lines() lines.Lines
	HasSuites() bool
	Suites() suites.Suites
	HasComment() bool
	Comment() string
}
