package rendering

import (
	"bytes"
	"fmt"
	"go/printer"
	"go/token"
	"strings"
	"unicode"

	"github.com/ryanmoran/faux/parsing"
)

func TitleString(name string) string {
	parts := strings.FieldsFunc(name, func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '_'
	})

	for i, part := range parts {
		var word string
		for j, r := range part {
			if j == 0 {
				r = unicode.ToUpper(r)
			}
			word = fmt.Sprintf("%s%c", word, r)
		}
		parts[i] = word
	}

	return strings.Join(parts, "")
}

func TypeName(t Type) string {
	switch s := t.(type) {
	case Slice:
		return fmt.Sprintf("%sSlice", TypeName(s.Elem))

	case BasicType:
		return s.Underlying.String()

	case NamedType:
		var targs []string
		for _, targ := range s.TypeArgs {
			targs = append(targs, TitleString(TypeName(targ)))
		}
		parts := strings.Split(s.Name, ".")
		return strings.Join(append([]string{parts[len(parts)-1]}, targs...), "")

	case Pointer:
		return TypeName(s.Elem)

	case Chan:
		return fmt.Sprintf("%sChannel", TypeName(s.Elem))

	default:
		buffer := bytes.NewBuffer([]byte{})
		printer.Fprint(buffer, token.NewFileSet(), s.Expr())
		return buffer.String()
	}
}

func FieldTypeName(args []parsing.Argument, index int, imports []parsing.Import) string {
	nameCounts := map[string]int{}
	counter := map[string]int{}
	for _, arg := range args {
		name := TypeName(NewType(arg.Type, arg.TypeArgs, imports))
		nameCounts[name]++
	}

	var indexedCounts []int
	for _, arg := range args {
		name := TypeName(NewType(arg.Type, arg.TypeArgs, imports))
		if nameCounts[name] > 1 {
			counter[name]++
			indexedCounts = append(indexedCounts, counter[name])
		} else {
			indexedCounts = append(indexedCounts, 0)
		}
	}

	typeName := TypeName(NewType(args[index].Type, args[index].TypeArgs, imports))
	if indexedCounts[index] > 0 {
		typeName = fmt.Sprintf("%s_%d", typeName, indexedCounts[index])
	}

	return typeName
}

func ParamName(index int) string {
	return fmt.Sprintf("param%d", index+1)
}
