package rendering

import (
	"go/types"
)

type Package struct {
	Name string
	Path string
}

func NewPackage(pkg *types.Package, pkgMap map[string]string) Package {
	name := pkgMap[pkg.Path()]
	if name == "." {
		name = ""
	}

	return Package{
		Name: name,
		Path: pkg.Path(),
	}
}
