package rendering

import (
	"go/ast"
)

type Struct struct {
	Fields []Field
}

func NewStruct(fields []Field) Struct {
	return Struct{
		Fields: fields,
	}
}

func (s Struct) Expr() ast.Expr {
	var fields []*ast.Field

	for _, field := range s.Fields {
		fields = append(fields, field.Field())
	}

	return &ast.StructType{
		Fields: &ast.FieldList{
			List: fields,
		},
	}
}

func (s Struct) isType() {}

func (s Struct) FieldWithName(name string) Field {
	name = TitleString(name)

	for _, field := range s.Fields {
		if field.Name == name {
			return field
		}
	}

	return Field{}
}
