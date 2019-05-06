package rendering

import "go/ast"

type Pointer struct {
	Elem Type
}

func NewPointer(elem Type) Pointer {
	return Pointer{
		Elem: elem,
	}
}

func (p Pointer) Expr() ast.Expr {
	return &ast.StarExpr{
		X: p.Elem.Expr(),
	}
}

func (p Pointer) isType() {}
