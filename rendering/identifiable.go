package rendering

import "go/ast"

type Identifiable interface {
	Ident() *ast.Ident
}

type Ident struct {
	Name string
}

func NewIdent(name string) Ident {
	return Ident{
		Name: name,
	}
}

func (i Ident) Ident() *ast.Ident {
	return ast.NewIdent(i.Name)
}

func (i Ident) Expr() ast.Expr {
	return i.Ident()
}
