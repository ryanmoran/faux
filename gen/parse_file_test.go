package gen_test

import (
	"strings"

	"github.com/ryanmoran/faux/gen"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParseFile", func() {
	It("parses the given file, returning a fake matching the given named interface", func() {
		source := strings.NewReader(`package main

import (
  "bytes"
	"io"
)

type SomeInterface interface{
	SomeMethod(someParam1 string, someParam2 *bytes.Buffer) (someResult1 int, someResult2 io.Reader)
}
`)

		fake, err := gen.ParseFile("some-file.go", source, "SomeInterface")
		Expect(err).NotTo(HaveOccurred())
		Expect(fake).To(Equal(gen.Fake{
			Name: "SomeInterface",
			Methods: []gen.Method{
				{
					Name:     "SomeMethod",
					Receiver: "SomeInterface",
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

	Context("failure cases", func() {
		Context("when the source file cannot be parsed", func() {
			It("returns an error", func() {
				source := strings.NewReader("%%%")
				_, err := gen.ParseFile("some-file.go", source, "SomeInterface")
				Expect(err).To(MatchError("could not parse source: some-file.go:1:1: expected 'package', found '%'"))
			})
		})

		Context("when the given interface name cannot be found", func() {
			It("returns an error", func() {
				source := strings.NewReader("package main")
				_, err := gen.ParseFile("some-file.go", source, "SomeInterface")
				Expect(err).To(MatchError("could not find interface \"SomeInterface\" in some-file.go"))
			})
		})
	})
})
