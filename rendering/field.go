package rendering

import (
	"go/ast"
)

type Field struct {
	Name string
	Type Type
}

func NewField(name string, t Type) Field {
	return Field{
		Name: name,
		Type: t,
	}
}

func (f Field) Field() *ast.Field {
	return &ast.Field{
		Names: []*ast.Ident{f.Ident()},
		Type:  f.Type.Expr(),
	}
}

func (f Field) Expr() ast.Expr {
	return f.Ident()
}

func (f Field) Ident() *ast.Ident {
	return ast.NewIdent(f.Name)
}
