package rendering

import (
	"go/ast"
	"go/token"
)

type IncrementStatement struct {
	X Expression
}

func NewIncrementStatement(expr Expression) IncrementStatement {
	return IncrementStatement{
		X: expr,
	}
}

func (is IncrementStatement) Stmt() ast.Stmt {
	return &ast.IncDecStmt{
		X:   is.X.Expr(),
		Tok: token.INC,
	}
}
