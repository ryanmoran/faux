package main_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("faux", func() {
	var (
		tempDir    string
		sourceFile string
		outputFile string
	)

	BeforeEach(func() {
		var err error
		tempDir, err = ioutil.TempDir("", "faux-test")
		Expect(err).NotTo(HaveOccurred())

		sourceFile = filepath.Join(tempDir, "source.go")
		outputFile = filepath.Join(tempDir, "fakes", "output.go")

		err = ioutil.WriteFile(sourceFile, []byte(`package main

import (
	"io"
	"bytes"
)

type SomeInterface interface {
	SomeMethod(someParam *bytes.Buffer) (someResult io.Reader)
}
`), 0644)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		err := os.Chmod(tempDir, 0777)
		Expect(err).NotTo(HaveOccurred())

		err = os.RemoveAll(tempDir)
		Expect(err).NotTo(HaveOccurred())
	})

	It("generates a fake", func() {
		command := exec.Command(executable,
			"--file", sourceFile,
			"--output", outputFile,
			"--interface", "SomeInterface")

		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		Eventually(session).Should(gexec.Exit(0))

		outputContent, err := ioutil.ReadFile(outputFile)
		Expect(err).NotTo(HaveOccurred())

		Expect(string(outputContent)).To(ContainSubstring(`package fakes

import (
	"bytes"
	"io"
)

type SomeInterface struct {
	SomeMethodCall struct {
		CallCount int
		Receives  struct {
			SomeParam *bytes.Buffer
		}
		Returns struct {
			SomeResult io.Reader
		}
	}
}

func (f *SomeInterface) SomeMethod(someParam *bytes.Buffer) io.Reader {
	f.SomeMethodCall.CallCount++
	f.SomeMethodCall.Receives.SomeParam = someParam
	return f.SomeMethodCall.Returns.SomeResult
}
`))
	})

	Context("when the source file is provided via an environment variable", func() {
		It("generates a fake", func() {
			command := exec.Command(executable,
				"--output", outputFile,
				"--interface", "SomeInterface")
			command.Env = append(command.Env, fmt.Sprintf("GOFILE=%s", sourceFile))

			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))

			outputContent, err := ioutil.ReadFile(outputFile)
			Expect(err).NotTo(HaveOccurred())

			Expect(string(outputContent)).To(ContainSubstring(`package fakes

import (
	"bytes"
	"io"
)

type SomeInterface struct {
	SomeMethodCall struct {
		CallCount int
		Receives  struct {
			SomeParam *bytes.Buffer
		}
		Returns struct {
			SomeResult io.Reader
		}
	}
}

func (f *SomeInterface) SomeMethod(someParam *bytes.Buffer) io.Reader {
	f.SomeMethodCall.CallCount++
	f.SomeMethodCall.Receives.SomeParam = someParam
	return f.SomeMethodCall.Returns.SomeResult
}
`))
		})
	})

	Context("failure cases", func() {
		Context("when a unknown flag is passed", func() {
			It("exits non-zero with an error", func() {
				command := exec.Command(executable, "--unknown-flag")

				session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				Eventually(session).Should(gexec.Exit(1))

				Expect(string(session.Err.Contents())).To(ContainSubstring("flag provided but not defined: -unknown-flag"))
			})
		})

		Context("when the source file does not exist", func() {
			It("exits non-zero with an error", func() {
				err := os.Remove(sourceFile)
				Expect(err).NotTo(HaveOccurred())

				command := exec.Command(executable,
					"--file", sourceFile,
					"--output", outputFile,
					"--interface", "SomeInterface")

				session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				Eventually(session).Should(gexec.Exit(1))

				Expect(string(session.Err.Contents())).To(ContainSubstring("could not open source file"))
			})
		})

		Context("when the source file cannot be parsed", func() {
			It("exits non-zero with an error", func() {
				err := ioutil.WriteFile(sourceFile, []byte(`garbage`), 0644)
				Expect(err).NotTo(HaveOccurred())

				command := exec.Command(executable,
					"--file", sourceFile,
					"--output", outputFile,
					"--interface", "SomeInterface")

				session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				Eventually(session).Should(gexec.Exit(1))

				Expect(string(session.Err.Contents())).To(ContainSubstring("could not parse source"))
			})
		})

		Context("when the output file directory cannot be created", func() {
			It("exits non-zero with an error", func() {
				err := os.Chmod(tempDir, 0555)
				Expect(err).NotTo(HaveOccurred())

				command := exec.Command(executable,
					"--file", sourceFile,
					"--output", outputFile,
					"--interface", "SomeInterface")

				session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				Eventually(session).Should(gexec.Exit(1))

				Expect(string(session.Err.Contents())).To(ContainSubstring("could not create directory"))
			})

			Context("when the output file cannot be created", func() {
				It("exits non-zero with an error", func() {
					err := os.Mkdir(filepath.Join(tempDir, "fakes"), 0555)
					Expect(err).NotTo(HaveOccurred())

					command := exec.Command(executable,
						"--file", sourceFile,
						"--output", outputFile,
						"--interface", "SomeInterface")

					session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
					Expect(err).NotTo(HaveOccurred())
					Eventually(session).Should(gexec.Exit(1))

					Expect(string(session.Err.Contents())).To(ContainSubstring("could not create output file"))
				})
			})
		})
	})
})
