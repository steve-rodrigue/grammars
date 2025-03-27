package asts

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/steve-care-software/grammars/domain/engine/grammars/blocks/lines/tokens/uniques"
)

type token struct {
	name     string
	elements Elements
	unique   uniques.Unique
}

func createToken(
	name string,
	elements Elements,
) Token {
	return createTokenInternally(
		name,
		elements,
		nil,
	)
}

func createTokenWithUnique(
	name string,
	elements Elements,
	unique uniques.Unique,
) Token {
	return createTokenInternally(
		name,
		elements,
		unique,
	)
}

func createTokenInternally(
	name string,
	elements Elements,
	unique uniques.Unique,
) Token {
	out := token{
		name:     name,
		elements: elements,
		unique:   unique,
	}

	return &out
}

// Validate validates token
func (obj *token) Validate(elementNameIndex map[string]BlockCount) (map[string]BlockCount, error) {
	if !obj.HasUnique() {
		return elementNameIndex, nil
	}

	expectedIndex := obj.unique.Index()
	elementName := obj.unique.Element().Name()
	if _, ok := elementNameIndex[elementName]; ok {
		str := fmt.Sprintf("the element (name: %s) was expected to be present to validate the uniqueness of token (name: %s, index: %d)", elementName, obj.name, expectedIndex)
		return nil, errors.New(str)
	}

	retElementNameIndex := elementNameIndex
	elementNameIndex[elementName].tokens[obj.name] = obj.Value()
	currentIndex := elementNameIndex[elementName].index
	if currentIndex != expectedIndex {
		return retElementNameIndex, nil
	}

	value, isContained := elementNameIndex[elementName].tokens[obj.name]
	if obj.unique.MustBe() && isContained {
		if bytes.Equal(obj.Value(), value) {
			str := fmt.Sprintf("the token (name: %s) was expected to be unique but it wasn't at element (name: %s, index: %d)", obj.name, elementName, elementNameIndex[elementName].index)
			return nil, errors.New(str)
		}

		str := fmt.Sprintf("the token (name: %s) was expected to be unique but it wasn't found at the provided location (name: %s, index: %d)", obj.name, elementName, elementNameIndex[elementName].index)
		return nil, errors.New(str)
	}

	if obj.unique.MustNot() && !isContained {
		if !bytes.Equal(obj.Value(), value) {
			str := fmt.Sprintf("the token (name: %s) was expected to NOT be unique but it wasn't at element (name: %s, index: %d)", obj.name, elementName, elementNameIndex[elementName].index)
			return nil, errors.New(str)
		}

		str := fmt.Sprintf("the token (name: %s) was expected to NOT be unique but it wasn't found at the provided location (name: %s, index: %d)", obj.name, elementName, elementNameIndex[elementName].index)
		return nil, errors.New(str)
	}

	return obj.elements.Validate(retElementNameIndex)
}

// Name returns the name
func (obj *token) Name() string {
	return obj.name
}

// Elements returns the elements
func (obj *token) Elements() Elements {
	return obj.elements
}

// Value returns the value of the token
func (obj *token) Value() []byte {
	return obj.elements.Value()
}

// HasUnique returns true if there is a unique, false otherwise
func (obj *token) HasUnique() bool {
	return obj.unique != nil
}

// Unique returns the unique, if any
func (obj *token) Unique() uniques.Unique {
	return obj.unique
}
