package rendering

import "go/ast"

type Chan struct {
	Elem Type
}

func NewChan(elem Type) Chan {
	return Chan{
		Elem: elem,
	}
}

func (c Chan) Expr() ast.Expr {
	return &ast.ChanType{
		Dir:   ast.SEND | ast.RECV,
		Value: c.Elem.Expr(),
	}
}

func (c Chan) isType() {}
