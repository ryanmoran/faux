package rendering

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRendering(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "rendering")
}
