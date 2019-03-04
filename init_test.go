package main_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFaux(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "faux")
}

var (
	executable string
	version    string
)

var _ = BeforeSuite(func() {
	var err error
	version = fmt.Sprintf("v%d", rand.Intn(30))
	Expect(err).NotTo(HaveOccurred())

	executable, err = gexec.Build("github.com/ryanmoran/faux",
		"-ldflags",
		fmt.Sprintf("-X main.version=%s", version))
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
