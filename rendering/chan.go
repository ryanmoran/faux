package rendering

import "go/ast"

type Chan struct {
	Elem Type
	Send bool
	Recv bool
}

func NewChan(elem Type, send, recv bool) Chan {
	return Chan{
		Elem: elem,
		Send: send,
		Recv: recv,
	}
}

func (c Chan) Expr() ast.Expr {
	var dir ast.ChanDir
	if c.Send {
		dir = dir | ast.SEND
	}
	if c.Recv {
		dir = dir | ast.RECV
	}

	return &ast.ChanType{
		Dir:   dir,
		Value: c.Elem.Expr(),
	}
}

func (c Chan) isType() {}
