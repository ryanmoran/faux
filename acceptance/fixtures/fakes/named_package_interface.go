package fakes

import (
	"encoding/base64"
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
		Returns struct {
			Encoding base64.Encoding
		}
		Stub func(clogger.Config, logger.LogLevel) base64.Encoding
	}
}

func (f *NamedPackageInterface) NamedPackageMethod(param1 clogger.Config, param2 logger.LogLevel) base64.Encoding {
	f.NamedPackageMethodCall.Lock()
	defer f.NamedPackageMethodCall.Unlock()
	f.NamedPackageMethodCall.CallCount++
	f.NamedPackageMethodCall.Receives.Config = param1
	f.NamedPackageMethodCall.Receives.Level = param2
	if f.NamedPackageMethodCall.Stub != nil {
		return f.NamedPackageMethodCall.Stub(param1, param2)
	}
	return f.NamedPackageMethodCall.Returns.Encoding
}
