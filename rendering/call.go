package rendering

import "go/ast"

type Call struct {
	Name string
}

func NewCall(name string) Call {
	return Call{
		Name: name,
	}
}

func (c Call) Ident() *ast.Ident {
	return ast.NewIdent(c.Name)
}
