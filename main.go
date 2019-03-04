package main

import (
	"bytes"
	"go/format"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pivotal-cf/jhanda"
	"github.com/ryanmoran/faux/gen"
	"golang.org/x/tools/imports"
)

var version string = "unknown"

func main() {
	var options struct {
		Help    bool `long:"help"    short:"h" description:"prints the usage"`
		Version bool `long:"version" short:"v" description:"prints the version"`

		File      string `long:"file"      short:"f" env:"GOFILE" description:"the name of the file to parse"`
		Output    string `long:"output"    short:"o"              description:"the name of the file to write"`
		Interface string `long:"interface" short:"i"              description:"the name of the interface to fake"`
	}

	stdout := log.New(os.Stdout, "", 0)
	stderr := log.New(os.Stderr, "", 0)

	_, err := jhanda.Parse(&options, os.Args[1:])
	if err != nil {
		stderr.Fatal(err)
	}

	if options.Help {
		flags, err := jhanda.PrintUsage(options)
		if err != nil {
			stderr.Fatal(err)
		}

		flags = strings.Join(strings.Split(flags, "\n"), "\n  ")

		stdout.Printf(`faux helps you generate fakes

Usage: faux --file <FILE> --output <FILE> --interface <INTERFACE-TO-FAKE> [--help]

Flags:
  %s
`, flags)
		os.Exit(0)
	}

	if options.Version {
		stdout.Print(version)
		os.Exit(0)
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
