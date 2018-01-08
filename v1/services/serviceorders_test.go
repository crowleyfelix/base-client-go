package services

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/bouk/monkey"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stone-payments/logistic-sdk-go"
	"github.com/stone-payments/logistic-sdk-go/errors"
	"github.com/stone-payments/logistic-sdk-go/http"
	"github.com/stone-payments/logistic-sdk-go/http/mocks"
	"github.com/stone-payments/logistic-sdk-go/v1/models"
)

var _ = Describe("Member", func() {

	var (
		baseURL   = "http://url"
		requester = new(mocks.Requestable)
		request   = new(mocks.Request)
		response  *mocks.Response
	)

	var service serviceOrder
	BeforeEach(func() {
		logisticsdk.RegisterRequester(requester)
		monkey.Patch(http.NewRequest, func() http.Request {
			return request
		})
		response = new(mocks.Response)

		service = serviceOrder{
			&manager{
				baseURL: baseURL,
			},
		}
	})
	AfterSuite(func() {
		monkey.UnpatchAll()
	})

	Describe("Get", func() {
		//input
		var (
			id string
		)

		var (
			requestError error
		)

		//output
		var (
			data *models.ServiceOrder
			err  error
		)
		JustBeforeEach(func() {
			data, err = service.Get(id)
		})
		BeforeEach(func() {
			id = strconv.Itoa(rand.Int())
			url := fmt.Sprintf("%s/v1/serviceorders/%s", baseURL, id)
			requester.On("Get", url, request).
				Return(response, requestError).Once()
		})
		Describe("When sending request", func() {
			Context("and failed", func() {
				BeforeEach(func() {
					requestError = fmt.Errorf("falhou")
				})
				It("should return an error", func() {
					Expect(err).ToNot(BeNil())
				})
			})
			Context("and success", func() {
				It("should return service order", func() {

				})
			})
		})
	})
	Describe("ByStoneCode", func() {
		//input
		var (
			stonecode int
		)

		//context
		var (
			requestError error
		)

		//output
		var (
			data []models.ServiceOrder
			err  error
		)

		JustBeforeEach(func() {
			data, err = service.ByStoneCode(stonecode)
		})
		BeforeEach(func() {
			stonecode = rand.Int()
			url := fmt.Sprintf("%s/v1/merchants/%d/serviceorders", baseURL, stonecode)
			requester.On("Get", url, request).
				Return(response).Once()
		})
		Context("When sending request", func() {
			Context("and failed", func() {
				BeforeEach(func() {
					requestError = fmt.Errorf("falhou")
				})
				It("should return an error", func() {
					Expect(err).ToNot(BeNil())
				})
			})
			Context("and success", func() {
				It("should return service orders", func() {

				})
			})
		})
	})
	Describe("List", func() {

		//input
		var (
			filters *models.ServiceOrderFilters
		)

		//context
		var (
			requestError error
		)

		//output
		var (
			data []models.ServiceOrder
			err  errors.Error
		)

		JustBeforeEach(func() {
			data, err = service.List(filters)
		})
		BeforeEach(func() {
			url := fmt.Sprintf("%s/v1/serviceorders", baseURL)
			requester.On("Get", url, request).
				Return(response).Once()
		})
		Context("When sending request", func() {
			Context("and failed", func() {
				BeforeEach(func() {
					requestError = fmt.Errorf("falhou")
				})
				It("should return an error", func() {
					Expect(err).ToNot(BeNil())
				})
			})
			Context("and success", func() {
				It("should return service orders", func() {

				})
			})
		})
	})
})
