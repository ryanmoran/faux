package main

import (
	"bytes"
	"io"

	"github.com/pivotal-cf/jhanda"
)

type SimpleInterface interface {
	SomeMethod(someParam *bytes.Buffer) (someResult io.Reader)
}

type ModuleInterface interface {
	SomeMethod(usage jhanda.Usage)
}
