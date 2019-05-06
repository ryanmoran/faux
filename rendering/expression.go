package rendering

import "go/ast"

type Expression interface {
	Expr() ast.Expr
}
