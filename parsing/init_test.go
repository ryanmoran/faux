package parsing

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestParsing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "parsing")
}
