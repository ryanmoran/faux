package parsing

import "go/types"

type Signature struct {
	Name    string
	Params  []Argument
	Results []Argument
}

func NewSignature(f *types.Func) Signature {
	signature := f.Type().(*types.Signature)

	var params []Argument
	for i := 0; i < signature.Params().Len(); i++ {
		var variadic bool
		if i == signature.Params().Len()-1 && signature.Variadic() {
			variadic = true
		}

		params = append(params, NewArgument(signature.Params().At(i), variadic))
	}

	var results []Argument
	for i := 0; i < signature.Results().Len(); i++ {
		results = append(results, NewArgument(signature.Results().At(i), false))
	}

	return Signature{
		Name:    f.Name(),
		Params:  params,
		Results: results,
	}
}
