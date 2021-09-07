package acceptance_test

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("faux", func() {
	var (
		tempDir    string
		outputFile string
	)

	BeforeEach(func() {
		var err error
		tempDir, err = ioutil.TempDir("", "faux-test")
		Expect(err).NotTo(HaveOccurred())

		outputFile = filepath.Join(tempDir, "fakes", "output.go")
	})

	AfterEach(func() {
		err := os.Chmod(tempDir, 0777)
		Expect(err).NotTo(HaveOccurred())

		err = os.RemoveAll(tempDir)
		Expect(err).NotTo(HaveOccurred())
	})

	DescribeTable("fake generation",
		func(fixture string, flags ...string) {
			flags = append(flags, "--output", outputFile)
			command := exec.Command(executable, flags...)
			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session, "10s").Should(gexec.Exit(0))

			outputContent, err := ioutil.ReadFile(outputFile)
			Expect(err).NotTo(HaveOccurred())

			expectedContent, err := ioutil.ReadFile(filepath.Join("fixtures", "fakes", fixture))
			Expect(err).NotTo(HaveOccurred())

			Expect(string(outputContent)).To(ContainSubstring(string(expectedContent)))
		},

		Entry("simple", "simple_interface.go", "--file", "./fixtures/interfaces.go", "--interface", "SimpleInterface"),
		Entry("channels", "chan_interface.go", "--file", "./fixtures/interfaces.go", "--interface", "ChanInterface"),
		Entry("duplicate arguments", "duplicate_argument_interface.go", "--file", "./fixtures/interfaces.go", "--interface", "DuplicateArgumentInterface"),
		Entry("gomod", "module_interface.go", "--file", "./fixtures/interfaces.go", "--interface", "ModuleInterface"),
		Entry("gopath", "jhanda_command.go", "--package", "github.com/pivotal-cf/jhanda", "--interface", "Command"),
		Entry("stdlib", "io_reader.go", "--package", "io", "--interface", "Reader"),
		Entry("variadic", "variadic_interface.go", "--file", "./fixtures/interfaces.go", "--interface", "VariadicInterface"),
		Entry("functions", "function_interface.go", "--file", "./fixtures/interfaces.go", "--interface", "FunctionInterface"),
		Entry("name", "named_interface.go", "--file", "./fixtures/interfaces.go", "--interface", "NamedInterface", "--name", "SomeNamedInterface"),
	)

	Context("when the source file is provided via an environment variable", func() {
		It("generates a fake", func() {
			command := exec.Command(executable,
				"--output", outputFile,
				"--interface", "SimpleInterface")
			command.Env = append(os.Environ(), "GOFILE=./fixtures/interfaces.go")

			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session, "10s").Should(gexec.Exit(0))

			outputContent, err := ioutil.ReadFile(outputFile)
			Expect(err).NotTo(HaveOccurred())

			expectedContent, err := ioutil.ReadFile("fixtures/fakes/simple_interface.go")
			Expect(err).NotTo(HaveOccurred())

			Expect(string(outputContent)).To(ContainSubstring(string(expectedContent)))
		})
	})
})
