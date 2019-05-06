package rendering

import (
	"go/ast"
	"go/token"
)

type Call struct {
	X      Expression
	Params []Param
}

func NewCall(expr Expression, params ...Param) Call {
	return Call{
		X:      expr,
		Params: params,
	}
}

func (c Call) Expr() ast.Expr {
	var (
		params      []ast.Expr
		ellipsisPos token.Pos
	)

	for _, param := range c.Params {
		if param.Variadic {
			ellipsisPos = 1000
		}
		params = append(params, param.Expr())
	}

	return &ast.CallExpr{
		Fun:      c.X.Expr(),
		Args:     params,
		Ellipsis: ellipsisPos,
	}
}
