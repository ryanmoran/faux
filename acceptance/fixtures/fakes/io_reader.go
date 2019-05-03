package fakes

type Reader struct {
	ReadCall struct {
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
	f.ReadCall.CallCount++
	f.ReadCall.Receives.P = param1
	return f.ReadCall.Returns.N, f.ReadCall.Returns.Err
}
