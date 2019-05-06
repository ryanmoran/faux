package fakes

import "sync"

type VariadicInterface struct {
	VariadicMethodCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			SomeParams []int
		}
		Stub func(...int)
	}
}

func (f *VariadicInterface) VariadicMethod(param1 ...int) {
	f.VariadicMethodCall.Lock()
	defer f.VariadicMethodCall.Unlock()
	f.VariadicMethodCall.CallCount++
	f.VariadicMethodCall.Receives.SomeParams = param1
	if f.VariadicMethodCall.Stub != nil {
		f.VariadicMethodCall.Stub(param1...)
	}
}
