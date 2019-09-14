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
	"github.com/ryanmoran/faux/parsing"
	"github.com/ryanmoran/faux/rendering"
	"golang.org/x/tools/imports"
)

var version string = "unknown"

func main() {
	var options struct {
		Help    bool `long:"help"    short:"h" description:"prints the usage"`
		Version bool `long:"version" short:"v" description:"prints the version"`

		Package   string `long:"package"   short:"p"              description:"the name of the package that contains the interface"`
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

	if options.Package == "" {
		path, err := filepath.Abs(options.File)
		if err != nil {
			stderr.Fatal(err)
		}

		options.Package = filepath.Dir(path)
	}

	iface, err := parsing.Parse(options.Package, options.Interface)
	if err != nil {
		stderr.Fatal(err)
	}

	err = os.MkdirAll(filepath.Dir(options.Output), 0755)
	if err != nil {
		stderr.Fatalf("could not create directory: %s", err)
	}

	output, err := os.OpenFile(options.Output, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		stderr.Fatalf("could not create output file: %s", err)
	}
	defer output.Close()

	buffer := bytes.NewBuffer([]byte{})
	context := rendering.NewContext()
	tree := context.Build(iface).AST()

	err = format.Node(buffer, token.NewFileSet(), tree)
	if err != nil {
		stderr.Fatalf("could not format fake ast: %s", err)
	}

	outputPath, err := filepath.Abs(output.Name())
	if err != nil {
		stderr.Fatalf("could not determine output absolute path: %s", err)
	}

	result, err := imports.Process(outputPath, buffer.Bytes(), &imports.Options{FormatOnly: false})
	if err != nil {
		stderr.Fatalf("could not process imports: %s\n\n%s", err, buffer.String())
	}

	_, err = output.Write(result)
	if err != nil {
		stderr.Fatalf("could not write output file: %s", err)
	}
}
