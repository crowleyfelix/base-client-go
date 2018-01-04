package services

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stone-payments/logistic-sdk-go/basetesting"
	"github.com/stone-payments/logistic-sdk-go/v1/models"
)

type TestSuite struct {
	basetesting.Suite
}

func (s *TestSuite) Members() []models.Member {
	config := s.Config()
	svc := memberService{
		&manager{
			config.URL,
			Credentials{config.SecretKey, config.AppKey, config.User},
		},
	}

	var (
		members []models.Member
		err     error
	)

	Describe("When listing merchants", func() {

		members, _, err = svc.List()

		It("should not return merchants empty", func() {
			Expect(members).ShouldNot(BeEmpty())
		})
		It("should not return an error", func() {
			Expect(err).Should(BeNil())
		})
	})

	return members
}

func (s *TestSuite) SystemConfigurations() []models.SystemConfiguration {
	config := s.Config()
	svc := systemConfigurationsService{
		&manager{
			config.URL,
			Credentials{config.SecretKey, config.AppKey, config.User},
		},
	}

	var (
		configs []models.SystemConfiguration
		err     error
	)

	Describe("When listing SystemConfigurations", func() {
		configs, err = svc.List()

		It("should not return empty", func() {
			Expect(configs).ShouldNot(BeEmpty())
		})
		It("should not return an error", func() {
			Expect(err).Should(BeNil())
		})
	})

	return configs
}
