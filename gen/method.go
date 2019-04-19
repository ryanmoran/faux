package gen

import (
	"go/ast"
	"go/token"
)

type Method struct {
	Name     string
	Receiver string
	Params   []Argument
	Results  []Argument
}

func (m Method) FieldStruct() *ast.Field {
	var fields []*ast.Field

	var receivesFields, returnsFields []*ast.Field

	for _, argument := range m.Params {
		receivesFields = append(receivesFields, argument.Field(true))
	}

	for _, argument := range m.Results {
		returnsFields = append(returnsFields, argument.Field(true))
	}

	fields = append(fields, &ast.Field{
		Names: []*ast.Ident{
			ast.NewIdent("CallCount"),
		},
		Type: ast.NewIdent("int"),
	})

	if len(receivesFields) > 0 {
		fields = append(fields, &ast.Field{
			Names: []*ast.Ident{
				ast.NewIdent("Receives"),
			},
			Type: &ast.StructType{
				Fields: &ast.FieldList{
					List: receivesFields,
				},
			},
		})
	}

	if len(returnsFields) > 0 {
		fields = append(fields, &ast.Field{
			Names: []*ast.Ident{
				ast.NewIdent("Returns"),
			},
			Type: &ast.StructType{
				Fields: &ast.FieldList{
					List: returnsFields,
				},
			},
		})
	}

	return &ast.Field{
		Names: []*ast.Ident{
			ast.NewIdent(m.Name + "Call"),
		},
		Type: &ast.StructType{
			Fields: &ast.FieldList{
				List: fields,
			},
		},
	}
}

func (m Method) MethodDeclaration() *ast.FuncDecl {
	var params, results []*ast.Field
	for _, argument := range m.Params {
		var argType ast.Expr = ast.NewIdent(argument.Type)
		if argument.Variadic {
			argType = &ast.Ellipsis{
				Elt: argType,
			}
		}

		params = append(params, &ast.Field{
			Names: []*ast.Ident{
				ast.NewIdent(argument.Name),
			},
			Type: argType,
		})
	}

	for _, argument := range m.Results {
		results = append(results, &ast.Field{
			Type: ast.NewIdent(argument.Type),
		})
	}

	var statements []ast.Stmt

	statements = append(statements, &ast.IncDecStmt{
		X: &ast.SelectorExpr{
			X: &ast.SelectorExpr{
				X:   ast.NewIdent("f"),
				Sel: ast.NewIdent(m.Name + "Call"),
			},
			Sel: ast.NewIdent("CallCount"),
		},
		Tok: token.INC,
	})

	for _, argument := range m.Params {
		statements = append(statements, argument.AssignStatement())
	}

	var returnValues []ast.Expr
	for _, argument := range m.Results {
		returnValues = append(returnValues, argument.ReturnValue())
	}

	if len(returnValues) > 0 {
		statements = append(statements, &ast.ReturnStmt{
			Results: returnValues,
		})
	}

	return &ast.FuncDecl{
		Recv: &ast.FieldList{
			List: []*ast.Field{
				&ast.Field{
					Names: []*ast.Ident{
						ast.NewIdent("f"),
					},
					Type: ast.NewIdent("*" + m.Receiver),
				},
			},
		},
		Name: ast.NewIdent(m.Name),
		Type: &ast.FuncType{
			Params: &ast.FieldList{
				List: params,
			},
			Results: &ast.FieldList{
				List: results,
			},
		},
		Body: &ast.BlockStmt{
			List: statements,
		},
	}
}
