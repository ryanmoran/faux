package rendering

import (
	"fmt"
	"go/ast"
	"go/types"
)

type Type interface {
	Expr() ast.Expr
}

func NewType(t types.Type) Type {
	switch s := t.(type) {
	case *types.Slice:
		return NewSlice(NewType(s.Elem()))

	case *types.Basic:
		return NewBasicType(s)

	case *types.Named:
		name := s.String()

		obj := s.Obj()
		pkg := obj.Pkg()
		if pkg != nil {
			name = fmt.Sprintf("%s.%s", pkg.Name(), obj.Name())
		}

		return NewNamedType(name, NewType(t.Underlying()))

	case *types.Interface:
		return Interface{}

	case *types.Pointer:
		return NewPointer(NewType(s.Elem()))

	case *types.Map:
		return NewMap(NewType(s.Key()), NewType(s.Elem()))

	case *types.Chan:
		return NewChan(NewType(s.Elem()))

	case *types.Struct:
		var fields []Field
		for i := 0; i < s.NumFields(); i++ {
			field := s.Field(i)
			fields = append(fields, NewField(field.Name(), NewType(field.Type())))
		}

		return NewStruct(fields)

	default:
		panic(fmt.Sprintf("unknown type: %#v\n", t))
	}
}
