package rendering

import "go/ast"

type CallStatement struct {
	Elem Type
}

func NewCallStatement(elem Type) CallStatement {
	return CallStatement{
		Elem: elem,
	}
}

func (cs CallStatement) Stmt() ast.Stmt {
	return &ast.ExprStmt{
		X: &ast.CallExpr{
			Fun: cs.Elem.Expr(),
		},
	}
}
