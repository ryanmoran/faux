package gen

import "go/ast"

func Build(fake Fake) *ast.File {
	var declarations []ast.Decl

	declarations = append(declarations, fake.StructDeclaration())

	for _, method := range fake.Methods {
		declarations = append(declarations, method.MethodDeclaration())
	}

	return &ast.File{
		Name:  ast.NewIdent("fakes"),
		Decls: declarations,
	}
}
