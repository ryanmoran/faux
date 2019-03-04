package gen

import (
	"go/ast"
	"go/token"
)

type Fake struct {
	Name    string
	Methods []Method
}

func (f Fake) StructDeclaration() *ast.GenDecl {
	var fields []*ast.Field

	for _, method := range f.Methods {
		fields = append(fields, method.FieldStruct())
	}

	return &ast.GenDecl{
		Tok: token.TYPE,
		Specs: []ast.Spec{
			&ast.TypeSpec{
				Name: ast.NewIdent(f.Name),
				Type: &ast.StructType{
					Fields: &ast.FieldList{
						List: fields,
					},
				},
			},
		},
	}
}
