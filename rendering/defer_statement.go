package rendering

import "go/ast"

type DeferStatement struct {
	Call Call
}

func NewDeferStatement(call Call) DeferStatement {
	return DeferStatement{
		Call: call,
	}
}

func (ds DeferStatement) Stmt() ast.Stmt {
	return &ast.DeferStmt{
		Call: ds.Call.Expr().(*ast.CallExpr),
	}
}
