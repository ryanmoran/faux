package rendering

import (
	"fmt"
	"go/types"

	"github.com/ryanmoran/faux/parsing"
)

type Type interface {
	Expression
	isType()
}

func NewType(t types.Type, targs []types.Type, imports []parsing.Import) Type {
	switch s := t.(type) {
	case *types.Slice:
		return NewSlice(NewType(s.Elem(), nil, imports))

	case *types.Basic:
		return NewBasicType(s)

	case *types.Named:
		name := s.String()

		obj := s.Obj()
		pkg := obj.Pkg()
		if pkg != nil {
			pname := pkg.Name()
			for _, p := range imports {
				if p.Path == pkg.Path() && p.Name != "" {
					pname = p.Name
				}
			}

			name = fmt.Sprintf("%s.%s", pname, obj.Name())
		}
		var targTypes []Type
		for _, targ := range targs {
			targTypes = append(targTypes, NewType(targ, nil, imports))
		}

		return NewDefinedType(name, targTypes)

	case *types.TypeParam:
		return NewNamedType(s.String(), NewType(s.Constraint(), nil, imports), nil)

	case *types.Interface:
		return Interface{}

	case *types.Pointer:
		return NewPointer(NewType(s.Elem(), nil, imports))

	case *types.Map:
		return NewMap(NewType(s.Key(), nil, imports), NewType(s.Elem(), nil, imports))

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

		return NewChan(NewType(s.Elem(), nil, imports), send, recv)

	case *types.Struct:
		var fields []Field
		for i := 0; i < s.NumFields(); i++ {
			field := s.Field(i)
			fields = append(fields, NewField(field.Name(), NewType(field.Type(), nil, imports)))
		}

		return NewStruct(fields)

	case *types.Signature:
		var params []Param
		for i := 0; i < s.Params().Len(); i++ {
			param := s.Params().At(i)
			params = append(params, NewParam("", NewType(param.Type(), nil, imports), false))
		}

		var results []Result
		for i := 0; i < s.Results().Len(); i++ {
			result := s.Results().At(i)
			results = append(results, NewResult(NewType(result.Type(), nil, imports)))
		}

		return NewFunc(s.String(), Receiver{}, params, results, nil)

	default:
		panic(fmt.Sprintf("unknown type: %#v\n", t))
	}
}
