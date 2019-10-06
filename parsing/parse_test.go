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
		iface, pkgMap, err := parsing.Parse("io", "Reader")
		Expect(err).NotTo(HaveOccurred())
		Expect(iface).To(Equal(parsing.Interface{
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
		}))
		Expect(pkgMap).To(HaveLen(15))
		Expect(pkgMap).To(HaveKey("bytes"))
		Expect(pkgMap).To(HaveKey("crypto/sha1"))
		Expect(pkgMap).To(HaveKey("errors"))
		Expect(pkgMap).To(HaveKey("fmt"))
		Expect(pkgMap).To(HaveKey("io"))
		Expect(pkgMap).To(HaveKey("io/ioutil"))
		Expect(pkgMap).To(HaveKey("log"))
		Expect(pkgMap).To(HaveKey("os"))
		Expect(pkgMap).To(HaveKey("runtime"))
		Expect(pkgMap).To(HaveKey("sort"))
		Expect(pkgMap).To(HaveKey("strings"))
		Expect(pkgMap).To(HaveKey("sync"))
		Expect(pkgMap).To(HaveKey("sync/atomic"))
		Expect(pkgMap).To(HaveKey("testing"))
		Expect(pkgMap).To(HaveKey("time"))
	})

	Context("when there is an embedded interface", func() {
		It("parses the package and returns a parsed interface", func() {
			iface, pkgMap, err := parsing.Parse("io", "ReadCloser")
			Expect(err).NotTo(HaveOccurred())
			Expect(iface).To(Equal(parsing.Interface{
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
			}))
			Expect(pkgMap).To(HaveLen(15))
			Expect(pkgMap).To(HaveKey("bytes"))
			Expect(pkgMap).To(HaveKey("crypto/sha1"))
			Expect(pkgMap).To(HaveKey("errors"))
			Expect(pkgMap).To(HaveKey("fmt"))
			Expect(pkgMap).To(HaveKey("io"))
			Expect(pkgMap).To(HaveKey("io/ioutil"))
			Expect(pkgMap).To(HaveKey("log"))
			Expect(pkgMap).To(HaveKey("os"))
			Expect(pkgMap).To(HaveKey("runtime"))
			Expect(pkgMap).To(HaveKey("sort"))
			Expect(pkgMap).To(HaveKey("strings"))
			Expect(pkgMap).To(HaveKey("sync"))
			Expect(pkgMap).To(HaveKey("sync/atomic"))
			Expect(pkgMap).To(HaveKey("testing"))
			Expect(pkgMap).To(HaveKey("time"))
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
				_, _, err := parsing.Parse("io", "Reader")
				Expect(err).To(MatchError(ContainSubstring("executable file not found in $PATH")))
			})
		})

		Context("when the name matches no object in scope", func() {
			It("returns an error", func() {
				_, _, err := parsing.Parse("io", "Banana")
				Expect(err).To(MatchError("failed to find named type: io.Banana"))
			})
		})

		Context("when the package has no files", func() {
			It("returns an error", func() {
				_, _, err := parsing.Parse("some-package", "SomeType")
				Expect(err).To(MatchError("failed to load package with any files: \"some-package\""))
			})
		})
	})
})
