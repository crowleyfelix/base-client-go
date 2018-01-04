// +build !integration

package models

import (
	"testing"

	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Member", func() {

	var (
		suite *TestSuite
	)

	BeforeEach(func() {
		suite = new(TestSuite)
	})

	Describe("MarshalJSON", func() {
		Context("When serializing", func() {
			It("should map enumerated ids", func() {
				member := suite.OutboundMember()
				member.Addresses = []Address{*suite.OutboundAddress()}
				member.BankAccounts = []BankAccount{*suite.OutboundBankAccount()}
				member.Benefits = []Benefit{*suite.OutboundBenefit()}
				member.Contacts = []Contact{*suite.OutboundContact()}
				member.Contacts[0].Emails = suite.OutboundEmail()
				member.Contacts[0].Phones = []Phone{*suite.OutboundPhone()}
				member.ExtraData = []ExtraData{*suite.OutboundExtraData()}
				member.Partners = []Partner{*suite.OutboundPartner()}
				member.Partners[0].JuridicalPerson.Addresses = []Address{*suite.OutboundAddress()}
				member.Partners[0].NaturalPerson.Addresses = []Address{*suite.OutboundAddress()}
				member.Tags = suite.OutboundTag()
				member.WorkSchedule = []WorkSchedule{*suite.OutboundWorkSchedule()}

				actual, err := json.Marshal(&member)
				Expect(actual).Should(MatchJSON(suite.OutboundMemberJSON()))
				Expect(err).Should(BeNil())
			})
		})
	})
})

func TestMember(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Member Suite")
}
