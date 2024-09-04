package fakes

import (
	"sync"

	"github.com/ryanmoran/faux/acceptance/fixtures"
)

type GenericInterface[T comparable, S comparable] struct {
	SomeMethodCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			MapTS map[T]S
		}
		Returns struct {
			ResultIntError fixtures.Result[int, error]
		}
		Stub func(map[T]S) fixtures.Result[int, error]
	}
}

func (f *GenericInterface[T, S]) SomeMethod(param1 map[T]S) fixtures.Result[int, error] {
	f.SomeMethodCall.mutex.Lock()
	defer f.SomeMethodCall.mutex.Unlock()
	f.SomeMethodCall.CallCount++
	f.SomeMethodCall.Receives.MapTS = param1
	if f.SomeMethodCall.Stub != nil {
		return f.SomeMethodCall.Stub(param1)
	}
	return f.SomeMethodCall.Returns.ResultIntError
}
