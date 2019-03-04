package main

import (
	"bytes"
	"go/format"
	"go/token"
	"log"
	"os"
	"path/filepath"

	"github.com/pivotal-cf/jhanda"
	"github.com/ryanmoran/faux/gen"
	"golang.org/x/tools/imports"
)

func main() {
	var options struct {
		File      string `long:"file"      short:"f" env:"GOFILE" description:"the name of the file to parse"     required:"true"`
		Output    string `long:"output"    short:"o"              description:"the name of the file to write"     required:"true"`
		Interface string `long:"interface" short:"i"              description:"the name of the interface to fake" required:"true"`
	}

	stderr := log.New(os.Stderr, "", 0)

	_, err := jhanda.Parse(&options, os.Args[1:])
	if err != nil {
		stderr.Fatal(err)
	}

	source, err := os.Open(options.File)
	if err != nil {
		stderr.Fatalf("could not open source file: %s", err)
	}

	fake, err := gen.Parse(options.File, source, options.Interface)
	if err != nil {
		stderr.Fatal(err)
	}

	err = os.MkdirAll(filepath.Dir(options.Output), 0755)
	if err != nil {
		stderr.Fatalf("could not create directory: %s", err)
	}

	output, err := os.OpenFile(options.Output, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		stderr.Fatalf("could not create output file: %s", err)
	}
	defer output.Close()

	buffer := bytes.NewBuffer([]byte{})
	err = format.Node(buffer, token.NewFileSet(), gen.Build(fake))
	if err != nil {
		stderr.Fatalf("could not format fake ast: %s", err)
	}

	imports.Debug = true
	result, err := imports.Process(output.Name(), buffer.Bytes(), nil)
	if err != nil {
		stderr.Fatalf("could not process imports: %s", err)
	}

	_, err = output.Write(result)
	if err != nil {
		stderr.Fatalf("could not write output file: %s", err)
	}
}
