package rendering

import (
	"go/ast"
	"go/token"
)

type NamedType struct {
	Name string
	Type Type
}

func NewNamedType(name string, t Type) NamedType {
	return NamedType{
		Name: name,
		Type: t,
	}
}

func (nt NamedType) Expr() ast.Expr {
	return ast.NewIdent(nt.Name)
}

func (nt NamedType) Decl() ast.Decl {
	return &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: ast.NewIdent(nt.Name),
				Type: nt.Type.Expr(),
			},
		},
	}
}
