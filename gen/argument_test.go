package gen_test

import (
	"go/ast"

	"github.com/ryanmoran/faux/gen"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Argument", func() {
	Describe("NewArgument", func() {
		It("creates an argument", func() {
			argument, err := gen.NewArgument(&ast.Field{
				Names: []*ast.Ident{ast.NewIdent("SomeName")},
				Type:  ast.NewIdent("SomeType"),
			}, "SomeMethod")
			Expect(err).NotTo(HaveOccurred())
			Expect(argument).To(Equal(gen.Argument{
				Method: "SomeMethod",
				Name:   "SomeName",
				Type:   "SomeType",
			}))
		})

		Context("when the field is invalid", func() {
			It("returns an error", func() {
				_, err := gen.NewArgument(&ast.Field{}, "SomeMethod")
				Expect(err).To(MatchError(ContainSubstring("could not format argument field type")))
			})
		})
	})
})
