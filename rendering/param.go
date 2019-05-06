package rendering

import (
	"go/ast"
)

type Param struct {
	Name     string
	Type     Type
	Variadic bool
}

func NewParam(name string, t Type, variadic bool) Param {
	return Param{
		Name:     name,
		Type:     t,
		Variadic: variadic,
	}
}

func (p Param) Expr() ast.Expr {
	return p.Ident()
}

func (p Param) Field() *ast.Field {
	expr := p.Type.Expr()

	if p.Variadic {
		if s, ok := p.Type.(Slice); ok {
			expr = &ast.Ellipsis{
				Elt: s.Elem.Expr(),
			}
		}
	}

	return &ast.Field{
		Names: []*ast.Ident{p.Ident()},
		Type:  expr,
	}
}

func (p Param) Ident() *ast.Ident {
	return ast.NewIdent(p.Name)
}
