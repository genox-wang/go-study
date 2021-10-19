package ginkgo_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AwesomeProject1", func() {
	var testValue int
	BeforeEach(func() {
		testValue = 3
	})
	Context("Correctness of testValue", func() {
		It(fmt.Sprintf("testValue should be %d", testValue), func() {
			Expect(testValue).To(Equal(3))
		})
	})
})

func TestGoStudy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoStudy Suite")
}
