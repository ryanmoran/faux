package gen_test

import (
	"go/ast"
	"go/types"

	"github.com/ryanmoran/faux/gen"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type FakeType struct {
	UnderlyingCall struct {
		Returns struct {
			Type types.Type
		}
	}

	StringCall struct {
		Returns struct {
			String string
		}
	}
}

func (f *FakeType) Underlying() types.Type {
	return f.UnderlyingCall.Returns.Type
}

func (f *FakeType) String() string {
	return f.StringCall.Returns.String
}

var _ = Describe("Argument", func() {
	Describe("NewArgument", func() {
		It("creates an argument", func() {
			fakeType := &FakeType{}
			fakeType.StringCall.Returns.String = "SomeType"

			Expect(gen.NewArgument(&ast.Field{
				Names: []*ast.Ident{ast.NewIdent("SomeName")},
				Type:  ast.NewIdent("SomeType"),
			}, "SomeMethod", fakeType, "fallbackName")).To(Equal(gen.Argument{
				Method: "SomeMethod",
				Name:   "someName",
				Type:   "SomeType",
			}))
		})
	})
})
