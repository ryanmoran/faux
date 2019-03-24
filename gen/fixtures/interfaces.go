package main

import (
	"bytes"
	"io"
)

type FullInterface interface {
	SomeMethod(someParam1 string, someParam2 *bytes.Buffer) (someResult1 int, someResult2 io.Reader)
}

type UnnamedFieldsInterface interface {
	SomeMethod(string, *bytes.Buffer) (int, io.Reader)
}

type ElidedTypesInterface interface {
	SomeMethod(someParam1, someParam2 string) (int, io.Reader)
}

type unexportedInterface interface {
	SomeMethod()
}
