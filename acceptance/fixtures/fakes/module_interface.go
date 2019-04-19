package fakes

import "github.com/pivotal-cf/jhanda"

type ModuleInterface struct {
	SomeMethodCall struct {
		CallCount int
		Receives  struct {
			Usage jhanda.Usage
		}
	}
}

func (f *ModuleInterface) SomeMethod(usage jhanda.Usage) {
	f.SomeMethodCall.CallCount++
	f.SomeMethodCall.Receives.Usage = usage
}
