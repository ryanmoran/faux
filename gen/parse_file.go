package gen

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io"
)

func ParseFile(filename string, source io.Reader, name string) (Fake, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, source, 0)
	if err != nil {
		return Fake{}, fmt.Errorf("could not parse source: %s", err)
	}

	info := types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
	}

	conf := types.Config{Importer: importer.Default()}
	_, err = conf.Check("banana", fset, []*ast.File{file}, &info)
	if err != nil {
		panic(err)
	}

	fake, found, err := parse(name, info.Types, file)
	if err != nil {
		panic(err)
	}
	if !found {
		return Fake{}, fmt.Errorf("could not find interface %q in %s", name, filename)
	}

	return fake, nil
}

func parse(name string, typesInfo map[ast.Expr]types.TypeAndValue, files ...*ast.File) (Fake, bool, error) {
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
										methodName := field.Names[0].Name

										methods = append(methods, Method{
											Name:     methodName,
											Receiver: typeSpec.Name.Name,
											Params:   parseArguments(methodName, typesInfo, funcType.Params),
											Results:  parseArguments(methodName, typesInfo, funcType.Results),
										})
									}
								}

								return Fake{
									Name:    typeSpec.Name.Name,
									Methods: methods,
								}, true, nil
							}
						}
					}
				}
			}
		}
	}

	return Fake{}, false, nil
}

func parseArguments(methodName string, typesInfo map[ast.Expr]types.TypeAndValue, fieldList *ast.FieldList) []Argument {
	argTypeCounts := map[string]int{}

	var args []Argument
	if fieldList != nil {
		for _, field := range fieldList.List {
			fallbackName := types.ExprString(field.Type)

			if len(field.Names) > 1 {
				for _, fieldName := range field.Names {
					argTypeCounts[fallbackName] = argTypeCounts[fallbackName] + 1

					if argTypeCounts[fallbackName] > 1 {
						fallbackName = fmt.Sprintf("%s%d", fallbackName, argTypeCounts[fallbackName])
					}

					args = append(args, NewArgument(types.ExprString(fieldName), field, methodName, typesInfo[field.Type].Type, fallbackName))
				}
			} else {
				argTypeCounts[fallbackName] = argTypeCounts[fallbackName] + 1

				if argTypeCounts[fallbackName] > 1 {
					fallbackName = fmt.Sprintf("%s%d", fallbackName, argTypeCounts[fallbackName])
				}

				args = append(args, NewArgument("", field, methodName, typesInfo[field.Type].Type, fallbackName))
			}
		}
	}

	return args
}
