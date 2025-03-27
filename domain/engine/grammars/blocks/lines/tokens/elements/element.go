package elements

import (
	"github.com/steve-care-software/grammars/domain/engine/grammars/blocks/lines/tokens/elements/references"
)

type element struct {
	rule      string
	block     string
	constant  string
	reference references.Reference
}

func createElementWithRule(rule string) Element {
	return createElementInternally(rule, "", "", nil)
}

func createElementWithBlock(block string) Element {
	return createElementInternally("", block, "", nil)
}

func createElementWithConstant(constant string) Element {
	return createElementInternally("", "", constant, nil)
}

func createElementWithReference(reference references.Reference) Element {
	return createElementInternally("", "", "", reference)
}

func createElementInternally(
	rule string,
	block string,
	constant string,
	reference references.Reference,
) Element {
	out := element{
		rule:      rule,
		block:     block,
		constant:  constant,
		reference: reference,
	}

	return &out
}

// Name returns the rule or block name
func (obj *element) Name() string {
	if obj.IsBlock() {
		return obj.block
	}

	if obj.IsConstant() {
		return obj.constant
	}

	if obj.IsReference() {
		return obj.reference.Name()
	}

	return obj.rule
}

// IsRule returns true if there is a rule, false otherwise
func (obj *element) IsRule() bool {
	return obj.rule != ""
}

// Rule returns the rule, if any
func (obj *element) Rule() string {
	return obj.rule
}

// IsBlock returns true if there is a block, false otherwise
func (obj *element) IsBlock() bool {
	return obj.block != ""
}

// Block returns the block, if any
func (obj *element) Block() string {
	return obj.block
}

// IsConstant returns true if there is a constant, false otherwise
func (obj *element) IsConstant() bool {
	return obj.constant != ""
}

// Constant returns the constant, if any
func (obj *element) Constant() string {
	return obj.constant
}

// IsReference returns true if there is a reference, false otherwise
func (obj *element) IsReference() bool {
	return obj.reference != nil
}

// Reference returns the reference, if any
func (obj *element) Reference() references.Reference {
	return obj.reference
}
