package rendering

import (
	"go/types"
)

type Imports struct {
	packages []Package
	used     []string
}

func (i *Imports) Add(pkg *types.Package) {
	for index, p := range i.packages {
		if pkg.Path() == p.Path && p.Name == "." {
			i.packages[index].Name = ""
		}
	}

	for _, path := range i.used {
		if pkg.Path() == path {
			return
		}
	}

	i.used = append(i.used, pkg.Path())
}

func (i Imports) Lookup(path string) string {
	for _, pkg := range i.packages {
		if pkg.Path == path {
			return pkg.Name
		}
	}

	return ""
}
