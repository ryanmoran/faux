package fakes

import (
	"sync"

	"github.com/pivotal-cf/jhanda"
)

type ModuleInterface struct {
	SomeMethodCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Usage jhanda.Usage
		}
		Stub func(jhanda.Usage)
	}
}

func (f *ModuleInterface) SomeMethod(param1 jhanda.Usage) {
	f.SomeMethodCall.mutex.Lock()
	defer f.SomeMethodCall.mutex.Unlock()
	f.SomeMethodCall.CallCount++
	f.SomeMethodCall.Receives.Usage = param1
	if f.SomeMethodCall.Stub != nil {
		f.SomeMethodCall.Stub(param1)
	}
}
