package rendering

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/ast/astutil"
)

type File struct {
	Package string
	Imports *Imports
	Types   []NamedType
	Funcs   []Func
}

func NewFile(pkg string, imports *Imports, types []NamedType, funcs []Func) File {
	return File{
		Package: pkg,
		Imports: imports,
		Types:   types,
		Funcs:   funcs,
	}
}

func (f File) AST() *ast.File {
	var decls []ast.Decl
	for _, ty := range f.Types {
		decls = append(decls, ty.Decl())
	}

	for _, fn := range f.Funcs {
		decls = append(decls, fn.Decl())
	}

	file := &ast.File{
		Name:  ast.NewIdent(f.Package),
		Decls: decls,
	}

	for _, path := range f.Imports.used {
		for _, pkg := range f.Imports.packages {
			if pkg.Path == path {
				astutil.AddNamedImport(token.NewFileSet(), file, pkg.Name, pkg.Path)
			}
		}
	}

	return file
}
