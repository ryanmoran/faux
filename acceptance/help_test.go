package acceptance_test

import (
	"os/exec"

	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("faux", func() {
	Context("when the help flag is provided", func() {
		It("prints the usage", func() {
			command := exec.Command(executable, "-h")

			session, err := gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))

			Expect(string(session.Out.Contents())).To(ContainSubstring("faux helps you generate fakes"))
		})
	})
})
