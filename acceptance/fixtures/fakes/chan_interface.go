package fakes

import "sync"

type ChanInterface struct {
	ChanMethodCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			IntChannel    chan int
			StringChannel <-chan string
		}
		Returns struct {
			BoolChannel chan<- bool
		}
		Stub func(chan int, <-chan string) chan<- bool
	}
}

func (f *ChanInterface) ChanMethod(param1 chan int, param2 <-chan string) chan<- bool {
	f.ChanMethodCall.Lock()
	defer f.ChanMethodCall.Unlock()
	f.ChanMethodCall.CallCount++
	f.ChanMethodCall.Receives.IntChannel = param1
	f.ChanMethodCall.Receives.StringChannel = param2
	if f.ChanMethodCall.Stub != nil {
		return f.ChanMethodCall.Stub(param1, param2)
	}
	return f.ChanMethodCall.Returns.BoolChannel
}
