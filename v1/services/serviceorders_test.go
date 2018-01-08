package services

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ServiceOrder", func() {

	service := NewServiceOrder("url", new(Credentials)).(*serviceOrder)

	Describe("When constructing service order service", func() {
		It("should set url attribute", func() {
			Expect(service.baseURL).To(Equal("url"))
		})
		It("should set credentials", func() {
			Expect(service.credentials).To(Equal(new(Credentials)))
		})
		It("should set response processor", func() {
			Expect(service.interceptors).To(HaveLen(1))
		})
	})

	Describe("When getting service order", func() {

	})
})
