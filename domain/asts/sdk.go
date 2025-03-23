package asts

import (
	"github.com/steve-care-software/grammars/domain/grammars"
	"github.com/steve-care-software/grammars/domain/grammars/blocks/lines/balances"
	"github.com/steve-care-software/grammars/domain/grammars/blocks/lines/balances/selectors"
	"github.com/steve-care-software/grammars/domain/grammars/blocks/lines/balances/selectors/chains"
	"github.com/steve-care-software/grammars/domain/grammars/blocks/lines/tokens/uniques"
)

// NewAdapter creates a new adapter
func NewAdapter(
	grammarRepository grammars.Repository,
) Adapter {
	grammarAdapter := grammars.NewAdapter()
	builder := NewBuilder()
	instructionBuilder := NewInstructionBuilder()
	tokensBuilder := NewTokensBuilder()
	tokenBuilder := NewTokenBuilder()
	elementsBuilder := NewElementsBuilder()
	elementBuilder := NewElementBuilder()
	constantBuilder := NewConstantBuilder()
	return createAdapter(
		grammarRepository,
		grammarAdapter,
		builder,
		instructionBuilder,
		tokensBuilder,
		tokenBuilder,
		elementsBuilder,
		elementBuilder,
		constantBuilder,
	)
}

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewInstructionBuilder creates a new instruction builder
func NewInstructionBuilder() InstructionBuilder {
	return createInstructionBuilder()
}

// NewTokensBuilder creates a new tokens builder
func NewTokensBuilder() TokensBuilder {
	return createTokensBuilder()
}

// NewTokenBuilder creates a new token builder
func NewTokenBuilder() TokenBuilder {
	return createTokenBuilder()
}

// NewElementsAdapter creates a new elements adapter
func NewElementsAdapter() ElementsAdapter {
	return createElementsAdapter()
}

// NewElementsBuilder creates a new elements builder
func NewElementsBuilder() ElementsBuilder {
	return createElementsBuilder()
}

// NewElementBuilder creates a new element builder
func NewElementBuilder() ElementBuilder {
	return createElementBuilder()
}

// NewConstantBuilder creates a new constant builder
func NewConstantBuilder() ConstantBuilder {
	return createConstantBuilder()
}

// Adapter represents the adapter
type Adapter interface {
	// ToAST takes the grammar and input and converts them to a ast instance and the remaining data
	ToAST(grammar grammars.Grammar, input []byte) (AST, []byte, error)

	// ToASTWithRoot creates a ast but changes the root block of the grammar
	ToASTWithRoot(grammar grammars.Grammar, rootBlockName string, input []byte) (AST, []byte, error)
}

// Builder represents the ast builder
type Builder interface {
	Create() Builder
	WithRoot(root Element) Builder
	Now() (AST, error)
}

// AST represents a ast
type AST interface {
	Root() Element
}

// InstructionBuilder represents the instruction builder
type InstructionBuilder interface {
	Create() InstructionBuilder
	WithBlock(block string) InstructionBuilder
	WithLine(line uint) InstructionBuilder
	WithTokens(tokens Tokens) InstructionBuilder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	Validate(elementNameIndex map[string]BlockCount) (map[string]BlockCount, error)
	Block() string
	Line() uint
	Tokens() Tokens
}

// TokensBuilder represents the tokens builder
type TokensBuilder interface {
	Create() TokensBuilder
	WithList(list []Token) TokensBuilder
	Now() (Tokens, error)
}

// Tokens represents tokens
type Tokens interface {
	Validate(elementNameIndex map[string]BlockCount) (map[string]BlockCount, error)
	List() []Token
	Value() []byte
	FetchAll(name string) ([]Token, error)
	Fetch(name string, index uint) (Token, error)
	Select(chain chains.Chain) ([]Token, Token, Element, error)
	IsBalanceValid(balance balances.Balance) bool
	IsSelectorValid(selector selectors.Selector) bool
	IsChainValid(chain chains.Chain) bool
}

// TokenBuilder represents the token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithName(name string) TokenBuilder
	WithElements(elements Elements) TokenBuilder
	WithUnique(unique uniques.Unique) TokenBuilder
	Now() (Token, error)
}

// Token represents a token
type Token interface {
	Validate(elementNameIndex map[string]BlockCount) (map[string]BlockCount, error)
	Name() string
	Elements() Elements
	Value() []byte
	HasUnique() bool
	Unique() uniques.Unique
}

// ElementsAdapter represents the elements adapter
type ElementsAdapter interface {
	// ToBytes takes an elements and returns its bytes
	ToBytes(elements Elements) ([]byte, error)
}

// ElementsBuilder represents the elements builder
type ElementsBuilder interface {
	Create() ElementsBuilder
	WithList(list []Element) ElementsBuilder
	Now() (Elements, error)
}

// Elements represents elements
type Elements interface {
	Validate(elementNameIndex map[string]BlockCount) (map[string]BlockCount, error)
	List() []Element
	Fetch(idx uint) (Element, error)
	Value() []byte
}

// ElementBuilder represents the element builder
type ElementBuilder interface {
	Create() ElementBuilder
	WithConstant(constant Constant) ElementBuilder
	WithInstruction(instruction Instruction) ElementBuilder
	WithAST(ast AST) ElementBuilder
	Now() (Element, error)
}

// Element represents an element
type Element interface {
	Search(name string, idx uint) (Token, error)
	Validate(elementNameIndex map[string]BlockCount) (map[string]BlockCount, error)
	Name() string
	Value() []byte
	IsChainValid(chain chains.Chain) bool
	IsConstant() bool
	Constant() Constant
	IsInstruction() bool
	Instruction() Instruction
	IsAST() bool
	AST() AST
}

// ConstantBuilder represents the constant builder
type ConstantBuilder interface {
	Create() ConstantBuilder
	WithName(name string) ConstantBuilder
	WithValue(value []byte) ConstantBuilder
	Now() (Constant, error)
}

// Constant represents a constant
type Constant interface {
	Name() string
	Value() []byte
	IsChainValid(chain chains.Chain) bool
}

// BlockCount represents a block count
type BlockCount struct {
	index  uint
	tokens map[string][]byte
}
