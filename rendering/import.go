package rendering

import (
	"go/ast"
	"go/token"
	"path/filepath"
	"strconv"
)

type Import struct {
	Name string
	Path string
}

func NewImport(name, path string) Import {
	return Import{
		Name: name,
		Path: path,
	}
}

func (i Import) Decl() ast.Decl {
	spec := &ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: strconv.Quote(i.Path),
		},
	}

	if filepath.Base(i.Path) != i.Name {
		spec.Name = ast.NewIdent(i.Name)
	}

	return &ast.GenDecl{
		Tok:   token.IMPORT,
		Specs: []ast.Spec{spec},
	}
}
