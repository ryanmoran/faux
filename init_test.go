package main_test

import (
	"testing"

	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFaux(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "faux")
}

var executable string

var _ = BeforeSuite(func() {
	var err error
	executable, err = gexec.Build("github.com/ryanmoran/faux")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
