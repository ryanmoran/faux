package rendering

import "go/ast"

type Identifiable interface {
	Ident() *ast.Ident
}

type Ident struct {
	Name string
}

func (i Ident) Expr() ast.Expr {
	return ast.NewIdent(i.Name)
}
