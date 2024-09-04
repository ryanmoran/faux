package parsing

import "go/types"

type TypeParam struct {
	Name       string
	Constraint types.Type
}
