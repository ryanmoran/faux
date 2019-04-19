package gen

import (
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"github.com/iancoleman/strcase"
)

type Argument struct {
	Method   string
	Name     string
	Type     string
	Variadic bool
}

func NewArgument(fieldName string, field *ast.Field, method string, fieldType types.Type, name string) Argument {
	if len(field.Names) > 0 {
		name = field.Names[0].Name
	}

	if fieldName != "" {
		name = fieldName
	}

	var variadic bool
	if _, ok := field.Type.(*ast.Ellipsis); ok {
		variadic = true
		if slice, ok := fieldType.(*types.Slice); ok {
			fieldType = slice.Elem()
		}
	}

	return Argument{
		Method: method,
		Name:   strcase.ToLowerCamel(name),
		Type: types.TypeString(fieldType, func(p *types.Package) string {
			return p.Name()
		}),
		Variadic: variadic,
	}
}

func (a Argument) Field(titleized bool) *ast.Field {
	name := a.Name
	if titleized {
		name = strings.Title(a.Name)
	}

	var fieldType ast.Expr = ast.NewIdent(a.Type)
	if a.Variadic {
		fieldType = &ast.ArrayType{
			Elt: fieldType,
		}
	}

	return &ast.Field{
		Names: []*ast.Ident{
			ast.NewIdent(name),
		},
		Type: fieldType,
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
