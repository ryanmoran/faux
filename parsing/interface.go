package parsing

import (
	"fmt"
	"go/types"
)

type Interface struct {
	Name       string
	TypeArgs   []*types.TypeParam
	Signatures []Signature
}

func NewInterface(n *types.Named) (Interface, error) {
	var signatures []Signature

	var targs []*types.TypeParam
	for i := 0; i < n.TypeParams().Len(); i++ {
		targs = append(targs, n.TypeParams().At(i))
	}

	underlying, ok := n.Underlying().(*types.Interface)
	if !ok {
		return Interface{}, fmt.Errorf("failed to load underlying type: %s is not an interface", n.Underlying())
	}

	for i := 0; i < underlying.NumMethods(); i++ {
		signatures = append(signatures, NewSignature(underlying.Method(i)))
	}

	return Interface{
		Name:       n.Obj().Name(),
		TypeArgs:   targs,
		Signatures: signatures,
	}, nil
}
