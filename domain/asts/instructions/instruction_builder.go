package instructions

import (
	"errors"
)

type builder struct {
	block  string
	pLine  *uint
	tokens Tokens
}

func createBuilder() Builder {
	out := builder{
		block:  "",
		pLine:  nil,
		tokens: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithBlock adds a block to the builder
func (app *builder) WithBlock(block string) Builder {
	app.block = block
	return app
}

// WithLine adds a line to the builder
func (app *builder) WithLine(line uint) Builder {
	app.pLine = &line
	return app
}

// WithTokens add tokens to the builder
func (app *builder) WithTokens(tokens Tokens) Builder {
	app.tokens = tokens
	return app
}

// Now builds a new Instruction instance
func (app *builder) Now() (Instruction, error) {
	if app.block == "" {
		return nil, errors.New("the block is mandatory in order to build an Instruction")
	}

	if app.pLine == nil {
		return nil, errors.New("the line is mandatory in order to build an Instruction")
	}

	if app.tokens == nil {
		return nil, errors.New("the tokens is mandatory in order to build an Instruction")
	}

	return createInstruction(app.block, *app.pLine, app.tokens), nil
}
