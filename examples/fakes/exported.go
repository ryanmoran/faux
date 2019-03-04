package fakes

type Exported struct {
	DoCall struct {
		Receives struct {
			Arg1 string
			Arg2 int
		}
		Returns struct {
			Output1 bool
			Output2 error
		}
	}
}

func (f *Exported) Do(Arg1 string, Arg2 int) (Output1 bool, Output2 error) {
	f.DoCall.Receives.Arg1 = arg1
	f.DoCall.Receives.Arg2 = arg2
	return f.DoCall.Returns.Output1, f.DoCall.Returns.Output2
}
