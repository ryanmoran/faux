package gen

import (
	"fmt"
	"go/ast"
	"go/types"
	"path/filepath"
	"strings"
)

func ParseFile(filename string, name string) (Fake, error) {
	filename, err := filepath.Abs(filename)
	if err != nil {
		panic(err)
	}

	return ParsePackage(filepath.Dir(filename), name)
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
								fakeName := strings.Title(typeSpec.Name.Name)

								for _, field := range interfaceType.Methods.List {
									if funcType, ok := field.Type.(*ast.FuncType); ok {
										methodName := field.Names[0].Name

										methods = append(methods, Method{
											Name:     methodName,
											Receiver: fakeName,
											Params:   parseArguments(methodName, typesInfo, funcType.Params),
											Results:  parseArguments(methodName, typesInfo, funcType.Results),
										})
									}
								}

								return Fake{
									Name:    fakeName,
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
