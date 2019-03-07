package gen

import (
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"github.com/iancoleman/strcase"
)

type Argument struct {
	Method string
	Name   string
	Type   string
}

func NewArgument(field *ast.Field, method string, fieldType types.Type, name string) Argument {
	if len(field.Names) > 0 {
		name = field.Names[0].Name
	}

	return Argument{
		Method: method,
		Name:   strcase.ToLowerCamel(name),
		Type: types.TypeString(fieldType, func(p *types.Package) string {
			return p.Name()
		}),
	}
}

func (a Argument) Field(titleized bool) *ast.Field {
	name := a.Name
	if titleized {
		name = strings.Title(a.Name)
	}

	return &ast.Field{
		Names: []*ast.Ident{
			ast.NewIdent(name),
		},
		Type: ast.NewIdent(a.Type),
	}
}

func (a Argument) AssignStatement() *ast.AssignStmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{
			&ast.SelectorExpr{
				X: &ast.SelectorExpr{
					X: &ast.SelectorExpr{
						X:   ast.NewIdent("f"),
						Sel: ast.NewIdent(a.Method + "Call"),
					},
					Sel: ast.NewIdent("Receives"),
				},
				Sel: ast.NewIdent(strings.Title(a.Name)),
			},
		},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{
			ast.NewIdent(a.Name),
		},
	}
}

func (a Argument) ReturnValue() *ast.SelectorExpr {
	return &ast.SelectorExpr{
		X: &ast.SelectorExpr{
			X: &ast.SelectorExpr{
				X:   ast.NewIdent("f"),
				Sel: ast.NewIdent(a.Method + "Call"),
			},
			Sel: ast.NewIdent("Returns"),
		},
		Sel: ast.NewIdent(strings.Title(a.Name)),
	}
}
