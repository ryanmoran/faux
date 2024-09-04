package parsing

import (
	"go/types"
)

type Argument struct {
	Name     string
	Type     types.Type
	TypeArgs []types.Type
	Variadic bool
	Package  string
}

func NewArgument(v *types.Var, variadic bool) Argument {
	var (
		pkg      string
		typeArgs []types.Type
	)

	if t, ok := v.Type().(*types.Named); ok {
		targs := t.TypeArgs()
		for i := 0; i < targs.Len(); i++ {
			typeArgs = append(typeArgs, targs.At(i))
		}

		if t.Obj().Pkg() != nil {
			pkg = t.Obj().Pkg().Path()
		}
	}

	return Argument{
		Name:     v.Name(),
		Type:     v.Type(),
		TypeArgs: typeArgs,
		Variadic: variadic,
		Package:  pkg,
	}
}
