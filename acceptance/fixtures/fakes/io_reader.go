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

func (f *Reader) Read(p []byte) (int, error) {
	f.ReadCall.CallCount++
	f.ReadCall.Receives.P = p
	return f.ReadCall.Returns.N, f.ReadCall.Returns.Err
}
