package rendering

import (
	"go/ast"
	"go/types"
)

var (
	BasicBool       = BasicLookup("bool")
	BasicInt        = BasicLookup("int")
	BasicInt8       = BasicLookup("int8")
	BasicInt16      = BasicLookup("int16")
	BasicInt32      = BasicLookup("int32")
	BasicInt64      = BasicLookup("int64")
	BasicUint       = BasicLookup("uint")
	BasicUint8      = BasicLookup("uint8")
	BasicUint16     = BasicLookup("uint16")
	BasicUint32     = BasicLookup("uint32")
	BasicUint64     = BasicLookup("uint64")
	BasicUintptr    = BasicLookup("uintptr")
	BasicFloat32    = BasicLookup("float32")
	BasicFloat64    = BasicLookup("float64")
	BasicComplex64  = BasicLookup("complex64")
	BasicComplex128 = BasicLookup("complex128")
	BasicString     = BasicLookup("string")
	BasicByte       = BasicLookup("byte")
	BasicRune       = BasicLookup("rune")
)

func BasicLookup(name string) types.Type {
	return types.Universe.Lookup(name).Type()
}

type BasicType struct {
	Underlying types.Type
}

func NewBasicType(t types.Type) Type {
	return BasicType{
		Underlying: t,
	}
}

func (bt BasicType) Expr() ast.Expr {
	return ast.NewIdent(bt.Underlying.String())
}

func (bt BasicType) isType() {}
