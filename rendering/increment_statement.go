package rendering

import (
	"go/ast"
	"go/token"
)

type IncrementStatement struct {
	Elem Type
}

func NewIncrementStatement(elem Type) IncrementStatement {
	return IncrementStatement{
		Elem: elem,
	}
}

func (is IncrementStatement) Stmt() ast.Stmt {
	return &ast.IncDecStmt{
		X:   is.Elem.Expr(),
		Tok: token.INC,
	}
}
