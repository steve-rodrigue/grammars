package elements

import "github.com/steve-care-software/grammars/domain/engine/grammars/blocks/lines/tokens/elements/references"

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
}

// Builder represents an elments list
type Builder interface {
	Create() Builder
	WithList(list []Element) Builder
	Now() (Elements, error)
}

// Elements represents elements
type Elements interface {
	List() []Element
	Fetch(name string) (Element, error)
}

// ElementBuilder represents the element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithRule(rule string) ElementBuilder
	WithBlock(block string) ElementBuilder
	WithConstant(constant string) ElementBuilder
	WithReference(reference references.Reference) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Name() string
	IsRule() bool
	Rule() string
	IsBlock() bool
	Block() string
	IsConstant() bool
	Constant() string
	IsReference() bool
	Reference() references.Reference
}
