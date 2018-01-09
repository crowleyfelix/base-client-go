package services

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/bouk/monkey"
	"github.com/stretchr/testify/mock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stone-payments/logistic-sdk-go/errors"
	"github.com/stone-payments/logistic-sdk-go/http"
	"github.com/stone-payments/logistic-sdk-go/http/mocks"
	"github.com/stone-payments/logistic-sdk-go/v1/models"
)

var _ = Describe("Member", func() {

	var (
		service  serviceOrder
		manager  = new(mocks.Manager)
		request  = new(mocks.Request)
		response *mocks.Response
	)

	BeforeEach(func() {
		monkey.Patch(http.NewRequest, func() http.Request {
			return request
		})
		response = new(mocks.Response)

		service = serviceOrder{manager}
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
			urlBuilded   = "builded"
			requestError error
		)

		//output
		var (
			data *models.ServiceOrder
			err  error
		)
		BeforeEach(func() {
			id = strconv.Itoa(rand.Int())
			manager.On("BuildURL", "/v1/serviceorders/%v", id).
				Return(urlBuilded).Once()
			manager.On("Request", mock.Anything, urlBuilded, mock.Anything).
				Return(response, requestError).Once()
		})
		JustBeforeEach(func() {
			data, err = service.Get(id)
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
			urlBuilded   = "builded"
			requestError error
		)

		//output
		var (
			data []models.ServiceOrder
			err  error
		)

		BeforeEach(func() {
			stonecode = rand.Int()
			manager.On("BuildURL", "/v1/merchants/%v/serviceorders", stonecode).
				Return(urlBuilded).Once()
			manager.On("Request", mock.Anything, urlBuilded, mock.Anything).
				Return(response, requestError).Once()
		})
		JustBeforeEach(func() {
			data, err = service.ByStoneCode(stonecode)
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
			urlBuilded   = "builded"
			requestError error
		)

		//output
		var (
			data []models.ServiceOrder
			err  errors.Error
		)

		BeforeEach(func() {
			manager.On("BuildURL", "/v1/serviceorders").
				Return(urlBuilded).Once()
			manager.On("Request", mock.Anything, urlBuilded, mock.Anything).
				Return(response, requestError).Once()
		})
		JustBeforeEach(func() {
			data, err = service.List(filters)
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
