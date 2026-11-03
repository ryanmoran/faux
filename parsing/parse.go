package parsing

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"go/types"
	"path/filepath"
	"strconv"

	"golang.org/x/tools/go/packages"
)

const PackagesParseMode = packages.NeedName | packages.NeedFiles | packages.NeedTypes | packages.NeedImports

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

	filename, err := sourceFile(pkgs[0], pkgs[0].Fset.File(object.Pos()).Name())
	if err != nil {
		return Fake{}, err
	}

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

// sourceFile resolves the on-disk path of the file that declares the named
// type. Positions recorded in export data are not always usable paths: they
// may be "$GOROOT"-relative, or module-relative when built with -trimpath.
// Since every file in a package shares a directory, the base name is enough
// to find the matching entry in the package's absolute GoFiles.
func sourceFile(pkg *packages.Package, position string) (string, error) {
	base := filepath.Base(position)
	for _, file := range pkg.GoFiles {
		if filepath.Base(file) == base {
			return file, nil
		}
	}

	return "", fmt.Errorf("failed to locate source file %q in package %q", position, pkg.PkgPath)
}
