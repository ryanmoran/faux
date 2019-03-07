package main_test

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

	Context("when the interface is a gopath package reference", func() {
		It("generates a fake", func() {
			command := exec.Command(executable,
				"--output", outputFile,
				"--package", "github.com/pivotal-cf/jhanda",
				"--interface", "Command")

			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))

			outputContent, err := ioutil.ReadFile(outputFile)
			Expect(err).NotTo(HaveOccurred())

			expectedContent, err := ioutil.ReadFile("fixtures/fakes/jhanda_command.go")
			Expect(err).NotTo(HaveOccurred())

			Expect(string(outputContent)).To(ContainSubstring(string(expectedContent)))
		})
	})
})
