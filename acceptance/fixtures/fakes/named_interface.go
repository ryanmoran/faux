package fakes

import (
	"bytes"
	"io"
	"sync"
)

type SomeNamedInterface struct {
	SomeMethodCall struct {
		mutex     sync.Mutex
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

func (f *SomeNamedInterface) SomeMethod(param1 *bytes.Buffer) io.Reader {
	f.SomeMethodCall.mutex.Lock()
	defer f.SomeMethodCall.mutex.Unlock()
	f.SomeMethodCall.CallCount++
	f.SomeMethodCall.Receives.SomeParam = param1
	if f.SomeMethodCall.Stub != nil {
		return f.SomeMethodCall.Stub(param1)
	}
	return f.SomeMethodCall.Returns.SomeResult
}
