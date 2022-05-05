package fixtures

import (
	"bytes"
	"io"

	"github.com/pelletier/go-toml/v2"
)

type SimpleInterface interface {
	SomeMethod(someParam *bytes.Buffer) (someResult io.Reader)
	OtherMethod(*bytes.Buffer) (io.Reader, error)
}

type PelletierParser struct {
	Key toml.Key
}
