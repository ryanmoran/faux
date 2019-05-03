package rendering_test

import (
	"go/ast"

	"github.com/ryanmoran/faux/rendering"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Selector", func() {
	Describe("Expr", func() {
		It("builds an *ast.SelectorExpr", func() {
			basicInt := rendering.NewBasicType(rendering.BasicInt)
			selector := rendering.Selector{
				Parts: []rendering.Identifiable{
					rendering.Field{Name: "a", Type: basicInt},
					rendering.Field{Name: "b", Type: basicInt},
					rendering.Field{Name: "c", Type: basicInt},
					rendering.Field{Name: "d", Type: basicInt},
				},
			}

			Expect(selector.Expr()).To(Equal(&ast.SelectorExpr{
				X: &ast.SelectorExpr{
					X: &ast.SelectorExpr{
						X:   ast.NewIdent("a"),
						Sel: ast.NewIdent("b"),
					},
					Sel: ast.NewIdent("c"),
				},
				Sel: ast.NewIdent("d"),
			}))
		})
	})
})
