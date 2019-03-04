package fakes

type Other struct {
	ExecuteCall struct {
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

func (f *Other) Execute(Arg1 string, Arg2 int) (Output1 bool, Output2 error) {
	f.ExecuteCall.Receives.Arg1 = arg1
	f.ExecuteCall.Receives.Arg2 = arg2
	return f.ExecuteCall.Returns.Output1, f.ExecuteCall.Returns.Output2
}
