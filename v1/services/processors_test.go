package services

import (
	"github.com/bouk/monkey"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stone-payments/logistic-sdk-go/errors"
	"github.com/stone-payments/logistic-sdk-go/http"
	"github.com/stone-payments/logistic-sdk-go/http/mocks"
)

var _ = Describe("Processors", func() {
	Describe("processResponse", func() {

		//input
		var (
			resp *mocks.Response
		)

		//output
		var (
			actualResponse http.Response
			err            errors.Error
		)

		BeforeEach(func() {
			resp = new(mocks.Response)
		})
		JustBeforeEach(func() {
			actualResponse, err = processResponse(resp)
		})
		AfterEach(func() { monkey.UnpatchAll() })

		Context("When response is not ok", func() {

			var (
				expectedError errors.Error
			)

			BeforeEach(func() {
				resp.On("Ok").Return(false).Once()
				monkey.Patch(trackError, func(_ http.Response) errors.Error {
					expectedError = errors.NewBadRequest()
					return expectedError
				})
			})
			It("should track error", func() {
				Expect(err).To(Equal(expectedError))
			})
		})
		Context("When response is ok", func() {

			BeforeEach(func() {
				resp.On("Ok").Return(true).Once()
			})
			It("should not return an error", func() {
				Expect(err).To(BeNil())
				Expect(actualResponse).To(Equal(resp))
			})
		})
	})

	Describe("trackError", func() {

		//input
		var (
			resp *mocks.Response
		)

		var (
			messages     = []string{"message"}
			errorBuilded errors.HTTPError
		)

		//output
		var (
			err errors.Error
		)

		BeforeEach(func() {
			resp = new(mocks.Response)
			resp.On("StatusCode").Return(500).Once()
			monkey.Patch(errorMessages, func(_ http.Response) []string {
				return messages
			})
			monkey.Patch(errors.Build, func(statusCode int, _ ...string) errors.HTTPError {
				Expect(statusCode).To(Equal(500))
				errorBuilded = errors.NewInternalServer()
				return errorBuilded
			})
		})
		JustBeforeEach(func() {
			err = trackError(resp)
		})
		AfterEach(func() { monkey.UnpatchAll() })

		It("should build error based on status code", func() {
			Expect(err).To(Equal(errorBuilded))
		})
	})
})
