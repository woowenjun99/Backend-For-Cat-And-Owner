package test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("A fake test", func() {
	It("should return 2 when 1 + 1", func() {
		Expect(1 + 1).To(Equal(2))
	})
})

func TestXxx(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "The mock test")
}
