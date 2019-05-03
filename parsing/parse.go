package parsing

import (
	"fmt"
	"go/types"

	"golang.org/x/tools/go/packages"
)

func Parse(path, name string) (Interface, error) {
	pkgs, err := packages.Load(&packages.Config{Mode: packages.LoadTypes}, path)
	if err != nil {
		return Interface{}, err
	}

	if len(pkgs) != 1 {
		return Interface{}, fmt.Errorf("failed to find package: %q", path)
	}
	pkg := pkgs[0].Types

	object := pkg.Scope().Lookup(name)
	if object == nil {
		return Interface{}, fmt.Errorf("failed to find named type: %s.%s", path, name)
	}

	namedType, ok := object.Type().(*types.Named)
	if !ok {
		return Interface{}, fmt.Errorf("failed to load named type: %s.%s", path, name)
	}

	return NewInterface(namedType)
}
