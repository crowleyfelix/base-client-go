// +build !integration

package models

import (
	"testing"

	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Contact", func() {

	var (
		suite *TestSuite
	)

	BeforeEach(func() {
		suite = new(TestSuite)
	})

	Describe("MarshalJSON", func() {
		Context("When serializing", func() {
			It("should map enumerated ids", func() {
				contact := suite.OutboundContact()
				contact.Phones = []Phone{*suite.OutboundPhone()}
				contact.Emails = suite.OutboundEmail()

				actual, err := json.Marshal(contact)

				Expect(actual).Should(MatchJSON(suite.OutboundContactJSON()))
				Expect(err).Should(BeNil())
			})
		})
	})
})

func TestContact(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Contact Suite")
}
