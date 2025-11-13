package fakes

import (
	"sync"

	"github.com/ryanmoran/faux/acceptance/fixtures"
)

type SimpleNestedGenericInterface[T any] struct {
	ProcessValueCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			ResultTError fixtures.Result[T, error]
		}
		Returns struct {
			T T
		}
		Stub func(fixtures.Result[T, error]) T
	}
	SimpleMethodCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Value T
		}
		Returns struct {
			ResultTError fixtures.Result[T, error]
		}
		Stub func(T) fixtures.Result[T, error]
	}
}

func (f *SimpleNestedGenericInterface[T]) ProcessValue(param1 fixtures.Result[T, error]) T {
	f.ProcessValueCall.mutex.Lock()
	defer f.ProcessValueCall.mutex.Unlock()
	f.ProcessValueCall.CallCount++
	f.ProcessValueCall.Receives.ResultTError = param1
	if f.ProcessValueCall.Stub != nil {
		return f.ProcessValueCall.Stub(param1)
	}
	return f.ProcessValueCall.Returns.T
}
func (f *SimpleNestedGenericInterface[T]) SimpleMethod(param1 T) fixtures.Result[T, error] {
	f.SimpleMethodCall.mutex.Lock()
	defer f.SimpleMethodCall.mutex.Unlock()
	f.SimpleMethodCall.CallCount++
	f.SimpleMethodCall.Receives.Value = param1
	if f.SimpleMethodCall.Stub != nil {
		return f.SimpleMethodCall.Stub(param1)
	}
	return f.SimpleMethodCall.Returns.ResultTError
}
