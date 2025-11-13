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

	switch t := v.Type().(type) {
	case *types.Named:
		targs := t.TypeArgs()
		for i := 0; i < targs.Len(); i++ {
			typeArgs = append(typeArgs, targs.At(i))
		}

		if t.Obj().Pkg() != nil {
			pkg = t.Obj().Pkg().Path()
		}

	case *types.Alias:
		// Handle type aliases - they can also have type arguments
		// For aliases, we need to look at the underlying type to get type arguments
		if named, ok := t.Underlying().(*types.Named); ok {
			targs := named.TypeArgs()
			for i := 0; i < targs.Len(); i++ {
				typeArgs = append(typeArgs, targs.At(i))
			}
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
