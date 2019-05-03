package fakes

type ChanInterface struct {
	ChanMethodCall struct {
		CallCount int
		Receives  struct {
			IntChannel chan int
		}
	}
}

func (f *ChanInterface) ChanMethod(param1 chan int) {
	f.ChanMethodCall.CallCount++
	f.ChanMethodCall.Receives.IntChannel = param1
}
