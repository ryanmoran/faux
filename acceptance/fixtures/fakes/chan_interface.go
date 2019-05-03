package fakes

import "sync"

type ChanInterface struct {
	ChanMethodCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			IntChannel chan int
		}
	}
}

func (f *ChanInterface) ChanMethod(param1 chan int) {
	f.ChanMethodCall.Lock()
	defer f.ChanMethodCall.Unlock()
	f.ChanMethodCall.CallCount++
	f.ChanMethodCall.Receives.IntChannel = param1
}
