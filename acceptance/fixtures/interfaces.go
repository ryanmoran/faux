package main

import (
	"bytes"
	"io"
)

type SomeInterface interface {
	SomeMethod(someParam *bytes.Buffer) (someResult io.Reader)
}
