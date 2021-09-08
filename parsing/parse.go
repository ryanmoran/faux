package parsing

import (
	"fmt"
	"go/types"

	"golang.org/x/tools/go/packages"
)

const PackagesParseMode = packages.NeedName | packages.NeedTypes | packages.NeedImports

func Parse(path, name string) (Fake, error) {
	pkgs, err := packages.Load(&packages.Config{Mode: PackagesParseMode}, path)
	if err != nil {
		return Fake{}, err
	}

	if len(pkgs) != 1 {
		return Fake{}, fmt.Errorf("failed to find package: %q", path)
	}
	pkg := pkgs[0].Types

	object := pkg.Scope().Lookup(name)
	if object == nil {
		return Fake{}, fmt.Errorf("failed to find named type: %s.%s", path, name)
	}

	namedType, ok := object.Type().(*types.Named)
	if !ok {
		return Fake{}, fmt.Errorf("failed to load named type: %s.%s", path, name)
	}

	var imports []Import
	for _, p := range pkg.Imports() {
		imports = append(imports, Import{
			Name: p.Name(),
			Path: p.Path(),
		})
	}

	iface, err := NewInterface(namedType)
	if err != nil {
		return Fake{}, err
	}

	return Fake{Interface: iface, Imports: imports}, nil
}
