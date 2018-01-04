package models

import (
	"github.com/stone-payments/logistic-sdk-go/basetesting"
)

const (
	testResources    = "./testresources"
	ouboundResources = testResources + "/oubound"
)

type TestSuite struct {
	basetesting.Suite
}

func (suite *TestSuite) OutboundAddressJSON() []byte {
	return suite.File(ouboundResources + "/address.json")
}

func (suite *TestSuite) OutboundAddress() *Address {
	var target Address
	suite.LoadJSON(ouboundResources+"/address.json", &target)
	return &target
}

func (suite *TestSuite) OutboundBankAccountJSON() []byte {
	return suite.File(ouboundResources + "/bankaccount.json")
}

func (suite *TestSuite) OutboundBankAccount() *BankAccount {
	var target BankAccount
	suite.LoadJSON(ouboundResources+"/bankaccount.json", &target)
	return &target
}

func (suite *TestSuite) OutboundBenefitJSON() []byte {
	return suite.File(ouboundResources + "/benefit.json")
}

func (suite *TestSuite) OutboundBenefit() *Benefit {
	var target Benefit
	suite.LoadJSON(ouboundResources+"/benefit.json", &target)
	return &target
}

func (suite *TestSuite) OutboundContactJSON() []byte {
	return suite.File(ouboundResources + "/contact.json")
}

func (suite *TestSuite) OutboundContact() *Contact {
	var target Contact
	suite.LoadJSON(ouboundResources+"/contact.json", &target)
	return &target
}

func (suite *TestSuite) OutboundExtraDataJSON() []byte {
	return suite.File(ouboundResources + "/extradata.json")
}

func (suite *TestSuite) OutboundExtraData() *ExtraData {
	var target ExtraData
	suite.LoadJSON(ouboundResources+"/extradata.json", &target)
	return &target
}

func (suite *TestSuite) OutboundEmailJSON() []byte {
	return suite.File(ouboundResources + "/email.json")
}

func (suite *TestSuite) OutboundEmail() []string {
	var target []string
	suite.LoadJSON(ouboundResources+"/email.json", &target)
	return target
}

func (suite *TestSuite) OutboundMemberJSON() []byte {
	return suite.File(ouboundResources + "/member.json")
}

func (suite *TestSuite) OutboundSystemConfigurationJSON() []byte {
	return suite.File(ouboundResources + "/system-configuration.json")
}

func (suite *TestSuite) OutboundMember() *Member {
	var target Member
	suite.LoadJSON(ouboundResources+"/member.json", &target)
	return &target
}

func (suite *TestSuite) OutboundPartnerJSON() []byte {
	return suite.File(ouboundResources + "/partner.json")
}

func (suite *TestSuite) OutboundPartner() *Partner {
	var target Partner
	suite.LoadJSON(ouboundResources+"/partner.json", &target)
	return &target
}

func (suite *TestSuite) OutboundPhoneJSON() []byte {
	return suite.File(ouboundResources + "/phone.json")
}

func (suite *TestSuite) OutboundPhone() *Phone {
	var target Phone
	suite.LoadJSON(ouboundResources+"/phone.json", &target)
	return &target
}

func (suite *TestSuite) OutboundTagJSON() []byte {
	return suite.File(ouboundResources + "/tag.json")
}

func (suite *TestSuite) OutboundTag() []string {
	var target []string
	suite.LoadJSON(ouboundResources+"/tag.json", &target)
	return target
}

func (suite *TestSuite) OutboundWorkScheduleJSON() []byte {
	return suite.File(ouboundResources + "/workschedule.json")
}

func (suite *TestSuite) OutboundWorkSchedule() *WorkSchedule {
	var target WorkSchedule
	suite.LoadJSON(ouboundResources+"/workschedule.json", &target)
	return &target
}
