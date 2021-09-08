package fakes

import (
	"sync"

	"github.com/pivotal-cf/jhanda"
)

type Command struct {
	ExecuteCall struct {
		mutex     sync.Mutex
		CallCount int
		Receives  struct {
			Args []string
		}
		Returns struct {
			Error error
		}
		Stub func([]string) error
	}
	UsageCall struct {
		mutex     sync.Mutex
		CallCount int
		Returns   struct {
			Usage jhanda.Usage
		}
		Stub func() jhanda.Usage
	}
}

func (f *Command) Execute(param1 []string) error {
	f.ExecuteCall.mutex.Lock()
	defer f.ExecuteCall.mutex.Unlock()
	f.ExecuteCall.CallCount++
	f.ExecuteCall.Receives.Args = param1
	if f.ExecuteCall.Stub != nil {
		return f.ExecuteCall.Stub(param1)
	}
	return f.ExecuteCall.Returns.Error
}
func (f *Command) Usage() jhanda.Usage {
	f.UsageCall.mutex.Lock()
	defer f.UsageCall.mutex.Unlock()
	f.UsageCall.CallCount++
	if f.UsageCall.Stub != nil {
		return f.UsageCall.Stub()
	}
	return f.UsageCall.Returns.Usage
}
