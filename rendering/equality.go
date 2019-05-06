package rendering

import (
	"go/ast"
	"go/token"
)

type Equality struct {
	Equal bool
	Left  Expression
	Right Expression
}

func NewEquality(equal bool, left, right Expression) Equality {
	return Equality{
		Equal: equal,
		Left:  left,
		Right: right,
	}
}

func (e Equality) Expr() ast.Expr {
	tok := token.NEQ
	if e.Equal {
		tok = token.EQL
	}

	return &ast.BinaryExpr{
		X:  e.Left.Expr(),
		Y:  e.Right.Expr(),
		Op: tok,
	}
}
