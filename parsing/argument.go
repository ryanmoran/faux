package parsing

import (
	"go/types"
)

type Argument struct {
	Name     string
	Type     types.Type
	Variadic bool
	Package  string
}

func NewArgument(v *types.Var, variadic bool) Argument {
	var pkg string
	if t, ok := v.Type().(*types.Named); ok {
		if t.Obj().Pkg() != nil {
			pkg = t.Obj().Pkg().Path()
		}
	}

	return Argument{
		Name:     v.Name(),
		Type:     v.Type(),
		Variadic: variadic,
		Package:  pkg,
	}
}
