package gen

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGen(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "gen")
}
