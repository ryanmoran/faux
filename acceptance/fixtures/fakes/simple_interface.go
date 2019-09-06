package fakes

import (
	"bytes"
	"io"
	"sync"
)

type SimpleInterface struct {
	OtherMethodCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Buffer *bytes.Buffer
		}
		Returns struct {
			Reader io.Reader
			Error  error
		}
		Stub func(*bytes.Buffer) (io.Reader, error)
	}
	SomeMethodCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			SomeParam *bytes.Buffer
		}
		Returns struct {
			SomeResult io.Reader
		}
		Stub func(*bytes.Buffer) io.Reader
	}
}

func (f *SimpleInterface) OtherMethod(param1 *bytes.Buffer) (io.Reader, error) {
	f.OtherMethodCall.Lock()
	defer f.OtherMethodCall.Unlock()
	f.OtherMethodCall.CallCount++
	f.OtherMethodCall.Receives.Buffer = param1
	if f.OtherMethodCall.Stub != nil {
		return f.OtherMethodCall.Stub(param1)
	}
	return f.OtherMethodCall.Returns.Reader, f.OtherMethodCall.Returns.Error
}
func (f *SimpleInterface) SomeMethod(param1 *bytes.Buffer) io.Reader {
	f.SomeMethodCall.Lock()
	defer f.SomeMethodCall.Unlock()
	f.SomeMethodCall.CallCount++
	f.SomeMethodCall.Receives.SomeParam = param1
	if f.SomeMethodCall.Stub != nil {
		return f.SomeMethodCall.Stub(param1)
	}
	return f.SomeMethodCall.Returns.SomeResult
}
