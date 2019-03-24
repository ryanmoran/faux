package gen_test

import (
	"github.com/ryanmoran/faux/gen"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParseFile", func() {
	It("parses the given file, returning a fake matching the given named interface", func() {
		fake, err := gen.ParseFile("./fixtures/interfaces.go", "FullInterface")
		Expect(err).NotTo(HaveOccurred())
		Expect(fake).To(Equal(gen.Fake{
			Name: "FullInterface",
			Methods: []gen.Method{
				{
					Name:     "SomeMethod",
					Receiver: "FullInterface",
					Params: []gen.Argument{
						{
							Method: "SomeMethod",
							Name:   "someParam1",
							Type:   "string",
						},
						{
							Method: "SomeMethod",
							Name:   "someParam2",
							Type:   "*bytes.Buffer",
						},
					},
					Results: []gen.Argument{
						{
							Method: "SomeMethod",
							Name:   "someResult1",
							Type:   "int",
						},
						{
							Method: "SomeMethod",
							Name:   "someResult2",
							Type:   "io.Reader",
						},
					},
				},
			},
		}))
	})

	Context("when the fields are unnamed", func() {
		It("parses the given file, returning a fake matching the given named interface", func() {
			fake, err := gen.ParseFile("./fixtures/interfaces.go", "UnnamedFieldsInterface")
			Expect(err).NotTo(HaveOccurred())
			Expect(fake).To(Equal(gen.Fake{
				Name: "UnnamedFieldsInterface",
				Methods: []gen.Method{
					{
						Name:     "SomeMethod",
						Receiver: "UnnamedFieldsInterface",
						Params: []gen.Argument{
							{
								Method: "SomeMethod",
								Name:   "string",
								Type:   "string",
							},
							{
								Method: "SomeMethod",
								Name:   "bytesBuffer",
								Type:   "*bytes.Buffer",
							},
						},
						Results: []gen.Argument{
							{
								Method: "SomeMethod",
								Name:   "int",
								Type:   "int",
							},
							{
								Method: "SomeMethod",
								Name:   "ioReader",
								Type:   "io.Reader",
							},
						},
					},
				},
			}))
		})
	})

	Context("when the types are elided", func() {
		It("parses the given file, returning a fake matching the given named interface", func() {
			fake, err := gen.ParseFile("./fixtures/interfaces.go", "ElidedTypesInterface")
			Expect(err).NotTo(HaveOccurred())
			Expect(fake).To(Equal(gen.Fake{
				Name: "ElidedTypesInterface",
				Methods: []gen.Method{
					{
						Name:     "SomeMethod",
						Receiver: "ElidedTypesInterface",
						Params: []gen.Argument{
							{
								Method: "SomeMethod",
								Name:   "someParam1",
								Type:   "string",
							},
							{
								Method: "SomeMethod",
								Name:   "someParam2",
								Type:   "string",
							},
						},
						Results: []gen.Argument{
							{
								Method: "SomeMethod",
								Name:   "int",
								Type:   "int",
							},
							{
								Method: "SomeMethod",
								Name:   "ioReader",
								Type:   "io.Reader",
							},
						},
					},
				},
			}))
		})
	})

	Context("when the interface is unexported", func() {
		It("parses the given file, returning a fake matching the given named interface", func() {
			fake, err := gen.ParseFile("./fixtures/interfaces.go", "unexportedInterface")
			Expect(err).NotTo(HaveOccurred())
			Expect(fake).To(Equal(gen.Fake{
				Name: "UnexportedInterface",
				Methods: []gen.Method{
					{
						Name:     "SomeMethod",
						Receiver: "UnexportedInterface",
					},
				},
			}))
		})
	})

	Context("failure cases", func() {
		Context("when the given interface name cannot be found", func() {
			It("returns an error", func() {
				_, err := gen.ParseFile("./fixtures/interfaces.go", "UndefinedInterface")
				Expect(err).To(MatchError(ContainSubstring("could not find interface \"UndefinedInterface\"")))
			})
		})
	})
})
