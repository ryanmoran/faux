package main_test

import (
	"fmt"
	"io"
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

		interfaceFixtureFile, err := os.Open("fixtures/interfaces.go")
		Expect(err).NotTo(HaveOccurred())

		sourceFile, err := os.OpenFile(sourceFile, os.O_RDWR|os.O_CREATE, 0644)
		Expect(err).NotTo(HaveOccurred())

		_, err = io.Copy(sourceFile, interfaceFixtureFile)
		Expect(err).NotTo(HaveOccurred())

		err = interfaceFixtureFile.Close()
		Expect(err).NotTo(HaveOccurred())

		err = sourceFile.Close()
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

		expectedContent, err := ioutil.ReadFile("fixtures/fakes/some_interface.go")
		Expect(err).NotTo(HaveOccurred())

		Expect(string(outputContent)).To(Equal(string(expectedContent)))
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

			expectedContent, err := ioutil.ReadFile("fixtures/fakes/some_interface.go")
			Expect(err).NotTo(HaveOccurred())

			Expect(string(outputContent)).To(Equal(string(expectedContent)))
		})
	})
})
