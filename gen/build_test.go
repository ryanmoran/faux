package gen_test

import (
	"go/ast"
	"go/token"

	"github.com/ryanmoran/faux/gen"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Build", func() {
	It("builds an AST describing the fake", func() {
		tree := gen.Build(gen.Fake{
			Name: "SomeInterface",
			Methods: []gen.Method{
				{
					Name:     "SomeMethod",
					Receiver: "SomeInterface",
					Params: []gen.Argument{
						{
							Method: "SomeMethod",
							Name:   "someParam",
							Type:   "string",
						},
					},
					Results: []gen.Argument{
						{
							Method: "SomeMethod",
							Name:   "someResult",
							Type:   "string",
						},
					},
				},
			},
		})
		Expect(tree).To(Equal(&ast.File{
			Name: ast.NewIdent("fakes"),
			Decls: []ast.Decl{
				&ast.GenDecl{
					Tok: token.TYPE,
					Specs: []ast.Spec{
						&ast.TypeSpec{
							Name: ast.NewIdent("SomeInterface"),
							Type: &ast.StructType{
								Fields: &ast.FieldList{
									List: []*ast.Field{
										&ast.Field{
											Names: []*ast.Ident{
												ast.NewIdent("SomeMethodCall"),
											},
											Type: &ast.StructType{
												Fields: &ast.FieldList{
													List: []*ast.Field{
														&ast.Field{
															Names: []*ast.Ident{
																ast.NewIdent("CallCount"),
															},
															Type: ast.NewIdent("int"),
														},
														&ast.Field{
															Names: []*ast.Ident{
																ast.NewIdent("Receives"),
															},
															Type: &ast.StructType{
																Fields: &ast.FieldList{
																	List: []*ast.Field{
																		&ast.Field{
																			Names: []*ast.Ident{
																				ast.NewIdent("SomeParam"),
																			},
																			Type: ast.NewIdent("string"),
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
																			Names: []*ast.Ident{
																				ast.NewIdent("SomeResult"),
																			},
																			Type: ast.NewIdent("string"),
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
				&ast.FuncDecl{
					Recv: &ast.FieldList{
						List: []*ast.Field{
							&ast.Field{
								Names: []*ast.Ident{
									ast.NewIdent("f"),
								},
								Type: ast.NewIdent("*SomeInterface"),
							},
						},
					},
					Name: ast.NewIdent("SomeMethod"),
					Type: &ast.FuncType{
						Params: &ast.FieldList{
							List: []*ast.Field{
								&ast.Field{
									Names: []*ast.Ident{
										ast.NewIdent("someParam"),
									},
									Type: ast.NewIdent("string"),
								},
							},
						},
						Results: &ast.FieldList{
							List: []*ast.Field{
								&ast.Field{
									Type: ast.NewIdent("string"),
								},
							},
						},
					},
					Body: &ast.BlockStmt{
						List: []ast.Stmt{
							&ast.IncDecStmt{
								X: &ast.SelectorExpr{
									X: &ast.SelectorExpr{
										X:   ast.NewIdent("f"),
										Sel: ast.NewIdent("SomeMethodCall"),
									},
									Sel: ast.NewIdent("CallCount"),
								},
								Tok: token.INC,
							},
							&ast.AssignStmt{
								Lhs: []ast.Expr{
									&ast.SelectorExpr{
										X: &ast.SelectorExpr{
											X: &ast.SelectorExpr{
												X:   ast.NewIdent("f"),
												Sel: ast.NewIdent("SomeMethodCall"),
											},
											Sel: ast.NewIdent("Receives"),
										},
										Sel: ast.NewIdent("SomeParam"),
									},
								},
								Tok: token.ASSIGN,
								Rhs: []ast.Expr{
									ast.NewIdent("someParam"),
								},
							},
							&ast.ReturnStmt{
								Results: []ast.Expr{
									&ast.SelectorExpr{
										X: &ast.SelectorExpr{
											X: &ast.SelectorExpr{
												X:   ast.NewIdent("f"),
												Sel: ast.NewIdent("SomeMethodCall"),
											},
											Sel: ast.NewIdent("Returns"),
										},
										Sel: ast.NewIdent("SomeResult"),
									},
								},
							},
						},
					},
				},
			},
		}))
	})
})
