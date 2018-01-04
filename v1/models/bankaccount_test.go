// +build !integration

package models

import (
	"testing"

	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BankAccount", func() {

	var (
		suite *TestSuite
	)

	BeforeEach(func() {
		suite = new(TestSuite)
	})

	Describe("MarshalJSON", func() {
		Context("When serializing", func() {
			It("should map enumerated ids", func() {
				bankAccount := suite.OutboundBankAccount()
				actual, err := json.Marshal(bankAccount)

				Expect(actual).Should(MatchJSON(suite.OutboundBankAccountJSON()))
				Expect(err).Should(BeNil())
			})
		})
	})
})

func TestBankAccount(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BankAccount Suite")
}
