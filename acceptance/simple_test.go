package acceptance_test

import (
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

	It("generates a fake", func() {
		command := exec.Command(executable,
			"--file", "./fixtures/interfaces.go",
			"--output", outputFile,
			"--interface", "SimpleInterface")

		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		Eventually(session, "10s").Should(gexec.Exit(0))

		outputContent, err := ioutil.ReadFile(outputFile)
		Expect(err).NotTo(HaveOccurred())

		expectedContent, err := ioutil.ReadFile("fixtures/fakes/simple_interface.go")
		Expect(err).NotTo(HaveOccurred())

		Expect(string(outputContent)).To(ContainSubstring(string(expectedContent)))
	})

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
