package rendering

import "go/ast"

type Statement interface {
	Stmt() ast.Stmt
}
