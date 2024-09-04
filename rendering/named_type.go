package rendering

import (
	"go/ast"
	"go/token"
)

type NamedType struct {
	Name     string
	Type     Type
	TypeArgs []Type
}

func NewNamedType(name string, t Type, targTypes []Type) NamedType {
	return NamedType{
		Name:     name,
		Type:     t,
		TypeArgs: targTypes,
	}
}

func NewDefinedType(name string, targTypes []Type) NamedType {
	return NewNamedType(name, Interface{}, targTypes)
}

func (nt NamedType) Expr() ast.Expr {
	switch len(nt.TypeArgs) {
	case 0:
		return ast.NewIdent(nt.Name)

	case 1:
		return &ast.IndexExpr{
			X:     ast.NewIdent(nt.Name),
			Index: nt.TypeArgs[0].Expr(),
		}

	default:
		var indices []ast.Expr
		for _, typeArg := range nt.TypeArgs {
			indices = append(indices, typeArg.Expr())
		}

		return &ast.IndexListExpr{
			X:       ast.NewIdent(nt.Name),
			Indices: indices,
		}
	}
}

func (nt NamedType) isType() {}

func (nt NamedType) Decl() ast.Decl {
	spec := &ast.TypeSpec{
		Name: ast.NewIdent(nt.Name),
		Type: nt.Type.Expr(),
	}

	if len(nt.TypeArgs) > 0 {
		var fields []*ast.Field
		for _, targ := range nt.TypeArgs {
			ntarg := targ.(NamedType)
			fields = append(fields, &ast.Field{
				Names: []*ast.Ident{ast.NewIdent(ntarg.Name)},
				Type:  ntarg.Type.Expr(),
			})
		}

		spec.TypeParams = &ast.FieldList{
			List: fields,
		}
	}

	return &ast.GenDecl{
		Tok:   token.TYPE,
		Specs: []ast.Spec{spec},
	}
}
