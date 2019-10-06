package parsing

import (
	"fmt"
	"go/parser"
	"go/token"
	"go/types"
	"path/filepath"
	"strconv"

	"golang.org/x/tools/go/packages"
)

func Parse(path, name string) (Interface, error) {
	pkgs, err := packages.Load(&packages.Config{Mode: packages.LoadAllSyntax}, path)
	if err != nil {
		return Interface{}, err
	}

	if len(pkgs) != 1 {
		return Interface{}, fmt.Errorf("failed to find package: %q", path)
	}
	pkg := pkgs[0].Types

	if len(pkgs[0].GoFiles) == 0 {
		return Interface{}, fmt.Errorf("failed to load package with any files: %q", path)
	}

	dir := filepath.Dir(pkgs[0].GoFiles[0])
	astPkgs, err := parser.ParseDir(token.NewFileSet(), dir, nil, parser.ImportsOnly)
	if err != nil {
		return Interface{}, fmt.Errorf("failed to parse directory: %s", dir)
	}

	pkgMap := map[string]string{}
	for _, pkg := range astPkgs {
		for _, file := range pkg.Files {
			for _, i := range file.Imports {
				var name string
				if i.Name != nil {
					name = i.Name.Name
				}

				path, err := strconv.Unquote(i.Path.Value)
				if err != nil {
					return Interface{}, err
				}

				pkgMap[path] = name
			}
		}
	}

	object := pkg.Scope().Lookup(name)
	if object == nil {
		return Interface{}, fmt.Errorf("failed to find named type: %s.%s", path, name)
	}

	namedType, ok := object.Type().(*types.Named)
	if !ok {
		return Interface{}, fmt.Errorf("failed to load named type: %s.%s", path, name)
	}

	return NewInterface(pkgMap, namedType)
}
