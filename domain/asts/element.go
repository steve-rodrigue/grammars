package asts

import (
	"github.com/steve-care-software/grammars/domain/grammars/blocks/lines/balances/selectors/chains"
)

type element struct {
	constant    Constant
	instruction Instruction
	ast         AST
}

func createElementWithConstant(constant Constant) Element {
	return createElementInternally(constant, nil, nil)
}

func createElementWithInstruction(instruction Instruction) Element {
	return createElementInternally(nil, instruction, nil)
}

func createElementWithAST(ast AST) Element {
	return createElementInternally(nil, nil, ast)
}

func createElementInternally(
	constant Constant,
	instruction Instruction,
	ast AST,
) Element {
	out := element{
		constant:    constant,
		instruction: instruction,
		ast:         ast,
	}

	return &out
}

// Validate validates an element
func (obj *element) Validate(elementNameIndex map[string]BlockCount) (map[string]BlockCount, error) {
	if obj.IsConstant() {
		return elementNameIndex, nil
	}

	return obj.instruction.Validate(elementNameIndex)
}

// Name returns the name
func (obj *element) Name() string {
	if obj.IsConstant() {
		return obj.constant.Name()
	}

	if obj.IsAST() {
		return obj.ast.Root().Name()
	}

	return obj.instruction.Block()
}

// IsConstant returns true if there is a constant, false otherwise
func (obj *element) IsConstant() bool {
	return obj.constant != nil
}

// Constant returns the constant, if any
func (obj *element) Constant() Constant {
	return obj.constant
}

// IsInstruction returns true if there is an instruction, false otherwise
func (obj *element) IsInstruction() bool {
	return obj.instruction != nil
}

// Instruction returns the instruction, if any
func (obj *element) Instruction() Instruction {
	return obj.instruction
}

// IsAST returns true if there is an ast, false otherwise
func (obj *element) IsAST() bool {
	return obj.ast != nil
}

// AST returns the ast, if any
func (obj *element) AST() AST {
	return obj.ast
}

// Value returns the value of the elements
func (obj *element) Value() []byte {
	if obj.IsConstant() {
		return obj.constant.Value()
	}

	return obj.instruction.Tokens().Value()
}

// Search searches inside the element
func (obj *element) Search(name string, idx uint) (Token, error) {
	if obj.IsConstant() {
		return nil, nil
	}

	if obj.IsAST() {
		retToken, err := obj.AST().Root().Search(name, idx)
		if err != nil {
			return nil, nil
		}

		return retToken, nil
	}

	retToken, err := obj.Instruction().Tokens().Fetch(name, idx)
	if err != nil {
		return nil, nil
	}

	return retToken, nil
}

// IsChainValid validates the element against the chain
func (obj *element) IsChainValid(chain chains.Chain) bool {
	if obj.IsInstruction() {
		return obj.instruction.Tokens().IsChainValid(chain)
	}

	if obj.IsAST() {
		return obj.ast.Root().IsChainValid(chain)
	}

	return obj.constant.IsChainValid(chain)
}
