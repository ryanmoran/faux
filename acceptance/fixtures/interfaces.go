package main

import (
	"bytes"
	"io"

	"github.com/pivotal-cf/jhanda"
)

type SimpleInterface interface {
	SomeMethod(someParam *bytes.Buffer) (someResult io.Reader)
	OtherMethod(*bytes.Buffer) (io.Reader, error)
}

type VariadicInterface interface {
	VariadicMethod(someParams ...int)
}

type ChanInterface interface {
	ChanMethod(chan int)
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
