package services

import (
	"encoding/json"
	"math/rand"
	"reflect"
	"strconv"

	"github.com/bouk/monkey"
	"github.com/stretchr/testify/mock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stone-payments/logistic-sdk-go/errors"
	"github.com/stone-payments/logistic-sdk-go/http"
	"github.com/stone-payments/logistic-sdk-go/http/mocks"
	. "github.com/stone-payments/logistic-sdk-go/testing"
	"github.com/stone-payments/logistic-sdk-go/v1/models"
)

var _ = Describe("ServiceOrder", func() {

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
	AfterEach(func() { monkey.UnpatchAll() })

	Describe("Get", func() {
		//input
		var (
			number int
		)

		var (
			urlBuilded   = "builded"
			requestError errors.Error
		)

		//output
		var (
			data *models.ServiceOrder
			err  errors.Error
		)
		BeforeEach(func() {
			number = rand.Int()
			manager.On("BuildURL", "/v1/serviceorders/%v", number).
				Return(urlBuilded).Once()
		})
		JustBeforeEach(func() {
			data, err = service.Get(number)
		})
		Describe("When sending request", func() {
			Context("and failed", func() {
				BeforeEach(func() {
					requestError = errors.NewSerializing()
					manager.On("Request", mock.Anything, urlBuilded, mock.Anything).
						Return(response, requestError).Once()
				})
				It("should return an error", func() {
					Expect(err).To(Equal(requestError))
				})
			})
			Context("and success", func() {

				var (
					expected models.ServiceOrder
				)

				BeforeEach(func() {
					requestError = nil
					manager.On("Request", mock.Anything, urlBuilded, mock.Anything).
						Return(response, requestError).Once()
					LoadJSON("testresources/serviceorder.json", &expected)

					monkey.PatchInstanceMethod(reflect.TypeOf(response), "JSON", func(_ *mocks.Response, obj interface{}) errors.Error {
						data := File("testresources/serviceorder.json")

						err := json.Unmarshal(data, &obj)
						Expect(err).To(BeNil())
						return nil
					})
				})
				It("should return service order", func() {
					Expect(data).To(Equal(&expected))
					Expect(err).To(BeNil())
				})
			})
		})
	})
	Describe("ByStoneCode", func() {
		//input
		var (
			stonecode string
		)

		//context
		var (
			urlBuilded   = "builded"
			requestError errors.Error
		)

		//output
		var (
			data *models.ServiceOrdersPage
			err  error
		)

		BeforeEach(func() {
			stonecode = strconv.Itoa(rand.Int())
			manager.On("BuildURL", "/v1/merchants/%v/serviceorders", stonecode).
				Return(urlBuilded).Once()
		})
		JustBeforeEach(func() {
			data, err = service.ByStoneCode(stonecode)
		})
		Context("When sending request", func() {
			Context("and failed", func() {
				BeforeEach(func() {
					requestError = errors.NewSerializing()
					manager.On("Request", mock.Anything, urlBuilded, mock.Anything).
						Return(response, requestError).Once()
				})
				It("should return an error", func() {
					Expect(err).To(Equal(requestError))
				})
			})
			Context("and success", func() {

				var (
					expected models.ServiceOrdersPage
				)

				BeforeEach(func() {
					requestError = nil
					manager.On("Request", mock.Anything, urlBuilded, mock.Anything).
						Return(response, requestError).Once()
					LoadJSON("testresources/serviceorders.json", &expected)

					monkey.PatchInstanceMethod(reflect.TypeOf(response), "JSON", func(_ *mocks.Response, obj interface{}) errors.Error {
						data := File("testresources/serviceorders.json")

						err := json.Unmarshal(data, &obj)
						Expect(err).To(BeNil())
						return nil
					})
				})
				It("should return service orders", func() {
					Expect(data).To(Equal(&expected))
					Expect(err).To(BeNil())
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
			requestError errors.Error
			params       map[string]string
		)

		//output
		var (
			data *models.ServiceOrdersPage
			err  errors.Error
		)

		BeforeEach(func() {
			manager.On("BuildURL", "/v1/serviceorders").
				Return(urlBuilded).Once()

			params = map[string]string{
				"param": "value",
			}
			monkey.Patch(http.ToMap, func(obj interface{}) map[string]string {
				Expect(obj).To(BeAssignableToTypeOf(filters))
				return params
			})
			request.On("SetParams", params).Once()
		})
		JustBeforeEach(func() {
			data, err = service.List(filters)
		})
		Context("When sending request", func() {
			Context("and failed", func() {
				BeforeEach(func() {
					requestError = errors.NewSerializing()
					manager.On("Request", mock.Anything, urlBuilded, mock.Anything).
						Return(response, requestError).Once()
				})
				It("should return an error", func() {
					Expect(err).To(Equal(requestError))
				})
			})
			Context("and success", func() {

				var (
					expected models.ServiceOrdersPage
				)

				BeforeEach(func() {
					requestError = nil
					manager.On("Request", mock.Anything, urlBuilded, mock.Anything).
						Return(response, requestError).Once()
					LoadJSON("testresources/serviceorders.json", &expected)

					monkey.PatchInstanceMethod(reflect.TypeOf(response), "JSON", func(_ *mocks.Response, obj interface{}) errors.Error {
						data := File("testresources/serviceorders.json")

						err := json.Unmarshal(data, &obj)
						Expect(err).To(BeNil())
						return nil
					})
				})
				It("should return service orders", func() {
					Expect(data).To(Equal(&expected))
					Expect(err).To(BeNil())
				})
			})
		})
	})
})
