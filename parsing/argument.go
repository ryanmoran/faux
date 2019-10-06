package parsing

import (
	"go/types"
)

type Argument struct {
	Name     string
	Type     types.Type
	Variadic bool
	Package  *types.Package
}

func NewArgument(v *types.Var, variadic bool) Argument {
	var pkg *types.Package
	switch t := v.Type().(type) {
	case *types.Named:
		if t.Obj().Pkg() != nil {
			pkg = t.Obj().Pkg()
		}
	case *types.Pointer:
		if e, ok := t.Elem().(*types.Named); ok {
			if e.Obj().Pkg() != nil {
				pkg = e.Obj().Pkg()
			}
		}
	}

	return Argument{
		Name:     v.Name(),
		Type:     v.Type(),
		Variadic: variadic,
		Package:  pkg,
	}
}
