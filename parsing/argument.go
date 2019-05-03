package parsing

import "go/types"

type Argument struct {
	Name     string
	Type     types.Type
	Variadic bool
}

func NewArgument(v *types.Var, variadic bool) Argument {
	return Argument{
		Name:     v.Name(),
		Type:     v.Type(),
		Variadic: variadic,
	}
}
