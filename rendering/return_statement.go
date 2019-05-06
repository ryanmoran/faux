package rendering

import "go/ast"

type ReturnStatement struct {
	Results []Expression
}

func NewReturnStatement(results ...Expression) ReturnStatement {
	return ReturnStatement{
		Results: results,
	}
}

func (rs ReturnStatement) Stmt() ast.Stmt {
	var results []ast.Expr
	for _, result := range rs.Results {
		results = append(results, result.Expr())
	}

	return &ast.ReturnStmt{
		Results: results,
	}
}
