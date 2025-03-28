package grammars

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"strconv"

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

type adapter struct {
	grammarBuilder                    Builder
	constantsBuilder                  constants.Builder
	constantBuilder                   constants.ConstantBuilder
	constantTokensBuilder             constant_tokens.Builder
	constantTokenBuilder              constant_tokens.TokenBuilder
	constantElementBuilder            constant_elements.Builder
	blocksBuilder                     blocks.Builder
	blockBuilder                      blocks.BlockBuilder
	suitesBuilder                     suites.Builder
	suiteBuilder                      suites.SuiteBuilder
	linesBuilder                      lines.Builder
	lineBuilder                       lines.LineBuilder
	balanceBuilder                    balances.Builder
	selectorsBuilder                  selectors.Builder
	selectorBuilder                   selectors.SelectorBuilder
	selectorChainBuilder              chains.Builder
	selectorChainTokenBuilder         chains.TokenBuilder
	selectorChainElementBuilder       chains.ElementBuilder
	tokensBuilder                     tokens.Builder
	tokenBuilder                      tokens.TokenBuilder
	reverseBuilder                    reverses.Builder
	uniqueBuilder                     uniques.Builder
	elementsBuilder                   elements.Builder
	elementBuilder                    elements.ElementBuilder
	rulesBuilder                      rules.Builder
	ruleBuilder                       rules.RuleBuilder
	cardinalityBuilder                cardinalities.Builder
	referenceBuilder                  references.Builder
	filterBytes                       []byte
	suiteSeparatorPrefix              []byte
	blockNameAfterFirstByteCharacters []byte
	possibleLowerCaseLetters          []byte
	possibleUpperCaseLetters          []byte
	possibleNumbers                   []byte
	possibleFuncNameCharacters        []byte
	omissionPrefix                    byte
	omissionSuffix                    byte
	versionPrefix                     byte
	versionSuffix                     byte
	rootPrefix                        byte
	rootSuffix                        byte
	blockSuffix                       byte
	suiteLineSuffix                   byte
	failSeparator                     byte
	blockDefinitionSeparator          byte
	linesSeparator                    byte
	lineSeparator                     byte
	tokenReversePrefix                byte
	tokenReverseEscapePrefix          byte
	tokenReverseEscapeSuffix          byte
	tokenMustBeUnique                 byte
	tokenMustNotBeUnique              byte
	tokenReferenceSeparator           byte
	ruleNameSeparator                 byte
	ruleNameValueSeparator            byte
	ruleValuePrefix                   byte
	ruleValueSuffix                   byte
	ruleValueEscape                   byte
	cardinalityOpen                   byte
	cardinalityClose                  byte
	cardinalitySeparator              byte
	cardinalityZeroPlus               byte
	cardinalityOnePlus                byte
	cardinalityOptional               byte
	constantNamePrefix                byte
	selectorChainElementPrefix        []byte
	selectorOperatorAnd               []byte
	selectorOperatorOr                []byte
	selectorOperatorXor               []byte
	openParenthesis                   byte
	closeParenthesis                  byte
	referenceBegin                    byte
	referenceEnd                      byte
	referencePathSeparator            byte
	referenceElementSeparator         byte
}

func createAdapter(
	grammarBuilder Builder,
	constantsBuilder constants.Builder,
	constantBuilder constants.ConstantBuilder,
	constantTokensBuilder constant_tokens.Builder,
	constantTokenBuilder constant_tokens.TokenBuilder,
	constantElementBuilder constant_elements.Builder,
	blocksBuilder blocks.Builder,
	blockBuilder blocks.BlockBuilder,
	suitesBuilder suites.Builder,
	suiteBuilder suites.SuiteBuilder,
	linesBuilder lines.Builder,
	lineBuilder lines.LineBuilder,
	balanceBuilder balances.Builder,
	selectorsBuilder selectors.Builder,
	selectorBuilder selectors.SelectorBuilder,
	selectorChainBuilder chains.Builder,
	selectorChainTokenBuilder chains.TokenBuilder,
	selectorChainElementBuilder chains.ElementBuilder,
	tokensBuilder tokens.Builder,
	tokenBuilder tokens.TokenBuilder,
	reverseBuilder reverses.Builder,
	uniqueBuilder uniques.Builder,
	elementsBuilder elements.Builder,
	elementBuilder elements.ElementBuilder,
	rulesBuilder rules.Builder,
	ruleBuilder rules.RuleBuilder,
	cardinalityBuilder cardinalities.Builder,
	referenceBuilder references.Builder,
	filterBytes []byte,
	suiteSeparatorPrefix []byte,
	blockNameAfterFirstByteCharacters []byte,
	possibleLowerCaseLetters []byte,
	possibleUpperCaseLetters []byte,
	possibleNumbers []byte,
	possibleFuncNameCharacters []byte,
	omissionPrefix byte,
	omissionSuffix byte,
	versionPrefix byte,
	versionSuffix byte,
	rootPrefix byte,
	rootSuffix byte,
	blockSuffix byte,
	suiteLineSuffix byte,
	failSeparator byte,
	blockDefinitionSeparator byte,
	linesSeparator byte,
	lineSeparator byte,
	tokenReversePrefix byte,
	tokenReverseEscapePrefix byte,
	tokenReverseEscapeSuffix byte,
	tokenMustBeUnique byte,
	tokenMustNotBeUnique byte,
	tokenReferenceSeparator byte,
	ruleNameSeparator byte,
	ruleNameValueSeparator byte,
	ruleValuePrefix byte,
	ruleValueSuffix byte,
	ruleValueEscape byte,
	cardinalityOpen byte,
	cardinalityClose byte,
	cardinalitySeparator byte,
	cardinalityZeroPlus byte,
	cardinalityOnePlus byte,
	cardinalityOptional byte,
	constantNamePrefix byte,
	selectorChainElementPrefix []byte,
	selectorOperatorAnd []byte,
	selectorOperatorOr []byte,
	selectorOperatorXor []byte,
	openParenthesis byte,
	closeParenthesis byte,
	referenceBegin byte,
	referenceEnd byte,
	referencePathSeparator byte,
	referenceElementSeparator byte,
) Adapter {
	out := adapter{
		grammarBuilder:                    grammarBuilder,
		constantsBuilder:                  constantsBuilder,
		constantBuilder:                   constantBuilder,
		constantTokensBuilder:             constantTokensBuilder,
		constantTokenBuilder:              constantTokenBuilder,
		constantElementBuilder:            constantElementBuilder,
		blocksBuilder:                     blocksBuilder,
		blockBuilder:                      blockBuilder,
		suitesBuilder:                     suitesBuilder,
		suiteBuilder:                      suiteBuilder,
		linesBuilder:                      linesBuilder,
		lineBuilder:                       lineBuilder,
		balanceBuilder:                    balanceBuilder,
		selectorsBuilder:                  selectorsBuilder,
		selectorBuilder:                   selectorBuilder,
		selectorChainBuilder:              selectorChainBuilder,
		selectorChainTokenBuilder:         selectorChainTokenBuilder,
		selectorChainElementBuilder:       selectorChainElementBuilder,
		tokensBuilder:                     tokensBuilder,
		tokenBuilder:                      tokenBuilder,
		reverseBuilder:                    reverseBuilder,
		uniqueBuilder:                     uniqueBuilder,
		elementsBuilder:                   elementsBuilder,
		elementBuilder:                    elementBuilder,
		rulesBuilder:                      rulesBuilder,
		ruleBuilder:                       ruleBuilder,
		cardinalityBuilder:                cardinalityBuilder,
		referenceBuilder:                  referenceBuilder,
		filterBytes:                       filterBytes,
		suiteSeparatorPrefix:              suiteSeparatorPrefix,
		blockNameAfterFirstByteCharacters: blockNameAfterFirstByteCharacters,
		possibleLowerCaseLetters:          possibleLowerCaseLetters,
		possibleUpperCaseLetters:          possibleUpperCaseLetters,
		possibleNumbers:                   possibleNumbers,
		possibleFuncNameCharacters:        possibleFuncNameCharacters,
		omissionPrefix:                    omissionPrefix,
		omissionSuffix:                    omissionSuffix,
		versionPrefix:                     versionPrefix,
		versionSuffix:                     versionSuffix,
		rootPrefix:                        rootPrefix,
		rootSuffix:                        rootSuffix,
		suiteLineSuffix:                   suiteLineSuffix,
		failSeparator:                     failSeparator,
		blockDefinitionSeparator:          blockDefinitionSeparator,
		blockSuffix:                       blockSuffix,
		linesSeparator:                    linesSeparator,
		lineSeparator:                     lineSeparator,
		tokenReversePrefix:                tokenReversePrefix,
		tokenReverseEscapePrefix:          tokenReverseEscapePrefix,
		tokenReverseEscapeSuffix:          tokenReverseEscapeSuffix,
		tokenMustBeUnique:                 tokenMustBeUnique,
		tokenMustNotBeUnique:              tokenMustNotBeUnique,
		tokenReferenceSeparator:           tokenReferenceSeparator,
		ruleNameSeparator:                 ruleNameSeparator,
		ruleNameValueSeparator:            ruleNameValueSeparator,
		ruleValuePrefix:                   ruleValuePrefix,
		ruleValueSuffix:                   ruleValueSuffix,
		ruleValueEscape:                   ruleValueEscape,
		cardinalityOpen:                   cardinalityOpen,
		cardinalityClose:                  cardinalityClose,
		cardinalitySeparator:              cardinalitySeparator,
		cardinalityZeroPlus:               cardinalityZeroPlus,
		cardinalityOnePlus:                cardinalityOnePlus,
		cardinalityOptional:               cardinalityOptional,
		constantNamePrefix:                constantNamePrefix,
		selectorChainElementPrefix:        selectorChainElementPrefix,
		selectorOperatorAnd:               selectorOperatorAnd,
		selectorOperatorOr:                selectorOperatorOr,
		selectorOperatorXor:               selectorOperatorXor,
		openParenthesis:                   openParenthesis,
		closeParenthesis:                  closeParenthesis,
		referenceBegin:                    referenceBegin,
		referenceEnd:                      referenceEnd,
		referencePathSeparator:            referencePathSeparator,
		referenceElementSeparator:         referenceElementSeparator,
	}

	return &out
}

// ToGrammar takes the input and converts it to a grammar instance and the remaining data
func (app *adapter) ToGrammar(input []byte) (Grammar, []byte, error) {
	input = filterPrefix(input, app.filterBytes)
	retVersion, retVersionRemaining, err := extractBetween(input, app.versionPrefix, app.versionSuffix, nil)
	if err != nil {
		return nil, nil, err
	}

	version, err := strconv.Atoi(string(retVersion))
	if err != nil {
		return nil, nil, err
	}

	retVersionRemaining = filterPrefix(retVersionRemaining, app.filterBytes)
	retRootBytes, retRootRemaining, err := extractBetween(retVersionRemaining, app.rootPrefix, app.rootSuffix, nil)
	if err != nil {
		return nil, nil, err
	}

	retRoot, _, err := app.bytesToElementReference(retRootBytes)
	if err != nil {
		return nil, nil, err
	}

	retRootRemaining = filterPrefix(retRootRemaining, app.filterBytes)
	remaining := retRootRemaining
	builder := app.grammarBuilder.Create().WithVersion(uint(version)).WithRoot(retRoot)
	retOmissionBytes, retOmissionRemaining, err := extractBetween(retRootRemaining, app.omissionPrefix, app.omissionSuffix, nil)
	if err == nil {
		retOmissions, _, err := app.bytesToElementReferences(retOmissionBytes)
		if err != nil {
			return nil, nil, err
		}

		builder.WithOmissions(retOmissions)
		remaining = retOmissionRemaining
	}

	retBlocks, retBlocksRemaining, err := app.bytesToBlocks(remaining)
	if err != nil {
		return nil, nil, err
	}

	remaining = retBlocksRemaining
	builder = builder.WithBlocks(retBlocks)
	retConstants, retConstantsRemaining, err := app.bytesToConstants(remaining)
	if err == nil {
		builder.WithConstants(retConstants)
		remaining = retConstantsRemaining
	}

	retRules, retRemaining, err := app.bytesToRules(remaining)
	if err != nil {
		return nil, nil, err
	}

	ins, err := builder.
		WithRules(retRules).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, filterPrefix(retRemaining, app.filterBytes), nil
}

func (app *adapter) bytesToConstants(input []byte) (constants.Constants, []byte, error) {
	cpt := 0
	remaining := input
	list := []constants.Constant{}
	for {
		retConstant, retRemaining, err := app.bytesToConstant(remaining)
		if err != nil {
			log.Printf("there was an error while creating the constant (idx: %d): %s", cpt, err.Error())
			break
		}

		list = append(list, retConstant)
		remaining = retRemaining
		cpt++
	}

	ins, err := app.constantsBuilder.Create().WithList(list).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, filterPrefix(remaining, app.filterBytes), nil
}

func (app *adapter) bytesToConstant(input []byte) (constants.Constant, []byte, error) {
	constantName, retConstantNameRemaining, err := app.bytesToConstantDefinition(input)
	if err != nil {
		return nil, nil, err
	}

	retTokens, retTokensRemaining, err := app.bytesToConstantTokens(retConstantNameRemaining)
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.constantBuilder.Create().WithName(constantName).WithTokens(retTokens).Now()
	if err != nil {
		return nil, nil, err
	}

	if retTokensRemaining[0] != app.blockSuffix {
		str := fmt.Sprintf("the constant was expected to contain the blockSuffix byte at its suffix, data: \n%s\n", string(retTokensRemaining))
		return nil, nil, errors.New(str)
	}

	return ins, filterPrefix(retTokensRemaining[1:], app.filterBytes), nil
}

func (app *adapter) bytesToConstantTokens(input []byte) (constant_tokens.Tokens, []byte, error) {
	cpt := 0
	remaining := input
	list := []constant_tokens.Token{}
	for {
		retConstant, retRemaining, err := app.bytesToConstantToken(remaining)
		if err != nil {
			log.Printf("there was an error while creating the constant token (idx: %d): %s", cpt, err.Error())
			break
		}

		list = append(list, retConstant)
		remaining = retRemaining
		cpt++
	}

	ins, err := app.constantTokensBuilder.Create().WithList(list).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, filterPrefix(remaining, app.filterBytes), nil
}

func (app *adapter) bytesToConstantToken(input []byte) (constant_tokens.Token, []byte, error) {
	retElement, retElementRemaining, err := app.bytesToConstantTokenElementReference(input)
	if err != nil {
		return nil, nil, err
	}

	remaining := retElementRemaining
	builder := app.constantTokenBuilder.Create().WithElement(retElement).WithAmount(1)
	retBracketIndex, retRemaining, err := bytesToBracketsIndex(
		retElementRemaining,
		app.possibleNumbers,
		app.cardinalityOpen,
		app.cardinalityClose,
		app.filterBytes,
	)

	if err == nil {
		builder.WithAmount(retBracketIndex)
		remaining = retRemaining
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, filterPrefix(remaining, app.filterBytes), nil
}

func (app *adapter) bytesToConstantTokenElementReference(input []byte) (constant_elements.Element, []byte, error) {
	input = filterPrefix(input, app.filterBytes)
	if len(input) <= 0 {
		return nil, nil, errors.New("the constantToken was expected to contain at least 1 byte")
	}

	if input[0] != app.tokenReferenceSeparator {
		return nil, nil, errors.New("the constantToken was expected to contain the tokenReference byte at its prefix")
	}

	input = filterPrefix(input[1:], app.filterBytes)
	return app.bytesToConstantTokenElement(input)
}

func (app *adapter) bytesToConstantTokenElement(input []byte) (constant_elements.Element, []byte, error) {
	// try to match a rule
	elementBuilder := app.constantElementBuilder.Create()
	ruleName, retRemaining, err := app.bytesToRuleName(input)
	if err != nil {
		// there is no rule, so try to match a constant
		constantName, retConstantRemaining, err := app.bytesToConstantName(input)
		if err == nil {
			elementBuilder.WithConstant(string(constantName))
			retRemaining = retConstantRemaining
		}
	} else {
		elementBuilder.WithRule(ruleName)
	}

	element, err := elementBuilder.Now()
	if err != nil {
		return nil, nil, err
	}

	return element, filterPrefix(retRemaining, app.filterBytes), nil
}

func (app *adapter) bytesToBlocks(input []byte) (blocks.Blocks, []byte, error) {
	cpt := 0
	remaining := input
	list := []blocks.Block{}
	for {
		retBlock, retRemaining, err := app.bytesToBlock(remaining)
		if err != nil {
			log.Printf("there was an error while creating the block (idx: %d): %s", cpt, err.Error())
			break
		}

		list = append(list, retBlock)
		remaining = retRemaining
		cpt++
	}

	ins, err := app.blocksBuilder.Create().WithList(list).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, filterPrefix(remaining, app.filterBytes), nil
}

func (app *adapter) bytesToBlock(input []byte) (blocks.Block, []byte, error) {
	blockName, retBlockNameRemaining, err := app.bytesToBlockDefinition(input)
	if err != nil {
		return nil, nil, err
	}

	retLines, retLinesRemaining, err := app.bytesToLines(retBlockNameRemaining)
	if err != nil {
		return nil, nil, err
	}

	remaining := retLinesRemaining
	builder := app.blockBuilder.Create().WithName(blockName).WithLines(retLines)
	retSuites, retSuitesRemaining, err := app.bytesToSuites(retLinesRemaining)
	if err == nil {
		builder.WithSuites(retSuites)
		remaining = retSuitesRemaining
	}

	if len(remaining) <= 0 {
		return nil, nil, errors.New("the block was expected to contain at least 1 byte at the end of its definition")
	}

	if remaining[0] != app.blockSuffix {
		str := fmt.Sprintf("the block was expected to contain the blockSuffix byte at its suffix, data: \n%s\n", string(remaining))
		return nil, nil, errors.New(str)
	}

	retIns, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return retIns, filterPrefix(remaining[1:], app.filterBytes), nil
}

func (app *adapter) bytesToSuites(input []byte) (suites.Suites, []byte, error) {
	input = filterPrefix(input, app.filterBytes)
	if !bytes.HasPrefix(input, app.suiteSeparatorPrefix) {
		return nil, nil, errors.New("the suite was expecting the suite prefix bytes as its prefix")
	}

	remaining := filterPrefix(input[len(app.suiteSeparatorPrefix):], app.filterBytes)
	list := []suites.Suite{}
	for {
		retSuite, retRemaining, err := app.bytesToSuite(remaining)
		if err != nil {
			break
		}

		list = append(list, retSuite)
		remaining = filterPrefix(retRemaining, app.filterBytes)
	}

	ins, err := app.suitesBuilder.Create().WithList(list).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, filterPrefix(remaining, app.filterBytes), nil
}

func (app *adapter) bytesToSuite(input []byte) (suites.Suite, []byte, error) {
	testName, retBlockNameRemaining, err := app.bytesToBlockDefinition(input)
	if err != nil {
		return nil, nil, err
	}

	remaining := retBlockNameRemaining
	builder := app.suiteBuilder.Create().WithName(testName)
	if len(retBlockNameRemaining) != 0 && retBlockNameRemaining[0] == app.failSeparator {
		builder.IsFail()
		remaining = retBlockNameRemaining[1:]
	}

	retSuiteInput, retRemainingAfterBetween, err := extractBetween(remaining, app.ruleValuePrefix, app.ruleValueSuffix, &app.ruleValueEscape)
	if err != nil {
		return nil, nil, err
	}

	retIns, err := builder.WithInput(retSuiteInput).Now()
	if err != nil {
		return nil, nil, err
	}

	if len(retRemainingAfterBetween) <= 0 {
		return nil, nil, errors.New("the suite was expected to contain at least 1 byte at the end of its instruction")
	}

	if retRemainingAfterBetween[0] != app.suiteLineSuffix {
		return nil, nil, errors.New("the suite was expected to contain the suiteLineSuffix byte at its suffix")
	}

	return retIns, filterPrefix(retRemainingAfterBetween[1:], app.filterBytes), nil
}

func (app *adapter) bytesToBlockDefinition(input []byte) (string, []byte, error) {
	blockName, retBlockRemaining, err := app.bytesToBlockName(input)
	if err != nil {
		return "", nil, err
	}

	if len(retBlockRemaining) <= 0 {
		return "", nil, errors.New("the blockDefinition was expected to contain at least 1 byte after fetching its name")
	}

	if retBlockRemaining[0] != app.blockDefinitionSeparator {
		return "", nil, errors.New("the blockDefinition was expected to contain the blockDefinitionSeparator byte at its suffix")
	}

	return blockName, filterPrefix(retBlockRemaining[1:], app.filterBytes), nil
}

func (app *adapter) bytesToConstantDefinition(input []byte) (string, []byte, error) {
	constantName, retBlockRemaining, err := app.bytesToConstantName(input)
	if err != nil {
		return "", nil, err
	}

	if len(retBlockRemaining) <= 0 {
		return "", nil, errors.New("the constantDefinition was expected to contain at least 1 byte after fetching its name")
	}

	if retBlockRemaining[0] != app.blockDefinitionSeparator {
		return "", nil, errors.New("the constantDefinition was expected to contain the blockDefinitionSeparator byte at its suffix")
	}

	return constantName, filterPrefix(retBlockRemaining[1:], app.filterBytes), nil
}

func (app *adapter) bytesToLines(input []byte) (lines.Lines, []byte, error) {
	remaining := input
	list := []lines.Line{}
	cpt := 0
	for {

		if len(remaining) <= 0 {
			break
		}

		isFirst := cpt <= 0
		if !isFirst && remaining[0] != app.linesSeparator {
			break
		}

		if !isFirst {
			remaining = filterPrefix(remaining[1:], app.filterBytes)
		}

		retLine, retRemaining, err := app.bytesToLine(remaining)
		if err != nil {
			break
		}

		list = append(list, retLine)
		remaining = filterPrefix(retRemaining, app.filterBytes)
		cpt++
	}

	ins, err := app.linesBuilder.Create().WithList(list).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, filterPrefix(remaining, app.filterBytes), nil
}

func (app *adapter) bytesToLine(input []byte) (lines.Line, []byte, error) {
	remaining := input
	builder := app.lineBuilder.Create()
	retTokens, retRemaining, err := app.bytesToTokens(remaining)
	if err != nil {
		return nil, nil, err
	}

	remaining = retRemaining
	retBalance, retRemaining, err := app.bytesToBalance(remaining)
	if err == nil {
		builder.WithBalance(retBalance)
		remaining = retRemaining
	}

	builder.WithTokens(retTokens)
	line, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return line, filterPrefix(remaining, app.filterBytes), nil
}

func (app *adapter) bytesToBalance(input []byte) (balances.Balance, []byte, error) {
	input = filterPrefix(input, app.filterBytes)
	if len(input) <= 0 {
		return nil, nil, errors.New("the balance was expected to contain at least 1 byte")
	}

	if input[0] != app.cardinalityOpen {
		return nil, nil, errors.New("the balance was expected to contain the cardinalityOpen byte at its prefix")
	}

	selectorsList, retRemaining, err := app.bytesToSelectorsList(input[1:])
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.balanceBuilder.Create().WithLines(selectorsList).Now()
	if err != nil {
		return nil, nil, err
	}

	if len(retRemaining) < 2 {
		return nil, nil, errors.New("the balance was expected to contain at least 2 bytes")
	}

	if retRemaining[0] != app.cardinalityClose {
		return nil, nil, errors.New("the balance was expected to contain the cardinalityClose byte at its suffix")
	}

	remainng := retRemaining[1:]
	if remainng[0] != app.blockSuffix {
		return nil, nil, errors.New("the balance was expected to contain the blockSuffix byte at its suffix")
	}

	return ins, filterPrefix(remainng[1:], app.filterBytes), nil
}

func (app *adapter) bytesToSelectorsList(input []byte) ([]selectors.Selectors, []byte, error) {
	list := []selectors.Selectors{}
	remaining := input
	for {
		retSelectors, retRemaining, err := app.bytesToSelectors(remaining)
		if err != nil {
			break
		}

		list = append(list, retSelectors)
		remaining = retRemaining
	}

	return list, filterPrefix(remaining, app.filterBytes), nil
}

func (app *adapter) bytesToSelectors(input []byte) (selectors.Selectors, []byte, error) {
	list := []selectors.Selector{}
	remaining := input
	for {

		retSelector, retRemaining, err := app.bytesToSelector(remaining)
		if err != nil {
			break
		}

		list = append(list, retSelector)
		remaining = retRemaining
		if len(remaining) > 0 && remaining[0] == app.blockSuffix {
			remaining = remaining[1:]
			break
		}

		if len(remaining) <= 0 {
			break
		}

		if remaining[0] != app.blockDefinitionSeparator {
			return nil, nil, errors.New("the selectors were expected to contain the blockDefinitionSeparator byte between them")
		}

		remaining = filterPrefix(remaining[1:], app.filterBytes)
	}

	ins, err := app.selectorsBuilder.Create().WithList(list).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, filterPrefix(remaining, app.filterBytes), nil
}

func (app *adapter) bytesToSelector(input []byte) (selectors.Selector, []byte, error) {
	remaining := filterPrefix(input, app.filterBytes)
	if len(remaining) <= 0 {
		return nil, nil, errors.New("the balance was expected to contain at least 1 byte")
	}

	builder := app.selectorBuilder.Create()
	if remaining[0] == app.tokenReversePrefix {
		builder.IsNot()
		remaining = remaining[1:]
	}

	if len(remaining) <= 0 {
		return nil, nil, errors.New("the selector was expected to contain at least 1 byte")
	}

	if remaining[0] != app.tokenReferenceSeparator {
		return nil, nil, errors.New("the selector was expected to contain the tokenReference byte at its prefix")
	}

	retChain, retRemaining, err := app.bytesToSelectorChain(remaining[1:])
	if err != nil {
		return nil, nil, err
	}

	ins, err := builder.WithChain(retChain).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, filterPrefix(retRemaining, app.filterBytes), nil
}

func (app *adapter) bytesToSelectorChain(input []byte) (chains.Chain, []byte, error) {
	retElement, retRemaining, err := app.bytesToElement(input)
	if err != nil {
		return nil, nil, err
	}

	remaining := retRemaining
	builder := app.selectorChainBuilder.Create().WithElement(retElement)
	retToken, retRemainingAfterToken, err := app.bytesToSelectorChainToken(remaining)
	if err == nil {
		builder.WithToken(retToken)
		remaining = retRemainingAfterToken
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, filterPrefix(remaining, app.filterBytes), nil
}

func (app *adapter) bytesToSelectorChainToken(input []byte) (chains.Token, []byte, error) {
	retIndex, retRemaining, err := bytesToBracketsIndex(
		input,
		app.possibleNumbers,
		app.cardinalityOpen,
		app.cardinalityClose,
		app.filterBytes,
	)

	if err != nil {
		return nil, nil, err
	}

	remaining := retRemaining
	builder := app.selectorChainTokenBuilder.Create().WithIndex(retIndex)
	retElement, retRemainingAfterElement, err := app.bytesToSelectorChainElement(remaining)
	if err == nil {
		remaining = retRemainingAfterElement
		builder.WithElement(retElement)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, filterPrefix(remaining, app.filterBytes), nil
}

func (app *adapter) bytesToSelectorChainElement(input []byte) (chains.Element, []byte, error) {
	retIndex, retRemaining, err := bytesToBracketsIndex(
		input,
		app.possibleNumbers,
		app.cardinalityOpen,
		app.cardinalityClose,
		app.filterBytes,
	)

	if err != nil {
		return nil, nil, err
	}

	remaining := retRemaining
	builder := app.selectorChainElementBuilder.Create().WithIndex(retIndex)
	if bytes.HasPrefix(remaining, app.selectorChainElementPrefix) {
		remaining = remaining[len(app.selectorChainElementPrefix):]
		retChain, retRemainingAfterChain, err := app.bytesToSelectorChain(remaining)
		if err != nil {
			return nil, nil, err
		}

		remaining = retRemainingAfterChain
		builder.WithChain(retChain)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, filterPrefix(remaining, app.filterBytes), nil
}

func (app *adapter) bytesToTokens(input []byte) (tokens.Tokens, []byte, error) {
	list, retRemaining, err := app.bytesToTokenList(input)
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.tokensBuilder.Create().
		WithList(list).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, filterPrefix(retRemaining, app.filterBytes), nil
}

func (app *adapter) bytesToTokenList(input []byte) ([]tokens.Token, []byte, error) {
	list := []tokens.Token{}
	remaining := input
	for {
		retToken, retRemaining, err := app.bytesToToken(remaining)
		if err != nil {
			break
		}

		list = append(list, retToken)
		remaining = retRemaining
	}

	return list, filterPrefix(remaining, app.filterBytes), nil
}

func (app *adapter) bytesToToken(input []byte) (tokens.Token, []byte, error) {
	remaining := filterPrefix(input, app.filterBytes)
	builder := app.tokenBuilder.Create()
	retUnique, retRemainingAfterUnique, err := app.bytesToTokenUnique(remaining)
	if err == nil {
		builder.WithUnique(retUnique)
		remaining = retRemainingAfterUnique
	}

	retReverse, retRemainingAfterReverse, err := app.bytesToTokenReverse(remaining)
	if err == nil {
		builder.WithReverse(retReverse)
		remaining = retRemainingAfterReverse
	}

	element, retRemaining, err := app.bytesToElementReference(remaining)
	if err != nil {
		return nil, nil, err
	}

	cardinalityIns, retRemainingAfterCardinality, err := app.bytesToCardinality(retRemaining)
	if err != nil {
		ins, err := app.cardinalityBuilder.Create().WithMin(1).WithMax(1).Now()
		if err != nil {
			return nil, nil, err
		}

		cardinalityIns = ins
	}

	if err == nil {
		retRemaining = retRemainingAfterCardinality
	}

	ins, err := builder.
		WithCardinality(cardinalityIns).
		WithElement(element).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, filterPrefix(retRemaining, app.filterBytes), nil
}

func (app *adapter) bytesToTokenUnique(input []byte) (uniques.Unique, []byte, error) {
	remaining := filterPrefix(input, app.filterBytes)
	if len(remaining) <= 0 {
		return nil, nil, errors.New("the remaining was NOT expected to be nil in order to create a Unique instance")
	}

	if remaining[0] != app.tokenMustBeUnique && remaining[0] != app.tokenMustNotBeUnique {
		return nil, remaining, errors.New("the mustBeUnique or mustNotBeUnique flag coult not be found in the provided input")
	}

	builder := app.uniqueBuilder.Create()
	if remaining[0] == app.tokenMustBeUnique {
		builder.MustBe()
	}

	if remaining[0] == app.tokenMustNotBeUnique {
		builder.MustNot()
	}

	retElement, retRemaining, err := app.bytesToElementReference(remaining[1:])
	if err != nil {
		return nil, nil, err
	}

	remaining = retRemaining
	builder.WithElement(retElement).WithIndex(0)
	retAmount, retRemainingAfterBracket, err := bytesToBracketsIndex(
		remaining,
		app.possibleNumbers,
		app.cardinalityOpen,
		app.cardinalityClose,
		app.filterBytes,
	)

	if err == nil {
		builder.WithIndex(retAmount)
		remaining = retRemainingAfterBracket
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, filterPrefix(remaining, app.filterBytes), nil
}

func (app *adapter) bytesToTokenReverse(input []byte) (reverses.Reverse, []byte, error) {
	remaining := filterPrefix(input, app.filterBytes)
	if len(remaining) <= 0 {
		return nil, nil, errors.New("the tokenReverse was expected to contain at least 1 byte")
	}

	if remaining[0] != app.tokenReversePrefix {
		return nil, nil, errors.New("the tokenReverse was expected to contain the tokenReversePrefix byte at its prefix")
	}

	remaining = remaining[1:]
	builder := app.reverseBuilder.Create()
	retEscape, retRemaining, err := app.bytesToTokenReverseEscape(remaining)
	if err == nil {
		remaining = retRemaining
		builder.WithEscape(retEscape)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, filterPrefix(remaining, app.filterBytes), nil
}

func (app *adapter) bytesToTokenReverseEscape(input []byte) (elements.Element, []byte, error) {
	remaining := filterPrefix(input, app.filterBytes)
	if len(remaining) <= 0 {
		return nil, nil, errors.New("the tokenReverseEscape was expected to contain at least 1 byte at its prefix")
	}

	if remaining[0] != app.tokenReverseEscapePrefix {
		return nil, nil, errors.New("the tokenReverseEscape was expected to contain the tokenReverseEscapePrefix byte at its prefix")
	}

	remaining = remaining[1:]
	retElement, retRemaining, err := app.bytesToElementReference(remaining)
	if err != nil {
		return nil, nil, err
	}

	if len(retRemaining) <= 0 {
		return nil, nil, errors.New("the tokenReverseEscape was expected to contain at least 1 byte at its suffix")
	}

	if retRemaining[0] != app.tokenReverseEscapeSuffix {
		return nil, nil, errors.New("the tokenReverseEscape was expected to contain the tokenReverseEscapeSuffix byte at its suffix")
	}

	return retElement, filterPrefix(retRemaining[1:], app.filterBytes), nil
}

func (app *adapter) bytesToElementReferences(input []byte) (elements.Elements, []byte, error) {
	list := []elements.Element{}
	remaining := input
	for {
		retElement, retRemaining, err := app.bytesToElementReference(remaining)
		if err != nil {
			break
		}

		list = append(list, retElement)
		remaining = retRemaining
	}

	ins, err := app.elementsBuilder.Create().WithList(list).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, filterPrefix(remaining, app.filterBytes), nil
}

func (app *adapter) bytesToElementReference(input []byte) (elements.Element, []byte, error) {
	input = filterPrefix(input, app.filterBytes)
	if len(input) <= 0 {
		return nil, nil, errors.New("the token was expected to contain at least 1 byte")
	}

	if input[0] != app.tokenReferenceSeparator {
		return nil, nil, errors.New("the token was expected to contain the tokenReference byte at its prefix")
	}

	input = filterPrefix(input[1:], app.filterBytes)
	return app.bytesToElement(input)
}

func (app *adapter) bytesToElement(input []byte) (elements.Element, []byte, error) {
	retReference, retRemaining, err := app.bytesToReference(input)
	if err != nil {
		return app.bytesToElementWithoutReference(input)
	}

	element, err := app.elementBuilder.Create().WithReference(retReference).Now()
	if err != nil {
		return nil, nil, err
	}

	return element, filterPrefix(retRemaining, app.filterBytes), nil
}

func (app *adapter) bytesToElementWithoutReference(input []byte) (elements.Element, []byte, error) {
	elementBuilder := app.elementBuilder.Create()

	// try to match a reference:
	remaining := input
	// try to match a rule
	ruleName, retRuleRemaining, err := app.bytesToRuleName(input)
	if err != nil {
		// there is no rule, so try to match a block
		blockName, retBlockRemaining, err := app.bytesToBlockName(input)
		if err != nil {
			// there is no rule or block, so try to match a constant
			constantName, retConstantRemaining, err := app.bytesToConstantName(input)
			if err != nil {
				return nil, nil, err
			}

			elementBuilder.WithConstant(string(constantName))
			remaining = retConstantRemaining

		} else {
			elementBuilder.WithBlock(string(blockName))
			remaining = retBlockRemaining
		}
	} else {
		elementBuilder.WithRule(ruleName)
		remaining = retRuleRemaining
	}

	element, err := elementBuilder.Now()
	if err != nil {
		return nil, nil, err
	}

	return element, filterPrefix(remaining, app.filterBytes), nil
}

func (app *adapter) bytesToReference(input []byte) (references.Reference, []byte, error) {
	retName, retRemaining, err := app.bytesToBlockName(input)
	if err != nil {
		return nil, nil, err
	}

	remaining := filterPrefix(retRemaining, app.filterBytes)
	if len(remaining) <= 0 {
		return nil, nil, errors.New("the token was expected to contain at least 1 byte")
	}

	if remaining[0] != app.referenceBegin {
		return nil, nil, errors.New("the token was expected to contain the referenceBegin byte at its prefix")
	}

	remaining = remaining[1:]
	endPathPos := bytes.Index(remaining, []byte{
		app.referenceElementSeparator,
	})

	if endPathPos == -1 {
		return nil, nil, errors.New("the token was expected to contain the referenceElementSeparator byte")
	}

	pathStr := string(remaining[:endPathPos])
	path := filepath.SplitList(pathStr)
	remaining = filterPrefix(remaining[endPathPos+1:], app.filterBytes)
	if len(remaining) <= 0 {
		return nil, nil, errors.New("the token was expected to contain at least 1 byte")
	}

	refEndPos := bytes.Index(remaining, []byte{
		app.referenceEnd,
	})

	if refEndPos == -1 {
		return nil, nil, errors.New("the token was expected to contain the referenceEnd byte")
	}

	version, err := strconv.Atoi(string(remaining[:refEndPos]))
	if err != nil {
		str := fmt.Sprintf("the reference (path: %s) does not contain a valid version, it should be a positive number", pathStr)
		return nil, nil, errors.New(str)
	}

	retIns, err := app.referenceBuilder.Create().WithPath(path).WithName(retName).WithVersion(uint(version)).Now()
	if err != nil {
		return nil, nil, err
	}

	return retIns, remaining[refEndPos+1:], nil
}

func (app *adapter) bytesToCardinality(input []byte) (cardinalities.Cardinality, []byte, error) {
	retMin, pRetMax, retRemaining, err := bytesToMinMax(
		input,
		app.possibleNumbers,
		app.cardinalityOpen,
		app.cardinalityClose,
		app.cardinalitySeparator,
		app.cardinalityZeroPlus,
		app.cardinalityOnePlus,
		app.cardinalityOptional,
		app.filterBytes,
	)

	if err != nil {
		return nil, nil, err
	}

	builder := app.cardinalityBuilder.Create().WithMin(retMin)
	if pRetMax != nil {
		builder.WithMax(*pRetMax)
	}

	retIns, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return retIns, filterPrefix(retRemaining, app.filterBytes), nil
}

func (app *adapter) bytesToRules(input []byte) (rules.Rules, []byte, error) {
	remaining := filterPrefix(input, app.filterBytes)
	list := []rules.Rule{}
	for {
		retRule, retRemaining, err := app.bytesToRule(remaining)
		if err != nil {
			break
		}

		list = append(list, retRule)
		remaining = filterPrefix(retRemaining, app.filterBytes)
	}

	ins, err := app.rulesBuilder.Create().WithList(list).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, filterPrefix(remaining, app.filterBytes), nil
}

func (app *adapter) bytesToRule(input []byte) (rules.Rule, []byte, error) {
	name, value, remaining, err := bytesToRuleNameAndValue(
		input,
		app.ruleNameValueSeparator,
		app.possibleUpperCaseLetters,
		app.ruleNameSeparator,
		app.ruleValuePrefix,
		app.ruleValueSuffix,
		app.ruleValueEscape,
		app.filterBytes,
	)

	if err != nil {
		return nil, nil, err
	}

	if len(remaining) <= 0 {
		return nil, nil, errors.New("the rule was expected to contain at least 1 byte at the end of its definition")
	}

	if remaining[0] != app.blockSuffix {
		return nil, nil, errors.New("the rule was expected to contain the blockSuffix byte at its suffix")
	}

	ins, err := app.ruleBuilder.Create().
		WithName(string(name)).
		WithBytes(value).
		Now()

	if err != nil {
		return nil, nil, err
	}

	return ins, filterPrefix(remaining[1:], app.filterBytes), nil
}

func (app *adapter) bytesToBlockName(input []byte) (string, []byte, error) {
	blockName, retBlockRemaining, err := blockName(input, app.possibleLowerCaseLetters, app.blockNameAfterFirstByteCharacters, app.filterBytes)
	if err != nil {
		return "", nil, err
	}

	return string(blockName), filterPrefix(retBlockRemaining, app.filterBytes), nil
}

func (app *adapter) bytesToConstantName(input []byte) (string, []byte, error) {
	blockName, retBlockRemaining, err := blockName(input, []byte{app.constantNamePrefix}, app.blockNameAfterFirstByteCharacters, app.filterBytes)
	if err != nil {
		return "", nil, err
	}

	return string(blockName), filterPrefix(retBlockRemaining, app.filterBytes), nil
}

func (app *adapter) bytesToRuleName(input []byte) (string, []byte, error) {
	retRuleName, retRemaining, err := bytesToRuleName(
		input,
		app.possibleUpperCaseLetters,
		app.ruleNameSeparator,
		app.filterBytes,
	)

	if err != nil {
		return "", nil, err
	}

	return string(retRuleName), filterPrefix(retRemaining, app.filterBytes), nil
}
