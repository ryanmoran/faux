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
		err := os.RemoveAll(tempDir)
		Expect(err).NotTo(HaveOccurred())
	})

	It("truncates the output file before writing", func() {
		var contents string
		for i := 0; i < 1000; i++ {
			contents = contents + "test\n"
		}

		err := os.Mkdir(filepath.Dir(outputFile), 0755)
		Expect(err).NotTo(HaveOccurred())

		err = ioutil.WriteFile(outputFile, []byte(contents), 0600)
		Expect(err).NotTo(HaveOccurred())

		command := exec.Command(executable,
			"--file", "./fixtures/interfaces.go",
			"--output", outputFile,
			"--interface", "SimpleInterface")

		session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
		Eventually(session).Should(gexec.Exit(0))

		outputContents, err := ioutil.ReadFile(outputFile)
		Expect(err).NotTo(HaveOccurred())

		Expect(string(outputContents)).NotTo(ContainSubstring(string("test\ntest")))
	})
})
