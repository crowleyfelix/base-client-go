// +build !integration

package models

import (
	"testing"

	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Benefit", func() {

	var (
		suite *TestSuite
	)

	BeforeEach(func() {
		suite = new(TestSuite)
	})

	Describe("MarshalJSON", func() {
		Context("When serializing", func() {
			It("should map enumerated ids", func() {
				benefit := suite.OutboundBenefit()
				actual, err := json.Marshal(benefit)

				Expect(actual).Should(MatchJSON(suite.OutboundBenefitJSON()))
				Expect(err).Should(BeNil())
			})
		})
	})
})

func TestBenefit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Benefit Suite")
}
