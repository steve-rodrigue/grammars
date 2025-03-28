package resources

import "github.com/steve-care-software/grammars/domain/engine/grammars/blocks/lines/balances/selectors/chains"

// Resources represents resources
type Resources interface {
	List() []Resource
}

// Resource represents a resource
type Resource interface {
	Name() string
	Path() []string
	HasChain() bool
	Chain() chains.Chain
}
