package gen_test

import (
	"go/ast"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ryanmoran/faux/gen"
)

var _ = Describe("Method", func() {
	Describe("FieldStruct", func() {
		It("returns a representation of the fake field struct for this method", func() {
			method := gen.Method{
				Name:     "SomeMethod",
				Receiver: "SomeReceiver",
				Params: []gen.Argument{
					{
						Method: "SomeMethod",
						Name:   "someParam",
						Type:   "SomeType",
					},
				},
				Results: []gen.Argument{
					{
						Method: "SomeMethod",
						Name:   "someResult",
						Type:   "SomeType",
					},
				},
			}

			Expect(method.FieldStruct()).To(Equal(&ast.Field{
				Names: []*ast.Ident{ast.NewIdent("SomeMethodCall")},
				Type: &ast.StructType{
					Fields: &ast.FieldList{
						List: []*ast.Field{
							&ast.Field{
								Names: []*ast.Ident{ast.NewIdent("CallCount")},
								Type:  ast.NewIdent("int"),
							},
							&ast.Field{
								Names: []*ast.Ident{
									ast.NewIdent("Receives"),
								},
								Type: &ast.StructType{
									Fields: &ast.FieldList{
										List: []*ast.Field{
											&ast.Field{
												Names: []*ast.Ident{ast.NewIdent("SomeParam")},
												Type:  ast.NewIdent("SomeType"),
											},
										},
									},
								},
							},
							&ast.Field{
								Names: []*ast.Ident{
									ast.NewIdent("Returns"),
								},
								Type: &ast.StructType{
									Fields: &ast.FieldList{
										List: []*ast.Field{
											&ast.Field{
												Names: []*ast.Ident{ast.NewIdent("SomeResult")},
												Type:  ast.NewIdent("SomeType"),
											},
										},
									},
								},
							},
						},
					},
				},
			}))
		})

		Context("when there are no params or results", func() {
			It("does not include those fields", func() {
				method := gen.Method{
					Name:     "SomeMethod",
					Receiver: "SomeReceiver",
				}

				Expect(method.FieldStruct()).To(Equal(&ast.Field{
					Names: []*ast.Ident{ast.NewIdent("SomeMethodCall")},
					Type: &ast.StructType{
						Fields: &ast.FieldList{
							List: []*ast.Field{
								&ast.Field{
									Names: []*ast.Ident{ast.NewIdent("CallCount")},
									Type:  ast.NewIdent("int"),
								},
							},
						},
					},
				}))
			})
		})
	})
})
