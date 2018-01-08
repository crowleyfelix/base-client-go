package services

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/bouk/monkey"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stone-payments/logistic-sdk-go"
	"github.com/stone-payments/logistic-sdk-go/http"
	"github.com/stone-payments/logistic-sdk-go/http/mocks"
)

var _ = Describe("Member", func() {

	var (
		baseURL   = "http://url"
		appKey    = "appKey"
		secretKey = "secretKey"
		user      = "user"
		signature = "signature"
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
				baseURL:     baseURL,
				credentials: &Credentials{secretKey, appKey, user},
			},
		}
	})

	Describe("Get", func() {
		//input
		var (
			id string
		)

		//output
		var (
			err error
		)

		BeforeEach(func() {
			id = strconv.Itoa(rand.Int())
			url := fmt.Sprintf("%s/serviceorders/%s", baseURL, id)
			requester.On("Get", url, request).
				Return(response, err).Once()
			service.Get(id)
		})
		Context("When sending request", func() {
			It("should pass url builded", func() {

				Expect(actualURL).Should(Equal(expect))
			})
			It("should pass credentials in options", func() {
				expect := map[string]string{
					"Contentoutput-Type": "application/json",
					"Authorization":      signature,
					"User-Identifier":    user,
				}
				Expect(actualOptions.Headers).Should(Equal(expect))
			})
		})
	})
	Describe("ByStoneCode", func() {
		//input
		var (
			stonecode int
		)

		//output
		var (
			err error
		)

		BeforeEach(func() {
			stonecode = rand.Int()
			url := fmt.Sprintf("%s/merchants/%d/serviceorders", baseURL, stonecode)
			requester.On("Get", url, request).
				Return(response).Once()
			service.ByStoneCode(stonecode)
		})
		Context("When sending request", func() {
			It("should pass url builded", func() {
				expect := fmt.Sprintf("%s/merchants/%d/serviceorders", baseURL, stonecode)
				Expect(actualURL).Should(Equal(expect))
			})
			It("should pass credentials in options", func() {
				expect := map[string]string{
					"Content-Type":    "application/json",
					"Authorization":   signature,
					"User-Identifier": user,
				}
				Expect(actualOptions.Headers).Should(Equal(expect))
			})
		})
	})
	Describe("List", func() {

		//input
		var ()

		//output
		var ()

		BeforeEach(func() {
			url := fmt.Sprintf("%s/serviceorders", baseURL)
			requester.On("Get", url, request).
				Return(response).Once()
			service.List()
		})
		Context("When sending request", func() {
			It("should pass url builded", func() {
				expect := fmt.Sprintf("%s/serviceorders", baseURL)
				Expect(actualURL).Should(Equal(expect))
			})
			It("should pass credentials in options", func() {
				expect := map[string]string{
					"Content-Type":    "application/json",
					"Authorization":   signature,
					"User-Identifier": user,
				}
				Expect(actualOptions.Headers).Should(Equal(expect))
			})
		})
	})
})
