package uniques

import "github.com/steve-care-software/grammars/domain/grammars/blocks/lines/tokens/elements"

type unique struct {
	element elements.Element
	mustBe  bool
	mustNot bool
	index   uint
}

func createUniqueWithMustBe(element elements.Element, index uint) Unique {
	return createUniqueInternally(element, true, false, index)
}

func createUniqueWithMustNot(element elements.Element, index uint) Unique {
	return createUniqueInternally(element, false, true, index)
}

func createUniqueInternally(
	element elements.Element,
	mustBe bool,
	mustNot bool,
	index uint,
) Unique {
	out := unique{
		element: element,
		mustBe:  mustBe,
		mustNot: mustNot,
		index:   index,
	}

	return &out
}

// Element returns the Element
func (obj *unique) Element() elements.Element {
	return obj.element
}

// MustBe returns true if must be unique, false otherwise
func (obj *unique) MustBe() bool {
	return obj.mustBe
}

// MustNot returns true if must NOT be unique, false otherwise
func (obj *unique) MustNot() bool {
	return obj.mustNot
}

// Index returns the index
func (obj *unique) Index() uint {
	return obj.index
}
