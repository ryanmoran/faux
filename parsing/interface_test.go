package parsing_test

import (
	"go/types"

	"github.com/ryanmoran/faux/parsing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Interface", func() {
	var (
		pkg        *types.Package
		typeName   *types.TypeName
		underlying *types.Interface
		namedType  *types.Named
	)

	BeforeEach(func() {
		pkg = types.NewPackage("some/package", "package")
		typeName = types.NewTypeName(0, pkg, "SomeType", nil)
		underlying = types.NewInterfaceType(nil, nil).Complete()
		namedType = types.NewNamed(typeName, underlying, nil)
	})

	It("parses an interface from a named type", func() {
		iface, err := parsing.NewInterface(nil, namedType)
		Expect(err).NotTo(HaveOccurred())
		Expect(iface).To(Equal(parsing.Interface{
			Name: "SomeType",
		}))
	})

	Context("when the interface has methods", func() {
		BeforeEach(func() {
			signature := types.NewSignature(nil, nil, nil, false)
			methods := []*types.Func{
				types.NewFunc(0, pkg, "SomeMethod", signature),
			}

			underlying = types.NewInterfaceType(methods, nil).Complete()
			namedType = types.NewNamed(typeName, underlying, nil)
		})

		It("includes those methods in the parsed interface", func() {
			iface, err := parsing.NewInterface(nil, namedType)
			Expect(err).NotTo(HaveOccurred())
			Expect(iface).To(Equal(parsing.Interface{
				Name: "SomeType",
				Signatures: []parsing.Signature{
					{
						Name: "SomeMethod",
					},
				},
			}))
		})
	})

	Context("when the underlying type is not interface", func() {
		BeforeEach(func() {
			intType := types.Universe.Lookup("int").Type()
			namedType = types.NewNamed(typeName, intType, nil)
		})

		It("returns an error", func() {
			_, err := parsing.NewInterface(nil, namedType)
			Expect(err).To(MatchError("failed to load underlying type: int is not an interface"))
		})
	})
})
