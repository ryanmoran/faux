package fakes

import (
	"sync"

	"github.com/ryanmoran/faux/acceptance/fixtures"
)

type AliasInterface struct {
	HandleErrorCall struct {
		mutex     sync.Mutex
		CallCount int
		Returns   struct {
			ErrorAlias fixtures.ErrorAlias
		}
		Stub func() fixtures.ErrorAlias
	}
	HandleRequestCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			HandlerFunc fixtures.HandlerFunc
		}
		Returns struct {
			MiddlewareFunc fixtures.MiddlewareFunc
		}
		Stub func(fixtures.HandlerFunc) fixtures.MiddlewareFunc
	}
	ProcessMapCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			MapAlias fixtures.MapAlias
		}
		Returns struct {
			SliceAlias fixtures.SliceAlias
		}
		Stub func(fixtures.MapAlias) fixtures.SliceAlias
	}
	ProcessStringCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			StringAlias fixtures.StringAlias
		}
		Returns struct {
			IntAlias fixtures.IntAlias
		}
		Stub func(fixtures.StringAlias) fixtures.IntAlias
	}
}

func (f *AliasInterface) HandleError() fixtures.ErrorAlias {
	f.HandleErrorCall.mutex.Lock()
	defer f.HandleErrorCall.mutex.Unlock()
	f.HandleErrorCall.CallCount++
	if f.HandleErrorCall.Stub != nil {
		return f.HandleErrorCall.Stub()
	}
	return f.HandleErrorCall.Returns.ErrorAlias
}
func (f *AliasInterface) HandleRequest(param1 fixtures.HandlerFunc) fixtures.MiddlewareFunc {
	f.HandleRequestCall.mutex.Lock()
	defer f.HandleRequestCall.mutex.Unlock()
	f.HandleRequestCall.CallCount++
	f.HandleRequestCall.Receives.HandlerFunc = param1
	if f.HandleRequestCall.Stub != nil {
		return f.HandleRequestCall.Stub(param1)
	}
	return f.HandleRequestCall.Returns.MiddlewareFunc
}
func (f *AliasInterface) ProcessMap(param1 fixtures.MapAlias) fixtures.SliceAlias {
	f.ProcessMapCall.mutex.Lock()
	defer f.ProcessMapCall.mutex.Unlock()
	f.ProcessMapCall.CallCount++
	f.ProcessMapCall.Receives.MapAlias = param1
	if f.ProcessMapCall.Stub != nil {
		return f.ProcessMapCall.Stub(param1)
	}
	return f.ProcessMapCall.Returns.SliceAlias
}
func (f *AliasInterface) ProcessString(param1 fixtures.StringAlias) fixtures.IntAlias {
	f.ProcessStringCall.mutex.Lock()
	defer f.ProcessStringCall.mutex.Unlock()
	f.ProcessStringCall.CallCount++
	f.ProcessStringCall.Receives.StringAlias = param1
	if f.ProcessStringCall.Stub != nil {
		return f.ProcessStringCall.Stub(param1)
	}
	return f.ProcessStringCall.Returns.IntAlias
}
