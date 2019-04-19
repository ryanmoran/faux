package fakes

import (
	"bytes"
	"io"
)

type SimpleInterface struct {
	SomeMethodCall struct {
		CallCount int
		Receives  struct {
			SomeParam *bytes.Buffer
		}
		Returns struct {
			SomeResult io.Reader
		}
	}
	VariadicMethodCall struct {
		CallCount int
		Receives  struct {
			SomeParams []int
		}
	}
}

func (f *SimpleInterface) SomeMethod(someParam *bytes.Buffer) io.Reader {
	f.SomeMethodCall.CallCount++
	f.SomeMethodCall.Receives.SomeParam = someParam
	return f.SomeMethodCall.Returns.SomeResult
}
func (f *SimpleInterface) VariadicMethod(someParams ...int) {
	f.VariadicMethodCall.CallCount++
	f.VariadicMethodCall.Receives.SomeParams = someParams
}
