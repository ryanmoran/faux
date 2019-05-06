package rendering

import "go/ast"

type Nil struct{}

func NewNil() Nil {
	return Nil{}
}

func (n Nil) Expr() ast.Expr {
	return ast.NewIdent("nil")
}
