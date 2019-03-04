package examples

//go:generate faux -i Exported -o fakes/exported.go
type Exported interface {
	Do(arg1 string, arg2 int) (output1 bool, output2 error)
}
