package fakes

type DuplicateArgumentInterface struct {
	DuplicatesCall struct {
		CallCount int
		Receives  struct {
			String_1 string
			String_2 string
			Int      int
		}
		Returns struct {
			String string
			Int_1  int
			Int_2  int
		}
	}
}

func (f *DuplicateArgumentInterface) Duplicates(param1 string, param2 string, param3 int) (string, int, int) {
	f.DuplicatesCall.CallCount++
	f.DuplicatesCall.Receives.String_1 = param1
	f.DuplicatesCall.Receives.String_2 = param2
	f.DuplicatesCall.Receives.Int = param3
	return f.DuplicatesCall.Returns.String, f.DuplicatesCall.Returns.Int_1, f.DuplicatesCall.Returns.Int_2
}
