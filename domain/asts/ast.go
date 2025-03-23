package asts

type ast struct {
	root Element
}

func createAST(
	root Element,
) AST {
	return createASTInternally(root)
}

func createASTInternally(
	root Element,
) AST {
	out := ast{
		root: root,
	}

	return &out
}

// Root returns the root
func (obj *ast) Root() Element {
	return obj.root
}
