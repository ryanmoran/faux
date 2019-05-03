package rendering

import "go/ast"

type Map struct {
	Key  Type
	Elem Type
}

func NewMap(key, elem Type) Map {
	return Map{
		Key:  key,
		Elem: elem,
	}
}

func (m Map) Expr() ast.Expr {
	return &ast.MapType{
		Key:   m.Key.Expr(),
		Value: m.Elem.Expr(),
	}
}
