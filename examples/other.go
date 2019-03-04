package examples

//go:generate faux -i Other -o fakes/other.go
type Other interface {
	Execute(arg1 string, arg2 int) (output1 bool, output2 error)
}
