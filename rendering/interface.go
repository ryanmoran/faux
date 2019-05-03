package rendering

import "go/ast"

type Interface struct{}

func (i Interface) Expr() ast.Expr {
	return &ast.InterfaceType{
		Methods: &ast.FieldList{},
	}
}
