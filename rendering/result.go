package rendering

import "go/ast"

type Result struct {
	Type Type
}

func NewResult(t Type) Result {
	return Result{
		Type: t,
	}
}

func (r Result) Field() *ast.Field {
	return &ast.Field{
		Type: r.Type.Expr(),
	}
}
