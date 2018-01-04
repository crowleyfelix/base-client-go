// +build !integration

package models

import (
	"testing"

	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Phone", func() {

	var (
		suite *TestSuite
	)

	BeforeEach(func() {
		suite = new(TestSuite)
	})

	Describe("MarshalJSON", func() {
		Context("When serializing", func() {
			It("should map enumerated ids", func() {
				phone := suite.OutboundPhone()
				actual, err := json.Marshal(phone)

				Expect(actual).Should(MatchJSON(suite.OutboundPhoneJSON()))
				Expect(err).Should(BeNil())
			})
		})
	})
})

func TestPhone(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Phone Suite")
}
