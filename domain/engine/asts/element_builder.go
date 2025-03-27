package asts

import (
	"errors"
)

type elementBuilder struct {
	constant    Constant
	instruction Instruction
	ast         AST
}

func createElementBuilder() ElementBuilder {
	out := elementBuilder{
		constant:    nil,
		instruction: nil,
		ast:         nil,
	}

	return &out
}

// Create initializes the elementBuilder
func (app *elementBuilder) Create() ElementBuilder {
	return createElementBuilder()
}

// WithConstant adds a constant to the elementBuilder
func (app *elementBuilder) WithConstant(constant Constant) ElementBuilder {
	app.constant = constant
	return app
}

// WithInstruction adds an instruction to the elementBuilder
func (app *elementBuilder) WithInstruction(instruction Instruction) ElementBuilder {
	app.instruction = instruction
	return app
}

// WithAST adds an ast to the elementBuilder
func (app *elementBuilder) WithAST(ast AST) ElementBuilder {
	app.ast = ast
	return app
}

// Now builds a new Element instance
func (app *elementBuilder) Now() (Element, error) {
	if app.constant != nil {
		return createElementWithConstant(app.constant), nil
	}

	if app.instruction != nil {
		return createElementWithInstruction(app.instruction), nil
	}

	if app.ast != nil {
		return createElementWithAST(app.ast), nil
	}

	return nil, errors.New("the Element is invalid")
}
