package rendering

import "go/ast"

type Slice struct {
	Elem Type
}

func NewSlice(elem Type) Slice {
	return Slice{
		Elem: elem,
	}
}

func (s Slice) Expr() ast.Expr {
	return &ast.ArrayType{
		Elt: s.Elem.Expr(),
	}
}

func (s Slice) isType() {}
