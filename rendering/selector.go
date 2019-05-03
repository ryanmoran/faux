package rendering

import (
	"go/ast"
)

type Selector struct {
	Parts []Identifiable
}

func NewSelector(parts ...Identifiable) Selector {
	return Selector{
		Parts: parts,
	}
}

func (s Selector) Expr() ast.Expr {
	expr := &ast.SelectorExpr{
		X:   s.Parts[0].Ident(),
		Sel: s.Parts[1].Ident(),
	}

	for _, part := range s.Parts[2:] {
		expr = &ast.SelectorExpr{
			X:   expr,
			Sel: part.Ident(),
		}
	}

	return expr
}
