package grammars

import (
	"github.com/steve-care-software/grammars/domain/engine/grammars/blocks"
	"github.com/steve-care-software/grammars/domain/engine/grammars/blocks/lines"
	"github.com/steve-care-software/grammars/domain/engine/grammars/blocks/lines/balances"
	"github.com/steve-care-software/grammars/domain/engine/grammars/blocks/lines/balances/selectors"
	"github.com/steve-care-software/grammars/domain/engine/grammars/blocks/lines/balances/selectors/chains"
	"github.com/steve-care-software/grammars/domain/engine/grammars/blocks/lines/tokens"
	"github.com/steve-care-software/grammars/domain/engine/grammars/blocks/lines/tokens/cardinalities"
	"github.com/steve-care-software/grammars/domain/engine/grammars/blocks/lines/tokens/elements"
	"github.com/steve-care-software/grammars/domain/engine/grammars/blocks/lines/tokens/elements/references"
	"github.com/steve-care-software/grammars/domain/engine/grammars/blocks/lines/tokens/reverses"
	"github.com/steve-care-software/grammars/domain/engine/grammars/blocks/lines/tokens/uniques"
	"github.com/steve-care-software/grammars/domain/engine/grammars/blocks/suites"
	"github.com/steve-care-software/grammars/domain/engine/grammars/constants"
	constant_tokens "github.com/steve-care-software/grammars/domain/engine/grammars/constants/tokens"
	constant_elements "github.com/steve-care-software/grammars/domain/engine/grammars/constants/tokens/elements"
	"github.com/steve-care-software/grammars/domain/engine/grammars/rules"
)

// CoreFn represents a core fn
type CoreFn func(input map[string][]byte) ([]byte, error)

const llA = "a"
const llB = "b"
const llC = "c"
const llD = "d"
const llE = "e"
const llF = "f"
const llG = "g"
const llH = "h"
const llI = "i"
const llJ = "j"
const llK = "k"
const llL = "l"
const llM = "m"
const llN = "n"
const llO = "o"
const llP = "p"
const llQ = "q"
const llR = "r"
const llS = "s"
const llT = "t"
const llU = "u"
const llV = "v"
const llW = "w"
const llX = "x"
const llY = "y"
const llZ = "z"

const ulA = "A"
const ulB = "B"
const ulC = "C"
const ulD = "D"
const ulE = "E"
const ulF = "F"
const ulG = "G"
const ulH = "H"
const ulI = "I"
const ulJ = "J"
const ulK = "K"
const ulL = "L"
const ulM = "M"
const ulN = "N"
const ulO = "O"
const ulP = "P"
const ulQ = "Q"
const ulR = "R"
const ulS = "S"
const ulT = "T"
const ulU = "U"
const ulV = "V"
const ulW = "W"
const ulX = "X"
const ulY = "Y"
const ulZ = "Z"

const nZero = "0"
const nOne = "1"
const nTwo = "2"
const nTree = "3"
const nFour = "4"
const nFive = "5"
const nSix = "6"
const nSeven = "7"
const nHeight = "8"
const nNine = "9"

const ruleValueEscape = "\\"
const ruleValuePrefix = "\""
const ruleValueSuffix = "\""
const ruleNameSeparator = "_"
const ruleNameValueSeparator = ":"
const cardinalityOpen = "["
const cardinalityClose = "]"
const cardinalitySeparator = ","
const cardinalityZeroPlus = "*"
const cardinalityOnePlus = "+"
const cardinalityOptional = "?"
const tokenReversePrefix = "!"
const tokenReverseEscapePrefix = "["
const tokenReverseEscapeSuffix = "]"
const tokenMustBeUnique = "#"
const tokenMustNotBeUnique = "$"
const tokenReference = "."
const linesSeparator = "|"
const lineSeparator = "-"
const funcNameSeparator = "_"
const blockDefinitionSeparator = ":"
const failSeparator = "!"
const suiteLineSuffix = ";"
const blockSuffix = ";"
const suiteSeparatorPrefix = "---"
const versionPrefix = "v"
const versionSuffix = ";"
const rootPrefix = ">"
const rootSuffix = ";"
const omissionPrefix = "#"
const omissionSuffix = ";"
const filterBytes = ` 	
` // space, tab and eol

const constantNamePrefix = "_"
const selectorChainElementPrefix = "->"
const selectorOperatorAnd = "&&"
const selectorOperatorOr = "||"
const selectorOperatorXor = "<>"
const openParenthesis = "("
const closeParenthesis = ")"

const referenceBegin = "["
const referenceEnd = "]"
const referencePathSeparator = "/"
const referenceElementSeparator = ","

// NewAdapter creates a new adapter
func NewAdapter() Adapter {
	grammarBuilder := NewBuilder()
	constantsBuilder := constants.NewBuilder()
	constantBuilder := constants.NewConstantBuilder()
	constantTokensBuilder := constant_tokens.NewBuilder()
	constantTokenBuilder := constant_tokens.NewTokenBuilder()
	constantElementBuilder := constant_elements.NewBuilder()
	blocksBuilder := blocks.NewBuilder()
	blockBuilder := blocks.NewBlockBuilder()
	suitesBuilder := suites.NewBuilder()
	suiteBuilder := suites.NewSuiteBuilder()
	linesBuilder := lines.NewBuilder()
	lineBuilder := lines.NewLineBuilder()
	balanceBuilder := balances.NewBuilder()
	selectorsBuilder := selectors.NewBuilder()
	selectorBuilder := selectors.NewSelectorBuilder()
	selectorChainBuilder := chains.NewBuilder()
	selectorChainTokenBuilder := chains.NewTokenBuilder()
	selectorChainElementBuilder := chains.NewElementBuilder()
	tokensBuilder := tokens.NewBuilder()
	tokenBuilder := tokens.NewTokenBuilder()
	reverseBuilder := reverses.NewBuilder()
	uniqueBuilder := uniques.NewBuilder()
	elementsBuilder := elements.NewBuilder()
	elementBuilder := elements.NewElementBuilder()
	rulesBuilder := rules.NewBuilder()
	ruleBuilder := rules.NewRuleBuilder()
	cardinalityBuilder := cardinalities.NewBuilder()
	referenceBuilder := references.NewBuilder()
	blockNameAfterFirstByteCharacters := createBlockNameCharacters()
	possibleLowerCaseLetters := createPossibleLowerCaseLetters()
	possibleUpperCaseLetters := createPossibleUpperCaseLetters()
	possibleNumbers := createPossibleNumbers()
	possibleFuncNameCharacters := createPossibleFuncNameCharacters()
	return createAdapter(
		grammarBuilder,
		constantsBuilder,
		constantBuilder,
		constantTokensBuilder,
		constantTokenBuilder,
		constantElementBuilder,
		blocksBuilder,
		blockBuilder,
		suitesBuilder,
		suiteBuilder,
		linesBuilder,
		lineBuilder,
		balanceBuilder,
		selectorsBuilder,
		selectorBuilder,
		selectorChainBuilder,
		selectorChainTokenBuilder,
		selectorChainElementBuilder,
		tokensBuilder,
		tokenBuilder,
		reverseBuilder,
		uniqueBuilder,
		elementsBuilder,
		elementBuilder,
		rulesBuilder,
		ruleBuilder,
		cardinalityBuilder,
		referenceBuilder,
		[]byte(filterBytes),
		[]byte(suiteSeparatorPrefix),
		blockNameAfterFirstByteCharacters,
		possibleLowerCaseLetters,
		possibleUpperCaseLetters,
		possibleNumbers,
		possibleFuncNameCharacters,
		[]byte(omissionPrefix)[0],
		[]byte(omissionSuffix)[0],
		[]byte(versionPrefix)[0],
		[]byte(versionSuffix)[0],
		[]byte(rootPrefix)[0],
		[]byte(rootSuffix)[0],
		[]byte(blockSuffix)[0],
		[]byte(suiteLineSuffix)[0],
		[]byte(failSeparator)[0],
		[]byte(blockDefinitionSeparator)[0],
		[]byte(linesSeparator)[0],
		[]byte(lineSeparator)[0],
		[]byte(tokenReversePrefix)[0],
		[]byte(tokenReverseEscapePrefix)[0],
		[]byte(tokenReverseEscapeSuffix)[0],
		[]byte(tokenMustBeUnique)[0],
		[]byte(tokenMustNotBeUnique)[0],
		[]byte(tokenReference)[0],
		[]byte(ruleNameSeparator)[0],
		[]byte(ruleNameValueSeparator)[0],
		[]byte(ruleValuePrefix)[0],
		[]byte(ruleValueSuffix)[0],
		[]byte(ruleValueEscape)[0],
		[]byte(cardinalityOpen)[0],
		[]byte(cardinalityClose)[0],
		[]byte(cardinalitySeparator)[0],
		[]byte(cardinalityZeroPlus)[0],
		[]byte(cardinalityOnePlus)[0],
		[]byte(cardinalityOptional)[0],
		[]byte(constantNamePrefix)[0],
		[]byte(selectorChainElementPrefix),
		[]byte(selectorOperatorAnd),
		[]byte(selectorOperatorOr),
		[]byte(selectorOperatorXor),
		[]byte(openParenthesis)[0],
		[]byte(closeParenthesis)[0],
		[]byte(referenceBegin)[0],
		[]byte(referenceEnd)[0],
		[]byte(referencePathSeparator)[0],
		[]byte(referenceElementSeparator)[0],
	)
}

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the adapter
type Adapter interface {
	// ToGrammar takes the input and converts it to a grammar instance and the remaining data
	ToGrammar(input []byte) (Grammar, []byte, error)
}

// NewRepositoryMemory creates a new reposiotry memory
func NewRepositoryMemory(
	grammars map[string]Grammar,
) Repository {
	return createRepositoryMemory(
		grammars,
	)
}

// Builder represents the grammar builder
type Builder interface {
	Create() Builder
	WithVersion(version uint) Builder
	WithRoot(root elements.Element) Builder
	WithRules(rules rules.Rules) Builder
	WithBlocks(blocks blocks.Blocks) Builder
	WithOmissions(omissions elements.Elements) Builder
	WithConstants(constants constants.Constants) Builder
	Now() (Grammar, error)
}

// Grammar represents a grammar
type Grammar interface {
	Version() uint
	Root() elements.Element
	Rules() rules.Rules
	Blocks() blocks.Blocks
	HasOmissions() bool
	Omissions() elements.Elements
	HasConstants() bool
	Constants() constants.Constants
}

// Repository represents a Grammar repository
type Repository interface {
	Init() error
	List() (map[string][]uint, error)
	Insert(path []string, grammar Grammar) error
	Retrieve(reference references.Reference) (Grammar, error)
	Delete(reference references.Reference) error
}
