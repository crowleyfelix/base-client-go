package http

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Manager", func() {

	var (
		url          = "url"
		interceptors []Interceptor
		manager      = NewManager("url", interceptors...).(*manager)
	)

	Describe("NewManager", func() {
		It("should set interceptors", func() {
			Expect(manager.interceptors).To(Equal(interceptors))
		})
		It("should set base url", func() {
			Expect(manager.baseURL).To(Equal(url))
		})
	})
	Describe("BuildURL", func() {
		It("should build url with params", func() {
			Expect(manager.BuildURL("/%v/%v/%v", 1, 2, 3)).To(Equal(url + "/1/2/3"))
		})
	})
	PDescribe("Request", func() {
		Context("When request failed", func() {
			It("should call OnRequest of external interceptors event", func() {

			})
			It("should call OnRequest of local interceptors event", func() {

			})
			It("should call OnRequest of parameterized interceptors event", func() {

			})
			It("should return an error", func() {

			})
		})
		Context("When request success", func() {
			Context("and interceptor return an error", func() {
				It("should not call next interceptors", func() {

				})
				It("should return an error", func() {

				})
			})
			Context("and executed with success", func() {
				It("should call OnResponse of external interceptors event", func() {

				})
				It("should call OnResponse of local interceptors event", func() {

				})
				It("should call OnResponse of parameterized interceptors event", func() {

				})
				It("should return response", func() {

				})
			})
		})
	})
})
