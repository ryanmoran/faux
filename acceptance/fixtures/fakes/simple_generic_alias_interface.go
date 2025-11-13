package fakes

import (
	"sync"

	"github.com/ryanmoran/faux/acceptance/fixtures"
)

type SimpleGenericAliasInterface struct {
	ProcessMapCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			MapAlias fixtures.MapAlias
		}
		Returns struct {
			ErrorAlias fixtures.ErrorAlias
		}
		Stub func(fixtures.MapAlias) fixtures.ErrorAlias
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

func (f *SimpleGenericAliasInterface) ProcessMap(param1 fixtures.MapAlias) fixtures.ErrorAlias {
	f.ProcessMapCall.mutex.Lock()
	defer f.ProcessMapCall.mutex.Unlock()
	f.ProcessMapCall.CallCount++
	f.ProcessMapCall.Receives.MapAlias = param1
	if f.ProcessMapCall.Stub != nil {
		return f.ProcessMapCall.Stub(param1)
	}
	return f.ProcessMapCall.Returns.ErrorAlias
}
func (f *SimpleGenericAliasInterface) ProcessString(param1 fixtures.StringAlias) fixtures.IntAlias {
	f.ProcessStringCall.mutex.Lock()
	defer f.ProcessStringCall.mutex.Unlock()
	f.ProcessStringCall.CallCount++
	f.ProcessStringCall.Receives.StringAlias = param1
	if f.ProcessStringCall.Stub != nil {
		return f.ProcessStringCall.Stub(param1)
	}
	return f.ProcessStringCall.Returns.IntAlias
}
