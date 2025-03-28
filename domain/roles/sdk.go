package roles

import (
	"github.com/steve-care-software/grammars/domain/roles/headers"
	"github.com/steve-care-software/grammars/domain/roles/permissions"
	"github.com/steve-care-software/grammars/domain/roles/resources"
	"github.com/steve-care-software/grammars/domain/roles/roles"
)

// Role represents a role
type Role interface {
	Header() headers.Header
	Roles() roles.Roles
	Permissions() permissions.Permissions
	Resources() resources.Resources
	HasComment() bool
	Comment() string
}
