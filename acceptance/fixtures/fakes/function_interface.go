package fakes

import "sync"

type FunctionInterface struct {
	FuncMethodCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			FuncStringError func(string) error
		}
		Returns struct {
			FuncIntBool func(int) bool
		}
		Stub func(func(string) error) func(int) bool
	}
}

func (f *FunctionInterface) FuncMethod(param1 func(string) error) func(int) bool {
	f.FuncMethodCall.mutex.Lock()
	defer f.FuncMethodCall.mutex.Unlock()
	f.FuncMethodCall.CallCount++
	f.FuncMethodCall.Receives.FuncStringError = param1
	if f.FuncMethodCall.Stub != nil {
		return f.FuncMethodCall.Stub(param1)
	}
	return f.FuncMethodCall.Returns.FuncIntBool
}
