package main

import (
	"bytes"
	"io"

	"github.com/pivotal-cf/jhanda"
)

type SimpleInterface interface {
	VariadicMethod(someParams ...int)
	SomeMethod(someParam *bytes.Buffer) (someResult io.Reader)
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
