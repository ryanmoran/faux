package fixtures

import (
	"bytes"
	"io"

	"github.com/BurntSushi/toml"
	"github.com/pivotal-cf/jhanda"
)

type VariadicInterface interface {
	VariadicMethod(someParams ...int)
}

type ChanInterface interface {
	ChanMethod(chan int, <-chan string) chan<- bool
}

type ModuleInterface interface {
	SomeMethod(usage jhanda.Usage)
}

type DuplicateArgumentInterface interface {
	Duplicates(string, string, int) (string, int, int)
}

type FunctionInterface interface {
	FuncMethod(func(string) error) func(int) bool
}

type NamedInterface interface {
	SomeMethod(someParam *bytes.Buffer) (someResult io.Reader)
}

type BurntSushiParser struct {
	Key toml.Key
}
