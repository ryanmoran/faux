package rendering

import "go/ast"

type CallStatement struct {
	Call Call
}

func NewCallStatement(call Call) CallStatement {
	return CallStatement{
		Call: call,
	}
}

func (cs CallStatement) Stmt() ast.Stmt {
	return &ast.ExprStmt{
		X: cs.Call.Expr(),
	}
}
