package rendering_test

import (
	"go/types"

	"github.com/ryanmoran/faux/parsing"
	"github.com/ryanmoran/faux/rendering"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Build", func() {
	It("builds a file representation", func() {
		file := rendering.Build(parsing.Interface{
			Name: "Reader",
			Signatures: []parsing.Signature{
				{
					Name: "Read",
					Params: []parsing.Argument{
						{
							Name:     "p",
							Type:     types.NewSlice(types.Universe.Lookup("byte").Type()),
							Variadic: true,
						},
					},
					Results: []parsing.Argument{
						{
							Name: "n",
							Type: types.Universe.Lookup("int").Type(),
						},
						{
							Name: "err",
							Type: types.Universe.Lookup("error").Type(),
						},
					},
				},
			},
		})

		Expect(file.Package).To(Equal("fakes"))
		Expect(file.Types).To(HaveLen(1))
		Expect(file.Funcs).To(HaveLen(1))

		fakeType := file.Types[0]
		Expect(fakeType.Name).To(Equal("Reader"))

		fakeStruct := fakeType.Type.(rendering.Struct)
		Expect(fakeStruct.Fields).To(HaveLen(1))

		By("checking the ReadCall field", func() {
			field := fakeStruct.Fields[0]
			Expect(field.Name).To(Equal("ReadCall"))

			fieldStruct := field.Type.(rendering.Struct)
			Expect(fieldStruct.Fields).To(HaveLen(5))

			By("checking the ReadCall.Mutex field", func() {
				field := fieldStruct.Fields[0]
				Expect(field.Name).To(Equal(""))
				Expect(field.Type).To(Equal(rendering.NamedType{
					Name: "sync.Mutex",
					Type: rendering.Struct{},
				}))
			})

			By("checking the ReadCall.CallCount field", func() {
				field := fieldStruct.Fields[1]
				Expect(field.Name).To(Equal("CallCount"))
				Expect(field.Type).To(Equal(rendering.BasicType{
					Underlying: rendering.BasicInt,
				}))
			})

			By("checking the ReadCall.Receives field", func() {
				field := fieldStruct.Fields[2]
				Expect(field.Name).To(Equal("Receives"))

				fieldStruct := field.Type.(rendering.Struct)
				Expect(fieldStruct.Fields).To(HaveLen(1))

				By("checking the ReadCall.Receives.P field", func() {
					field := fieldStruct.Fields[0]
					Expect(field.Name).To(Equal("P"))
					Expect(field.Type).To(Equal(rendering.Slice{
						Elem: rendering.BasicType{
							Underlying: rendering.BasicByte,
						},
					}))
				})
			})

			By("checking the ReadCall.Returns field", func() {
				field := fieldStruct.Fields[3]
				Expect(field.Name).To(Equal("Returns"))

				fieldStruct := field.Type.(rendering.Struct)
				Expect(fieldStruct.Fields).To(HaveLen(2))

				By("checking the ReadCall.Returns.N field", func() {
					field := fieldStruct.Fields[0]
					Expect(field.Name).To(Equal("N"))
					Expect(field.Type).To(Equal(rendering.BasicType{
						Underlying: rendering.BasicInt,
					}))
				})

				By("checking the ReadCall.Returns.Err field", func() {
					field := fieldStruct.Fields[1]
					Expect(field.Name).To(Equal("Err"))
					Expect(field.Type).To(Equal(rendering.NamedType{
						Name: "error",
						Type: rendering.Interface{},
					}))
				})
			})

			By("checking the ReadCall.Stub field", func() {
				field := fieldStruct.Fields[4]
				Expect(field.Name).To(Equal("Stub"))

				function := field.Type.(rendering.Func)
				Expect(function.Name).To(Equal(""))
				Expect(function.Receiver).To(Equal(rendering.Receiver{}))
				Expect(function.Params).To(HaveLen(1))
				Expect(function.Results).To(HaveLen(2))

				By("checking the Stub param", func() {
					param := function.Params[0]

					Expect(param.Name).To(Equal(""))
					Expect(param.Type).To(Equal(rendering.Slice{
						Elem: rendering.BasicType{
							Underlying: rendering.BasicByte,
						},
					}))
					Expect(param.Variadic).To(BeTrue())
				})

				By("checking the Stub results", func() {
					result1 := function.Results[0]

					Expect(result1.Type).To(Equal(rendering.BasicType{
						Underlying: rendering.BasicInt,
					}))

					result2 := function.Results[1]
					Expect(result2.Type).To(Equal(rendering.NamedType{
						Name: "error",
						Type: rendering.Interface{},
					}))
				})

				By("checking the Read body", func() {
					Expect(function.Body).To(HaveLen(0))
				})
			})
		})

		By("checking the Read func", func() {
			function := file.Funcs[0]
			Expect(function.Name).To(Equal("Read"))
			Expect(function.Receiver).To(Equal(rendering.Receiver{
				Name: "f",
				Type: rendering.Pointer{
					Elem: fakeType,
				},
			}))
			Expect(function.Params).To(HaveLen(1))
			Expect(function.Results).To(HaveLen(2))

			By("checking the Read.param1 param", func() {
				param := function.Params[0]

				Expect(param.Name).To(Equal("param1"))
				Expect(param.Type).To(Equal(rendering.Slice{
					Elem: rendering.BasicType{
						Underlying: rendering.BasicByte,
					},
				}))
				Expect(param.Variadic).To(BeTrue())
			})

			By("checking the Read.int result", func() {
				result := function.Results[0]

				Expect(result.Type).To(Equal(rendering.BasicType{
					Underlying: rendering.BasicInt,
				}))
			})

			By("checking the Read.error result", func() {
				result := function.Results[1]

				Expect(result.Type).To(Equal(rendering.NamedType{
					Name: "error",
					Type: rendering.Interface{},
				}))
			})

			By("checking the Read body", func() {
				Expect(function.Body).To(HaveLen(6))

				By("checking the ReadCall.Mutex.Lock call", func() {
					statement := function.Body[0].(rendering.CallStatement)
					selector := statement.Call.X.(rendering.Selector)

					Expect(selector.Parts).To(HaveLen(3))
					Expect(selector.Parts[0].(rendering.Receiver).Name).To(Equal("f"))
					Expect(selector.Parts[1].(rendering.Field).Name).To(Equal("ReadCall"))
					Expect(selector.Parts[2].(rendering.Ident).Name).To(Equal("Lock"))
				})

				By("checking the defered ReadCall.Mutex.Unlock call", func() {
					statement := function.Body[1].(rendering.DeferStatement)
					selector := statement.Call.X.(rendering.Selector)

					Expect(selector.Parts).To(HaveLen(3))
					Expect(selector.Parts[0].(rendering.Receiver).Name).To(Equal("f"))
					Expect(selector.Parts[1].(rendering.Field).Name).To(Equal("ReadCall"))
					Expect(selector.Parts[2].(rendering.Ident).Name).To(Equal("Unlock"))
				})

				By("checking the ReadCall.CallCount increment statement", func() {
					statement := function.Body[2].(rendering.IncrementStatement)
					selector := statement.X.(rendering.Selector)

					Expect(selector.Parts).To(HaveLen(3))
					Expect(selector.Parts[0].(rendering.Receiver).Name).To(Equal("f"))
					Expect(selector.Parts[1].(rendering.Field).Name).To(Equal("ReadCall"))
					Expect(selector.Parts[2].(rendering.Field).Name).To(Equal("CallCount"))
				})

				By("checking the ReadCall.Receives.P assign statement", func() {
					statement := function.Body[3].(rendering.AssignStatement)
					selector := statement.Left.(rendering.Selector)
					param := statement.Right.(rendering.Param)

					Expect(selector.Parts).To(HaveLen(4))
					Expect(selector.Parts[0].(rendering.Receiver).Name).To(Equal("f"))
					Expect(selector.Parts[1].(rendering.Field).Name).To(Equal("ReadCall"))
					Expect(selector.Parts[2].(rendering.Field).Name).To(Equal("Receives"))
					Expect(selector.Parts[3].(rendering.Field).Name).To(Equal("P"))

					Expect(param.Name).To(Equal("param1"))
				})

				By("checking the ReadCall stub conditional statement", func() {
					statement := function.Body[4].(rendering.IfStatement)
					condition := statement.Condition.(rendering.Equality)
					Expect(condition.Equal).To(BeFalse())

					left := condition.Left.(rendering.Selector)
					Expect(left.Parts).To(HaveLen(3))
					Expect(left.Parts[0].(rendering.Receiver).Name).To(Equal("f"))
					Expect(left.Parts[1].(rendering.Field).Name).To(Equal("ReadCall"))
					Expect(left.Parts[2].(rendering.Field).Name).To(Equal("Stub"))

					Expect(condition.Right).To(Equal(rendering.Nil{}))

					By("checking the ReadCall stub conditional statement body", func() {
						Expect(statement.Body).To(HaveLen(1))

						returnStatement := statement.Body[0].(rendering.ReturnStatement)
						Expect(returnStatement.Results).To(HaveLen(1))

						call := returnStatement.Results[0].(rendering.Call)
						selector := call.X.(rendering.Selector)
						Expect(selector.Parts).To(HaveLen(3))
						Expect(selector.Parts[0].(rendering.Receiver).Name).To(Equal("f"))
						Expect(selector.Parts[1].(rendering.Field).Name).To(Equal("ReadCall"))
						Expect(selector.Parts[2].(rendering.Field).Name).To(Equal("Stub"))

						Expect(call.Params).To(HaveLen(1))
						Expect(call.Params[0].Name).To(Equal("param1"))
					})
				})

				By("checking the ReadCall return statement", func() {
					statement := function.Body[5].(rendering.ReturnStatement)

					Expect(statement.Results).To(HaveLen(2))

					By("checking the ReadCall return statement N result", func() {
						selector := statement.Results[0].(rendering.Selector)

						Expect(selector.Parts).To(HaveLen(4))
						Expect(selector.Parts[0].(rendering.Receiver).Name).To(Equal("f"))
						Expect(selector.Parts[1].(rendering.Field).Name).To(Equal("ReadCall"))
						Expect(selector.Parts[2].(rendering.Field).Name).To(Equal("Returns"))
						Expect(selector.Parts[3].(rendering.Field).Name).To(Equal("N"))
					})

					By("checking the ReadCall return statement err result", func() {
						selector := statement.Results[1].(rendering.Selector)

						Expect(selector.Parts).To(HaveLen(4))
						Expect(selector.Parts[0].(rendering.Receiver).Name).To(Equal("f"))
						Expect(selector.Parts[1].(rendering.Field).Name).To(Equal("ReadCall"))
						Expect(selector.Parts[2].(rendering.Field).Name).To(Equal("Returns"))
						Expect(selector.Parts[3].(rendering.Field).Name).To(Equal("Err"))
					})
				})
			})
		})
	})

	Context("when the arguments are unnamed", func() {
		It("builds a file representation", func() {
			file := rendering.Build(parsing.Interface{
				Name: "unnamed",
				Signatures: []parsing.Signature{
					{
						Name: "Method",
						Params: []parsing.Argument{
							{
								Type: types.NewSlice(types.Universe.Lookup("byte").Type()),
							},
						},
						Results: []parsing.Argument{
							{
								Type: types.Universe.Lookup("int").Type(),
							},
						},
					},
				},
			})

			Expect(file.Package).To(Equal("fakes"))
			Expect(file.Types).To(HaveLen(1))
			Expect(file.Funcs).To(HaveLen(1))

			fakeType := file.Types[0]
			Expect(fakeType.Name).To(Equal("Unnamed"))

			fakeStruct := fakeType.Type.(rendering.Struct)
			Expect(fakeStruct.Fields).To(HaveLen(1))

			By("checking the MethodCall field", func() {
				field := fakeStruct.Fields[0]
				Expect(field.Name).To(Equal("MethodCall"))

				fieldStruct := field.Type.(rendering.Struct)
				Expect(fieldStruct.Fields).To(HaveLen(5))

				By("checking the MethodCall.Mutex field", func() {
					field := fieldStruct.Fields[0]
					Expect(field.Name).To(Equal(""))
					Expect(field.Type).To(Equal(rendering.NamedType{
						Name: "sync.Mutex",
						Type: rendering.Struct{},
					}))
				})

				By("checking the MethodCall.CallCount field", func() {
					field := fieldStruct.Fields[1]
					Expect(field.Name).To(Equal("CallCount"))
					Expect(field.Type).To(Equal(rendering.BasicType{
						Underlying: rendering.BasicInt,
					}))
				})

				By("checking the MethodCall.Receives field", func() {
					field := fieldStruct.Fields[2]
					Expect(field.Name).To(Equal("Receives"))

					fieldStruct := field.Type.(rendering.Struct)
					Expect(fieldStruct.Fields).To(HaveLen(1))

					By("checking the MethodCall.Receives.ByteSlice field", func() {
						field := fieldStruct.Fields[0]
						Expect(field.Name).To(Equal("ByteSlice"))
						Expect(field.Type).To(Equal(rendering.Slice{
							Elem: rendering.BasicType{
								Underlying: rendering.BasicByte,
							},
						}))
					})
				})

				By("checking the MethodCall.Returns field", func() {
					field := fieldStruct.Fields[3]
					Expect(field.Name).To(Equal("Returns"))

					fieldStruct := field.Type.(rendering.Struct)
					Expect(fieldStruct.Fields).To(HaveLen(1))

					By("checking the MethodCall.Returns.Int field", func() {
						field := fieldStruct.Fields[0]
						Expect(field.Name).To(Equal("Int"))
						Expect(field.Type).To(Equal(rendering.BasicType{
							Underlying: rendering.BasicInt,
						}))
					})
				})

				By("checking the MethodCall.Stub field", func() {
					field := fieldStruct.Fields[4]
					Expect(field.Name).To(Equal("Stub"))

					function := field.Type.(rendering.Func)
					Expect(function.Name).To(Equal(""))
					Expect(function.Receiver).To(Equal(rendering.Receiver{}))
					Expect(function.Params).To(HaveLen(1))
					Expect(function.Results).To(HaveLen(1))

					By("checking the Stub param", func() {
						param := function.Params[0]

						Expect(param.Name).To(Equal(""))
						Expect(param.Type).To(Equal(rendering.Slice{
							Elem: rendering.BasicType{
								Underlying: rendering.BasicByte,
							},
						}))
					})

					By("checking the Stub results", func() {
						result := function.Results[0]

						Expect(result.Type).To(Equal(rendering.BasicType{
							Underlying: rendering.BasicInt,
						}))
					})

					By("checking the Stub body", func() {
						Expect(function.Body).To(HaveLen(0))
					})
				})
			})

			By("checking the Method func", func() {
				function := file.Funcs[0]
				Expect(function.Name).To(Equal("Method"))
				Expect(function.Receiver).To(Equal(rendering.Receiver{
					Name: "f",
					Type: rendering.Pointer{
						Elem: fakeType,
					},
				}))
				Expect(function.Params).To(HaveLen(1))
				Expect(function.Results).To(HaveLen(1))

				By("checking the Method.param1 param", func() {
					param := function.Params[0]

					Expect(param.Name).To(Equal("param1"))
					Expect(param.Type).To(Equal(rendering.Slice{
						Elem: rendering.BasicType{
							Underlying: rendering.BasicByte,
						},
					}))
				})

				By("checking the Method.int result", func() {
					result := function.Results[0]

					Expect(result.Type).To(Equal(rendering.BasicType{
						Underlying: rendering.BasicInt,
					}))
				})

				By("checking the Method body", func() {
					Expect(function.Body).To(HaveLen(6))

					By("checking the MethodCall.Mutex.Lock call", func() {
						statement := function.Body[0].(rendering.CallStatement)
						selector := statement.Call.X.(rendering.Selector)

						Expect(selector.Parts).To(HaveLen(3))
						Expect(selector.Parts[0].(rendering.Receiver).Name).To(Equal("f"))
						Expect(selector.Parts[1].(rendering.Field).Name).To(Equal("MethodCall"))
						Expect(selector.Parts[2].(rendering.Ident).Name).To(Equal("Lock"))
					})

					By("checking the defered MethodCall.Mutex.Unlock call", func() {
						statement := function.Body[1].(rendering.DeferStatement)
						selector := statement.Call.X.(rendering.Selector)

						Expect(selector.Parts).To(HaveLen(3))
						Expect(selector.Parts[0].(rendering.Receiver).Name).To(Equal("f"))
						Expect(selector.Parts[1].(rendering.Field).Name).To(Equal("MethodCall"))
						Expect(selector.Parts[2].(rendering.Ident).Name).To(Equal("Unlock"))
					})

					By("checking the MethodCall.CallCount increment statement", func() {
						statement := function.Body[2].(rendering.IncrementStatement)
						selector := statement.X.(rendering.Selector)

						Expect(selector.Parts).To(HaveLen(3))
						Expect(selector.Parts[0].(rendering.Receiver).Name).To(Equal("f"))
						Expect(selector.Parts[1].(rendering.Field).Name).To(Equal("MethodCall"))
						Expect(selector.Parts[2].(rendering.Field).Name).To(Equal("CallCount"))
					})

					By("checking the MethodCall.Receives.ByteSlice assign statement", func() {
						statement := function.Body[3].(rendering.AssignStatement)
						selector := statement.Left.(rendering.Selector)
						param := statement.Right.(rendering.Param)

						Expect(selector.Parts).To(HaveLen(4))
						Expect(selector.Parts[0].(rendering.Receiver).Name).To(Equal("f"))
						Expect(selector.Parts[1].(rendering.Field).Name).To(Equal("MethodCall"))
						Expect(selector.Parts[2].(rendering.Field).Name).To(Equal("Receives"))
						Expect(selector.Parts[3].(rendering.Field).Name).To(Equal("ByteSlice"))

						Expect(param.Name).To(Equal("param1"))
					})

					By("checking the MethodCall stub conditional statement", func() {
						statement := function.Body[4].(rendering.IfStatement)
						condition := statement.Condition.(rendering.Equality)
						Expect(condition.Equal).To(BeFalse())

						left := condition.Left.(rendering.Selector)
						Expect(left.Parts).To(HaveLen(3))
						Expect(left.Parts[0].(rendering.Receiver).Name).To(Equal("f"))
						Expect(left.Parts[1].(rendering.Field).Name).To(Equal("MethodCall"))
						Expect(left.Parts[2].(rendering.Field).Name).To(Equal("Stub"))

						Expect(condition.Right).To(Equal(rendering.Nil{}))

						By("checking the ReadCall stub conditional statement body", func() {
							Expect(statement.Body).To(HaveLen(1))

							returnStatement := statement.Body[0].(rendering.ReturnStatement)
							Expect(returnStatement.Results).To(HaveLen(1))

							result := returnStatement.Results[0].(rendering.Call)
							selector := result.X.(rendering.Selector)
							Expect(selector.Parts).To(HaveLen(3))
							Expect(selector.Parts[0].(rendering.Receiver).Name).To(Equal("f"))
							Expect(selector.Parts[1].(rendering.Field).Name).To(Equal("MethodCall"))
							Expect(selector.Parts[2].(rendering.Field).Name).To(Equal("Stub"))
						})
					})

					By("checking the MethodCall return statement", func() {
						statement := function.Body[5].(rendering.ReturnStatement)

						Expect(statement.Results).To(HaveLen(1))

						By("checking the MethodCall return statement Int result", func() {
							selector := statement.Results[0].(rendering.Selector)

							Expect(selector.Parts).To(HaveLen(4))
							Expect(selector.Parts[0].(rendering.Receiver).Name).To(Equal("f"))
							Expect(selector.Parts[1].(rendering.Field).Name).To(Equal("MethodCall"))
							Expect(selector.Parts[2].(rendering.Field).Name).To(Equal("Returns"))
							Expect(selector.Parts[3].(rendering.Field).Name).To(Equal("Int"))
						})
					})
				})
			})
		})
	})
})
