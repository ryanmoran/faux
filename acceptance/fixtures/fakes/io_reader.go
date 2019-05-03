package fakes

import "sync"

type Reader struct {
	ReadCall struct {
		sync.Mutex
		CallCount int
		Receives  struct {
			P []byte
		}
		Returns struct {
			N   int
			Err error
		}
	}
}

func (f *Reader) Read(param1 []byte) (int, error) {
	f.ReadCall.Lock()
	defer f.ReadCall.Unlock()
	f.ReadCall.CallCount++
	f.ReadCall.Receives.P = param1
	return f.ReadCall.Returns.N, f.ReadCall.Returns.Err
}
