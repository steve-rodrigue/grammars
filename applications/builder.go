package applications

import (
	"github.com/steve-care-software/grammars/domain/asts"
	"github.com/steve-care-software/grammars/domain/walkers"
	"github.com/steve-care-software/grammars/domain/walkers/elements"
)

type builder struct {
	elementsAdapter asts.ElementsAdapter
	astAdapter      asts.Adapter
	elementAdapter  elements.Adapter
	tokensBuilder   asts.TokensBuilder
	pElement        *elements.Element
}

func createBuilder(
	elementsAdapter asts.ElementsAdapter,
	astAdapter asts.Adapter,
	elementAdapter elements.Adapter,
	tokensBuilder asts.TokensBuilder,
) Builder {
	out := builder{
		elementsAdapter: elementsAdapter,
		astAdapter:      astAdapter,
		elementAdapter:  elementAdapter,
		tokensBuilder:   tokensBuilder,
		pElement:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.elementsAdapter,
		app.astAdapter,
		app.elementAdapter,
		app.tokensBuilder,
	)
}

// WithElement adds an element to the builder
func (app *builder) WithElement(ins elements.Element) Builder {
	app.pElement = &ins
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	var walker walkers.Walker
	if app.pElement != nil {
		retWalker, err := app.elementAdapter.ToWalker(*app.pElement)
		if err != nil {
			return nil, err
		}

		walker = retWalker
	}

	return createApplication(
		app.elementsAdapter,
		app.astAdapter,
		app.tokensBuilder,
		walker,
	), nil
}
