package rendering

import "go/ast"

type DeferStatement struct {
	Elem Type
}

func NewDeferStatement(elem Type) DeferStatement {
	return DeferStatement{
		Elem: elem,
	}
}

func (ds DeferStatement) Stmt() ast.Stmt {
	return &ast.DeferStmt{
		Call: &ast.CallExpr{
			Fun: ds.Elem.Expr(),
		},
	}
}
