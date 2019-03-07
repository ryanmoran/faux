package gen

import (
	"fmt"

	"golang.org/x/tools/go/packages"
)

func ParsePackage(path, name string) (Fake, error) {
	pkgs, err := packages.Load(&packages.Config{Mode: packages.LoadSyntax}, path)
	if err != nil {
		panic(err)
	}

	for _, pkg := range pkgs {
		fake, found, err := parse(name, pkg.TypesInfo.Types, pkg.Syntax...)
		if err != nil {
			panic(err)
		}
		if found {
			return fake, nil
		}
	}

	return Fake{}, fmt.Errorf("could not find interface %q in %s", name, path)
}
