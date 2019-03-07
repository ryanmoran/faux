package gen_test

import (
	"github.com/ryanmoran/faux/gen"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParsePackage", func() {
	Context("when the package is in the standard library", func() {
		It("parses the given package, returning a fake matching the given named interface", func() {
			fake, err := gen.ParsePackage("io", "Reader")
			Expect(err).NotTo(HaveOccurred())
			Expect(fake).To(Equal(gen.Fake{
				Name: "Reader",
				Methods: []gen.Method{
					{
						Name:     "Read",
						Receiver: "Reader",
						Params: []gen.Argument{
							{
								Method: "Read",
								Name:   "p",
								Type:   "[]byte",
							},
						},
						Results: []gen.Argument{
							{
								Method: "Read",
								Name:   "n",
								Type:   "int",
							},
							{
								Method: "Read",
								Name:   "err",
								Type:   "error",
							},
						},
					},
				},
			}))
		})
	})

	Context("when the package is in the GOPATH", func() {
		It("parses the given package, returning a fake matching the given named interface", func() {
			fake, err := gen.ParsePackage("github.com/pivotal-cf/jhanda", "Command")
			Expect(err).NotTo(HaveOccurred())
			Expect(fake).To(Equal(gen.Fake{
				Name: "Command",
				Methods: []gen.Method{
					{
						Name:     "Execute",
						Receiver: "Command",
						Params: []gen.Argument{
							{
								Method: "Execute",
								Name:   "args",
								Type:   "[]string",
							},
						},
						Results: []gen.Argument{
							{
								Method: "Execute",
								Name:   "error",
								Type:   "error",
							},
						},
					},
					{
						Name:     "Usage",
						Receiver: "Command",
						Results: []gen.Argument{
							{
								Method: "Usage",
								Name:   "usage",
								Type:   "jhanda.Usage",
							},
						},
					},
				},
			}))
		})
	})
})
