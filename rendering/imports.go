package rendering

import (
	"go/ast"
	"go/token"
	"go/types"
)

type Imports []Package

func (i *Imports) Add(pkg *types.Package) {
	for _, path := range *i {
		if pkg.Path() == path.Path {
			return
		}
	}

	*i = append(*i, NewPackage(pkg))
}

func (i Imports) Spec() []*ast.ImportSpec {
	var spec []*ast.ImportSpec
	for _, pkg := range i {
		spec = append(spec, &ast.ImportSpec{
			Name: ast.NewIdent(pkg.Name),
			Path: &ast.BasicLit{
				Kind:  token.STRING,
				Value: pkg.Path,
			},
		})
	}

	return spec
}
