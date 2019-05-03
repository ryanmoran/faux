package fakes

import "github.com/pivotal-cf/jhanda"

type Command struct {
	ExecuteCall struct {
		CallCount int
		Receives  struct {
			Args []string
		}
		Returns struct {
			Error error
		}
	}
	UsageCall struct {
		CallCount int
		Returns   struct {
			Usage jhanda.Usage
		}
	}
}

func (f *Command) Execute(param1 []string) error {
	f.ExecuteCall.CallCount++
	f.ExecuteCall.Receives.Args = param1
	return f.ExecuteCall.Returns.Error
}
func (f *Command) Usage() jhanda.Usage {
	f.UsageCall.CallCount++
	return f.UsageCall.Returns.Usage
}
