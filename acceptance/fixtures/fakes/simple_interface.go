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
}

func (f *SimpleInterface) SomeMethod(someParam *bytes.Buffer) io.Reader {
	f.SomeMethodCall.CallCount++
	f.SomeMethodCall.Receives.SomeParam = someParam
	return f.SomeMethodCall.Returns.SomeResult
}
