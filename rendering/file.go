package rendering

import "go/ast"

type File struct {
	Package string
	Types   []NamedType
	Funcs   []Func
}

func NewFile(pkg string, types []NamedType, funcs []Func) File {
	return File{
		Package: pkg,
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

	return &ast.File{
		Name:  ast.NewIdent(f.Package),
		Decls: decls,
	}
}
