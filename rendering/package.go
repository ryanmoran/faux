package rendering

import (
	"go/types"
	"path"
)

type Package struct {
	Name string
	Path string
}

func NewPackage(pkg *types.Package) Package {
	var name string
	if path.Base(pkg.Path()) != pkg.Name() {
		name = pkg.Name()
	}

	return Package{
		Name: name,
		Path: pkg.Path(),
	}
}
