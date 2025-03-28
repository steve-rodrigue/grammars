package asts

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/grammars/domain/engine/grammars/blocks/lines/balances"
	"github.com/steve-care-software/grammars/domain/engine/grammars/blocks/lines/balances/selectors"
	"github.com/steve-care-software/grammars/domain/engine/grammars/blocks/lines/balances/selectors/chains"
)

type tokensStr struct {
	list []Token
	mp   map[string][]Token
}

func createTokens(
	list []Token,
	mp map[string][]Token,
) Tokens {
	out := tokensStr{
		list: list,
		mp:   mp,
	}

	return &out
}

// Validate validates tokens
func (obj *tokensStr) Validate(elementNameIndex map[string]BlockCount) (map[string]BlockCount, error) {
	lastBlockNameIndex := map[string]BlockCount{}
	for _, oneToken := range obj.list {
		retBlockNameIndex, err := oneToken.Validate(lastBlockNameIndex)
		if err != nil {
			return nil, err
		}

		lastBlockNameIndex = retBlockNameIndex
	}

	return lastBlockNameIndex, nil
}

// List returns the list of token
func (obj *tokensStr) List() []Token {
	return obj.list
}

// Value returns the value of the tokens
func (obj *tokensStr) Value() []byte {
	output := []byte{}
	for _, oneToken := range obj.list {
		output = append(output, oneToken.Value()...)
	}

	return output
}

// FetchAll fetches all tokens by name
func (obj *tokensStr) FetchAll(name string) ([]Token, error) {
	if ins, ok := obj.mp[name]; ok {
		return ins, nil
	}

	str := fmt.Sprintf("the token (name: %s) does not exists", name)
	return nil, errors.New(str)
}

// Fetch fetches a token by name and index
func (obj *tokensStr) Fetch(name string, idx uint) (Token, error) {
	ins, err := obj.FetchAll(name)
	if err != nil {
		return nil, err
	}

	length := uint(len(ins))
	if idx >= length {
		str := fmt.Sprintf("the token (%s) could not be found at index (%d), its length is: %d", name, idx, length)
		return nil, errors.New(str)
	}

	return ins[idx], nil
}

// Select executes a select query
func (obj *tokensStr) Select(chain chains.Chain) ([]Token, Token, Element, error) {
	elementName := chain.Element().Name()
	if chain.HasToken() {
		chainToken := chain.Token()
		tokenIndex := chainToken.Index()
		retToken, err := obj.Fetch(elementName, tokenIndex)
		if err != nil {
			return nil, nil, nil, err
		}

		if chainToken.HasElement() {
			chainElement := chainToken.Element()
			elementIndex := chainElement.Index()
			retElement, err := retToken.Elements().Fetch(elementIndex)
			if err != nil {
				return nil, nil, nil, err
			}

			if chainElement.HasChain() {
				retChain := chainElement.Chain()
				if retElement.IsConstant() {
					return nil, nil, nil, errors.New("the element was expected to contain an Instruction")
				}

				return retElement.Instruction().Tokens().Select(retChain)
			}

			return nil, nil, retElement, nil

		}

		return nil, retToken, nil, nil
	}

	retTokensList, err := obj.FetchAll(elementName)
	if err != nil {
		return nil, nil, nil, err
	}

	return retTokensList, nil, nil, nil
}

// Search search for instruction/token by name
func (obj *tokensStr) Search(name string, idx uint) (Token, error) {
	retToken, err := obj.Fetch(name, idx)
	if err == nil {
		return retToken, nil
	}

	for _, oneToken := range obj.list {
		elementList := oneToken.Elements().List()
		for _, oneElement := range elementList {
			retToken, err := oneElement.Search(name, idx)
			if err != nil {
				continue
			}

			if retToken == nil {
				continue
			}

			return retToken, nil
		}
	}

	str := fmt.Sprintf("the token (name: %s, index: %d) could not be found", name, idx)
	return nil, errors.New(str)
}

// IsBalanceValid validates the tokens against the balance
func (obj *tokensStr) IsBalanceValid(balance balances.Balance) bool {
	list := balance.Lines()
	for _, oneSelectors := range list {
		operationIsValid := true
		selectorsList := oneSelectors.List()
		for _, oneSelector := range selectorsList {
			isValid := obj.IsSelectorValid(oneSelector)
			if !isValid {
				operationIsValid = false
			}
		}

		if operationIsValid {
			return true
		}
	}

	return false
}

// IsSelectorValid validates the tokens against the selector
func (obj *tokensStr) IsSelectorValid(selector selectors.Selector) bool {
	chain := selector.Chain()
	isChainValid := obj.IsChainValid(chain)
	if selector.IsNot() {
		return !isChainValid
	}

	return isChainValid
}

// IsChainValid validates the tokens against the chain
func (obj *tokensStr) IsChainValid(chain chains.Chain) bool {
	name := chain.Element().Name()
	if chain.HasToken() {
		token := chain.Token()
		tokenIndex := token.Index()
		retASTToken, err := obj.Search(name, tokenIndex)
		if err != nil {
			return false
		}

		if token.HasElement() {
			element := token.Element()
			elementIndex := element.Index()
			retASTElement, err := retASTToken.Elements().Fetch(elementIndex)
			if err != nil {
				return false
			}

			if element.HasChain() {
				chain := element.Chain()
				return retASTElement.IsChainValid(chain)
			}

			return true
		}

		return true
	}

	_, err := obj.Fetch(name, 0)
	return err == nil
}
