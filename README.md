# Faux

Generates fakes for use in testing Go programs.

## Usage
```
faux helps you generate fakes

Usage: faux --file <FILE> --output <FILE> --interface <INTERFACE-TO-FAKE> [--help]

Flags:
  --file, -f, GOFILE  string  the name of the file to parse
  --help, -h          bool    prints the usage
  --interface, -i     string  the name of the interface to fake
  --output, -o        string  the name of the file to write
```

## Example

Say I have an interface in my codebase that I would like to generate a fake for.
I can add a `go:generate` comment like the following to generate a fake.
```Go
package main

//go:generate faux -i SomeInterface -o fakes/some_interface.go
type SomeInterface interface {
  SomeMethod(someParam bool) (someResult int)
}
```

This will output a `fakes/some_interface.go` file with the following contents.
```
package fakes

type SomeInterface struct {
  SomeMethodCall struct {
    CallCount int
    Receives  struct {
      SomeParam bool
    }
    Returns struct {
      SomeResult int
    }
  }
}

func (f *SomeInterface) SomeMethod(someParam bool) int {
  f.SomeMethodCall.CallCount++
  f.SomeMethodCall.Receives.SomeParam = someParam
  return f.SomeMethodCall.Returns.SomeResult
}
```

## Installation

To download `faux` go to [Releases](https://github.com/ryanmoran/faux/releases).

Alternatively, you can install `faux` via `homebrew`
```sh
brew tap ryanmoran/tools
brew install faux
```

You can also build from source.

### Building from Source
You'll need at least Go 1.11, as
`om` uses Go Modules to manage dependencies.

To build from source, after you've cloned the repo, run these commands from the top level of the repo:

```bash
GO111MODULE=on go mod download
GO111MODULE=on go build
```

Go 1.11 uses some heuristics to determine if Go Modules should be used.
The process above overrides those herusitics
to ensure that Go Modules are _always_ used.
If you have cloned this repo outside of your GOPATH,
`GO111MODULE=on` can be excluded from the above steps.
