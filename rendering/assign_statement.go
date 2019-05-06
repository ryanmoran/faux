package rendering

import (
	"go/ast"
	"go/token"
)

type AssignStatement struct {
	Left  Expression
	Right Expression
}

func NewAssignStatement(left, right Expression) AssignStatement {
	return AssignStatement{
		Left:  left,
		Right: right,
	}
}

func (as AssignStatement) Stmt() ast.Stmt {
	return &ast.AssignStmt{
		Tok: token.ASSIGN,
		Lhs: []ast.Expr{as.Left.Expr()},
		Rhs: []ast.Expr{as.Right.Expr()},
	}
}
