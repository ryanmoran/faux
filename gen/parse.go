package gen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
)

func Parse(filename string, source io.Reader, name string) (Fake, error) {
	var files []*ast.File

	file, err := parser.ParseFile(token.NewFileSet(), filename, source, 0)
	if err != nil {
		return Fake{}, fmt.Errorf("could not parse source: %s", err)
	}

	files = append(files, file)

	for _, file := range files {
		for _, declaration := range file.Decls {
			if generalDeclaration, ok := declaration.(*ast.GenDecl); ok {
				for _, spec := range generalDeclaration.Specs {
					if typeSpec, ok := spec.(*ast.TypeSpec); ok {
						if interfaceType, ok := typeSpec.Type.(*ast.InterfaceType); ok {
							if typeSpec.Name.Name == name {
								var methods []Method

								for _, field := range interfaceType.Methods.List {
									if funcType, ok := field.Type.(*ast.FuncType); ok {
										var params []Argument
										methodName := field.Names[0].Name

										for _, field := range funcType.Params.List {
											argument, err := NewArgument(field, methodName)
											if err != nil {
												return Fake{}, err
											}

											params = append(params, argument)
										}

										var results []Argument

										for _, field := range funcType.Results.List {
											argument, err := NewArgument(field, methodName)
											if err != nil {
												return Fake{}, err
											}

											results = append(results, argument)
										}

										methods = append(methods, Method{
											Name:     methodName,
											Receiver: typeSpec.Name.Name,
											Params:   params,
											Results:  results,
										})
									}
								}

								return Fake{
									Name:    typeSpec.Name.Name,
									Methods: methods,
								}, nil
							}
						}
					}
				}
			}
		}
	}

	return Fake{}, fmt.Errorf("could not find interface %q in %s", name, filename)
}
