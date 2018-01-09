// +build integration

package services

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stone-payments/logistic-sdk-go/errors"
	. "github.com/stone-payments/logistic-sdk-go/testing"
	"github.com/stone-payments/logistic-sdk-go/v1/models"
)

var _ = Describe("ServiceOrder Integration", func() {

	var (
		config  = Config()
		service ServiceOrder
		page    *models.ServiceOrdersPage
	)

	BeforeSuite(func() {
		service = NewServiceOrder(config.URL)

		var err error
		page, err = service.List(new(models.ServiceOrderFilters))

		Expect(err).To(BeNil())
		Expect(page.ServiceOrders).ToNot(BeEmpty())
	})

	Describe("Get", func() {
		var (
			actual *models.ServiceOrder
			err    errors.Error
		)
		BeforeEach(func() {
			actual, err = service.Get(page.ServiceOrders[0].ServiceNumber)
		})
		Context("When id passed", func() {
			It("Should return service order", func() {
				Expect(actual).ShouldNot(BeNil())
				Expect(err).Should(BeNil())
			})
		})
	})

	Describe("ByStoneCode", func() {
		var (
			actual *models.ServiceOrdersPage
			err    errors.Error
		)
		BeforeEach(func() {
			actual, err = service.ByStoneCode(page.ServiceOrders[0].Client.StoneCode)
		})
		Context("When stonecode passed", func() {
			It("Should return service orders", func() {
				Expect(actual).ShouldNot(BeNil())
				Expect(err).Should(BeNil())
			})
		})
	})
})
