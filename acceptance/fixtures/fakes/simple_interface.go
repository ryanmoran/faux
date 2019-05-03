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

func (f *SimpleInterface) SomeMethod(param1 *bytes.Buffer) io.Reader {
	f.SomeMethodCall.CallCount++
	f.SomeMethodCall.Receives.SomeParam = param1
	return f.SomeMethodCall.Returns.SomeResult
}
func (f *SimpleInterface) VariadicMethod(param1 ...int) {
	f.VariadicMethodCall.CallCount++
	f.VariadicMethodCall.Receives.SomeParams = param1
}
