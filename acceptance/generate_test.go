package acceptance_test

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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
			command.Dir = "./fixtures"

			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))

			outputContent, err := ioutil.ReadFile(outputFile)
			Expect(err).NotTo(HaveOccurred())

			expectedContent, err := ioutil.ReadFile(filepath.Join("fixtures", "fakes", fixture))
			Expect(err).NotTo(HaveOccurred())

			var lines []interface{}
			for _, line := range strings.Split(string(expectedContent), "\n") {
				lines = append(lines, line)
			}

			Expect(string(outputContent)).To(ContainLines(lines...))
		},

		Entry("simple", "simple_interface.go", "--file", "./interfaces.go", "--interface", "SimpleInterface"),
		Entry("channels", "chan_interface.go", "--file", "./interfaces.go", "--interface", "ChanInterface"),
		Entry("duplicate arguments", "duplicate_argument_interface.go", "--file", "./interfaces.go", "--interface", "DuplicateArgumentInterface"),
		Entry("gomod", "module_interface.go", "--file", "./interfaces.go", "--interface", "ModuleInterface"),
		Entry("gopath", "jhanda_command.go", "--package", "github.com/pivotal-cf/jhanda", "--interface", "Command"),
		Entry("stdlib", "io_reader.go", "--package", "io", "--interface", "Reader"),
		Entry("variadic", "variadic_interface.go", "--file", "./interfaces.go", "--interface", "VariadicInterface"),
		Entry("functions", "function_interface.go", "--file", "./interfaces.go", "--interface", "FunctionInterface"),
		Entry("name", "named_interface.go", "--file", "./interfaces.go", "--interface", "NamedInterface", "--name", "SomeNamedInterface"),
		Entry("generic", "generic_interface.go", "--file", "./interfaces.go", "--interface", "GenericInterface"),
		Entry("package conflict", "package_conflict_interface.go", "--file", "./interfaces.go", "--interface", "PackageConflictInterface"),
	)

	Context("when the source file is provided via an environment variable", func() {
		It("generates a fake", func() {
			command := exec.Command(executable,
				"--output", outputFile,
				"--interface", "SimpleInterface")
			command.Env = append(os.Environ(), "GOFILE=./interfaces.go")
			command.Dir = "./fixtures"

			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))

			outputContent, err := ioutil.ReadFile(outputFile)
			Expect(err).NotTo(HaveOccurred())

			expectedContent, err := ioutil.ReadFile("fixtures/fakes/simple_interface.go")
			Expect(err).NotTo(HaveOccurred())

			var lines []interface{}
			for _, line := range strings.Split(string(expectedContent), "\n") {
				lines = append(lines, line)
			}

			Expect(string(outputContent)).To(ContainLines(lines...))
		})
	})
})
