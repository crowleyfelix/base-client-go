package http

import (
	"io"
	"io/ioutil"
	"strings"
	"net/http"
	"github.com/crowleyfelix/base-client-go/errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Response", func() {

	var (
		resp *Response
	)

	BeforeEach(func() {
		resp = &Response{
			Response: new(http.Response),
		}
	})

	Describe("Ok", func() {
		Context("When status code is minor than 200", func () {
			It("should return false", func() {
				resp.StatusCode = 199
				Expect(resp.Ok()).To(BeFalse())
			})
		})
		Context("When status code is between 200 and 400", func () {
			It("should return true", func() {
				resp.StatusCode = 200
				Expect(resp.Ok()).To(BeTrue())
			})
		})
		Context("When status code is greater than 399", func () {
			It("should return true", func() {
				resp.StatusCode = 400
				Expect(resp.Ok()).To(BeFalse())
			})
		})
	})
	Describe("JSON", func() {
		var (
			err errors.Error
			body io.ReadCloser
			actual string
		)

		JustBeforeEach(func(){
			resp.Body = body
			err = resp.JSON(&actual)
		})
		
		Context("When body is json desserializable", func () {
			BeforeEach(func() {
				body = ioutil.NopCloser(strings.NewReader("\"teste\""))
			})
			It("should load json data", func() {
				Expect(err).To(BeNil())
				Expect(actual).To(Equal("teste"))		
			})
		})
	})
})
