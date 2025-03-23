package asts

import (
	"errors"
	"fmt"
)

type elementsStr struct {
	list []Element
}

func createElements(
	list []Element,
) Elements {
	out := elementsStr{
		list: list,
	}

	return &out
}

// Validate validates elements
func (obj *elementsStr) Validate(elementNameIndex map[string]BlockCount) (map[string]BlockCount, error) {
	lastBlockNameIndex := map[string]BlockCount{}
	for _, oneElement := range obj.list {
		retBlockNameIndex, err := oneElement.Validate(lastBlockNameIndex)
		if err != nil {
			return nil, err
		}

		lastBlockNameIndex = retBlockNameIndex
	}

	return lastBlockNameIndex, nil
}

// List returns the list of element
func (obj *elementsStr) List() []Element {
	return obj.list
}

// Fetch fetches an element by index
func (obj *elementsStr) Fetch(idx uint) (Element, error) {
	length := len(obj.list)
	if idx >= uint(length) {
		str := fmt.Sprintf("the provided index (%d) must be smaller than the length (%d) of the list", idx, length)
		return nil, errors.New(str)
	}

	return obj.list[idx], nil
}

// Value returns the value of the elements
func (obj *elementsStr) Value() []byte {
	output := []byte{}
	for _, oneElement := range obj.list {
		output = append(output, oneElement.Value()...)
	}

	return output
}
