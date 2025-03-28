package permissions

import "github.com/steve-care-software/grammars/domain/roles/permissions/lines"

// Permissions represents permissions
type Permissions interface {
	List() []Permission
}

// Permission represents a permission
type Permission interface {
	Name() string
	Lines() lines.Lines
	HasComment() bool
	Comment() string
}
