package rendering

import "go/ast"

type Receiver struct {
	Name string
	Type Type
}

func NewReceiver(name string, t Type) Receiver {
	return Receiver{
		Name: name,
		Type: t,
	}
}

func (r Receiver) Expr() ast.Expr {
	return r.Ident()
}

func (r Receiver) Ident() *ast.Ident {
	return ast.NewIdent(r.Name)
}

func (r Receiver) Field() *ast.Field {
	return &ast.Field{
		Names: []*ast.Ident{r.Ident()},
		Type:  r.Type.Expr(),
	}
}
