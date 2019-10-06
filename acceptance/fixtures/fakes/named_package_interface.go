package fakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-utils/logger"
	clogger "github.com/hashicorp/consul/logger"
)

type NamedPackageInterface struct {
	NamedPackageMethodCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			Config clogger.Config
			Level  logger.LogLevel
		}
		Stub func(clogger.Config, logger.LogLevel)
	}
}

func (f *NamedPackageInterface) NamedPackageMethod(param1 clogger.Config, param2 logger.LogLevel) {
	f.NamedPackageMethodCall.Lock()
	defer f.NamedPackageMethodCall.Unlock()
	f.NamedPackageMethodCall.CallCount++
	f.NamedPackageMethodCall.Receives.Config = param1
	f.NamedPackageMethodCall.Receives.Level = param2
	if f.NamedPackageMethodCall.Stub != nil {
		f.NamedPackageMethodCall.Stub(param1, param2)
	}
}
