package rendering

import (
	"fmt"
	"go/types"
)

type Type interface {
	Expression
	isType()
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

		return NewDefinedType(name)

	case *types.Interface:
		return Interface{}

	case *types.Pointer:
		return NewPointer(NewType(s.Elem()))

	case *types.Map:
		return NewMap(NewType(s.Key()), NewType(s.Elem()))

	case *types.Chan:
		var send, recv bool
		switch s.Dir() {
		case types.SendRecv:
			send = true
			recv = true
		case types.SendOnly:
			send = true
		case types.RecvOnly:
			recv = true
		}

		return NewChan(NewType(s.Elem()), send, recv)

	case *types.Struct:
		var fields []Field
		for i := 0; i < s.NumFields(); i++ {
			field := s.Field(i)
			fields = append(fields, NewField(field.Name(), NewType(field.Type())))
		}

		return NewStruct(fields)

	case *types.Signature:
		var params []Param
		for i := 0; i < s.Params().Len(); i++ {
			param := s.Params().At(i)
			params = append(params, NewParam("", NewType(param.Type()), false))
		}

		var results []Result
		for i := 0; i < s.Results().Len(); i++ {
			result := s.Results().At(i)
			results = append(results, NewResult(NewType(result.Type())))
		}

		return NewFunc(s.String(), Receiver{}, params, results, nil)

	default:
		panic(fmt.Sprintf("unknown type: %#v\n", t))
	}
}
