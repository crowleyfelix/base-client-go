// +build !integration

package models

import (
	"testing"

	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Address", func() {

	var (
		suite *TestSuite
	)

	BeforeEach(func() {
		suite = new(TestSuite)
	})

	Describe("MarshalJSON", func() {
		Context("When serializing", func() {
			It("should map enumerated ids", func() {
				address := suite.OutboundAddress()
				actual, err := json.Marshal(address)

				Expect(actual).Should(MatchJSON(suite.OutboundAddressJSON()))
				Expect(err).Should(BeNil())
			})
		})
	})
})

func TestAddress(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Address Suite")
}
