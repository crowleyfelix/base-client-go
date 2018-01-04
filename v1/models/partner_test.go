// +build !integration

package models

import (
	"testing"

	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Partner", func() {

	var (
		suite *TestSuite
	)

	BeforeEach(func() {
		suite = new(TestSuite)
	})

	Describe("MarshalJSON", func() {
		Context("When serializing", func() {
			It("should map enumerated ids", func() {
				partner := suite.OutboundPartner()
				partner.JuridicalPerson.Addresses = []Address{*suite.OutboundAddress()}
				partner.NaturalPerson.Addresses = []Address{*suite.OutboundAddress()}

				actual, err := json.Marshal(partner)

				Expect(actual).Should(MatchJSON(suite.OutboundPartnerJSON()))
				Expect(err).Should(BeNil())
			})
		})
	})
})

func TestPartner(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Partner Suite")
}
