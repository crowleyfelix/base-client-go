// +build !integration

package models

import (
	"testing"

	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SystemConfiguration", func() {

	var (
		suite *TestSuite
	)

	BeforeEach(func() {
		suite = new(TestSuite)
	})

	Describe("MarshalJSON", func() {
		Context("When serializing", func() {
			It("should map enumerated ids", func() {
				var systemConfiguration SystemConfiguration

				err := json.Unmarshal(suite.OutboundSystemConfigurationJSON(), &systemConfiguration)
				Expect(err).Should(BeNil())

				actual, err := json.Marshal(&systemConfiguration)
				Expect(actual).Should(MatchJSON(suite.OutboundSystemConfigurationJSON()))
				Expect(err).Should(BeNil())
			})
		})
	})
})

func TestSystemConfiguration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SystemConfiguration Suite")
}
