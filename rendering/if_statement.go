package rendering

import "go/ast"

type IfStatement struct {
	Condition Expression
	Body      []Statement
}

func NewIfStatement(condition Expression, body []Statement) IfStatement {
	return IfStatement{
		Condition: condition,
		Body:      body,
	}
}

func (is IfStatement) Stmt() ast.Stmt {
	var body []ast.Stmt
	for _, statement := range is.Body {
		body = append(body, statement.Stmt())
	}

	return &ast.IfStmt{
		Cond: is.Condition.Expr(),
		Body: &ast.BlockStmt{
			List: body,
		},
	}
}
