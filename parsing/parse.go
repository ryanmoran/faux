package parsing

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"go/types"
	"os"
	"runtime"
	"strconv"

	"golang.org/x/tools/go/packages"
)

const PackagesParseMode = packages.NeedName | packages.NeedTypes | packages.NeedImports

func Parse(path, name string) (Fake, error) {
	pkgs, err := packages.Load(&packages.Config{Mode: PackagesParseMode}, path)
	if err != nil {
		return Fake{}, err
	}

	if len(pkgs) != 1 {
		return Fake{}, fmt.Errorf("failed to find package: %q", path)
	}
	pkg := pkgs[0].Types

	object := pkg.Scope().Lookup(name)
	if object == nil {
		return Fake{}, fmt.Errorf("failed to find named type: %s.%s", path, name)
	}

	namedType, ok := object.Type().(*types.Named)
	if !ok {
		return Fake{}, fmt.Errorf("failed to load named type: %s.%s", path, name)
	}

	filename := os.Expand(pkgs[0].Fset.File(object.Pos()).Name(), func(name string) string {
		if name == "GOROOT" {
			return runtime.GOROOT()
		}

		return os.Getenv(name)
	})

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, parser.ImportsOnly)
	if err != nil {
		return Fake{}, fmt.Errorf("failed to parse scope file: %w", err)
	}

	imports := []Import{{Path: pkg.Path()}}
	for _, i := range file.Imports {
		buffer := bytes.NewBuffer(nil)

		var name string
		if i.Name != nil {
			err = printer.Fprint(buffer, fset, i.Name)
			if err != nil {
				return Fake{}, fmt.Errorf("failed to print import package name: %w", err)
			}
		}
		name = buffer.String()

		buffer.Reset()
		err = printer.Fprint(buffer, fset, i.Path)
		if err != nil {
			return Fake{}, fmt.Errorf("failed to print import package path: %w", err)
		}

		path, err := strconv.Unquote(buffer.String())
		if err != nil {
			return Fake{}, fmt.Errorf("failed to unquote import package path: %w", err)
		}

		imports = append(imports, Import{Name: name, Path: path})
	}

	iface, err := NewInterface(namedType)
	if err != nil {
		return Fake{}, err
	}

	return Fake{Interface: iface, Imports: imports}, nil
}
