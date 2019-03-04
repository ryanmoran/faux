package gen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"strings"
)

type Argument struct {
	Method string
	Name   string
	Type   string
}

func NewArgument(field *ast.Field, method string) (Argument, error) {
	var buf bytes.Buffer
	err := format.Node(&buf, token.NewFileSet(), field.Type)
	if err != nil {
		return Argument{}, fmt.Errorf("could not format argument field type: %s", err)
	}

	return Argument{
		Method: method,
		Name:   field.Names[0].Name,
		Type:   buf.String(),
	}, nil
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
