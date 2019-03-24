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

		Context("when the source file cannot be parsed", func() {
			It("exits non-zero with an error", func() {
				command := exec.Command(executable,
					"--file", "./fixtures/garbage/interface.go",
					"--output", outputFile,
					"--interface", "SimpleInterface")

				session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				Eventually(session).Should(gexec.Exit(1))

				Expect(string(session.Err.Contents())).To(ContainSubstring("could not find interface \"SimpleInterface\""))
			})
		})

		Context("when the output file directory cannot be created", func() {
			It("exits non-zero with an error", func() {
				err := os.Chmod(tempDir, 0555)
				Expect(err).NotTo(HaveOccurred())

				command := exec.Command(executable,
					"--file", "./fixtures/interfaces.go",
					"--output", outputFile,
					"--interface", "SimpleInterface")

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
						"--file", "./fixtures/interfaces.go",
						"--output", outputFile,
						"--interface", "SimpleInterface")

					session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
					Expect(err).NotTo(HaveOccurred())
					Eventually(session).Should(gexec.Exit(1))

					Expect(string(session.Err.Contents())).To(ContainSubstring("could not create output file"))
				})
			})
		})
	})
})
