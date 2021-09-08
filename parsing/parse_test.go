package parsing_test

import (
	"go/types"
	"os"

	"github.com/ryanmoran/faux/parsing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parse", func() {
	It("parses the package and returns a parsed interface", func() {
		fake, err := parsing.Parse("io", "Reader")
		Expect(err).NotTo(HaveOccurred())
		Expect(fake).To(Equal(parsing.Fake{
			Imports: []parsing.Import{
				{Name: "sync", Path: "sync"},
			},
			Interface: parsing.Interface{
				Name: "Reader",
				Signatures: []parsing.Signature{
					{
						Name: "Read",
						Params: []parsing.Argument{
							{
								Name: "p",
								Type: types.NewSlice(types.Universe.Lookup("byte").Type()),
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
			},
		}))
	})

	Context("when there is an embedded interface", func() {
		It("parses the package and returns a parsed interface", func() {
			fake, err := parsing.Parse("io", "ReadCloser")
			Expect(err).NotTo(HaveOccurred())
			Expect(fake).To(Equal(parsing.Fake{
				Imports: []parsing.Import{
					{Name: "sync", Path: "sync"},
				},
				Interface: parsing.Interface{
					Name: "ReadCloser",
					Signatures: []parsing.Signature{
						{
							Name: "Close",
							Results: []parsing.Argument{
								{
									Name: "",
									Type: types.Universe.Lookup("error").Type(),
								},
							},
						},
						{
							Name: "Read",
							Params: []parsing.Argument{
								{
									Name: "p",
									Type: types.NewSlice(types.Universe.Lookup("byte").Type()),
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
				},
			}))
		})
	})

	Context("failure cases", func() {
		Context("when the package loader errors", func() {
			BeforeEach(func() {
				os.Setenv("GOPACKAGESDRIVER", "garbage") // overriding this causes packages.Load to error
			})

			AfterEach(func() {
				os.Unsetenv("GOPACKAGESDRIVER")
			})

			It("returns an error", func() {
				_, err := parsing.Parse("io", "Reader")
				Expect(err).To(MatchError(ContainSubstring("executable file not found in $PATH")))
			})
		})

		Context("when the name matches no object in scope", func() {
			It("returns an error", func() {
				_, err := parsing.Parse("some-package", "SomeType")
				Expect(err).To(MatchError("failed to find named type: some-package.SomeType"))
			})
		})
	})
})
