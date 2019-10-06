package parsing_test

import (
	"go/types"

	"github.com/ryanmoran/faux/parsing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Signature", func() {
	var (
		pkg *types.Package
		sig *types.Signature
	)

	BeforeEach(func() {
		pkg = types.NewPackage("some/package", "package")
		sig = types.NewSignature(nil, nil, nil, false)
	})

	It("parses a signature from a func", func() {
		method := types.NewFunc(0, pkg, "SomeMethod", sig)

		signature := parsing.NewSignature(nil, method)
		Expect(signature).To(Equal(parsing.Signature{
			Name: "SomeMethod",
		}))
	})

	Context("when the signature has params", func() {
		var (
			intType, boolType types.Type
			params, results   *types.Tuple
		)

		BeforeEach(func() {
			intType = types.Universe.Lookup("int").Type()
			boolType = types.Universe.Lookup("bool").Type()

			params = types.NewTuple(types.NewParam(0, pkg, "arg1", types.NewSlice(intType)))
			results = types.NewTuple(types.NewParam(0, pkg, "result1", boolType))
			sig = types.NewSignature(nil, params, results, false)
		})

		It("parses those params", func() {
			method := types.NewFunc(0, pkg, "SomeMethod", sig)

			signature := parsing.NewSignature(nil, method)
			Expect(signature).To(Equal(parsing.Signature{
				Name: "SomeMethod",
				Params: []parsing.Argument{
					{
						Name: "arg1",
						Type: types.NewSlice(intType),
					},
				},
				Results: []parsing.Argument{
					{
						Name: "result1",
						Type: boolType,
					},
				},
			}))
		})

		Context("when the signature is variadic", func() {
			BeforeEach(func() {
				sig = types.NewSignature(nil, params, results, true)
			})

			It("marks the last param as variadic", func() {
				method := types.NewFunc(0, pkg, "SomeMethod", sig)

				signature := parsing.NewSignature(nil, method)
				Expect(signature).To(Equal(parsing.Signature{
					Name: "SomeMethod",
					Params: []parsing.Argument{
						{
							Name:     "arg1",
							Type:     types.NewSlice(intType),
							Variadic: true,
						},
					},
					Results: []parsing.Argument{
						{
							Name: "result1",
							Type: boolType,
						},
					},
				}))
			})
		})
	})
})
