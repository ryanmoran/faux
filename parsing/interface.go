package parsing

import (
	"fmt"
	"go/types"
)

type Interface struct {
	Name       string
	Signatures []Signature
}

func NewInterface(pkgMap map[string]string, n *types.Named) (Interface, error) {
	var signatures []Signature

	underlying, ok := n.Underlying().(*types.Interface)
	if !ok {
		return Interface{}, fmt.Errorf("failed to load underlying type: %s is not an interface", n.Underlying())
	}

	for i := 0; i < underlying.NumMethods(); i++ {
		signatures = append(signatures, NewSignature(pkgMap, underlying.Method(i)))
	}

	return Interface{
		Name:       n.Obj().Name(),
		Signatures: signatures,
	}, nil
}
